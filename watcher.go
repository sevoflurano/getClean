package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)

var extensions = map[string]string{
	".jpg":  "Pictures/jpeg",
	".jpeg": "Pictures/jpeg",
	".png":  "Pictures/png",
	".gif":  "Pictures/gif",
	".webp": "Pictures/webp",
	".svg":  "Pictures/svg",
	".bmp":  "Pictures/bmp",
	".ico":  "Pictures/ico",

	".mp4":  "Videos/mp4",
	".mkv":  "Videos/mkv",
	".avi":  "Videos/avi",
	".mov":  "Videos/mov",
	".webm": "Videos/webm",

	".mp3":  "Music/mp3",
	".wav":  "Music/wav",
	".flac": "Music/flac",
	".ogg":  "Music/ogg",

	".pdf":  "Documents/pdf",
	".docx": "Documents/word",
	".xlsx": "Documents/excel",
	".pptx": "Documents/powerpoint",
	".txt":  "Documents/txt",

	".zip": "Archives",
	".rar": "Archives",
	".7z":  "Archives",
	".tar": "Archives",
	".gz":  "Archives",

	".exe": "Programs",
	".msi": "Programs",
	".dmg": "Programs",
	".iso": "Programs",
}

var ignoredExtensions = map[string]struct{}{
	".lnk":        {},
	".tmp":        {},
	".part":       {},
	".crdownload": {},
}

func WatchDownloads() error {
	downloads, err := getDownloadsPath()
	if err != nil {
		return err
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	if err := watcher.Add(downloads); err != nil {
		return err
	}

	var processing sync.Map

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return nil
			}

			if !shouldHandle(event) {
				continue
			}

			path := event.Name

			if _, loaded := processing.LoadOrStore(path, struct{}{}); loaded {
				continue
			}

			go func(path string) {
				defer processing.Delete(path)

				if err := organizeFile(path, downloads); err != nil {
					log.Println(err)
				}
			}(path)

		case err, ok := <-watcher.Errors:
			if !ok {
				return nil
			}
			log.Println(err)
		}
	}
}

func getDownloadsPath() (string, error) {
	custom := strings.TrimSpace(os.Getenv("GETCLEAN_DOWNLOADS"))
	if custom != "" {
		return custom, nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, "Downloads"), nil
}

func shouldHandle(event fsnotify.Event) bool {
	return event.Has(fsnotify.Create) || event.Has(fsnotify.Write)
}

func organizeFile(path, downloads string) error {
	ext := strings.ToLower(filepath.Ext(path))
	if ext == "" {
		return nil
	}

	if _, ignored := ignoredExtensions[ext]; ignored {
		return nil
	}

	folder, ok := extensions[ext]
	if !ok {
		return nil
	}

	ready, err := waitUntilReady(path, 2*time.Minute)
	if err != nil {
		return err
	}
	if !ready {
		return nil
	}

	destDir := filepath.Join(downloads, folder)
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return err
	}

	destPath := uniquePath(filepath.Join(destDir, filepath.Base(path)))

	if err := os.Rename(path, destPath); err != nil {
		return err
	}

	return nil
}

func waitUntilReady(path string, timeout time.Duration) (bool, error) {
	deadline := time.Now().Add(timeout)

	var lastSize int64 = -1
	stableChecks := 0

	for time.Now().Before(deadline) {
		info, err := os.Stat(path)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				return false, nil
			}
			time.Sleep(time.Second)
			continue
		}

		if info.IsDir() {
			return false, nil
		}

		size := info.Size()

		if size == lastSize {
			stableChecks++
			if stableChecks >= 2 {
				return true, nil
			}
		} else {
			lastSize = size
			stableChecks = 0
		}

		time.Sleep(2 * time.Second)
	}

	return false, nil
}

func uniquePath(path string) string {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return path
	}

	dir := filepath.Dir(path)
	ext := filepath.Ext(path)
	base := strings.TrimSuffix(filepath.Base(path), ext)

	for i := 1; ; i++ {
		candidate := filepath.Join(dir, fmt.Sprintf("%s (%d)%s", base, i, ext))
		if _, err := os.Stat(candidate); errors.Is(err, os.ErrNotExist) {
			return candidate
		}
	}
}

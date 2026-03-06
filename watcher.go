package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
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
}

func Watcher() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer watcher.Close()

	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	downloads := filepath.Join(home, "Downloads")
	err = watcher.Add(downloads)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}

			if event.Has(fsnotify.Create) {
				ext := filepath.Ext(event.Name)
				ext = strings.ToLower(ext)

				if ext == ".lnk" {
					continue
				}

				time.Sleep(2 * time.Second)

				mafketel, vinden := extensions[ext]
				if vinden {

					destino := filepath.Join(home, mafketel)
					os.MkdirAll(destino, 0755)
					err := os.Rename(event.Name, filepath.Join(destino, filepath.Base(event.Name)))
					if err != nil {
						fmt.Println("error: Could not move:", err)
						continue
					}
				}
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			fmt.Println("error:", err)

		}
	}
}

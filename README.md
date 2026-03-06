# getClean

`getClean` is a small Go utility that automatically organizes your **Downloads** folder.

It watches your Downloads directory and moves files into categorized folders based on their file extension.

The goal of the project is simple: keep the Downloads folder clean without manual sorting.

## Features

* Watches the Downloads folder in real time
* Automatically moves files based on extension
* Creates destination folders if they don't exist
* Runs as a background service
* Lightweight and minimal resource usage

## Example

Before:

Downloads/
file.zip
setup.exe
movie.mp4

After:

Downloads/
Archives/file.zip
Programs/setup.exe
Videos/movie.mp4

## Categories

Files are grouped into folders depending on their extension:

Pictures → jpg, jpeg, png, gif
Videos → mp4, mkv, avi
Music → mp3, wav, flac
Documents → pdf, docx, txt
Archives → zip, rar, 7z
Programs → exe, msi

## Installation

Clone the repository:

```
git clone https://github.com/sevoflurano/getClean
cd getClean
```

Install dependencies:

```
go mod tidy
```

Run the program:

```
go run .
```

Build a binary:

```
go build
```

## Notes

This project was built mainly as a learning exercise while exploring Go and filesystem automation.

The design intentionally favors simplicity over complex configuration.

## License

MIT

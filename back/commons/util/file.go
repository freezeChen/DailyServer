package util

import (
	"errors"
	"os"
	"path"
)

// Mkdir create directory.
func Mkdir(Path string) error {
	p, _ := path.Split(Path)
	if p == "" {
		return errors.New("invalid path")
	}
	d, err := os.Stat(p)
	if err != nil || !d.IsDir() {
		err = os.MkdirAll(p, 0777)
	}
	return err
}

// IsDirExists returns true if the path is exists.
func IsDirExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	} else {
		return fi.IsDir()
	}
	return false
}

// IsImage returns true if the file is image.
func IsImage(filetype string) bool {
	switch filetype {
	case ".jpeg", ".jpg", ".gif", ".png", ".bmp":
		return true
	default:
		return false
	}
}

// IsAudio returns true if the file is audio.
func IsAudio(filetype string) bool {
	switch filetype {
	case ".wav", ".mp3", ".wma", ".ogg", ".ape", ".acc":
		return true
	default:
		return false
	}
}

// IsAudio returns true if the file is video.
func IsVideo(filetype string) bool {
	switch filetype {
	case ".avi", ".mov", ".asf", ".wmv", ".flv", ".mp4":
		return true
	default:
		return false
	}
}

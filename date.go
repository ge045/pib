package pib

import (
	"fmt"
	"github.com/xiam/exif"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	data, _ := exif.Read("_examples/resources/test.jpg")
	for key, val := range data.Tags {
		fmt.Printf("%s = %s\n", key, val)
	}
}

func Extend(slice []string, element string) []string {
	n := len(slice)
	if n == cap(slice) {
		// Slice is full; must grow.
		// We double its size and add 1, so if the size is zero we still
		// grow.
		newSlice := make([]string, len(slice), 2*len(slice)+1)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}

func HasSuffix(s string) bool {
	suffices := [...]string{".png", ".jpg", ".JPG"}
	for _, suffix := range suffices {
		if strings.HasSuffix(s, suffix) {
			return true
		}
	}
	return false
}

func PictureScan(root string, ext string) []string {
	pictureSlice := make([]string, 0, 10)
	filepath.Walk(root, func(path string, _ os.FileInfo, _ error) error {
		if HasSuffix(path) {
			pictureSlice = Extend(pictureSlice, path)
		}
		return nil
	})
	return pictureSlice
}
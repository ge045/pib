package main

import (
	"flag"
	"fmt"
	"github.com/xiam/exif"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	currDir, _ := os.Getwd()
	inPtr := flag.String("in", currDir,
		"The path to be scanned recursively.")

	outdir, err := ioutil.TempDir("/tmp", "pid")
	if err != nil {
		log.Fatal(err)
	}

	flag.Parse()
	fmt.Printf("IN: %s\n", *inPtr)
	fmt.Printf("TMP: %s\n", outdir)

	pics := PictureScan(*inPtr)
	for _, pic := range pics {
		fmt.Printf("processing %s\n", pic)
		GetExifDate(pic)
	}
}

func GetExifDate(pic string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Failed to read EXIF data for file ", pic)
		}
	}()
	data, _ := exif.Read(pic)
	for key, val := range data.Tags {
		fmt.Printf("\t%s = %s\n", key, val)
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
	suffices := [...]string{".JPG", ".png", ".jpg"}
	for _, suffix := range suffices {
		if strings.HasSuffix(s, suffix) {
			return true
		}
	}
	return false
}

func PictureScan(root string) []string {
	pictureSlice := make([]string, 0, 10)
	filepath.Walk(root, func(path string, _ os.FileInfo, _ error) error {
		if HasSuffix(path) {
			pictureSlice = Extend(pictureSlice, path)
		}
		return nil
	})
	return pictureSlice
}

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
	"time"
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
	const format = "2006:01:02 15:04:05"
	var datetime, current, uninizialized time.Time
	var final_key string
	data, _ := exif.Read(pic)
	for key, val := range data.Tags {
		if strings.Contains(key, "Date and Time") {
			fmt.Printf("\t\t%s %s\n", key, val)
			current, _ = time.Parse(format, val)
			if (datetime == uninizialized) || current.Before(datetime) {
				datetime = current
				final_key = key
			}
		}
	}
	fmt.Printf("\t%s = %s\n", final_key, datetime)
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
	suffices := [...]string{".JPG", ".png", ".jpg", ".jpeg", ".JPEG"}
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

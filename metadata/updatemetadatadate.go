package metadata

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func UpdateMetadataDate(path string, info os.FileInfo, fileDateTime *time.Time, override bool, replaces map[string]string) {
	if !info.IsDir() && strings.HasSuffix(strings.ToLower(path), ".jpg") ||
		strings.HasSuffix(strings.ToLower(path), ".jpeg") ||
		strings.HasSuffix(strings.ToLower(path), ".gif") {
		updateMetadataDateJpg(path, fileDateTime, override, replaces)
	} else if !info.IsDir() && strings.HasSuffix(strings.ToLower(path), ".mpg") ||
		strings.HasSuffix(strings.ToLower(path), ".mpeg") {
		updateMetadataDateMpeg(path, fileDateTime, override, replaces)
	}
}
func updateMetadataDateMpeg(filepath string, fileDateTime *time.Time, override bool, replaces map[string]string) {
	fmt.Println("    - file: ", filepath)
	var existingFileDateTime = getModTime(filepath)
	var newDate = getNewDate(existingFileDateTime, fileDateTime, override, replaces)
	if newDate != nil {
		os.Chtimes(filepath, *newDate, *newDate)
	}
}
func updateMetadataDateJpg(filepath string, fileDateTime *time.Time, override bool, replaces map[string]string) {
	fmt.Println("    - file: ", filepath)
	var existingFileDateTime = extractExifMetadataDate(filepath)
	var newDate = getNewDate(existingFileDateTime, fileDateTime, override, replaces)
	if newDate == nil && override {
		newDate = getModTime(filepath)
	}

	if newDate != nil {
		setExifMetadataDate(filepath, *newDate)
	}
}
func getNewDate(existingFileDateTime *time.Time, fileDateTime *time.Time, override bool, replaces map[string]string) *time.Time {
	// existing metadata
	if existingFileDateTime != nil {
		fmt.Println("      existingDateTime: ", existingFileDateTime)
	} else {
		return fileDateTime
	}

	if fileDateTime != nil && override {
		fmt.Println("      set - newDateTime: ", fileDateTime)
		return fileDateTime
	}

	if existingFileDateTime, override = processReplaces(existingFileDateTime, replaces); override {
		fmt.Println("      set - newDateTime: ", existingFileDateTime)
		return existingFileDateTime
	}
	fmt.Println("      keep - existingDateTime")
	return nil
}

func getModTime(filepath string) *time.Time {
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		log.Fatal(err)
	}
	result := fileInfo.ModTime()
	return &result
}

func processReplaces(dateTime *time.Time, replaces map[string]string) (*time.Time, bool) {
	layout := "2006-01-02T15:04:05Z07:00"

	key := dateTime.Format("2006-01-02")
	if val, ok := replaces[key]; ok {
		orig := dateTime.Format(layout)
		newVal := val + orig[10:]

		k, err := time.Parse(layout, newVal)
		if err != nil {
			log.Fatal(err)
		}
		return &k, true
	}
	return dateTime, false
}

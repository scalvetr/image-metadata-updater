package metadata

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func UpdateMetadataDate(path string, info os.FileInfo, fileDateTime *time.Time, override bool, replaces map[string]string, dateFilePatterns []string) {
	if !info.IsDir() && strings.HasSuffix(strings.ToLower(path), ".jpg") ||
		strings.HasSuffix(strings.ToLower(path), ".jpeg") ||
		strings.HasSuffix(strings.ToLower(path), ".gif") {
		updateMetadataDateJpg(path, fileDateTime, override, replaces, dateFilePatterns)
	} else if !info.IsDir() && strings.HasSuffix(strings.ToLower(path), ".mpg") ||
		strings.HasSuffix(strings.ToLower(path), ".mpeg") {
		updateMetadataDateMpeg(path, fileDateTime, override, replaces, dateFilePatterns)
	}
}
func updateMetadataDateMpeg(filePath string, fileDateTime *time.Time, override bool, replaces map[string]string, dateFilePatterns []string) {
	fmt.Println("    - file: ", filePath)
	var existingFileDateTime = getModTime(filePath)
	var newDate = getNewDate(filePath, existingFileDateTime, fileDateTime, override, replaces, dateFilePatterns)
	if newDate != nil {
		os.Chtimes(filePath, *newDate, *newDate)
	}
}
func updateMetadataDateJpg(filePath string, fileDateTime *time.Time, override bool, replaces map[string]string, dateFilePatterns []string) {
	fmt.Println("    - file: ", filePath)
	var existingFileDateTime = extractExifMetadataDate(filePath)
	var newDate = getNewDate(filePath, existingFileDateTime, fileDateTime, override, replaces, dateFilePatterns)
	if newDate == nil && override {
		newDate = getModTime(filePath)
	}

	if newDate != nil {
		setExifMetadataDate(filePath, *newDate)
	}
}
func getNewDate(
	fileName string,
	existingFileDateTime *time.Time,
	fileDateTime *time.Time,
	override bool,
	replaces map[string]string,
	dateFilePatterns []string,
) *time.Time {

	var result *time.Time
	// there's already a date in the metadata
	if existingFileDateTime != nil {
		fmt.Println("      existingDateTime: ", existingFileDateTime)
		// replaces are always on the existing date in the metadata.
		var replaced bool
		if result, replaced = processReplaces(existingFileDateTime, replaces); replaced {
			fmt.Println("      set - newDateTime (replaces): ", result)
		}
	}

	// result = existing date + replaces
	if result != nil && !override {
		// if override is set to false, just return what we got
		return result
	}

	// try first to extract the date from the filePatterns
	var extracted bool
	if result, extracted = processDateFilePatterns(fileName, dateFilePatterns); extracted {
		fmt.Println("      set - newDateTime (filePatterns): ", result)
		return result
	}

	if fileDateTime != nil {
		fmt.Println("      set - newDateTime (fixed): ", fileDateTime)
		return fileDateTime
	}

	fmt.Println("      keep - existingDateTime")
	return nil
}

func processDateFilePatterns(fileName string, patterns []string) (*time.Time, bool) {
	return nil, false
}

func getModTime(filePath string) *time.Time {
	fileInfo, err := os.Stat(filePath)
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

package metadata

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func UpdateMetadataDate(filePath string, info os.FileInfo, defaultDate *time.Time, override bool, fileDate bool, replaces map[string]string, dateFilePatterns []string) {
	if !info.IsDir() && strings.HasSuffix(strings.ToLower(filePath), ".jpg") ||
		strings.HasSuffix(strings.ToLower(filePath), ".jpeg") ||
		strings.HasSuffix(strings.ToLower(filePath), ".gif") {
		updateMetadataDateJpg(filePath, defaultDate, override, fileDate, replaces, dateFilePatterns)
	} else if !info.IsDir() && strings.HasSuffix(strings.ToLower(filePath), ".mpg") ||
		strings.HasSuffix(strings.ToLower(filePath), ".mpeg") ||
		strings.HasSuffix(strings.ToLower(filePath), ".mp4") {
		updateMetadataDateMpeg(filePath, defaultDate, override, fileDate, replaces, dateFilePatterns)
	}
}
func updateMetadataDateMpeg(filePath string, defaultDate *time.Time, override bool, fileDate bool, replaces map[string]string, dateFilePatterns []string) {
	fmt.Println("    - file: ", filePath)
	if fileDate {
		defaultDate = getModTime(filePath)
	}
	var existingFileDateTime = getModTime(filePath)
	var newDate = getNewDate(filePath, existingFileDateTime, defaultDate, override, replaces, dateFilePatterns)
	if newDate != nil {
		os.Chtimes(filePath, *newDate, *newDate)
	}
}
func updateMetadataDateJpg(filePath string, defaultDate *time.Time, override bool, fileDate bool, replaces map[string]string, dateFilePatterns []string) {
	fmt.Println("    - file: ", filePath)
	if fileDate {
		defaultDate = getModTime(filePath)
	}
	var existingFileDateTime = extractExifMetadataDate(filePath)
	var newDate = getNewDate(filePath, existingFileDateTime, defaultDate, override, replaces, dateFilePatterns)
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
	defaultDate *time.Time,
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

	if defaultDate != nil {
		fmt.Println("      set - newDateTime (fixed): ", defaultDate)
		return defaultDate
	}

	fmt.Println("      keep - existingDateTime")
	return nil
}

func processDateFilePatterns(fileName string, patterns []string) (*time.Time, bool) {
	loc, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		loc = time.Local
	}
	for _, pattern := range patterns {
		pattern := regexp.MustCompile(pattern)
		if pattern.MatchString(fileName) {
			matches := pattern.FindStringSubmatch(fileName)
			result := make(map[string]string)
			for i, name := range pattern.SubexpNames() {
				if i != 0 && name != "" {
					result[name] = matches[i]
				}
			}
			var err error
			year, err := strconv.Atoi(result["year"])
			if err != nil {
				log.Fatal(err)
			}
			month, err := strconv.Atoi(result["month"])
			if err != nil {
				log.Fatal(err)
			}
			day, err := strconv.Atoi(result["day"])
			if err != nil {
				log.Fatal(err)
			}
			hour, err := strconv.Atoi(result["hour"])
			if err != nil {
				log.Println(err)
				hour = 00
			}
			minute, err := strconv.Atoi(result["minute"])
			if err != nil {
				log.Println(err)
				minute = 01
			}
			second, err := strconv.Atoi(result["second"])
			if err != nil {
				log.Println(err)
				second = 01
			}

			returnDate := time.Date(year, time.Month(month), day, hour, minute, second, 00, loc)
			return &returnDate, true
		}
	}
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

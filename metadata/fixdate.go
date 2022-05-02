package metadata

import (
	"fmt"
	"strings"
)

func FixDate(filepath string, year int, month int) {
	if strings.HasSuffix(strings.ToLower(filepath), ".jpg") ||
		strings.HasSuffix(strings.ToLower(filepath), ".jpeg") ||
		strings.HasSuffix(strings.ToLower(filepath), ".gif") {
		fixDateDateJpg(filepath, year, month)
	} else if strings.HasSuffix(strings.ToLower(filepath), ".mpg") ||
		strings.HasSuffix(strings.ToLower(filepath), ".mpeg") {
		fixDateDateMpeg(filepath, year, month)
	}
}
func fixDateDateMpeg(filepath string, year int, month int) {
	fmt.Println("    - file: ", filepath)
	var existingFileDateTime = getModTime(filepath)
	fmt.Printf("year => \nalbum: %s \nfile: %s", year, existingFileDateTime.Year())
	fmt.Printf("month => \nalbum: %s \nfile: %s", month, existingFileDateTime.Month())
}
func fixDateDateJpg(filepath string, year int, month int) {
	fmt.Println("    - file: ", filepath)
	var existingFileDateTime = extractExifMetadataDate(filepath)
	if existingFileDateTime == nil {
		fmt.Println("no date specified")
	} else {
		fmt.Printf("year => \n  album: %s \n  file: %s\n", year, existingFileDateTime.Year())
		fmt.Printf("month => \n  album: %s \n  file: %s\n", month, int(existingFileDateTime.Month()))
	}
}

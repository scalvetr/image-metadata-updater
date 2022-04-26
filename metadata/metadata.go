package metadata

import (
	"fmt"
	"github.com/dsoprea/go-exif/v3"
	exifcommon "github.com/dsoprea/go-exif/v3/common"
	"github.com/dsoprea/go-logging"
	album "image-metadata-updater/album"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func UpdateDateFromMetadata(albumInfo album.AlbumInfo, path string, info os.FileInfo) {
	if strings.HasSuffix(strings.ToLower(path), ".jpg") ||
		strings.HasSuffix(strings.ToLower(path), ".jpeg") ||
		strings.HasSuffix(strings.ToLower(path), ".gif") {
		processJpg(albumInfo, path, info)
	} else if strings.HasSuffix(strings.ToLower(path), ".mpg") ||
		strings.HasSuffix(strings.ToLower(path), ".mpeg") {
		processMpg(albumInfo, path, info)
	}
}

func processMpg(info album.AlbumInfo, path string, info2 os.FileInfo) {

}
func processJpg(albumInfo album.AlbumInfo, filepath string, info os.FileInfo) {
	metadata := extractExifMetadata(albumInfo, filepath, info)
	var fileDateTime *time.Time
	layout := "2006:01:02 15:04:05 -07"
	for _, ifdEntry := range metadata {
		fmt.Println(ifdEntry.TagName, ifdEntry.Value)
		valueStr, _ := ifdEntry.Value.(string)

		if ifdEntry.TagName == "DateTime" {
			k, _ := time.Parse(layout, valueStr+" +02")
			fileDateTime = &k
			break
		}
	}
	if fileDateTime != nil {
		fmt.Println("DateTime", fileDateTime)
		err := os.Chtimes(filepath, *fileDateTime, *fileDateTime)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func extractExifMetadata(albumInfo album.AlbumInfo, filepath string, info os.FileInfo) []IfdEntry {
	f, err := os.Open(filepath)
	log.PanicIf(err)
	data, err := ioutil.ReadAll(f)
	log.PanicIf(err)
	rawExif, err := exif.SearchAndExtractExif(data)
	if err != nil {
		return nil
	}

	im, err := exifcommon.NewIfdMappingWithStandard()
	log.PanicIf(err)

	ti := exif.NewTagIndex()

	_, index, err := exif.Collect(im, ti, rawExif)
	log.PanicIf(err)

	entries := make([]IfdEntry, 0)
	for _, thisIte := range index.RootIfd.Entries() {
		ite := thisIte
		value, err := ite.Value()
		if err != nil {
			continue
		}

		entry := IfdEntry{
			IfdPath:     ite.IfdPath(),
			FqIfdPath:   ite.ChildFqIfdPath(),
			TagId:       ite.TagId(),
			TagName:     ite.TagName(),
			TagTypeId:   ite.TagType(),
			UnitCount:   ite.UnitCount(),
			Value:       value,
			ValueString: ite.String(),
		}
		entries = append(entries, entry)
	}
	return entries

}

type IfdEntry struct {
	IfdPath     string                      `json:"ifd_path"`
	FqIfdPath   string                      `json:"fq_ifd_path"`
	IfdIndex    int                         `json:"ifd_index"`
	TagId       uint16                      `json:"tag_id"`
	TagName     string                      `json:"tag_name"`
	TagTypeId   exifcommon.TagTypePrimitive `json:"tag_type_id"`
	TagTypeName string                      `json:"tag_type_name"`
	UnitCount   uint32                      `json:"unit_count"`
	Value       interface{}                 `json:"value"`
	ValueString string                      `json:"value_string"`
}

package metadata

import (
	"fmt"
	"github.com/dsoprea/go-exif/v3"
	exifcommon "github.com/dsoprea/go-exif/v3/common"
	"github.com/dsoprea/go-logging"
	"io/ioutil"
	"os"
	"time"
)

func extractExifMetadata(filepath string, info os.FileInfo) []IfdEntry {
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

func extractExifMetadataDate(filepath string, info os.FileInfo) *time.Time {
	metadata := extractExifMetadata(filepath, info)
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
	return fileDateTime
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

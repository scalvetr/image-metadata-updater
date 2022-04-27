package metadata

import (
	"bytes"
	"fmt"
	exif "github.com/dsoprea/go-exif/v3"
	exifcommon "github.com/dsoprea/go-exif/v3/common"
	jpeg "github.com/dsoprea/go-jpeg-image-structure/v2"
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
		fmt.Println("      exifMetadata: nil")
		return nil
	}
	fmt.Println("      exifMetadata: found")

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

func setExifMetadataDate(filepath string, t time.Time) error {
	parser := jpeg.NewJpegMediaParser()
	intfc, err := parser.ParseFile(filepath)
	if err != nil {
		return fmt.Errorf("Failed to parse JPEG file: %v", err)
	}

	sl := intfc.(*jpeg.SegmentList)

	rootIb, err := sl.ConstructExifBuilder()
	if err != nil {
		fmt.Println("No EXIF; creating it from scratch")

		im, err := exifcommon.NewIfdMappingWithStandard()
		if err != nil {
			return fmt.Errorf("Failed to create new IFD mapping with standard tags: %v", err)
		}
		ti := exif.NewTagIndex()
		if err := exif.LoadStandardTags(ti); err != nil {
			return fmt.Errorf("Failed to load standard tags: %v", err)
		}

		rootIb = exif.NewIfdBuilder(im, ti, exifcommon.IfdStandardIfdIdentity,
			exifcommon.EncodeDefaultByteOrder)
		rootIb.AddStandardWithName("ProcessingSoftware", "photos-uploader")
	}

	//TODO check if DateTime is already set.

	// Form our timestamp string
	ts := exifcommon.ExifFullTimestampString(t)

	// Set DateTime
	ifdPath := "IFD0"
	if err := setExifTag(rootIb, ifdPath, "DateTime", ts); err != nil {
		return fmt.Errorf("Failed to set tag %v: %v", "DateTime", err)
	}

	// Set DateTimeOriginal
	ifdPath = "IFD/Exif"
	if err := setExifTag(rootIb, ifdPath, "DateTimeOriginal", ts); err != nil {
		return fmt.Errorf("Failed to set tag %v: %v", "DateTimeOriginal", err)
	}

	// Update the exif segment.
	if err := sl.SetExif(rootIb); err != nil {
		return fmt.Errorf("Failed to set EXIF to jpeg: %v", err)
	}

	// Write the modified file
	b := new(bytes.Buffer)
	if err := sl.Write(b); err != nil {
		return fmt.Errorf("Failed to create JPEG data: %v", err)
	}

	fmt.Printf("Number of image bytes: %v\n", len(b.Bytes()))

	// Save the file
	if err := ioutil.WriteFile(filepath, b.Bytes(), 0644); err != nil {
		return fmt.Errorf("Failed to write JPEG file: %v", err)
	}

	fmt.Printf("Wrote %v\n", filepath)

	return nil
}
func setExifTag(rootIB *exif.IfdBuilder, ifdPath, tagName, tagValue string) error {
	fmt.Printf("setTag(): ifdPath: %v, tagName: %v, tagValue: %v",
		ifdPath, tagName, tagValue)

	ifdIb, err := exif.GetOrCreateIbFromRootIb(rootIB, ifdPath)
	if err != nil {
		return fmt.Errorf("Failed to get or create IB: %v", err)
	}

	if err := ifdIb.SetStandardWithName(tagName, tagValue); err != nil {
		return fmt.Errorf("failed to set DateTime tag: %v", err)
	}

	return nil
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

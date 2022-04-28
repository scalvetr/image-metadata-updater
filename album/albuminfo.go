package config

import (
	"fmt"
	"io/fs"
	"strings"
)

type AlbumInfo struct {
	Name  string
	Year  string
	Month string
}

func (a AlbumInfo) GetName() string {
	return fmt.Sprintf("%s-%s - %s", a.Year, a.Month, a.Name)
}
func ExtractAlbumInfo(directory fs.FileInfo) AlbumInfo {
	name := directory.Name()
	nameArr := strings.Split(name, "-")
	year := name[:4]
	month := name[7:9]
	if month[0] != '0' && month[0] != '1' {
		month = "01"
	}

	return AlbumInfo{
		Name:  strings.TrimSpace(nameArr[len(nameArr)-1]),
		Year:  year,
		Month: month,
	}

}

package album

import (
	"bytes"
	"log"
	"os"
	c "photo-manager-cli/config"
	"regexp"
	"strconv"
	"text/template"
)

type AlbumInfo struct {
	Name  string
	Year  int
	Month int
}

func (a AlbumInfo) GetName(config c.AlbumInfoConfig) string {
	albumNamePattern := config.GetAlbumNamePattern()

	albumNameTemplate := template.Must(template.New("albumNameTemplate").Parse(albumNamePattern))

	var tpl bytes.Buffer
	if err := albumNameTemplate.Execute(&tpl, a); err != nil {
		log.Fatal(err)
	}
	return tpl.String()
}

func ExtractAlbumInfo(directory os.DirEntry, config c.AlbumInfoConfig) AlbumInfo {
	folderRegexp := config.GetFolderRegexp()
	pattern := regexp.MustCompile(folderRegexp)
	matches := pattern.FindStringSubmatch(directory.Name())
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
	name := result["name"]

	return AlbumInfo{
		Name:  name,
		Year:  year,
		Month: month,
	}
}

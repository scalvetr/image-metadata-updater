# photo-manager-cli


## Build

```shell
go build

````

## Configure

`photo-manager-cli` requires a `config.yaml` to work.
The file consists on a list of actions to be performed.

### `config.yaml`

Create a configuration file:
```yaml
- action: "UPDATE_METADATA"
  path: "/Users/scalvetr/Pictures/upload/2002 - 12 Desembre - Canaries"
  regexp: "(.*)\\.JPG"
  update_metadata_config:
    override: false
    date: "2005-05-17T14:00:00+01:00"
    date_replaces:
      - day: "2005-04-03"
        new_day: "2006-05-16"
      - day: "2005-04-04"
        new_day: "2006-05-17"
- action: "UPDATE_DATE_FROM_METADATA"
  path: "/Users/scalvetr/Pictures/upload"
- action: "UPLOAD_ALBUMS"
  path: "/Users/scalvetr/Pictures/upload"
- action: "CHECK_ALBUM_DATE_MISMATCH"
  path: "/Users/scalvetr/Pictures/upload"
  report_file: "check_album_date_report.txt"
  album_info:
    folder_regexp: "(?P<year>\d{4}) - (?P<month>\d{2})(.*) - (?P<name>.*)"
    album_name_pattern: {{printf "%04d" .Year}}-{{printf "%02d" .Month}} - {{.Name}}
- action: "INCREASE_DATE"
  path: "/Users/scalvetr/Pictures/upload"
  increase_date_config:
    date_range_from: "2015-01-01"
    date_range_to: "2017-01-01"
    increase_seconds: 215568000
  

```

Accepts 3 kind of actions:
* UPDATE_METADATA: sets a specific date to the `jpeg` `exif` metadata. Requires a `update_metadata_config`.
* UPDATE_DATE_FROM_METADATA: read the `exif` metadata and updates the file system date accordingly (creation and updated dates)
* UPLOAD_ALBUMS: requires `google_client.json` configuration file. You can get more information [here](https://developers.google.com/photos/library/guides/get-started#configure-app)

## Run
```shell
./photo-manager-cli

````

# Videos

Videos are not supported for now, so we suggest using an alternate method. For instance `exiftool`.

```shell
brew install exiftool
cd ~/Pictures/fix/2006\ -\ 05\ Maig\ -\ Final\ Champions
# check the dates
exiftool -a -G1 MOV00342.MPG
# update
exiftool -AllDates="2005:05:16 14:00:00" MOV00342.MPG
```
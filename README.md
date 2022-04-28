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
- action: "UPDATE_METADATA_DATE"
  path: "/Users/scalvetr/Pictures/upload/2002 - 12 Desembre - Canaries"
  date: "2002-12-30T14:00:00+02:00"
- action: "UPDATE_DATE_FROM_METADATA"
  path: "/Users/scalvetr/Pictures/upload"
- action: "UPLOAD_ALBUMS"
  path: "/Users/scalvetr/Pictures/upload"
```

Accepts 3 kind of actions:
* UPDATE_METADATA_DATE: sets a specific date to the `jpeg` `exif` metadata. Requires a `date`.
* UPDATE_DATE_FROM_METADATA: read the `exif` metadata and updates the file system date accordingly (creation and updated dates)
* UPLOAD_ALBUMS: requires `google_client.json` configuration file. You can get more information [here](https://developers.google.com/photos/library/guides/get-started#configure-app)

## Run
```shell
./photo-manager-cli

````

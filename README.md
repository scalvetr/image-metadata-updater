# photo-manager-cli


## Build

```shell
go build

````

## Run

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
Run 
```shell
./photo-manager-cli

````

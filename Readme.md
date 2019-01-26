# S3 Bucket Read API

A small configurable service written in Go consisting of two endpoints:

`GET /files` lists all files in the bucket specified in the `` environment variable.  

`GET /files/:filename` streams bytes of the specified file to the browser.

## Configuration example

Here's an example of the required environment vars:
```shell
"PORT":                  "8080",
"AWS_ACCESS_KEY_ID":     "XXXXXXXXXXXX",
"AWS_SECRET_ACCESS_KEY": "xxxxxxxxxxxxxxxx",
"BUCKET_NAME":           "my-bucket-name"
```
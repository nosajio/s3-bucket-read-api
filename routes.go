package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-chi/chi"
	"net/http"
)

type listResult struct {
	Filename string `json:"filename"`
	URL      string `json:"url"`
}

// RegisterRoutes will associate api routes with handlers
func RegisterRoutes(router *chi.Mux, awsSession *s3.S3) {
	bucketName := EnvDefault("BUCKET_NAME", "").(string)
	if bucketName == "" {
		Error.Println("The \"BUCKET_NAME\" environment variable is required")
		return
	}

	router.Get("/files", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Add("Content-Type", "application/json")
		fileList := ListFilesForBucket(awsSession, bucketName, []string{".png"})
		files := fileList.Contents

		var jsonResults []listResult

		for _, object := range files {
			filename := *object.Key
			filePath := RelativeURLPath(fmt.Sprintf("files/%s", filename))
			fileURL := fmt.Sprintf("//%s", filePath)
			jsonResults =
				append(jsonResults, listResult{Filename: filename, URL: fileURL})
		}

		resJSON, err := json.Marshal(jsonResults)
		if err != nil {
			res.Write([]byte(fmt.Sprint("{\"error\": true")))
			return
		}
		res.Write(resJSON)
	})

	router.Get("/files/{filename}", func(res http.ResponseWriter, req *http.Request) {
		filename := chi.URLParam(req, "filename")
		contentType, objectBytes := GetObjectBytes(awsSession, bucketName, filename)
		res.Header().Add("Content-Type", contentType)
		res.Write(objectBytes)
	})

}

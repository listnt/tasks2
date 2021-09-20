package mymodule

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
)

func WGet(url string, outputName string) error {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()

	outputFile, err := getOutputFile(outputName, resp.Header.Get("Content-Type"))
	if err != nil {
		return err
	}
	defer outputFile.Close()
	_, err = io.Copy(outputFile, resp.Body)
	return err
}

func getOutputFile(fileName string, contentType string) (outputFile *os.File, err error) {
	if fileName == "" {
		fileName = "res"
		var arr []string
		arr, err = mime.ExtensionsByType(contentType)
		if err != nil {
			return
		}
		if len(arr) > 0 {
			fileName += arr[0]
		}
	}

	outputFile, err = os.Create(fileName)
	return
}

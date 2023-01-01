package utils

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

type BasicCredentials struct {
	Username string
	Password string
}

type RequestModel struct {
	URL            string            // Mandatory: URL to send request to
	Method         string            // Mandatory: http method
	Body           []byte            // Optional: body of request
	Headers        map[string]string // Optional: headers
	Params         map[string]string // Optional: params
	BasicAuth      *BasicCredentials // Optional: basic auth1
	Context        context.Context   // Optional: context
	TimeoutSeconds int               // Optional: timeout in seconds
	Files          []string          // Optional: list of paths to files
}

func HTTPRequest(reqModel *RequestModel) (string, int, error) {
	// recover from panic
	defer func() {
		if r := recover(); r != nil {
			time.Sleep(5 * time.Second)
			fmt.Println("Recovered in SimpleHTTPRequestWithModel", r)
		}
	}()

	if reqModel.URL == "" {
		return "", -1, fmt.Errorf("url is empty")
	}

	if reqModel.Method == "" {
		return "", -1, fmt.Errorf("method is empty")
	}

	httpClient := &http.Client{}
	if reqModel.TimeoutSeconds > 0 {
		httpClient.Timeout = time.Duration(reqModel.TimeoutSeconds) * time.Second
	}
	var err error
	var req *http.Request

	switch {
	case reqModel.Body != nil && len(reqModel.Files) == 0:
		// Normal request with only json body
		req, err = http.NewRequest(reqModel.Method, reqModel.URL, bytes.NewBuffer(reqModel.Body))
		req.Header.Set("Content-Type", "application/json")

	case len(reqModel.Body) == 0 && len(reqModel.Files) > 0:
		// Multipart request with files
		// Create a new multipart request
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		// Create the file fields
		for i, file := range reqModel.Files {
			// Open the file
			f, err := os.Open(file)
			if err != nil {
				return "", -1, err
			}
			defer f.Close()

			formFile, err := writer.CreateFormFile(fmt.Sprintf("file%d", i), file)
			if err != nil {
				return "", -1, err
			}

			io.Copy(formFile, f)
		}

		// Close the writer
		writer.Close()

		// Create the request
		req, err = http.NewRequest(reqModel.Method, reqModel.URL, body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

	default:
		// Normal request without body or files
		req, err = http.NewRequest(reqModel.Method, reqModel.URL, nil)
	}

	if err != nil {
		return "", -1, err
	}

	// add basic auth
	if reqModel.BasicAuth != nil {
		req.SetBasicAuth(reqModel.BasicAuth.Username, reqModel.BasicAuth.Password)
	}
	if reqModel.Context != nil {
		req = req.WithContext(reqModel.Context)
	}
	// add params
	for key, value := range reqModel.Params {
		q := req.URL.Query()
		q.Add(key, value)
		req.URL.RawQuery = q.Encode()
	}

	// add headers
	for key, value := range reqModel.Headers {
		req.Header.Set(key, value)
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", -1, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", -1, err
	}

	return string(data), resp.StatusCode, nil
}

func DownloadGithubZip(url string, path string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

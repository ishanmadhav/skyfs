package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/go-resty/resty/v2"
	"github.com/ishanmadhav/skyfs/client/api"
)

const MASTER_URL string = "http://localhost:8080"

type Client struct {
	cli *resty.Client
}

// Creates a new object storage client
// Auth will also be handled by this
func NewClient() *Client {
	httpCli := resty.New()
	return &Client{
		cli: httpCli,
	}
}

// Creates a new distributed file system bucket
func (c *Client) CreateBucket(bucketName string) {

	url := MASTER_URL + "/api/bucket"
	// Define the request parameters
	params := map[string]string{
		"bucketName": bucketName,
	}

	// Send the POST request
	resp, err := c.cli.R().
		SetQueryParams(params). // Set the response type
		Post(url)

	if err != nil {
		log.Fatalf("Error while sending request: %v", err)
	}

	// Parse the response
	fmt.Println(resp.String())
}

// Puts/Creates a new object
// Akin to adding a file to the file system
func (c *Client) PutObject(objectName string, bucketName string, filePath string) error {
	url := MASTER_URL + "/api/object"
	fileSize := getFileSize(filePath)
	newObject := api.Object{
		Bucket:   bucketName,
		FileName: objectName,
		FileSize: fileSize,
	}
	fmt.Print(newObject)
	req, err := json.Marshal(newObject)
	if err != nil {
		return err
	}

	resp, err := c.cli.R().SetHeader("Content-Type", "application/json").SetBody(req).Post(url)
	if err != nil {
		return err
	}
	var objectResponse api.ObjectResponseDTO
	err = json.Unmarshal(resp.Body(), &objectResponse)
	if err != nil {
		return err
	}
	fmt.Print(objectResponse)
	err = c.uploadFile(objectName, filePath, bucketName, objectResponse.ChunkGraph)
	if err != nil {
		return err
	}
	return nil
}

// Gets object from FS.
// Akin to downloading a file
func (c *Client) GetObject(objectName string, bucketName string, downloadPath string) error {
	url := MASTER_URL + "/api/object" + "?bucket=" + bucketName + "&fileName=" + objectName
	newObject := api.Object{
		Bucket:   bucketName,
		FileName: objectName,
		FileSize: 0,
	}
	fmt.Print(newObject)
	_, err := json.Marshal(newObject)
	if err != nil {
		return err
	}

	resp, err := c.cli.R().SetHeader("Content-Type", "application/json").Get(url)
	if err != nil {
		return err
	}
	var objectResponse api.ObjectResponseDTO
	err = json.Unmarshal(resp.Body(), &objectResponse)
	if err != nil {
		return err
	}
	//fmt.Print(objectResponse)
	c.downloadFile("obj5", "downloads", "Buck", downloadPath, objectResponse.ChunkGraph)
	return nil
}

func (c *Client) downloadFile(objectName string, downloadPath string, bucketName string, outPutFilePath string, chunkGraph api.ChunkGraph) error {
	var downloadPaths []string
	for i, item := range chunkGraph.Graph {
		chunkName := fmt.Sprintf("%s_%s_%d", bucketName, objectName, i)
		chunkDownloadPath := downloadPath + "/" + chunkName
		downloadUrl := "http://" + item[0].Address + "/download?fileName=" + chunkName
		err := c.downloadChunk(chunkName, chunkDownloadPath, bucketName, downloadUrl)
		if err != nil {
			return err
		}
		downloadPaths = append(downloadPaths, chunkDownloadPath)
	}
	objectDownloadPath := "downloads/" + outPutFilePath
	c.recreateFileFromChunks(downloadPaths, objectDownloadPath)
	return nil
}

// Function that does the downloading from different chunk servers
func (c *Client) downloadChunk(objectName string, downloadPath string, bucketName string, url string) error {
	//url := "http://localhost:3000/download?filename=obj5_0"
	resp, err := c.cli.R().
		SetDoNotParseResponse(true). // Important for downloading files
		Get(url)
	if err != nil {
		return err
	}
	defer resp.RawBody().Close()
	// Ensure the download directory exists
	err = os.MkdirAll(filepath.Dir(downloadPath), os.ModePerm)
	if err != nil {
		return err
	}

	// Create the file
	outFile, err := os.Create(downloadPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Write the response body to file
	_, err = io.Copy(outFile, resp.RawBody())
	if err != nil {
		return err
	}

	fmt.Println("File downloaded successfully to:", downloadPath)
	return nil
}

// Function that does the actual distributed upload to different chunk servers
func (c *Client) uploadFile(objectName string, filePath string, bucketName string, chunkGraph api.ChunkGraph) error {
	fmt.Println("Printing chunk graph")
	fmt.Println(chunkGraph)
	fmt.Println("Printed chunk graph")
	chunkSize := int64(1024 * 1024) // 1 MB
	chunkPaths, err := c.breakFileIntoChunks(filePath, chunkSize, objectName, bucketName)
	if err != nil {
		fmt.Println("Error breaking into chunks")
		return err
	}
	for i, chunkPath := range chunkPaths {
		fmt.Print("Chunk" + chunkPath)
		addr := chunkGraph.Graph[i][0].Address
		err = c.uploadChunk(chunkPath, objectName, bucketName, addr)
		if err != nil {
			fmt.Print(err)
			return err
		}
	}

	return nil
}

// Uploads individiual chunk to its mapped chunk server
func (c *Client) uploadChunk(chunkPath string, objectName string, bucketName string, address string) error {

	filePath := chunkPath
	fmt.Print("File Path" + filePath)
	// Send a POST request to upload the file with additional data
	resp, err := c.cli.R().
		SetFile("file", filePath).
		SetFormData(map[string]string{
			"objectName": objectName,
			"chunkName":  chunkPath,
			"bucketName": bucketName,
		}).
		Post("http://" + address + "/upload")
	if err != nil {
		return err
	}
	fmt.Print(resp.String())
	return nil
}

func (cli *Client) GetBuckets() {

}

func getFileSize(filePath string) float64 {
	// Open the file to get its size
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	// Get the file info to retrieve the size
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatalf("Error getting file info: %v", err)
	}

	// Get the file size
	fileSize := fileInfo.Size()
	fileSizeMB := float64(fileSize) / (1024 * 1024)
	return fileSizeMB

}

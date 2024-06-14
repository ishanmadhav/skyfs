package client

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func (c *Client) breakFileIntoChunks(filePath string, chunkSize int64, objectName string, bucketName string) ([]string, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Get the file info
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to get file info: %w", err)
	}

	// Create a directory to store the chunks
	chunkDir := "chunks"
	err = os.MkdirAll(chunkDir, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to create chunks directory: %w", err)
	}

	var chunkPaths []string
	buffer := make([]byte, chunkSize)
	for i := int64(0); i < fileInfo.Size(); i += chunkSize {
		bytesRead, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return nil, fmt.Errorf("failed to read file: %w", err)
		}
		if bytesRead == 0 {
			break
		}

		chunkPath := filepath.Join(chunkDir, fmt.Sprintf("%s_%s_%d", bucketName, objectName, i/chunkSize))
		chunkFile, err := os.Create(chunkPath)
		if err != nil {
			return nil, fmt.Errorf("failed to create chunk file: %w", err)
		}

		_, err = chunkFile.Write(buffer[:bytesRead])
		if err != nil {
			return nil, fmt.Errorf("failed to write to chunk file: %w", err)
		}
		chunkFile.Close()
		chunkPaths = append(chunkPaths, chunkPath)
	}
	return chunkPaths, nil
}

func (c *Client) recreateFileFromChunks(chunkPaths []string, outputPath string) error {
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outputFile.Close()

	for _, chunkPath := range chunkPaths {
		chunkFile, err := os.Open(chunkPath)
		if err != nil {
			return fmt.Errorf("failed to open chunk file: %w", err)
		}

		_, err = io.Copy(outputFile, chunkFile)
		if err != nil {
			chunkFile.Close()
			return fmt.Errorf("failed to write chunk to output file: %w", err)
		}
		chunkFile.Close()
	}
	return nil
}

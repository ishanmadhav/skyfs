package chunkserver

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

const MASTER_URL string = "http://localhost:8080"

type ChunkServer struct {
	serverName string
	serverID   string
	address    string
	cli        *resty.Client
	app        *fiber.App
}

func NewChunkServer(serverName string, address string) ChunkServer {
	httpCli := resty.New()
	app := fiber.New()
	return ChunkServer{
		serverName: serverName,
		serverID:   "",
		address:    address,
		cli:        httpCli,
		app:        app,
	}
}

func (cs *ChunkServer) Start() {
	msg := fmt.Sprintf("Chunk server runnign on address %s", cs.address)
	fmt.Println(msg)
	err := cs.connectToMaster()
	if err != nil {
		panic(err)
	}
	cs.setupRoutes()
	cs.app.Listen(cs.address)
}

func (cs *ChunkServer) setupRoutes() {
	cs.app.Post("/upload", func(c *fiber.Ctx) error {
		// Parse the multipart form
		if form, err := c.MultipartForm(); err == nil {
			// Get additional form fields
			objectName := c.FormValue("objectName")
			chunkName := c.FormValue("chunkName")
			bucketName := c.FormValue("bucketName")

			fmt.Printf("Received - objectName: %s, chunkName: %s, bucketName: %s\n", objectName, chunkName, bucketName)

			// Get all files from "file" key
			files := form.File["file"]

			for _, file := range files {
				// Ensure the directory to save the file exists, create it if not
				if err := os.MkdirAll("./uploads", os.ModePerm); err != nil {
					fmt.Println(err)
					return err
				}

				// Save each file with a unique name in the "uploads" directory
				err := c.SaveFile(file, fmt.Sprintf("./uploads/%s", file.Filename))
				if err != nil {
					fmt.Println(err)
					return err
				}
			}
		} else {
			return err
		}

		// Send a response
		return c.SendString("File(s) uploaded successfully with additional data!")
	})

	cs.app.Get("/download", func(c *fiber.Ctx) error {
		// Get the filename from the query parameters
		filename := c.Query("filename")
		if filename == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Filename query parameter is required")
		}

		// Construct the full file path
		filepath := filepath.Join("./uploads", filename)

		// Check if the file exists
		if _, err := os.Stat(filepath); os.IsNotExist(err) {
			return c.Status(fiber.StatusNotFound).SendString("File not found")
		}

		// Send the file to the client
		return c.SendFile(filepath)
	})

}

func (cs *ChunkServer) connectToMaster() error {
	req := ChunkServerRequestDTO{
		Address: cs.address,
	}
	url := MASTER_URL + "/api/chunkserver"
	reqBody, err := json.Marshal(req)
	if err != nil {
		fmt.Println("Error marshalling JSON", err)
		return err
	}
	resp, err := cs.cli.R().SetHeader("Content-Type", "application/json").SetBody(reqBody).Post(url)
	if err != nil {
		fmt.Println("Error sending post request to master", err)
		return err
	}
	fmt.Println(resp.StatusCode())
	var newChunkServer ChunkServerResponseDTO
	err = json.Unmarshal(resp.Body(), &newChunkServer)
	if err != nil {
		return err
	}
	fmt.Println(newChunkServer)
	return nil
}

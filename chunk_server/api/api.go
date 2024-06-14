package api

type Bucket struct {
	Name string `json:"name"`
}

type Object struct {
	Bucket   string  `json:"bucket"`
	FileName string  `json:"fileName"`
	FileSize float64 `json:"fileSize"`
}

// ChunkServer
type ChunkServer struct {
	Id      string `json:"id"`
	Address string `json:"address"`
}

// ChunkGraph represents the graph structure
type ChunkGraph struct {
	Graph [][]ChunkServer `json:"graph"`
}

// ObjectResponseDTO represents the structure of the JSON response
type ObjectResponseDTO struct {
	Bucket      string     `json:"bucket"`
	FileName    string     `json:"fileName"`
	FileSize    float64    `json:"fileSize"`
	NumOfChunks int        `json:"numOfChunks"`
	ChunkGraph  ChunkGraph `json:"chunkGraph"`
}

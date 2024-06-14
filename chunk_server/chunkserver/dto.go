package chunkserver

type ChunkServerRequestDTO struct {
	Address string `json:"address"`
}

type ChunkServerResponseDTO struct {
	Id      string `json:"chunkserverid"`
	Status  bool   `json:"status"`
	Address string `json:"address"`
}

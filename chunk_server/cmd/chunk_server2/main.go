package main

import "github.com/ishanmadhav/skyfs/chunkserver"

func main() {
	cs := chunkserver.NewChunkServer("DemoServer2", "localhost:3001")
	cs.Start()
}

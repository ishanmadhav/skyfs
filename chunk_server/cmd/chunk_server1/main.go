package main

import "github.com/ishanmadhav/skyfs/chunkserver"

func main() {
	cs := chunkserver.NewChunkServer("DemoServer", "localhost:3000")
	cs.Start()
}

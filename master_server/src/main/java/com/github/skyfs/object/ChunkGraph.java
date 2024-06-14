package com.github.skyfs.object;

import com.github.skyfs.chunkservice.ChunkServer;

import java.util.ArrayList;
import java.util.List;

public class ChunkGraph {
    public List <List<ChunkServer>> graph;

    public ChunkGraph() {
        this.graph = new ArrayList<>();
    }
}

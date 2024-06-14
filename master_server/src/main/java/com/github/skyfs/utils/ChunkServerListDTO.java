package com.github.skyfs.utils;

import com.github.skyfs.chunkservice.ChunkServer;

import java.util.List;

public class ChunkServerListDTO {
    public List<ChunkServer> chunkServerList;

    public ChunkServerListDTO(List<ChunkServer> chunkServerList) {
        this.chunkServerList=chunkServerList;
    }
}

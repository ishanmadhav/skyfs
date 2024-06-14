package com.github.skyfs.utils;

import com.github.skyfs.object.ChunkGraph;

public class ObjectResponseDTO {
    public String bucket;
    public String fileName;
    public double fileSize;
    public int numOfChunks;
    public ChunkGraph chunkGraph;

    public ObjectResponseDTO(String bucket, String fileName, double fileSize, int numOfChunks, ChunkGraph chunkGraph) {
        this.bucket=bucket;
        this.fileName=fileName;
        this.fileSize=fileSize;
        this.numOfChunks=numOfChunks;
        this.chunkGraph=chunkGraph;
    }
}

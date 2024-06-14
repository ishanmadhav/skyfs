package com.github.skyfs.object;

import java.util.UUID;

//Validation for object creation will happen at top level
public class Object {
    public String name;
    public String bucketName;
    public double fileSize;
    private String id;
    private ChunkGraph chunkGraph;

    public Object(String objectName, String bucketName, ChunkGraph chunkGraph, double fileSize) {
        this.name=objectName;
        this.bucketName=bucketName;
        UUID uuid=UUID.randomUUID();
        this.id=uuid.toString();
        this.chunkGraph=chunkGraph;
    }

    public ChunkGraph getChunkGraph() {
        return this.chunkGraph;
    }
    public String getId() {
        return id;
    }
    public String getObjectName() {
        return name;
    }

}

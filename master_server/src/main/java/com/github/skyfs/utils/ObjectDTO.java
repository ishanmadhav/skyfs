package com.github.skyfs.utils;

public class ObjectDTO {
    public String bucket;
    public String fileName;
    public double fileSize;

    public ObjectDTO(String bucket, String fileName, double fileSize) {
        this.bucket=bucket;
        this.fileName=fileName;
        this.fileSize=fileSize;
    }
}

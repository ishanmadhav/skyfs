package com.github.skyfs.utils;

public class BucketResponseDTO {
    public String bucketName;
    public boolean status;

    public BucketResponseDTO(String bucketName, boolean status) {
        this.bucketName=bucketName;
        this.status=status;
    }
}

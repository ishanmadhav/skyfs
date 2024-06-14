package com.github.skyfs.bucket;

import org.springframework.stereotype.Component;

import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

@Component
public class BucketService {
    private List<Bucket> buckets;

    public BucketService() {
        buckets=new ArrayList<Bucket>();
    }

    public List<Bucket> getBuckets() {
        return buckets;
    }

    public Bucket getBucket(String bucketName) {
        if (buckets.isEmpty()) {
            // throw an error
        }

        Bucket bucketToBeReturned=new Bucket("random");
        for (Bucket bucket:buckets) {
            if (Objects.equals(bucket.name, bucketName)) {
                bucketToBeReturned=bucket;
                break;
            }
        }
        return bucketToBeReturned;
    }

    public Bucket addBucket(String bucketName) {
        System.out.println(bucketName+"from add bucket");
        Bucket newBucket=new Bucket(bucketName);
        buckets.add(newBucket);
        return newBucket;
    }

    public boolean bucketExists(String bucketName) {
        System.out.println(buckets);
        if (buckets.isEmpty()) {
            return false;
        }
        for (Bucket bucket : buckets) {
            if (Objects.equals(bucket.name, bucketName)) {
                return true;
            }
        }
        return false;
    }
}

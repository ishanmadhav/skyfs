package com.github.skyfs.bucket;

import com.github.skyfs.utils.BucketResponseDTO;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/api/bucket")
public class BucketController {
    @Autowired
    private BucketService bucketService;

    @PostMapping
    public BucketResponseDTO createBucket(@RequestParam(value="bucketName", defaultValue = "default") String bucketName) {
        boolean exists=bucketService.bucketExists(bucketName);
        if (exists) {
            return new BucketResponseDTO(bucketName, false);
        }

        bucketService.addBucket(bucketName);
        return new BucketResponseDTO(bucketName, true);
    }

    @GetMapping
    public void getBuckets() {
        System.out.println(bucketService.getBuckets());
    }
}

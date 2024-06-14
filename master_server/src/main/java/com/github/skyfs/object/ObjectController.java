package com.github.skyfs.object;

import com.github.skyfs.bucket.Bucket;
import com.github.skyfs.bucket.BucketService;
import com.github.skyfs.chunkservice.ChunkServer;
import com.github.skyfs.chunkservice.ChunkService;
import com.github.skyfs.utils.ObjectDTO;
import com.github.skyfs.utils.ObjectResponseDTO;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.ArrayList;
import java.util.List;
import java.util.Random;

@RestController
@RequestMapping("/api/object")
public class ObjectController {
    @Autowired
    private BucketService bucketService;
    @Autowired
    private ChunkService chunkService;

    @PostMapping
    public ObjectResponseDTO putObject(@RequestBody ObjectDTO body) {
        System.out.println(body.bucket);
        System.out.println(body.fileSize);
        System.out.println(body.fileName);
        //Check if this bucket already exists
        boolean bucketExists=bucketService.bucketExists(body.bucket);
        if (!bucketExists) {
            // return and throw error here
        }
        // Check if object with same name already exists in given bucket
        Bucket bucket=bucketService.getBucket(body.bucket);
        boolean objectExists=bucket.objectExists(body.fileName);
        if (objectExists) {
            // return and throw error here
        }
        // Generate chunk graph
        double size=body.fileSize;
        int numOfChunks=((int)size/1)+(size%1==0?0:1);
        ChunkGraph chunkGraph=mapToGraph(numOfChunks);
        // Create object and persist it to memory
        bucket.putObject(body.fileName, size, chunkGraph);
        //return the object to the client, another controller will confirm the status of this object's
        // persistence to the file systme once all the operations complte. it will be done via a differnet api call
        return new ObjectResponseDTO(body.bucket,  body.fileName, body.fileSize, numOfChunks, chunkGraph);
    }


    @GetMapping
    public ObjectResponseDTO getObject(@RequestParam String bucket, @RequestParam String fileName) {
        System.out.println("Bucket Name "+bucket);
        System.out.println("Object Name "+fileName);

        // Check if this bucket already exists
        boolean bucketExists = bucketService.bucketExists(bucket);
        if (!bucketExists) {
            System.out.println("Bucket does not exist");
            // throw an appropriate exception or return an error response
        }

        // Check if object with the same name already exists in the given bucket
        Bucket bucketObj = bucketService.getBucket(bucket);
        boolean objectExists = bucketObj.objectExists(fileName);
        List <Object> objectList=bucketObj.getObjects();
        System.out.println("Printing object list in bucket");
        for (Object obj:objectList) {
            System.out.println(obj.name);
        }
        if (!objectExists) {
            System.out.println("Object does not exist");
            // throw an appropriate exception or return an error response
        }

        Object objectToReturned = bucketObj.getObject(fileName, bucket);
        ChunkGraph chunkGraph = objectToReturned.getChunkGraph();
        return new ObjectResponseDTO(objectToReturned.bucketName, objectToReturned.name, objectToReturned.fileSize, chunkGraph.graph.size(), chunkGraph);
    }

    // Map file chunks for an object to different servers
    public ChunkGraph mapToGraph(int numOfChunks) {
        List<ChunkServer> chunkServerListFromService=chunkService.getChunkServerList();
        System.out.println(chunkServerListFromService);
        Random random=new Random();
        ChunkGraph chunkGraph=new ChunkGraph();
        for (int i=0;i<numOfChunks;i++) {
            int randomIndex=random.nextInt(chunkServerListFromService.size());
            List <ChunkServer> tempList=new ArrayList<ChunkServer>();
            tempList.add(chunkServerListFromService.get(randomIndex));
            chunkGraph.graph.add(tempList);
        }
        System.out.println(chunkGraph);
        return chunkGraph;
    }
}

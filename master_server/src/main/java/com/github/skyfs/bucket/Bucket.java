package com.github.skyfs.bucket;

import com.github.skyfs.object.ChunkGraph;
import com.github.skyfs.object.Object;

import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

//Validation for bucket creation will happen at top level
public class Bucket {
    public String name;
    private List<Object> objects;

    public Bucket(String bucketName) {
        this.name=bucketName;
        objects=new ArrayList<Object>();
    }

    // Checks if object with same name already exists in bucket
    public boolean objectExists(String objectName) {
        for (Object tempObject : objects) {
            if (tempObject.name==objectName) {
                return true;
            }
        }
        return false;
    }

    public void putObject(String objectName, double fileSize, ChunkGraph chunkGraph) {
        boolean exits=objectExists(objectName);
        if (exits) {
            return;
        }
        Object newObject=new Object(objectName, this.name, chunkGraph, fileSize);
        objects.add(newObject);
    }

    public Object getObject(String objectName, String bucketName) {
        Object tempObject = null;
        for (Object object:objects) {
            if (Objects.equals(object.name, objectName)) {
                return object;
            }
        }
        return tempObject;
    }

    public List<Object> getObjects() {
        return this.objects;
    }

}

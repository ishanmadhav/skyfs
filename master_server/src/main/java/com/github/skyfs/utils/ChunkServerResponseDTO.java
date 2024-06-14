package com.github.skyfs.utils;

public class ChunkServerResponseDTO {
    public String chunkServerID;
    public boolean status;
    public String address;

    public ChunkServerResponseDTO(String chunkServerID, boolean status, String address) {
        this.chunkServerID=chunkServerID;
        this.status=status;
        this.address=address;
    }

}

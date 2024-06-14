package com.github.skyfs.chunkservice;

import org.springframework.stereotype.Component;

import java.util.ArrayList;
import java.util.List;
import java.util.Objects;
import java.util.UUID;

@Component
public class ChunkService {
    private List<ChunkServer> servers;

    public ChunkService() {
        this.servers=new ArrayList<>();
    }

    public List<ChunkServer> getChunkServerList() {
        return servers;
    }

    public ChunkServer addChunkServer(String address) {
        UUID uuid=UUID.randomUUID();
        String uuidStr=uuid.toString();
        ChunkServer newChunkServer=new ChunkServer(uuidStr, address);
        servers.add(newChunkServer);
        return newChunkServer;
    }

    // Check if chunk server with this address is already added
    public boolean chunkServerExists(String address) {
        for (ChunkServer server:servers) {
            if (Objects.equals(server.address, address)) {
                return true;
            }
        }
        return false;
    }
}

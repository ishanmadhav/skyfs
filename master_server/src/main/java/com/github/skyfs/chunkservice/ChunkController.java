package com.github.skyfs.chunkservice;

import com.github.skyfs.utils.ChunkServerListDTO;
import com.github.skyfs.utils.ChunkServerRequestDTO;
import com.github.skyfs.utils.ChunkServerResponseDTO;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("api/chunkserver")
public class ChunkController {
    @Autowired
    private ChunkService chunkService;

    @PostMapping
    public ChunkServerResponseDTO addChunkServer(@RequestBody ChunkServerRequestDTO body) {
        boolean exists=chunkService.chunkServerExists(body.address);
        if (exists) {
            return new ChunkServerResponseDTO("", false, "");
        }
        ChunkServer chunkServer=chunkService.addChunkServer(body.address);
        return new ChunkServerResponseDTO(chunkServer.id, true, chunkServer.address);
    }

    @GetMapping
    public ChunkServerListDTO getChunkServerList() {
        ChunkServerListDTO resp=new ChunkServerListDTO(chunkService.getChunkServerList());
        return resp;
    }
}

package com.github.skyfs.wal;

import org.springframework.stereotype.Component;

import java.io.File;
import java.util.List;

@Component
public class WALService {
    private List<LogEntry> log;
    private File logFile;
}

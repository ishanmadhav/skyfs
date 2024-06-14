package com.github.skyfs.wal;

public class LogEntry {
    public String operation;
    public String data;
    public int seqNum;

    public LogEntry(String operation, String data, int seqNum) {
        this.operation=operation;
        this.data=data;
        this.seqNum=seqNum;
    }
}

# 🧰 Note Taker CLI Installation Guide

This document explains how to download, install, and run **Note Taker CLI** for your operating system.

---

## 🐧-linux-amd64
```bash 
curl -L https://github.com/thissidemayur/3-golang-CLI_Note_Taker/releases/download/v1.0.0/note-taker-linux-amd64.tar.gz \
  -o /tmp/note-taker-linux-amd64.tar.gz && \
sudo tar -xzf /tmp/note-taker-linux-amd64.tar.gz -C /usr/local/bin && \
sudo mv /usr/local/bin/note-taker-linux-amd64 /usr/local/bin/note-taker && \
note-taker
```



## 🍎-macosamd64
```bash
curl -L  https://github.com/thissidemayur/3-golang-CLI_Note_Taker/releases/download/v1.0.0/note-taker-darwin-amd64.tar.gz  \
    -o /tmp/note-taker-darwin-amd64.tar.gz && \
sudo tar -xzf /tmp/note-taker-darwin-amd64.tar.gz -C /usr/local/bin && \
sudo mv /usr/local/bin/note-taker-darwin-amd64 /usr/local/bin/note-taker && \
note-taker
```

## 🪟-windows
```powershell
Invoke-WebRequest -Uri "https://github.com/thissidemayur/3-golang-CLI_Note_Taker/releases/download/v1.0.0/note-taker-windows-amd64.tar.gz" -OutFile "$env:TEMP\note-taker.tar.gz"
tar -xzf "$env:TEMP\note-taker.tar.gz" -C "$env:TEMP"
Move-Item "$env:TEMP\note-taker-windows-amd64.exe" "C:\Program Files\note-taker\note-taker.exe"
setx PATH "$env:PATH;C:\Program Files\note-taker"
note-taker
```

--

## ✅ Verify Installation
```bash
note-taker --version
```
If you see version output `(e.g. Note Taker v1.0.0)`, installation succeeded.


# 📝 Note Taker CLI — Built with Go + Bubble Tea

A modern, responsive **terminal-based Note Taking CLI** built using [Bubble Tea](https://github.com/charmbracelet/bubbletea) and [Lipgloss](https://github.com/charmbracelet/lipgloss).  
Designed for developers who love **speed, minimalism, and good UX** — no mouse required.

---

## 🚀 Features

- ⚡ **Interactive TUI** using Bubble Tea  
- 💾 Create, edit, and delete notes with instant feedback  
- 📂 Browse and filter notes in real time  
- 🎨 Styled with Lipgloss — vivid colors and message boxes (Error, Info, Success, Warning)  
- 📱 **Responsive layout** — adjusts automatically when you resize the terminal  
- 🧰 Built with clean, layered Go architecture (`cmd/`, `internal/`, `ui/`, etc.)  
- 🐳 Ready-to-deploy via Docker & Makefile  

---

## 🧩 Directory Structure
```
note-taker/
├── bin/
│   └── app
├── cmd/
│   └── note-taker/
│       └── main.go              # just setup & start the app
├── internal/
│   ├── ui/                      # Bubble Tea UI (model, update, view)
│   │   ├── model.go
│   │   ├── update.go
│   │   ├── view.go
│   │   └── styles.go
│   ├── notes/                   # core logic (file handling)
│   │   ├── manager.go
│   │   └── list.go
│   └── config/
│       └── paths.go    # vaultDir setup
├── Dockerfile
├── Makefile          
```

## ⚙️ Installation
### 1️⃣ Clone the repository
```bash
git clone https://github.com/yourusername/note-taker.git
cd note-taker
```
### 2️⃣ Build the binary
 1. Using Makefile 
 `make build `

 2. or manually: 
  `go build -o bin/app ./cmd/note-taker`

## 🐳 Run with Docker
### Build the container:
 `docker build -t note-taker . `
### Run the app interactively:
`docker run -it note-taker`

## 💻 Usage
Start the app:
`./bin/app`

## 🧠 Architecture Overview
```
| Layer                | Description                                 |
| -------------------- | ------------------------------------------- |
| **cmd/**             | Entry point — starts the Bubble Tea program |
| **internal/ui/**     | TUI layer — model, update, view, styles     |
| **internal/notes/**  | File management & note operations           |
| **internal/config/** | Directory setup & app paths                 |
| **Makefile**         | Build & run automation                      |
| **Dockerfile**       | Containerized deployment                    |

```


# 📝 Note Taker CLI — Built with Go + Bubble Tea

A modern, responsive **terminal-based Note Taking CLI** built using [Bubble Tea](https://github.com/charmbracelet/bubbletea) and [Lipgloss](https://github.com/charmbracelet/lipgloss).  
Designed for developers who love **speed, minimalism, and good UX** — no mouse required.

---

## ⚡ Features

- ⚡ **Interactive TUI** using Bubble Tea
- 💾 Create, edit, and delete notes with instant feedback
- 📂 Browse and filter notes in real time
- 🎨 Styled with Lipgloss — vivid colors and message boxes (Error, Info, Success, Warning)
- 📱 Responsive layout — adjusts automatically when you resize the terminal
- 🧰 Built with clean, layered Go architecture (`cmd/`, `internal/`, `ui/`, etc.)
- 🐳 Ready-to-deploy via Docker or Podman
- 🗂 Supports cross-platform builds: Linux, macOS, Windows
- 🔒 Optional encryption support (planned)
- 🎨 Theme switcher support (planned)

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
## 🐳 Run with Docker / Podman

### Build the container:
```bash
docker build -t note-taker .           # or using Podman: podman build -t note-taker .
```
### Run interactively:
```
docker run -it note-taker              # or podman run -it note-taker
```
--
## 💻 Usage
Start the app:
`./bin/app`

--

## 🌐 GHCR / Container Registry
```
podman pull ghcr.io/thissidemayur/note-taker:v1.0.0
podman run -it --rm ghcr.io/thissidemayur/note-taker:v1.0.0
```
--



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

--

## 🛠 Automation & Releases

We provide a **release script** `release.sh` to automate your workflow:

- Build **cross-platform binaries** (Linux/macOS/Windows)  
- Package binaries as `.tar.gz`  
- Build **Docker/Podman image**  
- Push image to **GHCR/DockerHub** (optional)  
- Create a **GitHub release** with binaries  

Example usage:
```bash
./release.sh v1.0.0
```

--
## 🔮 Future Improvements

We plan to enhance Note Taker CLI with the following features:

- 🔐 **Note Encryption**  
  Ensure privacy by encrypting notes on disk.

- 🎨 **Theme Switcher**  
  Allow users to toggle between dark/light or custom themes.

- 📄 **Configuration File Support**  
  Customize app paths, themes, and other preferences via a config file.

- 🧪 **Unit & Integration Tests**  
  Improve reliability and maintainability of the codebase.

- 📦 **Makefile Automation**  
  Add commands for building, testing, and releasing the project easily.

- 🌍 **Pre-Built Binaries**  
  Publish compiled binaries for Linux, macOS, and Windows on GitHub Releases.

- 📈 **Logging & Analytics**  
  Collect CLI usage stats to improve UX and detect issues.

- 🐛 **Enhanced Error Handling & Validation**  
  Catch edge cases and provide informative messages.

- 🚀 **Automate Multi-Arch Docker/Podman Builds**  
  Build and push images for multiple architectures automatically.

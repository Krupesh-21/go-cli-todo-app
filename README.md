# Go CLI Todo App 🗒️

A cross-platform, colorful command-line todo application built with Go. Simple, fast, and persistent across sessions.

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go&logoColor=white)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Cross-Platform](https://img.shields.io/badge/Platform-Windows%20%7C%20Linux%20%7C%20macOS-brightgreen)](https://golang.org/doc/install)

## ✨ Features

- ✅ **Add todos** with natural language descriptions
- ✅ **List todos** with beautiful color-coded output (completed ✅ / incomplete 🟣)
- ✅ **Complete todos** by ID with visual strikethrough
- ✅ **Remove todos** by ID
- ✅ **Persistent storage** using `todos.json`
- ✅ **Cross-platform builds** (Windows, Linux, macOS)
- ✅ **Colorful terminal output** using `github.com/fatih/color`
- ✅ **Zero dependencies** for end-users (single binary)

## 🎮 Quick Start

```
# Clone and build
git clone <repo-url>
cd go-cli-todo-app
make build

# Or just run directly
make run ARGS="add Buy groceries"
```

## 📋 Commands

| Command    | Example                                | Description               |
| ---------- | -------------------------------------- | ------------------------- |
| `add`      | `./go-cli-todo-app add "Walk the dog"` | Add a new todo            |
| `list`     | `./go-cli-todo-app list`               | Show all todos            |
| `complete` | `./go-cli-todo-app complete 1`         | Mark todo #1 as completed |
| `remove`   | `./go-cli-todo-app remove 2`           | Delete todo #2            |

### Example Usage

```
# Add todos
$ make run ARGS="add Buy groceries"
$ make run ARGS="add Walk the dog"
Added todo: {ID:1 Description:Buy groceries Completed:false}
Added todo: {ID:2 Description:Walk the dog Completed:false}

# List todos
$ make run ARGS="list"
Completed    Incomplete

1: ~~Buy groceries~~
2: Walk the dog

# Complete a todo
$ make run ARGS="complete 1"
Marked todo 1 as completed

# Remove a todo
$ make run ARGS="remove 2"
```

## 🛠️ Build & Install

```
# Build for current platform
make build

# Build for all platforms
make all-platforms

# Run with arguments
make run ARGS="list"

# Clean binaries
make clean

# See all options
make help
```

### Cross-Platform Binaries

```
go-cli-todo-app      # Current OS
go-cli-todo-app.exe  # Windows
go-cli-todo-app-linux # Linux
go-cli-todo-app-macos # macOS (arm64)
```

## 🚀 Installation

1. **Pre-built binary**: Download from [Releases](https://github.com/Krupesh-21/go-cli-todo-app/releases)
2. **From source**:
   ```
   git clone https://github.com/Krupesh-21/go-cli-todo-app.git
   cd go-cli-todo-app
   make build
   sudo mv go-cli-todo-app /usr/local/bin/todo
   ```

## 📁 Project Structure

```
go-cli-todo-app/
├── main.go          # Entry point & command routing
├── utils.go         # Core todo logic
├── todos.json       # Persistent storage (auto-created)
├── Makefile         # Build automation
└── README.md
```

## 🌈 Output Preview

```
Completed    Incomplete

1: ~~Buy groceries~~
2: Walk the dog
3: Code new feature
```

## 🔧 Tech Stack

- **Go 1.21+** - Single binary, cross-compilation
- **github.com/fatih/color** - Beautiful terminal colors
- **JSON persistence** - Simple, human-readable storage

## 🤝 Contributing

1. Fork the repo
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to branch (`git push origin feature/amazing-feature`)
5. Open Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙌 Acknowledgments

Built with ❤️ using Go's amazing standard library and cross-compilation capabilities.

---

**⭐ Star this repo if you found it useful!**

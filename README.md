[![progress-banner](https://backend.codecrafters.io/progress/shell/bd5c2931-915e-4dc2-bda5-a5681647b3f3)](https://app.codecrafters.io/users/codecrafters-bot?r=2qF)

# Build Your Own Shell in Go

Welcome to my solution for the ["Build Your Own Shell" Challenge](https://app.codecrafters.io/courses/shell/overview) by Codecrafters! 🚀 

In this challenge, I've implemented a POSIX-compliant shell capable of interpreting shell commands, executing external programs, and handling built-in commands like `cd`, `pwd`, and `echo`. This project dives deep into concepts like shell command parsing, REPLs (Read-Eval-Print Loops), and the mechanics of how a shell interacts with the operating system. It's been a fantastic way to explore Go's standard library while building something practical.

> **Note**: If you're viewing this repo on GitHub, check out [codecrafters.io](https://codecrafters.io) to try the challenge yourself!

---

## Getting Started

### Prerequisites
- Go 1.19 or later installed on your system.
- A Unix-like environment (Linux or macOS). For Windows, use WSL or a compatible terminal emulator.

### How to Run the Shell
1. Clone this repository:
   ```sh
   git clone https://github.com/your-username/codecrafters-shell.git
   cd codecrafters-shell
2. Run the shell program:
   ```sh
   ./your_program.sh
3. Start entering shell commands:
   ```sh
   $ echo "Hello, Shell!"
   Hello, Shell!
4. To exit the shell, use:
   ```sh
   exit

## How It Works

- The shell is implemented as a **REPL** in Go. 
- Commands are read from standard input, parsed into components (handling quotes and escape sequences), and executed either as built-in or external commands.
- Key features implemented:
  - **Built-in Commands**: Includes `cd`, `pwd`, `echo`, etc.
  - **Quoting Support**: Handles single quotes, double quotes, and backslashes.
  - **External Command Execution**: Uses Go's `os/exec` package to run external programs.

---

## Progress

This implementation follows the stages outlined in the Codecrafters challenge:
- **Stage 1**: Basic shell functionality (REPL setup and external command execution).
- **Stage 2**: Support for built-in commands like `echo`, `cd`, and `pwd`.
- **Stage 3 and Beyond**: Advanced features like quoting, backslash handling, and argument parsing.

Follow the progress banner at the top to see how far this project has come!

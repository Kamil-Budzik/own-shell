package main

func main() {
	registerCommand("exit", handleExit)
	registerCommand("echo", handleEcho)
	registerCommand("type", handleType)
	registerCommand("pwd", handlePwd)

	for {
		handleCommand()
	}
}

package main

import (
	"fmt"
	"log"
	"os/exec"
)

/**
LookPath searches for an executable named file in the directories named by the PATH environment variable.
If file contains a slash, it is tried directly and the PATH is not consulted.
The result may be an absolute path or a path relative to the current directory.
找是否存在可执行得文件或命令
 */
func main() {
	path, err := exec.LookPath("cmd")
	if err != nil {
		log.Fatal("installing fortune is in your future")
	}
	fmt.Printf("fortune is available at %s\n", path)
}

// Contains some util functions for the other CLI scripts
package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//copyFile copies (more like links) the currently traversed file to the destination
func copyFile(source, dest string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		d := strings.Replace(path, source, dest, 1)
		if info.IsDir() {
			//if it's a dir, make a new dir
			os.Mkdir(d, os.ModePerm)
		} else {
			//if it's not a dir, create a hard link
			os.Link(path, d)
		}

		return nil
	}
}

//CopyDir copies (more like hard links) the contents of source into dest
func CopyDir(source, dest string) {

	log.Printf("Copying contents of %s in %s", source, dest)

	if err := os.RemoveAll(dest); err != nil {
		log.Fatal(err)
	}

	if err := filepath.Walk(source, copyFile(source, dest)); err != nil {
		log.Fatal(err)
	}
}

//RunCommand executes a given command with the provided args and tails the output
func RunCommand(command string, args []string) {

	cmd := exec.Command(command, args...)
	log.Println("Executing: ", cmd)

	//create a pipe for the output of the script
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal("Error creating StdoutPipe", err)
	}

	//tail the output of the command
	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Printf("\t > %s\n", scanner.Text())
		}
	}()

	//start and wait for the command to finish
	if err := cmd.Start(); err != nil {
		log.Fatal("Error starting command", err)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal("Error waiting for command", err)
	}
}

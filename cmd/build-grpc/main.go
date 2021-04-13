//Script to compile the protobuf files and replicate them across each service's directory
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//compileCommand is the command used for generating Golang code from the proto files
var compileCommand string = "protoc"

//compileArgs are the args passed to complieCommand
var compileArgs string = "--go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative %s"

//sourceDir is the dir where the grpc contract files is contained
var sourceDir string = "grpc-contract"

//desinationtDirs are the dirs where the compiled proto files will be copied to
var desinationtDirs []string = []string{
	filepath.Join("tracker-service", "grpc-contract"),
	filepath.Join("gateway-service", "grpc-contract"),
}

//protoPaths are the exact path to the proto files
var protoPaths []string = []string{
	filepath.Join(sourceDir, "tracker-service", "v1", "issue", "issue.proto"),
	filepath.Join(sourceDir, "tracker-service", "v1", "project", "project.proto"),
}

//compileProto runs the proto compilation commands for the proto files
func compileProto() {

	for _, proto := range protoPaths {
		log.Println("Compiling: ", proto)
		args := strings.Split(fmt.Sprintf(compileArgs, proto), " ")
		cmd := exec.Command(compileCommand, args...)
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}

}

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

//copyDir copies (more like hard links) the contents of source into dest
func copyDir(source, dest string) {

	log.Printf("Copying contents of %s in %s", source, dest)

	if err := os.RemoveAll(dest); err != nil {
		log.Fatal(err)
	}

	if err := filepath.Walk(source, copyFile(source, dest)); err != nil {
		log.Fatal(err)
	}
}

func main() {

	//chdir to the root of the project
	os.Chdir(filepath.Join("..", ".."))

	//compile the proto files
	compileProto()

	//copy the compiled files into the root of each service
	for _, dest := range desinationtDirs {
		copyDir(sourceDir, dest)
	}

}

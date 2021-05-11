//Script to compile the protobuf files and replicate them across each service's directory
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Peshowe/issue-tracker/cmd/utils"
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
	filepath.Join("mail-service", "grpc-contract"),
}

//protoPaths are the exact path to the proto files
var protoPaths []string = []string{
	filepath.Join(sourceDir, "tracker-service", "v1", "issue", "issue.proto"),
	filepath.Join(sourceDir, "tracker-service", "v1", "project", "project.proto"),
	filepath.Join(sourceDir, "mail-service", "v1", "mailer", "mailer.proto"),
}

//compileProto runs the proto compilation commands for the proto files
func compileProto() {

	for _, proto := range protoPaths {
		log.Println("Compiling: ", proto)
		args := strings.Split(fmt.Sprintf(compileArgs, proto), " ")
		utils.RunCommand(compileCommand, args)
	}

}

func main() {

	//chdir to the root of the project
	os.Chdir(filepath.Join("..", ".."))

	//compile the proto files
	compileProto()

	//copy the compiled files into the root of each service
	for _, dest := range desinationtDirs {
		utils.CopyDir(sourceDir, dest)
	}

}

//Script to build the react frontend and copy the build into the gateway service directory
package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/Peshowe/issue-tracker/cmd/utils"
)

//compileCommand is the command used for building the react app
var compileCommand string = "npm"

//compileArgs are the args passed to complieCommand
var compileArgs string = "--prefix frontend-client run build"

//sourceDir is the dir where the built react app is contained
var sourceDir string = filepath.Join("frontend-client", "build")

//desinationtDir is the dir where the build will be copied to
var destinationDir string = filepath.Join("gateway-service", "frontend", "build")

//buildReact runs the build commands for app
func buildReact() {

	args := strings.Split(compileArgs, " ")
	utils.RunCommand(compileCommand, args)

}

func main() {

	//chdir to the root of the project
	os.Chdir(filepath.Join("..", ".."))

	//build the app
	buildReact()

	//copy the build to the gateway service
	utils.CopyDir(sourceDir, destinationDir)
}

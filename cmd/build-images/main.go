//Script to build the images for each of the services
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Peshowe/issue-tracker/cmd/utils"
)

var dockerCommand string = "docker"

//dockerBuildArgs are the arguements for the build command
var dockerBuildArgs string = "build -t %s --rm %s"

//serviceDirs are the directories of each of the services
var serviceDirs []string = []string{
	"tracker-service",
	"gateway-service",
	"mail-service",
}

//readArgs reads the CLI args passed (should be names of services to build)
func readArgs() []string {
	return os.Args[1:]
}

//buildImages build the images for the passed services
func buildImages(services []string) {

	for _, service := range services {
		log.Println("Building: ", service)
		args := strings.Split(fmt.Sprintf(dockerBuildArgs, service, service), " ")
		utils.RunCommand(dockerCommand, args)
	}

}

func main() {

	//chdir to the root of the project
	os.Chdir(filepath.Join("..", ".."))

	serviceArgs := readArgs()

	//build the images for the given services, or for all servies if no args were passed
	if len(serviceArgs) > 0 {
		buildImages(serviceArgs)
	} else {
		buildImages(serviceDirs)
	}

}

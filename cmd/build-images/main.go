//Script to build the images for each of the services
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var dockerCommand string = "docker"

//dockerBuildArgs are the arguements for the build command
var dockerBuildArgs string = "build -t %s --rm %s"

//serviceDirs are the directories of each of the services
var serviceDirs []string = []string{
	"tracker-service",
	"gateway-service",
}

//readArgs reads the CLI args passed (should be names of services to build)
func readArgs() []string {
	return os.Args[1:]
}

//runCommand executes the dockerCommand with the provided args and tails the output
func runCommand(args []string) {

	cmd := exec.Command(dockerCommand, args...)
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

//buildImages build the images for the passed services
func buildImages(services []string) {

	for _, service := range services {
		log.Println("Building: ", service)
		args := strings.Split(fmt.Sprintf(dockerBuildArgs, service, service), " ")
		runCommand(args)
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

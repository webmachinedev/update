package main

import (
	"io/fs"
	"log"
	"os"
	"os/exec"
)

var containerIDFile string = ".core_container_id"
var dockerImage string = "ghcr.io/webmachinedev/core:main"

func main() {
	var err error
	err = pullLatestImage()
	if err != nil {
		log.Fatal(err)
	}
	err = stopRunningContainer()
	if err != nil {
		log.Fatal(err)
	}
	err = startNewContainer()
	if err != nil {
		log.Fatal(err)
	}
}

func pullLatestImage() error {
	c := exec.Command("docker", "pull", dockerImage)
	return c.Run()
}


func stopRunningContainer() error {
	id, err := containerID()
	if err != nil {
		return err
	}
	c := exec.Command("docker", "stop", id)
	return c.Run()
}

func containerID() (string, error) {
	id, err := os.ReadFile(containerIDFile)
	if err != nil {
		return "", err
	}
	return string(id), nil
}

func writeContainerID(id []byte) error {
	return os.WriteFile(containerIDFile, id, fs.ModePerm)
}

func startNewContainer() error {
	c := exec.Command("docker", "run", "-d", "-p", "80:80", dockerImage)
	containerID, err := c.Output()
	if err != nil {
		return err
	}
	return writeContainerID(containerID)
}
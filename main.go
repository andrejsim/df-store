package main

import (
	"fmt"
	"os/exec"
)

func main(){
	// come out of package b and then go inside package a to run the executable file as
	cmd := exec.Command("../datafabric/app/call.go")
	if err := cmd.Run(); err != nil{
		fmt.Println(err)
	}
}


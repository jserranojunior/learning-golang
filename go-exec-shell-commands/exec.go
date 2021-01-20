package main

import "os/exec"
import . "fmt"

func main() {
    command := "ls"
    cmd := exec.Command(command)
    stdout, err := cmd.Output()

    if err != nil {
        Println(err.Error())
        return
    }

    Print(string(stdout))
}
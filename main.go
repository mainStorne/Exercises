package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"os/exec"
)

func main() {

	fi, err := os.OpenFile("io.in", os.O_CREATE&os.O_WRONLY, 0)
	if err != nil {
		panic(err)
	}
	fi.WriteString(fmt.Sprint(rand.Int()))
	defer fi.Close()
	cmd := exec.Command("powershell", "/c", "cat", "io.in")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fo, _ := os.Create("io.out")
	buffer := make([]byte, 4)
	for {
		n, err := fi.Read(buffer)
		log.Println(n)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		if _, err := fo.Write(buffer[:n]); err != nil {
			panic(err)
		}
	}

}

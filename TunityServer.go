package main

import (
	"fmt"
	"os"
	"strconv"
	"errors"
)

type UDP_Server struct {
	portIn int
	portOut int
}

func main(){
	portIn, portOut, err := GetPortsFromArgs()

	if err != 0{
		fmt.Println(err)
		os.Exit(err)
	}

	fmt.Println(portIn, portOut)
	server := UDP_Server{portIn,portOut}
	fmt.Println(server)
	//server := UDP_Server{int(os.Args[1]), os.Args[2]}
}

func GetPortsFromArgs() (portIn int ,portOut int,error int){
	args := os.Args
	portIn = 0
	portOut = 0
	if len(args) != 3 {
		error = -1
		return
	}
	error = -2
	portIn, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("input port:",args[1],"is not a number")
		fmt.Println(err)
		return
	}
	portOut, err = strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("output port:",args[2],"is not a number")
		fmt.Println(err)
		return
	}

	if portIn == portOut{
		err = errors.New("input port and output port cannot be equal")
		fmt.Println(err)
		return
	}
	error = 0

	return
}

func PrintUsage() {
	fmt.Println("go run TunityServer.go <portIn> <portOut>")
}

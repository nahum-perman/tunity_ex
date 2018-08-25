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
	argc := len(os.Args)
	if argc != 3 {
		PrintUsage()
		os.Exit(-1)
	}

	portIn, portOut, err := GetPortsFromArgs(os.Args)

	if err != nil{
		fmt.Println(err)
		os.Exit(-2)
	}

	fmt.Println(portIn, portOut)
	server := UDP_Server{portIn,portOut}
	fmt.Println(server)
	//server := UDP_Server{int(os.Args[1]), os.Args[2]}
}

func GetPortsFromArgs(args []string) (portIn int ,portOut int,err error){
	fmt.Println(args)
	portIn = 0
	portOut = 0
	portIn, err = strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("input port:",args[1],"is not a number")
		return
	}
	portOut, err = strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("output port:",args[2],"is not a number")
	}

	if portIn == portOut{
		err = errors.New("input port and output port cannot be equal")
	}
	return
}

func PrintUsage() {
	fmt.Println("go run TunityServer.go <portIn> <portOut>")
}

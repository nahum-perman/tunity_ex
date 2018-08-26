package main

import (
	"fmt"
	"os"
	"strconv"
	"errors"
	"net"
	"time"
)

type UDP_Server struct {
	portIn int
	portOut int
	uc *net.UDPConn
	isClientAConnected bool
}

const (
	ClientTypeA = iota
	ClientTypeB = iota
)

type ClientType int

type IncomingMessage struct {
	client ClientType
	port int
}

func main(){
	portIn, portOut, err := GetPortsFromArgs()

	if err != 0{
		fmt.Println(err)
		os.Exit(err)
	}

	fmt.Println(portIn, portOut)
	var server = UDP_Server{portIn,portOut, nil , false}
	fmt.Println(server)
	go startServer(&server)
	//server := UDP_Server{int(os.Args[1]), os.Args[2]}

	fmt.Println("Press the Enter Key to terminate the console screen!")
	var input string
	fmt.Scanln(&input)
	if server.uc != nil {
		server.uc.Close()
		time.Sleep(100)
	}
}

func startServer(server *UDP_Server) (err error){
	clientAUdpAddr, err := net.ResolveUDPAddr("udp",fmt.Sprintf("localhost:%d",server.portIn))
	if err != nil {
		fmt.Println(err)
		return
	}
	uc, err := net.ListenUDP("udp", clientAUdpAddr)

	if err != nil {
		fmt.Println(err)
		return
	}

	server.uc = uc
	buffer := make([]byte,1024)
	for ; err == nil; {
		n, err := uc.Read(buffer)
		if n > 0 {
			fmt.Println(n, buffer[:n])
			var addr string
			addr = string(buffer[:n])
			fmt.Println(addr)
			server.isClientAConnected = true
			uClientAAddr, err := net.ResolveUDPAddr("udp", addr)
			uc.WriteToUDP(buffer, uClientAAddr)
			fmt.Println(err)
			uClientACon ,err := net.DialUDP("udp",nil, uClientAAddr)
			fmt.Println(err)
			uClientACon.Write(buffer)
			fmt.Println(uc.RemoteAddr())
		}
		if err != nil {
			fmt.Println(err)
			break
		}
	}

	fmt.Println("closing listner to client A")
	uc.Close()
	return
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

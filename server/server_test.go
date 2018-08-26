package main

import (
	"testing"
	"fmt"
	"os"
	"net"
	"time"
	"bytes"
	"encoding/binary"
)

func TestPrintUsage(t *testing.T) {
	PrintUsage()
}

func ExamplePrintUsage() {
	PrintUsage()

	//Output: go run TunityServer.go <portIn> <portOut>
}

func TestGetPortsFromArgs(t *testing.T) {
	argStr := []string{"path","2356","324e"}
	fmt.Println(argStr)
	os.Args = argStr
	pin, pout, err := GetPortsFromArgs()
	if err == 0 {
		t.Error("GetPortFromArgs should fail")
	}
	if pout != 0 {
		t.Error("pout default should be zero instead it was", pout)
	}
	os.Args = []string{"path","aaaa","3243"}
	fmt.Println(os.Args)
	pin, pout, err = GetPortsFromArgs()
	if err == 0 {
		t.Error("GetPortFromArgs should fail")
	}
	if pin != 0 {
		t.Error("pin default should be zero instead it was", pin)
	}

	os.Args = []string{"path","1234","1234"}
	fmt.Println(os.Args)
	pin, pout, err = GetPortsFromArgs()
	if err == 0 {
		t.Error("GetPortFromArgs should fail when in port and out port are equal")
	}

	os.Args = []string{"path","1234","1235"}
	fmt.Println(os.Args)
	pin, pout, err = GetPortsFromArgs()
	if err != 0 {
		t.Error("GetPortFromArgs should success")
	}

}

func connectToServer(port int) (uc *net.UDPConn) {
	add := fmt.Sprintf("localhost:%d",port)
	ua, err := net.ResolveUDPAddr("udp",add)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	uc, err = net.DialUDP("udp",nil, ua)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return
}

func TestClientAListener(t *testing.T){
	fmt.Println("testing server listener")
	serverCtx := UDP_Server{4567, 7654, nil, false}
	go startServer(&serverCtx)
	time.Sleep(200)
	uc := connectToServer(serverCtx.portIn)
	buff := make([]byte, 1024)
	var b bytes.Buffer
	bCon := []byte(uc.LocalAddr().String())
	fmt.Printf("%s %s",fmt.Sprintf("%v", uc), uc.LocalAddr().String())
	binary.Write(&b, binary.LittleEndian, uc.LocalAddr())
	fmt.Println(b.Bytes())
	n, err := uc.Write(bCon)
	fmt.Println(n, err)
	uc.SetDeadline(time.Now().Add(time.Second))
	n, err = uc.Read(buff)
	fmt.Println(n, buff)
	uc.Close()
	serverCtx.uc.Close()
	if !serverCtx.isClientAConnected {
		t.Error("failed to open connection to server")
	} else {
		fmt.Println("Success")
	}


}

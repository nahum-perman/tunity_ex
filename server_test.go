package main

import (
	"testing"
	"fmt"
	"os"
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

package main

import (
	"testing"
	"fmt"
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
	pin, pout, err := GetPortsFromArgs(argStr)
	if err == nil {
		t.Error("GetPortFromArgs should fail")
	}
	if pout != 0 {
		t.Error("pout default should be zero instead it was", pout)
	}
	argStr = []string{"path","aaaa","3243"}
	fmt.Println(argStr)
	pin, pout, err = GetPortsFromArgs(argStr)
	if err == nil {
		t.Error("GetPortFromArgs should fail")
	}
	if pin != 0 {
		t.Error("pin default should be zero instead it was", pin)
	}

	argStr = []string{"path","1234","1234"}
	fmt.Println(argStr)
	pin, pout, err = GetPortsFromArgs(argStr)
	if err == nil {
		t.Error("GetPortFromArgs should fail when in port and out port are equal")
	}

	argStr = []string{"path","1234","1235"}
	fmt.Println(argStr)
	pin, pout, err = GetPortsFromArgs(argStr)
	if err != nil {
		t.Error("GetPortFromArgs should success")
	}

}

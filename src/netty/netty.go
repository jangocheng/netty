package main

import (
	"fmt"

	"lib"
)

func main() {
	ipCmd := &lib.IpCmd{
		Namespace: "ns0",
	}

	//fmt.Printf(ipCmd.GetNamespace())
	ipCmd.SetNamespace("")
	//fmt.Printf(ipCmd.GetNamespace())
	//r, e := ipCmd.GetInterfacesName()
	//fmt.Print(r)
	//fmt.Print(e)
	r, e := ipCmd.GetInterfaceDetails("enp0s10")
	fmt.Print(r)
	fmt.Print(e)
}

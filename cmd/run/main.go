package main

import (
	n2n "github.com/ivyxjc/n2n-go-cli/internal"
	log "github.com/sirupsen/logrus"
)

func createEdge() *n2n.Edge {
	edge := new(n2n.Edge)
	edge.CommunityName = "c1"
	edge.SuperNodeNum = 0
	edge.RegisterInterval = 20
	edge.DeviceName = "host1"
	edge.DeviceIPMode = "static"
	edge.DeviceIP = "100.64.0.1"
	edge.DeviceMask = "255.255.255.0"
	edge.DisablePMTUDiscovery = true
	edge.EncryptKey = "secret1"
	edge.SuperNodeHostPort = "106.14.27.58:8089"
	edge.TransopId = 2
	edge.DeviceMac = "asdf"
	edge.MTU = 1500
	return edge
}

func main() {
	edge := createEdge()
	if err := edge.Configure(); err != nil {
		log.Fatal(err)
	}
	//if err := edge.OpenTunTapDevice(); err != nil {
	//	log.Fatal(err)
	//}
}

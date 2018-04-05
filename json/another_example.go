//https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/07.2.html

package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type ServerSlice struct {
	Servers []Server
}

func main() {
	//json to struct
	fmt.Println("json to struct:")
	var s ServerSlice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)

	// struct to json
	fmt.Println("\nstruct to json:")
	var z ServerSlice
	z.Servers = append(z.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	z.Servers = append(z.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
	b, err := json.Marshal(z)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}

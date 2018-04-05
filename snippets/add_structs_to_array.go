package main

import (
	"fmt"
)

func main() {
	type Result struct {
		SessionID int
		Hostname  string
		Check     bool
		URL       string
	}

	somedevice01 := Result{13, "some_device_01", true, "https://devicecheck.scapp-corp.swisscom.com/report"}
	somedevice02 := Result{13, "some_device_02", false, "https://devicecheck.scapp-corp.swisscom.com/report"}

	var response []Result
	response = append(response, somedevice01, somedevice02)
	fmt.Println(response)
}

//Use Case for ahc_sim to
//This adds the the urls to the already saved job (the cpe do not contain the urls).

package main

import (
	"fmt"
	"strconv"
)

func main() {
	type cpe struct {
		id  int
		url string
	}
	type job struct {
		jobid   int
		cpeList []cpe
	}
	//Creating a cpe list and adding to a job struct
	var cpeList []cpe
	for i := 0; i < 5; i++ {
		cpeList = append(cpeList, cpe{i, ""})
	}
	newJob := job{1, cpeList}

	//Extract all cpe from the job and add the matching url to each of them
	for i, v := range newJob.cpeList {
		url := "https://fakeurl/"
		url = url + strconv.Itoa(v.id)
		// v.url = url // This does not work
		newJob.cpeList[i].url = url
	}
	fmt.Println(newJob)
}

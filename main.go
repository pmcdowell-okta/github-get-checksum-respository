package main

import (
	"os"
	"strings"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"time"
)

type ShaStruct struct {
	Sha string `json:"sha"`
}

func main() {

	if len(os.Args) != 1 {

		gitHubRespository := strings.ToLower(os.Args[1]) //"https://hardCodeYourOktaOrg.oktapreview.com"
		currentCheckSum:=getChecksum(gitHubRespository)
		duration := time.Second*1

		for true {

			newChecksum := getChecksum(gitHubRespository) //"https://hardCodeYourOktaOrg.oktapreview.com"
			fmt.Println("Current Checksum: "+newChecksum)

			if (newChecksum!=currentCheckSum) {
				os.Exit(0)
			}

			time.Sleep(duration)


		}

	} else {
		fmt.Println("Usage: command https://yourgithubrepository.git")

	}
}

func getChecksum ( githubUrl string)  string {

	gitHubRespository:=githubUrl
	dat := ShaStruct{}


	if (  strings.HasPrefix(gitHubRespository, "https://github.com/")) {
		gitHubRespository = strings.Replace(gitHubRespository, "https://github.com/", "", -1)
		gitHubRespository = strings.Replace(gitHubRespository, ".git", "", -1)
	}

	gitHubRespository = fmt.Sprintf("https://api.github.com/repos/%s/commits/master", gitHubRespository)

	response, _ := http.Get(gitHubRespository)
	if response.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(response.Body)
		if err := json.Unmarshal(bodyBytes, &dat); err != nil {
			panic(err)
		}


	}

	return (dat.Sha)

}

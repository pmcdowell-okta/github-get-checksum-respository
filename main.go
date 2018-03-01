package main

import (
	"os"
	"strings"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type ShaStruct struct {
	Sha string `json:"sha"`
}

func main() {

	if len(os.Args) != 1 {

		gitHubRespository := strings.ToLower(os.Args[1]) //"https://hardCodeYourOktaOrg.oktapreview.com"

		if (  strings.HasPrefix(gitHubRespository, "https://github.com/")) {
			gitHubRespository = strings.Replace(gitHubRespository, "https://github.com/", "", -1)
			gitHubRespository = strings.Replace(gitHubRespository, ".git", "", -1)
		}

		gitHubRespository = fmt.Sprintf("https://api.github.com/repos/%s/commits/master", gitHubRespository)

		fmt.Println(gitHubRespository)

		response, _ := http.Get(gitHubRespository)
		if response.StatusCode == http.StatusOK {
			bodyBytes, _ := ioutil.ReadAll(response.Body)
			//bodyString := string(bodyBytes)
			//fmt.Println(bodyString)

			dat := ShaStruct{}
			if err := json.Unmarshal(bodyBytes, &dat); err != nil {
				panic(err)
			}

			fmt.Println(dat.Sha)

		}
	}
}

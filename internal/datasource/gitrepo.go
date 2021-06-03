package datasource

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetValueFromGitRepo(url string) string {

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln("error reaching repo")
	}
	//req.Header.Add("Accept", "application/vnd.github.v3.raw")
	//req.Header.Add("authorization", "token 929f19719c9c9aac8c37c3a3766ebfce211cf5a9")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("error reaching repo")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	return sb

}

package datasource

import (
	//"context"
	"io/ioutil"
	"log"
	"net/http"
)

func (source *DataSource) GetFeature(partialUrl string) (interface{}, bool) {

	client := &http.Client{}
	finalURL := source.Url + partialUrl

	req, err := http.NewRequest("GET", finalURL, nil)
	if err != nil {
		log.Panicf("http error object %s", finalURL)
		return nil, true
	}
	//req.Header.Add("Accept", "application/vnd.github.v3.raw")
	//req.Header.Add("authorization", "token 929f19719c9c9aac8c37c3a3766ebfce211cf5a9")
	resp, err := client.Do(req)
	if err != nil {
		log.Panicf("error reaching repo %s", finalURL)
		return nil, true
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("error processing answer")
	}
	sb := string(body)
	return sb, true

}

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	client "github.com/akamai/AkamaiOPEN-edgegrid-golang/client-v1"
)

func errorCheck(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func printJSON(str interface{}) {
	jsonRes, _ := json.MarshalIndent(str, "", "  ")
	fmt.Printf("%+v\n", string(jsonRes))
}

func fetchData(urlPath, method string, body io.Reader) (result string) {
	req, err := client.NewRequest(edgeConfig, method, urlPath, body)
	errorCheck(err)

	resp, err := client.Do(edgeConfig, req)
	errorCheck(err)

	defer resp.Body.Close()
	byt, _ := ioutil.ReadAll(resp.Body)

	return string(byt)
}

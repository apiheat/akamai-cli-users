package main

import (
	"io"
	"io/ioutil"

	client "github.com/akamai/AkamaiOPEN-edgegrid-golang/client-v1"
	common "github.com/apiheat/akamai-cli-common"
)

func fetchData(urlPath, method string, body io.Reader) (result string) {
	req, err := client.NewRequest(edgeConfig, method, urlPath, body)
	common.ErrorCheck(err)

	resp, err := client.Do(edgeConfig, req)
	common.ErrorCheck(err)

	defer resp.Body.Close()
	byt, _ := ioutil.ReadAll(resp.Body)

	return string(byt)
}

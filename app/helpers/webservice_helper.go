package helpers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	lo "log"
	"net/http"
)

const msgFailedExecWebService = "failed create request"

func ExecWebService(url string, method string, headers http.Header, payload *bytes.Buffer, respond interface{}) {
	req, err := http.NewRequest(method, url, payload)
	Panics("error creating request", err)

	req.Header = headers
	client := &http.Client{}

	resp, err := client.Do(req)
	Panics("error making request", err)
	defer resp.Body.Close()

	reqResp, err := ioutil.ReadAll(resp.Request.Body)
	body, err := ioutil.ReadAll(resp.Body)
	Panics("error reading response", err)

	lo.Println("REQUEST : ", string(reqResp))
	lo.Println("RESPONSE : ", string(body))
	lo.Println("HTTP STATUS : ", resp.StatusCode)
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		err = errors.New(string(body))
		Panics("error descripting response", err)
		return
	}
	err = json.Unmarshal([]byte(string(body)), &respond)
	Panics("error marshaling response", err)
}

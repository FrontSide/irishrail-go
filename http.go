package irishrail

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func sendHTTPRequestXML(request *http.Request, responseBody interface{}) (int, error) {

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return -1, err
	}

	defer response.Body.Close()
	responseBodyRaw, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return response.StatusCode, err
	}

	log.Printf("raw response body: %s", responseBodyRaw)

	err = xml.Unmarshal(responseBodyRaw, &responseBody)
	if err != nil {
		return response.StatusCode, err
	}

	return response.StatusCode, nil

}

func SendHTTPGetRequestXML(url string, urlParameters url.Values, responseBody interface{}) error {

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	request.URL.RawQuery = urlParameters.Encode()

	log.Printf("get request: %s", request.URL.RawQuery)

	statusCode, err := sendHTTPRequestXML(request, responseBody)
	if err != nil {
		return err
	}

	if statusCode != 200 {
		return fmt.Errorf("server responded with unaccepted code: %d", statusCode)
	}

	return nil

}

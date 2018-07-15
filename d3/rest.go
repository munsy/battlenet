package d3

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Convert an HTTP response from a given endpoint to the supplied interface.
// This function expects the body to contain the associated JSON response
// from the given endpoint and will return an error if it fails to properly unmarshal.
func (c *D3Client) get(endpoint string, v interface{}) error {
	if nil == v {
		return battlenet.ErrorNoInterfaceSupplied
	}

	response, err := http.Get(endpoint)
	if nil != err {
		return err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if nil != err {
		return err
	}

	err = json.Unmarshal([]byte(body), &v)
	if nil != err {
		return err
	}

	return nil
}

package quote

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/cloudfoundry/libbuildpack"
)

type Law []struct {
	Name  string `json: "name"`
	Quote string `json: "name"`
}

type LawRetriever struct {
	Log    *libbuildpack.Logger
	Client HttpClient
}

func (lawRetriever LawRetriever) RetrieveLaw(source string) (Law, error) {
	log := lawRetriever.Log
	client := lawRetriever.Client

	if len(source) <= 0 {
		log.Error("invalid source provided")
		return nil, errors.New("source must be provided")
	}

	req, err := http.NewRequest(http.MethodGet, source, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	laws := Law{}
	err = json.Unmarshal(body, &laws)
	if err != nil {
		return nil, err
	}

	//	law := laws[0].Name + laws[0].quote

	return laws, nil
}

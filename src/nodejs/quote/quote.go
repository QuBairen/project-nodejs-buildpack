package quote

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"errors"

	"github.com/cloudfoundry/libbuildpack"
)

type Law []struct {
	Name  string `json: "name"`
	Quote string `json: "name"`
}

type LawRetriever struct {
	Log *libbuildpack.Logger
}

func (lawRetriever LawRetriever) RetrieveLaw(source string) error {
	log := lawRetriever.Log
	//TODO if source is nil

	if (len(source) <= 0) {
		log.Error("invalid source provided")
		return errors.New("source must be provided")
	}
	httpClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, source, nil)
	if err != nil {
		return err
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	laws := Law{}
	err = json.Unmarshal(body, &laws)
	if err != nil {
		return err
	}
	fmt.Println(laws[0].Name)

	//TODO return laws as well
	return nil
}

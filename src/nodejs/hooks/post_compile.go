package hooks

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/cloudfoundry/nodejs-buildpack/src/nodejs/quote"

	"github.com/cloudfoundry/libbuildpack"
)

type PostCompileHook struct {
	libbuildpack.DefaultHook
	Log       *libbuildpack.Logger
	LawClient quote.LawClient
}

func init() {
	logger := libbuildpack.NewLogger(os.Stdout)
	httpClient := &http.Client{
		Timeout: time.Second * 2,
	}
	lawClient := quote.LawRetriever{
		Log:    logger,
		Client: httpClient,
	}
	libbuildpack.AddHook(PostCompileHook{
		Log:       logger,
		LawClient: &lawClient,
	})
}

func (h PostCompileHook) AfterCompile(stager *libbuildpack.Stager) error {
	lawRetriever := h.LawClient
	if strings.EqualFold(os.Getenv("PCH_ENABLED"), "TRUE") {
		h.Log.Info("Legality check...")
		laws, err := lawRetriever.RetrieveLaw("http://goins.me/laws.json")

		if err != nil {
			return err
		}

		rand.Seed(time.Now().UnixNano())
		lawIdx := rand.Intn(len(laws))
		fmt.Println(lawIdx)
		law := laws[lawIdx].Name + " ::: " + laws[lawIdx].Quote
		h.Log.Info(law)

		// Something must go wrong
		if strings.EqualFold("Murphy's law", laws[lawIdx].Name) {
			return errors.New(law)
		}
	}
	return nil
}

package quote_test

import (
	"bytes"
	"errors"
	"net/http"
	"time"

	"github.com/cloudfoundry/libbuildpack"

	"github.com/cloudfoundry/nodejs-buildpack/src/nodejs/quote"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("test quote client", func() {
	var (
		logger       *libbuildpack.Logger
		lawRetriever quote.LawRetriever
		buffer       *bytes.Buffer
		source       string
		httpClient   quote.HttpClient
	)

	BeforeEach(func() {
		buffer = new(bytes.Buffer)
		logger = libbuildpack.NewLogger(buffer)
	})

	Context("with a broken httpclient", func() {
		It("should catch client err", func() {

			mockDo := func(req *http.Request) (*http.Response, error) {
				return nil, errors.New("client do call returns error")
			}
			lawRetriever = quote.LawRetriever{
				Log: logger,
				Client: &MockClient{
					MockDo: mockDo,
				},
			}
			source = "http://goins.me/laws.json"
			laws, err := lawRetriever.RetrieveLaw(source)
			Expect(err).ToNot(BeNil())
			Expect(laws).To(BeNil())

		})

	})
	Context("with a working httpclient", func() {
		BeforeEach(func() {
			httpClient = &http.Client{
				Timeout: time.Second * 2,
			}
			lawRetriever = quote.LawRetriever{
				Log:    logger,
				Client: httpClient,
			}
		})
		Context("Source is not provided", func() {
			It("should error", func() {
				source = ""
				laws, err := lawRetriever.RetrieveLaw(source)
				Expect(laws).To(BeNil())
				Expect(err).To(MatchError("source must be provided"))
			})
		})

//TODO Integration test should not be included...
		Context("read from webserice", func() {
			It("should not error", func() {
				source = "http://goins.me/laws.json"
				laws, err := lawRetriever.RetrieveLaw(source)
				Expect(err).To(BeNil())
				Expect(laws).ToNot(BeNil())
			})
		})

	})

})

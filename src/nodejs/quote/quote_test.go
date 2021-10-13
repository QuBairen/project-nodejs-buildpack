package quote_test

import (
	"bytes"

	"github.com/cloudfoundry/libbuildpack"

	"github.com/cloudfoundry/nodejs-buildpack/src/nodejs/quote"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("test quote client", func() {
	var (
		logger          *libbuildpack.Logger
		lawRetriever    quote.LawRetriever
		buffer          *bytes.Buffer
		source          string
	)

	BeforeEach(func() {
		buffer = new(bytes.Buffer)
		logger = libbuildpack.NewLogger(buffer)
		lawRetriever = quote.LawRetriever{
			Log: logger,
		}
	})

	Context("Source is not provided",func(){
		It("should error", func(){
			source =""
			err := lawRetriever.RetrieveLaw(source)
			Expect(err).To(MatchError("source must be provided"))
		})
	})

	Context("read from webserice",func(){
		It("should not error", func(){
			source ="http://goins.me/laws.json"
			err := lawRetriever.RetrieveLaw(source)
			Expect(err).To(BeNil())
		})
	})


})

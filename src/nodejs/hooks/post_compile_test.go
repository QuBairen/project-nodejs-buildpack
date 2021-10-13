package hooks_test

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/cloudfoundry/libbuildpack"

	"github.com/cloudfoundry/nodejs-buildpack/src/nodejs/hooks"
	"github.com/cloudfoundry/nodejs-buildpack/src/nodejs/quote"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type MockLawClient struct {
	MockRetrieve func(source string) (quote.Law, error)
}

func (m *MockLawClient) RetrieveLaw(source string) (quote.Law, error) {
	return m.MockRetrieve(source)

}

var _ = Describe("postCompileHook", func() {
	var (
		logger          *libbuildpack.Logger
		stager          *libbuildpack.Stager
		buffer          *bytes.Buffer
		postCompileHook hooks.PostCompileHook
	)
	// create hook
	BeforeEach(func() {
		buffer = new(bytes.Buffer)
		logger = libbuildpack.NewLogger(buffer)
	})
	// create stager
	JustBeforeEach(func() {
		args := []string{"", "", "", ""}
		stager = libbuildpack.NewStager(args, logger, &libbuildpack.Manifest{})
	})
	Context("with ws dependencies", func() {
		BeforeEach(func() {
			lawClient := quote.LawRetriever{
				Log: logger,
				Client: &http.Client{
					Timeout: time.Second * 2,
				},
			}
			postCompileHook = hooks.PostCompileHook{
				Log:       logger,
				LawClient: &lawClient,
			}

			It("should echo results", func() {
				os.Setenv("PCH_ENABLED", "TRUE")
				err := postCompileHook.AfterCompile(stager)
				// if on the off chance this pulls back murphy (integration test)
				//Integration test should not really be included...
				if err != nil {
					Expect(fmt.Sprint(err)).To(ContainSubstring("Murphy"))
				}
				bufferString := buffer.String()
				Expect(bufferString).To(ContainSubstring("Legality Check"))
				Expect(bufferString).To(ContainSubstring(" ::: "))
			})
		})
	})
	Context("with no dependencies", func() {
		BeforeEach(func() {
			mockRetrieve := func(source string) (quote.Law, error) {
				laws := quote.Law{
					{Name: "Murphy's law", Quote: "Anything that can go wrong will go wrong."},
				}
				return laws, nil
			}
			postCompileHook = hooks.PostCompileHook{
				Log: logger,
				LawClient: &MockLawClient{
					MockRetrieve: mockRetrieve,
				},
			}
		})
		It("should error", func() {
			os.Setenv("PCH_ENABLED", "TRUE")
			err := postCompileHook.AfterCompile(stager)
			Expect(err).ToNot(BeNil())
		})
	})
	Context("with no dependencies", func() {
		BeforeEach(func() {
			mockRetrieve := func(source string) (quote.Law, error) {
				laws := quote.Law{
					{Name: "law one", Quote: "something"},
				}
				return laws, nil
			}
			postCompileHook = hooks.PostCompileHook{
				Log: logger,
				LawClient: &MockLawClient{
					MockRetrieve: mockRetrieve,
				},
			}
		})
		Context("PostCompileHook(PCH) is not enabled", func() {
			It("should not write messages", func() {
				os.Setenv("PCH_ENABLED", "FALSE")
				err := postCompileHook.AfterCompile(stager)
				Expect(err).To(BeNil())
				Expect(buffer.String()).To(Equal(""))
			})
		})
		Context("PostCompileHook(PCH) is enabled", func() {
			It("should write messages", func() {
				os.Setenv("PCH_ENABLED", "TRUE")
				err := postCompileHook.AfterCompile(stager)
				Expect(err).To(BeNil())
				Expect(buffer.String()).To(ContainSubstring("law one"))
			})
		})
	})
})

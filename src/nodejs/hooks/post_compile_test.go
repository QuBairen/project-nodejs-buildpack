package hooks_test

import (
	"bytes"
	"os"

	"github.com/cloudfoundry/libbuildpack"

	"github.com/cloudfoundry/nodejs-buildpack/src/nodejs/hooks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

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
		postCompileHook = hooks.PostCompileHook{
			Log: logger,
		}
	})
	// create stager
	JustBeforeEach(func() {
		args := []string{"", "", "", ""}
		stager = libbuildpack.NewStager(args, logger, &libbuildpack.Manifest{})
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
			Expect(buffer.String()).To(ContainSubstring("I was here in the logs"))
		})
	})
})

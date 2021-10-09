package hooks_test

import (
	"os"
	"bytes"

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

	Context("post compile does nothing", func() {
// Since i'm using a logger to carry the message
		os.Setenv("BP_DEBUG", "TRUE")
		It("should not error & write messages", func() {
			postCompileHook.AfterCompile(stager)
			Expect(buffer.String()).To(ContainSubstring("I was here in the logs"))
		})
	})
})

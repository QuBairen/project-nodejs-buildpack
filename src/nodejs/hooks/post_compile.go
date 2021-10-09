package hooks

import (
	"os"
	"strings"

	"github.com/cloudfoundry/libbuildpack"
)

type PostCompileHook struct {
	libbuildpack.DefaultHook
	Log *libbuildpack.Logger
}

func init() {
	logger := libbuildpack.NewLogger(os.Stdout)
	libbuildpack.AddHook(PostCompileHook{
		Log: logger,
	})
}

func (h PostCompileHook) AfterCompile(stager *libbuildpack.Stager) error {
	if strings.EqualFold(os.Getenv("PCH_ENABLED"), "TRUE") {
		h.Log.Info("I was here in the logs")
	}
	return nil
}

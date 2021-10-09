package hooks

import (
	"os"
	"github.com/cloudfoundry/libbuildpack"
)

type PostCompileHook struct {
	libbuildpack.DefaultHook
	Log               *libbuildpack.Logger
}

func init() {
	logger := libbuildpack.NewLogger(os.Stdout)
	libbuildpack.AddHook(PostCompileHook{
		Log: logger,
	})
}

func (h PostCompileHook) AfterCompile(stager *libbuildpack.Stager) error {
	h.Log.Debug("I was here in the logs")
	return nil
}

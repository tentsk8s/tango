package deploy

import (
	"github.com/spf13/cobra"
	"github.com/tentsk8s/tango/pkg/decoder"
	"github.com/tentsk8s/tango/pkg/log"
	"github.com/tentsk8s/tango/pkg/translate"
)

type runner struct{}

func (r *runner) run(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		return log.Error(true, "Don't pass any arguments to this command")
	}
	svc, err := decoder.Svc()
	if err != nil {
		return log.Err(true, err)
	}
	deployment := translate.Service(svc)
	log.Statusf(true, "deployment: %s", deployment)

	return nil
}

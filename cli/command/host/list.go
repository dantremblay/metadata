package host

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	log "github.com/Sirupsen/logrus"
	"github.com/juliengk/go-utils"
	"github.com/kassisol/metadata/cli/command"
	"github.com/kassisol/metadata/storage"
	"github.com/spf13/cobra"
)

var hostListFilter []string

func newListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ls",
		Aliases: []string{"list"},
		Short:   "List hosts",
		Long:    listDescription,
		Run:     runList,
	}

	flags := cmd.Flags()
	flags.StringSliceVarP(&hostListFilter, "filter", "f", []string{}, "Filter output based on conditions provided")

	return cmd
}

func runList(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", command.DBFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer s.End()

	filters := utils.ConvertSliceToMap("=", hostListFilter)

	hosts := s.ListHost(filters)

	if len(hosts) > 0 {
		w := tabwriter.NewWriter(os.Stdout, 20, 1, 2, ' ', 0)
		fmt.Fprintln(w, "ID\tENABLED\tNAME\tFQDN\tUUID\tINTERFACES\tPROFILE")

		for _, host := range hosts {
			fmt.Fprintf(w, "%d\t%t\t%s\t%s\t%s\t%s\t%s\n", host.ID, host.Enabled, host.Name, host.FQDN, host.UUID, strings.Join(host.Interfaces, ", "), host.Profile)
		}

		w.Flush()
	}
}

var listDescription = `
List hosts
`

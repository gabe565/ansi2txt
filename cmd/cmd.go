package cmd

import (
	"io"
	"os"

	"gabe565.com/ansi2txt/pkg/ansi2txt"
	"gabe565.com/utils/cobrax"
	"gabe565.com/utils/termx"
	"github.com/spf13/cobra"
)

func New(options ...cobrax.Option) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ansi2txt [file]",
		Short: "Drop ANSI control codes",
		Long: `Convert text containing ANSI control codes into plain ASCII text.
It works as a filter, reading from stdin or a file, removing all ANSI codes, and sending the output to stdout.`,
		Args: cobra.MaximumNArgs(1),
		RunE: run,
	}

	for _, option := range options {
		option(cmd)
	}

	return cmd
}

func run(cmd *cobra.Command, args []string) error {
	w := ansi2txt.NewWriter(cmd.OutOrStdout())

	if len(args) == 0 || args[0] == "-" {
		if termx.IsTerminal(cmd.InOrStdin()) {
			return cmd.Usage()
		}

		_, err := io.Copy(w, cmd.InOrStdin())
		return err
	}

	f, err := os.Open(args[0])
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()

	_, err = io.Copy(w, f)
	return err
}

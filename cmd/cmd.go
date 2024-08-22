package cmd

import (
	"io"
	"os"

	"github.com/gabe565/ansi2txt/pkg/ansi2txt"
	"github.com/mattn/go-isatty"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ansi2txt",
		Short: "Drop ANSI control codes",
		Long: `Convert text containing ANSI control codes into plain ASCII text.
It works as a filter, reading from stdin, removing all ANSI codes, and sending the output to stdout.`,
		RunE: run,
	}
	initVersion(cmd)
	return cmd
}

func run(cmd *cobra.Command, _ []string) error {
	if isatty.IsTerminal(os.Stdin.Fd()) || isatty.IsCygwinTerminal(os.Stdin.Fd()) {
		return cmd.Help()
	}

	w := ansi2txt.NewWriter(cmd.OutOrStdout())
	_, err := io.Copy(w, cmd.InOrStdin())
	return err
}

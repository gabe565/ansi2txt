package cmd

import (
	"runtime/debug"

	"github.com/spf13/cobra"
)

func initVersion(cmd *cobra.Command) {
	var commit string
	var modified bool
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			switch setting.Key {
			case "vcs.revision":
				commit = setting.Value
			case "vcs.modified":
				if setting.Value == "true" {
					modified = true
				}
			}
		}
	}

	if commit != "" {
		if len(commit) > 8 {
			commit = commit[:8]
		}
		if modified {
			commit = "*" + commit
		}
		cmd.Version = commit
		cmd.SetVersionTemplate(`{{with .Name}}{{.}}{{end}} {{printf "commit %s" .Version}}
`)
		cmd.InitDefaultVersionFlag()
	}
}

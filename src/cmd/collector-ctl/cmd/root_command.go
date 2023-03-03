package cmd

import "github.com/spf13/cobra"

func Execute() {
	root := newRootCommand(config{
		SourceURL: defaultCardCacheURL,
	})
	err := root.Execute()
	cobra.CheckErr(err)
}

type config struct {
	SourceURL string
}

func newRootCommand(c config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "collector-ctl",
		Short: "Magic: The Gathering card collection manager.",
	}

	cmd.PersistentFlags().String("path", "~/.collector", "specify the directory where collection data is stored")

	cmd.AddCommand(newCacheCommand(c))

	return cmd
}

package cmd

import (
	"context"
	"fmt"
	"github.com/Legit-Labs/legitify/cmd/common_options"
	"github.com/Legit-Labs/legitify/internal/clients/github"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(newListOrgsCommand())
}

var listOrgsArgs args

func newListOrgsCommand() *cobra.Command {
	listOrgsCmd := &cobra.Command{
		Use:          "list-orgs",
		Short:        `List GitHub organizations associated with a PAT`,
		RunE:         executeListOrgsCommand,
		SilenceUsage: true,
	}

	viper.AutomaticEnv()
	flags := listOrgsCmd.Flags()
	flags.StringVarP(&listOrgsArgs.Token, common_options.ArgToken, "t", "", "token to authenticate with github (required unless environment variable GITHUB_TOKEN is set)")
	flags.StringVarP(&listOrgsArgs.OutputFile, common_options.ArgOutputFile, "o", "", "output file, defaults to stdout")
	flags.StringVarP(&listOrgsArgs.ErrorFile, common_options.ArgErrorFile, "e", "error.log", "error log path")

	return listOrgsCmd
}

func validateListOrgsArgs() error {
	if err := github.IsTokenValid(listOrgsArgs.Token); err != nil {
		return err
	}

	return nil
}

func executeListOrgsCommand(cmd *cobra.Command, _args []string) error {
	if listOrgsArgs.Token == "" {
		listOrgsArgs.Token = viper.GetString(common_options.EnvToken)
	}

	err := validateListOrgsArgs()
	if err != nil {
		return err
	}

	if err = setErrorFile(listOrgsArgs.ErrorFile); err != nil {
		return err
	}

	err = setOutputFile(listOrgsArgs.OutputFile)
	if err != nil {
		return err
	}

	ctx := context.Background()

	githubClient, err := github.NewClient(ctx, listOrgsArgs.Token, []string{})
	if err != nil {
		return err
	}

	orgs, err := githubClient.CollectOrganizations()
	if err != nil {
		return err
	}

	if len(orgs) == 0 {
		fmt.Printf("No organizations are associated with this PAT.\n")
	} else {
		fmt.Printf("Organizations:\n")
		for _, org := range orgs {
			fmt.Printf("- %s (%s)\n", *org.Login, org.Role)
		}
	}

	return nil
}
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(viewPrCommand)
}

var viewPrCommand = &cobra.Command{
	Use:   "view-pr [repository-name]",
	Args:  cobra.MaximumNArgs(1),
	Short: "View pull requests for a repository",
	Long:  "View the pull reposts for a specific repository. If no repository is specified, it takes the checked out repository",
	Run: func(_ *cobra.Command, _ []string) {
		repoName, err := getRepositoryName()
		if err != nil {
			fmt.Println(err)
		}

		status := "OPEN"
		repoUrl := fmt.Sprintf("https://us-west-2.console.aws.amazon.com/codesuite/codecommit/repositories/%s/pull-requests?status=%s", repoName, status)
		err = openURL(repoUrl)
		if err != nil {
			fmt.Println(err)
		}
	},
}

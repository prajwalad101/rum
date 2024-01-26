package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/prajwalad101/rum/util"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(prCommand)
}

var prCommand = &cobra.Command{
	Use:   "pr [source-branch] destination-branch",
	Args:  cobra.MinimumNArgs(1),
	Short: "Open pull requests in the browser",
	Long:  "Open pull requests by specifying source and destination branches. If no source branch is specified, it takes the currently checked out branch",
	Run: func(_ *cobra.Command, args []string) {
		var sourceBranch string
		var destinationBranch string

		// if only destination branch is specified
		if len(args) == 1 {
			destinationBranch = args[0]
			var err error
			sourceBranch, err = getBranchName()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			// get source branch from user
			sourceBranch = args[0]
			destinationBranch = args[1]
		}

		destinationBranch = strings.Trim(destinationBranch, "\n")
		sourceBranch = strings.Trim(sourceBranch, "\n")

		repositoryName, err := getRepositoryName()

		repoUrl := fmt.Sprintf("https://us-west-2.console.aws.amazon.com/codesuite/codecommit/repositories/%s/pull-requests/new/refs/heads/%s/.../refs/heads/%s?region=us-west-2", repositoryName, destinationBranch, sourceBranch)
		openURL(repoUrl)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func getBranchName() (string, error) {
	out, errout, err := util.Shellout("git branch --show-current")
	if err != nil {
		return "", err
	}

	if errout != "" {
		return "", errors.New(errout)
	} else {
		if out == "" {
			return "", fmt.Errorf("No branch checked out. Please checkout a branch.")
		}
	}

	return out, nil
}

func openURL(url string) error {
	command := fmt.Sprintf("xdg-open %s &", url)
	out, errout, err := util.Shellout(command)
	if err != nil {
		return err
	}

	if errout != "" {
		return errors.New(errout)
	} else {
		fmt.Println(out)
	}

	fmt.Println("Opening url ...")

	return nil
}

func getRepositoryName() (string, error) {
	remoteURL, errout, err := util.Shellout("git config --get remote.origin.url")
	if err != nil {
		return "", err
	}

	if errout != "" {
		return "", errors.New(errout)
	}

	splitRemoteURL := strings.Split(remoteURL, "@")
	repositoryName := splitRemoteURL[1]
	repositoryName = strings.Trim(repositoryName, "\n")

	return repositoryName, nil
}

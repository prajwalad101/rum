package cmd

import (
	"errors"
	"fmt"

	"github.com/prajwalad101/za/util"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(prCommand)
}

var prCommand = &cobra.Command{
	Use:  "pr",
	Args: cobra.MinimumNArgs(1),
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

		repoUrl := fmt.Sprintf("https://us-west-2.console.aws.amazon.com/codesuite/codecommit/repositories/veribom-fe/pull-requests/new/refs/heads/%s/.../refs/heads/%s?region=us-west-2", destinationBranch, sourceBranch)
		openURL(repoUrl)
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
	command := fmt.Sprintf("xdg-open %s", url)
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

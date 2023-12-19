package cli

import (
	"github.com/sirupsen/logrus"
)

type ProjectCmd struct {
	Key   string           `arg:"-k,required"`
	List  *ProjectListCmd  `arg:"subcommand:list"`
	Clone *ProjectCloneCmd `arg:"subcommand:clone"`
}

func (b *BitbucketCLI) RunProjectCmd(cmd *ProjectCmd) {
	if cmd == nil {
		b.logger.Fatal("Unable to execute command")
		return
	}

	if cmd.Key == "" {
		logrus.Fatal("A project key must be provided")
	}

	if cmd.List != nil {
		b.projectList(cmd)
		return
	} else if cmd.Clone != nil {
		b.projectClone(cmd)
		return
	}

	b.logger.Fatal(errSpecifySubcommand)
}

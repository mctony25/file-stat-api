package command

import (
	"file-stat/formatter"
	"file-stat/stat"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

type CliCommand struct{}

var (
	directoryName string
	cliCmd        = &cobra.Command{
		Use:   "list",
		Short: "List the files of a given directory",
		Run:   ListFile,
	}
)

func (cc CliCommand) GetCommand() *cobra.Command {

	cliCmd.PersistentFlags().StringVar(&directoryName, "directory", "/tmp", "config file (default is '/tmp')")

	return cliCmd
}

func ListFile(cmd *cobra.Command, args []string) {

	fileLister := stat.FileLister{}
	files, err := fileLister.List(directoryName)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, file := range files {
		fmt.Println(fmt.Sprintf("%d %s %s", file.Size(), file.ModTime().Format(time.RFC1123), file.Name()))
	}

	statAgg := stat.FileAggregate{}
	stats := statAgg.AggregateFromList(files)

	count := stats.TotalElement
	size := stats.TotalSize

	fmt.Println(fmt.Sprintf(
		"Total items: %d, totalling : %s",
		count,
		(formatter.SizeFormatter{}).Format(size),
	))
}

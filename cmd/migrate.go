package cmd

import (
	"errors"
	"fmt"

	"github.com/canbefree/magazine/cmd/mariadb"
	"github.com/canbefree/magazine/utils"
	"github.com/canbefree/magazine/vars"

	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
)

var (
	ErrEngine = errors.New("err engine")
	ErrorFile = errors.New("file or path not exist")
)

// migrateCommand
type migrateCommand struct {
	//
	output string
	// 执行
	excute bool
	//
	engine string
	// define you flags here
	*baseCommand
}

// assert
var _ cmder = (*migrateCommand)(nil)

// NewmigrateCommand
func NewMigrateCommand() *migrateCommand {
	cc := &migrateCommand{}
	cc.baseCommand = &baseCommand{
		cmd: &cobra.Command{
			Use: "migrate",
			RunE: func(cmd *cobra.Command, args []string) error {
				return cmd.Help()
			},
		},
	}

	cc.baseCommand.cmd.AddCommand(&cobra.Command{
		Use: "dbinit",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cc.DBInit()
		},
	})

	cc.cmd.Flags().BoolVarP(&cc.excute, "excute", "s", false, "excute")
	cc.cmd.Flags().StringVarP(&cc.output, "output", "o", "./", "output path")
	cc.cmd.Flags().StringVarP(&cc.engine, "engine", "e", "mariadb", "default mairadb")

	return cc
}

func (c *migrateCommand) DBInit() error {
	var sqls []string
	switch c.engine {
	case "mariadb":
		sqls = mariadb.GetSqlsFiles()
		jww.TRACE.Printf("%#v", sqls)
	default:
		return ErrEngine
	}
	//
	if c.excute {
		switch c.engine {
		case "mariadb":
			for _, v := range sqls {
				vars.DB.Exec(v)
			}
		default:
			return ErrEngine
		}
	} else {
		if !utils.FileExists(c.output) {
			return fmt.Errorf("%v:%v", ErrorFile, c.output)
		}
	}
	return nil

}

// Send
func (c *migrateCommand) Send() error {
	return nil
}

// Up 执行数据库
func (c *migrateCommand) Up() error {
	//
	return nil
}

func (m *migrateCommand) getCommand() *cobra.Command {
	return m.cmd
}

func init() {
	AddCommands(NewMigrateCommand())
}

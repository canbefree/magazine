package cmd

import (
	"github.com/canbefree/magazine/utils"
	"github.com/canbefree/magazine/vars"
	"github.com/spf13/cobra"
)

// config 配置自动生成
type configCommand struct {
	config string // config path
	*baseCommand
}

func NewConfigCommand() *configCommand {
	c := &configCommand{}
	c.baseCommand = &baseCommand{}
	c.baseCommand.cmd = &cobra.Command{
		Use: "config example",
		Run: func(cmd *cobra.Command, args []string) {
			c._OutExample(cmd)
		},
	}
	c.baseCommand.cmd.Flags().StringVarP(&c.config, "config", "c", "./.config", "default config path")
	return c
}

func (m *configCommand) _OutExample(cmd *cobra.Command) {
	vars.ConfigPath = m.config

	// TODO  代码应该保留一份默认配置
	for k, v := range vars.ConfigMap {
		err := vars.ConfigManage.Loader.SaveConfig(k, v)
		utils.PanicIfErr(err)
	}
}

func (m *configCommand) getCommand() *cobra.Command {
	return m.cmd
}

func init() {
	AddCommands(NewConfigCommand())
}

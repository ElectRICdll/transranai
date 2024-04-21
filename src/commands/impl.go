package commands

import (
	"translator/config"
	"translator/logic"
)

type Command interface {
	Execute(...string) (string, error)

	Help() string
}

type SetCmd struct {
	toSet *config.Config
	value string
}

func (c *SetCmd) Execute(args ...string) (string, error) {
	return "", nil
}

func (c *SetCmd) Help() string {
	return config.Msg("set")
}

type CheckoutCmd struct {
	toCheckout *config.Config
}

func (c *CheckoutCmd) Execute(args ...string) (string, error) {
	return "", nil
}

func (c *CheckoutCmd) Help() string {
	return config.Msg("checkout")
}

type HelpCmd struct {
	cmds []Command
}

func (c *HelpCmd) Execute(args ...string) (string, error) {
	return "", nil
}

func (c *HelpCmd) Help() string {
	return config.Msg("help")
}

type TranslateCmd struct {
	*logic.Prop
}

func (c *TranslateCmd) Execute(args ...string) (string, error) {
	return "", nil
}

func (c *TranslateCmd) Help() string {
	return config.Msg("translate")
}
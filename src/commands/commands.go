package commands

const (
	SET      = "set"
	CHECKOUT = "checkout"
	HELP     = "help"
)

var CommandSet = map[string]Command{
	SET:      setCommand(),
	CHECKOUT: checkoutCommand(),
	HELP:     helpCommand(),
}

func CommandSelector(cmdArgs ...string) Command {
	command := cmdArgs[0]

	if CommandSet[command] == nil {
		return translateCommand()
	}

	return CommandSet[command]
}

func setCommand() *SetCmd {
	return &SetCmd{}
}

func checkoutCommand() *CheckoutCmd {
	return &CheckoutCmd{}
}

func helpCommand() *HelpCmd {
	return &HelpCmd{}
}

func translateCommand() *TranslateCmd {
	return &TranslateCmd{}
}

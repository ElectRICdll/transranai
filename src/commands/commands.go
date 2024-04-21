package commands

const (
	SET      = "set"
	CHECKOUT = "checkout"
	HELP     = "help"
)

var CommandSet = map[string]Command{
	SET:      Set(),
	CHECKOUT: Checkout(),
	HELP:     Help(),
}

func Set() *SetCmd {
	return &SetCmd{}
}

func Checkout() *CheckoutCmd {
	return &CheckoutCmd{}
}

func Help() *HelpCmd {
	return &HelpCmd{}
}

func Translate() *TranslateCmd {
	return &TranslateCmd{}
}

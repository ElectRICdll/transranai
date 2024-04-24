package logic

var (
	CurrentTranslator Translator
	manager           = map[string]Constructor{}
)
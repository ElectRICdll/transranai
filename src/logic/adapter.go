package logic

import "fmt"

// Constructor should describe the function to create a Translator.
type Constructor func() (Translator, error)

// Translator abstracts translate APIs.
type Translator interface {
	Translate(string, Prop) (string, error)

	LangSupported() []Lang
}

// Reg registers a new Constructor with a alias.
func Reg(alias string, c Constructor) {
	if manager[alias] != nil {
		panic(fmt.Sprintf("%s duplicate registered! check your adapters", alias))
	}
	manager[alias] = c
}

func TranslatorFactory(name string) (Translator, error) {
	return manager[name]()
}
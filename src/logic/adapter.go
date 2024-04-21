package logic

type Adapter interface {
	Translate(string, Prop) (string, error)
}
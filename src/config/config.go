package config

type Config struct {
	Lang          string `yaml:"lang"`
	MaxTextLength int    `yaml:"max_text_length"`
}

func init() {

}
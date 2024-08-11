package models

type CohereModel struct{}

func (c CohereModel) Command() string {
	return "command"
}

func (c CohereModel) CommandNightly() string {
	return "command-nightly"
}

func (c CohereModel) CommandLight() string {
	return "command-light"
}

func (c CohereModel) CommandLightNightly() string {
	return "command-light-nightly"
}

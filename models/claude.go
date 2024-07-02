package models

type CluadeModel struct{}

func (c CluadeModel) Claude3Opus() string {
	return "claude-3-opus-20240229"
}

func (c CluadeModel) Claude3Sonnet() string {
	return "claude-3-5-sonnet-20240620"
}

func (c CluadeModel) Claude3Haiku() string {
	return "claude-3-haiku-20240307"
}

func (c CluadeModel) Claude2Legacy() string {
	return "claude-2.1"
}

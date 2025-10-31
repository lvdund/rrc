package ies

// MCCH-MessageType ::= CHOICE
const (
	McchMessagetypeChoiceNothing = iota
	McchMessagetypeChoiceC1
	McchMessagetypeChoiceLater
)

type McchMessagetype struct {
	Choice uint64
	C1     *McchMessagetypeC1
	Later  *McchMessagetypeLater
}

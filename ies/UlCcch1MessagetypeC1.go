package ies

// UL-CCCH1-MessageType-c1 ::= CHOICE
const (
	UlCcch1MessagetypeC1ChoiceNothing = iota
	UlCcch1MessagetypeC1ChoiceRrcresumerequest1
	UlCcch1MessagetypeC1ChoiceSpare3
	UlCcch1MessagetypeC1ChoiceSpare2
	UlCcch1MessagetypeC1ChoiceSpare1
)

type UlCcch1MessagetypeC1 struct {
	Choice            uint64
	Rrcresumerequest1 *Rrcresumerequest1
	Spare3            *struct{}
	Spare2            *struct{}
	Spare1            *struct{}
}

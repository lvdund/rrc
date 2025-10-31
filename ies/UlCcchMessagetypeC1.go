package ies

// UL-CCCH-MessageType-c1 ::= CHOICE
const (
	UlCcchMessagetypeC1ChoiceNothing = iota
	UlCcchMessagetypeC1ChoiceRrcsetuprequest
	UlCcchMessagetypeC1ChoiceRrcresumerequest
	UlCcchMessagetypeC1ChoiceRrcreestablishmentrequest
	UlCcchMessagetypeC1ChoiceRrcsysteminforequest
)

type UlCcchMessagetypeC1 struct {
	Choice                    uint64
	Rrcsetuprequest           *Rrcsetuprequest
	Rrcresumerequest          *Rrcresumerequest
	Rrcreestablishmentrequest *Rrcreestablishmentrequest
	Rrcsysteminforequest      *Rrcsysteminforequest
}

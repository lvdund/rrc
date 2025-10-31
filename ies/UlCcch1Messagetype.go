package ies

// UL-CCCH1-MessageType ::= CHOICE
const (
	UlCcch1MessagetypeChoiceNothing = iota
	UlCcch1MessagetypeChoiceC1
	UlCcch1MessagetypeChoiceMessageclassextension
)

type UlCcch1Messagetype struct {
	Choice                uint64
	C1                    *UlCcch1MessagetypeC1
	Messageclassextension *UlCcch1MessagetypeMessageclassextension
}

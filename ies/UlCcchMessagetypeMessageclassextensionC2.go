package ies

// UL-CCCH-MessageType-messageClassExtension-c2 ::= CHOICE
const (
	UlCcchMessagetypeMessageclassextensionC2ChoiceNothing = iota
	UlCcchMessagetypeMessageclassextensionC2ChoiceRrcconnectionresumerequestR13
)

type UlCcchMessagetypeMessageclassextensionC2 struct {
	Choice                        uint64
	RrcconnectionresumerequestR13 *RrcconnectionresumerequestR13
}

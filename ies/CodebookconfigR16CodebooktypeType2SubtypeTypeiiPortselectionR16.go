package ies

import "rrc/utils"

// CodebookConfig-r16-codebookType-type2-subType-typeII-PortSelection-r16 ::= SEQUENCE
type CodebookconfigR16CodebooktypeType2SubtypeTypeiiPortselectionR16 struct {
	PortselectionsamplingsizeR16        CodebookconfigR16CodebooktypeType2SubtypeTypeiiPortselectionR16PortselectionsamplingsizeR16
	TypeiiPortselectionriRestrictionR16 utils.BITSTRING `lb:4,ub:4`
}

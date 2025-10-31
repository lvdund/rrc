package ies

import "rrc/utils"

// CodebookConfig-codebookType-type2-subType-typeII-PortSelection ::= SEQUENCE
type CodebookconfigCodebooktypeType2SubtypeTypeiiPortselection struct {
	Portselectionsamplingsize        *CodebookconfigCodebooktypeType2SubtypeTypeiiPortselectionPortselectionsamplingsize
	TypeiiPortselectionriRestriction utils.BITSTRING `lb:2,ub:2`
}

package ies

import "rrc/utils"

// T-ReassemblyExt-r17 ::= ENUMERATED
type TReassemblyextR17 struct {
	Value utils.ENUMERATED
}

const (
	TReassemblyextR17EnumeratedNothing = iota
	TReassemblyextR17EnumeratedMs210
	TReassemblyextR17EnumeratedMs220
	TReassemblyextR17EnumeratedMs340
	TReassemblyextR17EnumeratedMs350
	TReassemblyextR17EnumeratedMs550
	TReassemblyextR17EnumeratedMs1100
	TReassemblyextR17EnumeratedMs1650
	TReassemblyextR17EnumeratedMs2200
)

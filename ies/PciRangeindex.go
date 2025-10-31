package ies

import "rrc/utils"

// PCI-RangeIndex ::= utils.INTEGER (1..maxNrofPCI-Ranges)
type PciRangeindex struct {
	Value utils.INTEGER `lb:0,ub:maxNrofPCIRanges`
}

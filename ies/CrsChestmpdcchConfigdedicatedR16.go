package ies

import "rrc/utils"

// CRS-ChEstMPDCCH-ConfigDedicated-r16 ::= SEQUENCE
type CrsChestmpdcchConfigdedicatedR16 struct {
	PowerratioR16           *utils.ENUMERATED
	LocalizedmappingtypeR16 utils.ENUMERATED
}

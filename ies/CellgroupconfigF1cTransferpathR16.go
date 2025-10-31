package ies

import "rrc/utils"

// CellGroupConfig-f1c-TransferPath-r16 ::= ENUMERATED
type CellgroupconfigF1cTransferpathR16 struct {
	Value utils.ENUMERATED
}

const (
	CellgroupconfigF1cTransferpathR16EnumeratedNothing = iota
	CellgroupconfigF1cTransferpathR16EnumeratedLte
	CellgroupconfigF1cTransferpathR16EnumeratedNr
	CellgroupconfigF1cTransferpathR16EnumeratedBoth
)

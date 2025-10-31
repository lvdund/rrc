package ies

import "rrc/utils"

// CellGroupConfig-f1c-TransferPathNRDC-r17 ::= ENUMERATED
type CellgroupconfigF1cTransferpathnrdcR17 struct {
	Value utils.ENUMERATED
}

const (
	CellgroupconfigF1cTransferpathnrdcR17EnumeratedNothing = iota
	CellgroupconfigF1cTransferpathnrdcR17EnumeratedMcg
	CellgroupconfigF1cTransferpathnrdcR17EnumeratedScg
	CellgroupconfigF1cTransferpathnrdcR17EnumeratedBoth
)

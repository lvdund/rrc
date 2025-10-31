package ies

import "rrc/utils"

// IntraFreqCAG-CellListPerPLMN-r16 ::= SEQUENCE
type IntrafreqcagCelllistperplmnR16 struct {
	PlmnIdentityindexR16 utils.INTEGER `lb:0,ub:maxPLMN`
	CagCelllistR16       []PciRange    `lb:1,ub:maxCAGCellR16`
}

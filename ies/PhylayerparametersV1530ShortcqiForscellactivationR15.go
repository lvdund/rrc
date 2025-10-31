package ies

import "rrc/utils"

// PhyLayerParameters-v1530-shortCQI-ForSCellActivation-r15 ::= ENUMERATED
type PhylayerparametersV1530ShortcqiForscellactivationR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530ShortcqiForscellactivationR15EnumeratedNothing = iota
	PhylayerparametersV1530ShortcqiForscellactivationR15EnumeratedSupported
)

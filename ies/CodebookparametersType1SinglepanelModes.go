package ies

import "rrc/utils"

// CodebookParameters-type1-singlePanel-modes ::= ENUMERATED
type CodebookparametersType1SinglepanelModes struct {
	Value utils.ENUMERATED
}

const (
	CodebookparametersType1SinglepanelModesEnumeratedNothing = iota
	CodebookparametersType1SinglepanelModesEnumeratedMode1
	CodebookparametersType1SinglepanelModesEnumeratedMode1andmode2
)

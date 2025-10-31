package ies

import "rrc/utils"

// CodebookParameters-type1-multiPanel-modes ::= ENUMERATED
type CodebookparametersType1MultipanelModes struct {
	Value utils.ENUMERATED
}

const (
	CodebookparametersType1MultipanelModesEnumeratedNothing = iota
	CodebookparametersType1MultipanelModesEnumeratedMode1
	CodebookparametersType1MultipanelModesEnumeratedMode2
	CodebookparametersType1MultipanelModesEnumeratedBoth
)

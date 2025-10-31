package ies

import "rrc/utils"

// CodebookParameters-type1-multiPanel-nrofPanels ::= ENUMERATED
type CodebookparametersType1MultipanelNrofpanels struct {
	Value utils.ENUMERATED
}

const (
	CodebookparametersType1MultipanelNrofpanelsEnumeratedNothing = iota
	CodebookparametersType1MultipanelNrofpanelsEnumeratedN2
	CodebookparametersType1MultipanelNrofpanelsEnumeratedN4
)

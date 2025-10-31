package ies

import "rrc/utils"

// Phy-ParametersCommon-maxNumberSearchSpaces ::= ENUMERATED
type PhyParameterscommonMaxnumbersearchspaces struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonMaxnumbersearchspacesEnumeratedNothing = iota
	PhyParameterscommonMaxnumbersearchspacesEnumeratedN10
)

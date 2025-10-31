package ies

import "rrc/utils"

// PhyLayerParameters-NB-v1440-interferenceRandomisation-r14 ::= ENUMERATED
type PhylayerparametersNbV1440InterferencerandomisationR14 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersNbV1440InterferencerandomisationR14EnumeratedNothing = iota
	PhylayerparametersNbV1440InterferencerandomisationR14EnumeratedSupported
)

package ies

import "rrc/utils"

// PhyLayerParameters-NB-v1430-twoHARQ-Processes-r14 ::= ENUMERATED
type PhylayerparametersNbV1430TwoharqProcessesR14 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersNbV1430TwoharqProcessesR14EnumeratedNothing = iota
	PhylayerparametersNbV1430TwoharqProcessesR14EnumeratedSupported
)

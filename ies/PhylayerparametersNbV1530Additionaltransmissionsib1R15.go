package ies

import "rrc/utils"

// PhyLayerParameters-NB-v1530-additionalTransmissionSIB1-r15 ::= ENUMERATED
type PhylayerparametersNbV1530Additionaltransmissionsib1R15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersNbV1530Additionaltransmissionsib1R15EnumeratedNothing = iota
	PhylayerparametersNbV1530Additionaltransmissionsib1R15EnumeratedSupported
)

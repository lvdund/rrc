package ies

import "rrc/utils"

// PhyLayerParameters-NB-v1530-sr-WithHARQ-ACK-r15 ::= ENUMERATED
type PhylayerparametersNbV1530SrWithharqAckR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersNbV1530SrWithharqAckR15EnumeratedNothing = iota
	PhylayerparametersNbV1530SrWithharqAckR15EnumeratedSupported
)

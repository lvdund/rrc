package ies

import "rrc/utils"

// PhyLayerParameters-NB-v1530-sr-WithoutHARQ-ACK-r15 ::= ENUMERATED
type PhylayerparametersNbV1530SrWithoutharqAckR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersNbV1530SrWithoutharqAckR15EnumeratedNothing = iota
	PhylayerparametersNbV1530SrWithoutharqAckR15EnumeratedSupported
)

package ies

import "rrc/utils"

// PhyLayerParameters-NB-v1530-nprach-Format2-r15 ::= ENUMERATED
type PhylayerparametersNbV1530NprachFormat2R15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersNbV1530NprachFormat2R15EnumeratedNothing = iota
	PhylayerparametersNbV1530NprachFormat2R15EnumeratedSupported
)

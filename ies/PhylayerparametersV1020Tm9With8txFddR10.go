package ies

import "rrc/utils"

// PhyLayerParameters-v1020-tm9-With-8Tx-FDD-r10 ::= ENUMERATED
type PhylayerparametersV1020Tm9With8txFddR10 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1020Tm9With8txFddR10EnumeratedNothing = iota
	PhylayerparametersV1020Tm9With8txFddR10EnumeratedSupported
)

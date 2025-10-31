package ies

import "rrc/utils"

// PhyLayerParameters-NB-v1610-multiTB-HARQ-AckBundling-r16 ::= ENUMERATED
type PhylayerparametersNbV1610MultitbHarqAckbundlingR16 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersNbV1610MultitbHarqAckbundlingR16EnumeratedNothing = iota
	PhylayerparametersNbV1610MultitbHarqAckbundlingR16EnumeratedSupported
)

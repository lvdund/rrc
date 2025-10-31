package ies

import "rrc/utils"

// PhyLayerParameters-NB-v1610-slotSymbolResourceResvDL-r16 ::= ENUMERATED
type PhylayerparametersNbV1610SlotsymbolresourceresvdlR16 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersNbV1610SlotsymbolresourceresvdlR16EnumeratedNothing = iota
	PhylayerparametersNbV1610SlotsymbolresourceresvdlR16EnumeratedSupported
)

package ies

import "rrc/utils"

// PhyLayerParameters-NB-v1610-slotSymbolResourceResvUL-r16 ::= ENUMERATED
type PhylayerparametersNbV1610SlotsymbolresourceresvulR16 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersNbV1610SlotsymbolresourceresvulR16EnumeratedNothing = iota
	PhylayerparametersNbV1610SlotsymbolresourceresvulR16EnumeratedSupported
)

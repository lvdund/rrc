package ies

import "rrc/utils"

// PhyLayerParameters-v1130-txDiv-PUCCH1b-ChSelect-r11 ::= ENUMERATED
type PhylayerparametersV1130TxdivPucch1bChselectR11 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1130TxdivPucch1bChselectR11EnumeratedNothing = iota
	PhylayerparametersV1130TxdivPucch1bChselectR11EnumeratedSupported
)

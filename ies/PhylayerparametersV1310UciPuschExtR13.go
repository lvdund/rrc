package ies

import "rrc/utils"

// PhyLayerParameters-v1310-uci-PUSCH-Ext-r13 ::= ENUMERATED
type PhylayerparametersV1310UciPuschExtR13 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1310UciPuschExtR13EnumeratedNothing = iota
	PhylayerparametersV1310UciPuschExtR13EnumeratedSupported
)

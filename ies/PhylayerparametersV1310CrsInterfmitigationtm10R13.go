package ies

import "rrc/utils"

// PhyLayerParameters-v1310-crs-InterfMitigationTM10-r13 ::= ENUMERATED
type PhylayerparametersV1310CrsInterfmitigationtm10R13 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1310CrsInterfmitigationtm10R13EnumeratedNothing = iota
	PhylayerparametersV1310CrsInterfmitigationtm10R13EnumeratedSupported
)

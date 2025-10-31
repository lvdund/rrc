package ies

import "rrc/utils"

// PhyLayerParameters-v1330-cch-InterfMitigation-RefRecTypeB-r13 ::= ENUMERATED
type PhylayerparametersV1330CchInterfmitigationRefrectypebR13 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1330CchInterfmitigationRefrectypebR13EnumeratedNothing = iota
	PhylayerparametersV1330CchInterfmitigationRefrectypebR13EnumeratedSupported
)

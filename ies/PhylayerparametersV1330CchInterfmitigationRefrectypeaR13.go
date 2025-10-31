package ies

import "rrc/utils"

// PhyLayerParameters-v1330-cch-InterfMitigation-RefRecTypeA-r13 ::= ENUMERATED
type PhylayerparametersV1330CchInterfmitigationRefrectypeaR13 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1330CchInterfmitigationRefrectypeaR13EnumeratedNothing = iota
	PhylayerparametersV1330CchInterfmitigationRefrectypeaR13EnumeratedSupported
)

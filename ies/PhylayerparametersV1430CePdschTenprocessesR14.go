package ies

import "rrc/utils"

// PhyLayerParameters-v1430-ce-PDSCH-TenProcesses-r14 ::= ENUMERATED
type PhylayerparametersV1430CePdschTenprocessesR14 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1430CePdschTenprocessesR14EnumeratedNothing = iota
	PhylayerparametersV1430CePdschTenprocessesR14EnumeratedSupported
)

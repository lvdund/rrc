package ies

import "rrc/utils"

// CA-ParametersNRDC-v1700-beamManagementType-CBM-r17 ::= ENUMERATED
type CaParametersnrdcV1700BeammanagementtypeCbmR17 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrdcV1700BeammanagementtypeCbmR17EnumeratedNothing = iota
	CaParametersnrdcV1700BeammanagementtypeCbmR17EnumeratedSupported
)

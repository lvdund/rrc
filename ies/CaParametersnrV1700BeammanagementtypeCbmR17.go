package ies

import "rrc/utils"

// CA-ParametersNR-v1700-beamManagementType-CBM-r17 ::= ENUMERATED
type CaParametersnrV1700BeammanagementtypeCbmR17 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1700BeammanagementtypeCbmR17EnumeratedNothing = iota
	CaParametersnrV1700BeammanagementtypeCbmR17EnumeratedSupported
)

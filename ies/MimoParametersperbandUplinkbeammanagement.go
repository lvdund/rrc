package ies

import "rrc/utils"

// MIMO-ParametersPerBand-uplinkBeamManagement ::= SEQUENCE
type MimoParametersperbandUplinkbeammanagement struct {
	MaxnumbersrsResourcepersetBm MimoParametersperbandUplinkbeammanagementMaxnumbersrsResourcepersetBm
	MaxnumbersrsResourceset      utils.INTEGER `lb:0,ub:8`
}

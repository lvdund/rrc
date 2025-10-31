package ies

import "rrc/utils"

// CA-ParametersNR-v1540-csi-RS-IM-ReceptionForFeedbackPerBandComb ::= SEQUENCE
type CaParametersnrV1540CsiRsImReceptionforfeedbackperbandcomb struct {
	MaxnumbersimultaneousnzpCsiRsActbwpAllcc        *utils.INTEGER `lb:0,ub:64`
	TotalnumberportssimultaneousnzpCsiRsActbwpAllcc *utils.INTEGER `lb:0,ub:256`
}

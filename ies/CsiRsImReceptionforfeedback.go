package ies

import "rrc/utils"

// CSI-RS-IM-ReceptionForFeedback ::= SEQUENCE
type CsiRsImReceptionforfeedback struct {
	MaxconfignumbernzpCsiRsPercc              utils.INTEGER `lb:0,ub:64`
	MaxconfignumberportsacrossnzpCsiRsPercc   utils.INTEGER `lb:0,ub:256`
	MaxconfignumbercsiImPercc                 CsiRsImReceptionforfeedbackMaxconfignumbercsiImPercc
	MaxnumbersimultaneousnzpCsiRsPercc        utils.INTEGER `lb:0,ub:64`
	TotalnumberportssimultaneousnzpCsiRsPercc utils.INTEGER `lb:0,ub:256`
}

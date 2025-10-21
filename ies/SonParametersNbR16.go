package ies

import "rrc/utils"

// SON-Parameters-NB-r16 ::= SEQUENCE
type SonParametersNbR16 struct {
	AnrReportR16  *utils.ENUMERATED
	RachReportR16 *utils.ENUMERATED
}

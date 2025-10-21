package ies

import "rrc/utils"

// SON-Parameters-r9 ::= SEQUENCE
type SonParametersR9 struct {
	RachReportR9 *utils.ENUMERATED
}

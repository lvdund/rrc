package ies

import "rrc/utils"

// VarConnEstFailReport-r11 ::= SEQUENCE
type VarconnestfailreportR11 struct {
	ConnestfailreportR11 ConnestfailreportR11
	PlmnIdentityR11      PlmnIdentity
}

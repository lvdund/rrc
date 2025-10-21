package ies

import "rrc/utils"

// PDCP-Parameters-v1610 ::= SEQUENCE
type PdcpParametersV1610 struct {
	PdcpVersionchangewithouthoR16 *utils.ENUMERATED
	EhcR16                        *utils.ENUMERATED
	ContinueehcContextR16         *utils.ENUMERATED
	MaxnumberehcContextsR16       *utils.ENUMERATED
	JointehcRohcConfigR16         *utils.ENUMERATED
}

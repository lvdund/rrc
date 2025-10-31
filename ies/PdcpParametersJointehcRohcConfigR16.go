package ies

import "rrc/utils"

// PDCP-Parameters-jointEHC-ROHC-Config-r16 ::= ENUMERATED
type PdcpParametersJointehcRohcConfigR16 struct {
	Value utils.ENUMERATED
}

const (
	PdcpParametersJointehcRohcConfigR16EnumeratedNothing = iota
	PdcpParametersJointehcRohcConfigR16EnumeratedSupported
)

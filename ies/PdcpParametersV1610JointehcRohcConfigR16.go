package ies

import "rrc/utils"

// PDCP-Parameters-v1610-jointEHC-ROHC-Config-r16 ::= ENUMERATED
type PdcpParametersV1610JointehcRohcConfigR16 struct {
	Value utils.ENUMERATED
}

const (
	PdcpParametersV1610JointehcRohcConfigR16EnumeratedNothing = iota
	PdcpParametersV1610JointehcRohcConfigR16EnumeratedSupported
)

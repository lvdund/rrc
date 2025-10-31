package ies

import "rrc/utils"

// SL-Parameters-r12-discSupportedProc-r12 ::= ENUMERATED
type SlParametersR12DiscsupportedprocR12 struct {
	Value utils.ENUMERATED
}

const (
	SlParametersR12DiscsupportedprocR12EnumeratedNothing = iota
	SlParametersR12DiscsupportedprocR12EnumeratedN50
	SlParametersR12DiscsupportedprocR12EnumeratedN400
)

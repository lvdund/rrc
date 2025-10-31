package ies

import "rrc/utils"

// SL-Parameters-v1310-discSysInfoReporting-r13 ::= ENUMERATED
type SlParametersV1310DiscsysinforeportingR13 struct {
	Value utils.ENUMERATED
}

const (
	SlParametersV1310DiscsysinforeportingR13EnumeratedNothing = iota
	SlParametersV1310DiscsysinforeportingR13EnumeratedSupported
)

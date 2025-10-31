package ies

import "rrc/utils"

// MAC-Parameters-v1610-earlyData-UP-5GC-r16 ::= ENUMERATED
type MacParametersV1610EarlydataUp5gcR16 struct {
	Value utils.ENUMERATED
}

const (
	MacParametersV1610EarlydataUp5gcR16EnumeratedNothing = iota
	MacParametersV1610EarlydataUp5gcR16EnumeratedSupported
)

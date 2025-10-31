package ies

import "rrc/utils"

// PTRS-DownlinkConfig-maxNrofPorts-r16 ::= ENUMERATED
type PtrsDownlinkconfigMaxnrofportsR16 struct {
	Value utils.ENUMERATED
}

const (
	PtrsDownlinkconfigMaxnrofportsR16EnumeratedNothing = iota
	PtrsDownlinkconfigMaxnrofportsR16EnumeratedN1
	PtrsDownlinkconfigMaxnrofportsR16EnumeratedN2
)

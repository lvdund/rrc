package ies

import "rrc/utils"

// LAA-Parameters-r13-downlinkLAA-r13 ::= ENUMERATED
type LaaParametersR13DownlinklaaR13 struct {
	Value utils.ENUMERATED
}

const (
	LaaParametersR13DownlinklaaR13EnumeratedNothing = iota
	LaaParametersR13DownlinklaaR13EnumeratedSupported
)

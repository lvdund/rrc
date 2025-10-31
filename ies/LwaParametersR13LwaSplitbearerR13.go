package ies

import "rrc/utils"

// LWA-Parameters-r13-lwa-SplitBearer-r13 ::= ENUMERATED
type LwaParametersR13LwaSplitbearerR13 struct {
	Value utils.ENUMERATED
}

const (
	LwaParametersR13LwaSplitbearerR13EnumeratedNothing = iota
	LwaParametersR13LwaSplitbearerR13EnumeratedSupported
)

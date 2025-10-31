package ies

import "rrc/utils"

// SL-Parameters-r12-discScheduledResourceAlloc-r12 ::= ENUMERATED
type SlParametersR12DiscscheduledresourceallocR12 struct {
	Value utils.ENUMERATED
}

const (
	SlParametersR12DiscscheduledresourceallocR12EnumeratedNothing = iota
	SlParametersR12DiscscheduledresourceallocR12EnumeratedSupported
)

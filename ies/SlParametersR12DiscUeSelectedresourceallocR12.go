package ies

import "rrc/utils"

// SL-Parameters-r12-disc-UE-SelectedResourceAlloc-r12 ::= ENUMERATED
type SlParametersR12DiscUeSelectedresourceallocR12 struct {
	Value utils.ENUMERATED
}

const (
	SlParametersR12DiscUeSelectedresourceallocR12EnumeratedNothing = iota
	SlParametersR12DiscUeSelectedresourceallocR12EnumeratedSupported
)

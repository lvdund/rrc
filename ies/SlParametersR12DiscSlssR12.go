package ies

import "rrc/utils"

// SL-Parameters-r12-disc-SLSS-r12 ::= ENUMERATED
type SlParametersR12DiscSlssR12 struct {
	Value utils.ENUMERATED
}

const (
	SlParametersR12DiscSlssR12EnumeratedNothing = iota
	SlParametersR12DiscSlssR12EnumeratedSupported
)

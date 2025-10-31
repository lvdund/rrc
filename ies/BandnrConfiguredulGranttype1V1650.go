package ies

import "rrc/utils"

// BandNR-configuredUL-GrantType1-v1650 ::= ENUMERATED
type BandnrConfiguredulGranttype1V1650 struct {
	Value utils.ENUMERATED
}

const (
	BandnrConfiguredulGranttype1V1650EnumeratedNothing = iota
	BandnrConfiguredulGranttype1V1650EnumeratedSupported
)

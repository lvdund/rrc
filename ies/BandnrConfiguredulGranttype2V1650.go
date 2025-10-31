package ies

import "rrc/utils"

// BandNR-configuredUL-GrantType2-v1650 ::= ENUMERATED
type BandnrConfiguredulGranttype2V1650 struct {
	Value utils.ENUMERATED
}

const (
	BandnrConfiguredulGranttype2V1650EnumeratedNothing = iota
	BandnrConfiguredulGranttype2V1650EnumeratedSupported
)

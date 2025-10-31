package ies

import "rrc/utils"

// BandSidelink-r16-rx-IUC-Scheme1-SCI-r17 ::= ENUMERATED
type BandsidelinkR16RxIucScheme1SciR17 struct {
	Value utils.ENUMERATED
}

const (
	BandsidelinkR16RxIucScheme1SciR17EnumeratedNothing = iota
	BandsidelinkR16RxIucScheme1SciR17EnumeratedSupported
)

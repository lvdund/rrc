package ies

import "rrc/utils"

// BandSidelink-r16-rx-IUC-Scheme1-PreferredMode2Sidelink-r17 ::= ENUMERATED
type BandsidelinkR16RxIucScheme1Preferredmode2sidelinkR17 struct {
	Value utils.ENUMERATED
}

const (
	BandsidelinkR16RxIucScheme1Preferredmode2sidelinkR17EnumeratedNothing = iota
	BandsidelinkR16RxIucScheme1Preferredmode2sidelinkR17EnumeratedSupported
)

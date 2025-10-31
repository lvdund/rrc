package ies

import "rrc/utils"

// BandSidelinkPC5-r16-rx-IUC-Scheme1-NonPreferredMode2Sidelink-r17 ::= ENUMERATED
type Bandsidelinkpc5R16RxIucScheme1Nonpreferredmode2sidelinkR17 struct {
	Value utils.ENUMERATED
}

const (
	Bandsidelinkpc5R16RxIucScheme1Nonpreferredmode2sidelinkR17EnumeratedNothing = iota
	Bandsidelinkpc5R16RxIucScheme1Nonpreferredmode2sidelinkR17EnumeratedSupported
)

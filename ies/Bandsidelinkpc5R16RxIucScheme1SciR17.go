package ies

import "rrc/utils"

// BandSidelinkPC5-r16-rx-IUC-Scheme1-SCI-r17 ::= ENUMERATED
type Bandsidelinkpc5R16RxIucScheme1SciR17 struct {
	Value utils.ENUMERATED
}

const (
	Bandsidelinkpc5R16RxIucScheme1SciR17EnumeratedNothing = iota
	Bandsidelinkpc5R16RxIucScheme1SciR17EnumeratedSupported
)

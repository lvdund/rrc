package ies

import "rrc/utils"

// BandSidelinkPC5-r16-rx-IUC-Scheme1-SCI-ExplicitReq-r17 ::= ENUMERATED
type Bandsidelinkpc5R16RxIucScheme1SciExplicitreqR17 struct {
	Value utils.ENUMERATED
}

const (
	Bandsidelinkpc5R16RxIucScheme1SciExplicitreqR17EnumeratedNothing = iota
	Bandsidelinkpc5R16RxIucScheme1SciExplicitreqR17EnumeratedSupported
)

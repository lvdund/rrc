package ies

import "rrc/utils"

// BandSidelink-r16-rx-IUC-Scheme1-SCI-ExplicitReq-r17 ::= ENUMERATED
type BandsidelinkR16RxIucScheme1SciExplicitreqR17 struct {
	Value utils.ENUMERATED
}

const (
	BandsidelinkR16RxIucScheme1SciExplicitreqR17EnumeratedNothing = iota
	BandsidelinkR16RxIucScheme1SciExplicitreqR17EnumeratedSupported
)

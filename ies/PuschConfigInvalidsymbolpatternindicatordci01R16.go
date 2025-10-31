package ies

import "rrc/utils"

// PUSCH-Config-invalidSymbolPatternIndicatorDCI-0-1-r16 ::= ENUMERATED
type PuschConfigInvalidsymbolpatternindicatordci01R16 struct {
	Value utils.ENUMERATED
}

const (
	PuschConfigInvalidsymbolpatternindicatordci01R16EnumeratedNothing = iota
	PuschConfigInvalidsymbolpatternindicatordci01R16EnumeratedEnabled
)

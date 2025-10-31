package ies

import "rrc/utils"

// PUSCH-Config-invalidSymbolPatternIndicatorDCI-0-2-r16 ::= ENUMERATED
type PuschConfigInvalidsymbolpatternindicatordci02R16 struct {
	Value utils.ENUMERATED
}

const (
	PuschConfigInvalidsymbolpatternindicatordci02R16EnumeratedNothing = iota
	PuschConfigInvalidsymbolpatternindicatordci02R16EnumeratedEnabled
)

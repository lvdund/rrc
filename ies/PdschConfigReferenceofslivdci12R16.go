package ies

import "rrc/utils"

// PDSCH-Config-referenceOfSLIVDCI-1-2-r16 ::= ENUMERATED
type PdschConfigReferenceofslivdci12R16 struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigReferenceofslivdci12R16EnumeratedNothing = iota
	PdschConfigReferenceofslivdci12R16EnumeratedEnabled
)

package ies

import "rrc/utils"

// PUSCH-Config-secondTPCFieldDCI-0-2-r17 ::= ENUMERATED
type PuschConfigSecondtpcfielddci02R17 struct {
	Value utils.ENUMERATED
}

const (
	PuschConfigSecondtpcfielddci02R17EnumeratedNothing = iota
	PuschConfigSecondtpcfielddci02R17EnumeratedEnabled
)

package ies

import "rrc/utils"

// PUSCH-Config-secondTPCFieldDCI-0-1-r17 ::= ENUMERATED
type PuschConfigSecondtpcfielddci01R17 struct {
	Value utils.ENUMERATED
}

const (
	PuschConfigSecondtpcfielddci01R17EnumeratedNothing = iota
	PuschConfigSecondtpcfielddci01R17EnumeratedEnabled
)

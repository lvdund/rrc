package ies

import "rrc/utils"

// PUCCH-Config-secondTPCFieldDCI-1-2-r17 ::= ENUMERATED
type PucchConfigSecondtpcfielddci12R17 struct {
	Value utils.ENUMERATED
}

const (
	PucchConfigSecondtpcfielddci12R17EnumeratedNothing = iota
	PucchConfigSecondtpcfielddci12R17EnumeratedEnabled
)

package ies

import "rrc/utils"

// PUCCH-Config-secondTPCFieldDCI-1-1-r17 ::= ENUMERATED
type PucchConfigSecondtpcfielddci11R17 struct {
	Value utils.ENUMERATED
}

const (
	PucchConfigSecondtpcfielddci11R17EnumeratedNothing = iota
	PucchConfigSecondtpcfielddci11R17EnumeratedEnabled
)

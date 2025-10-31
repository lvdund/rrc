package ies

import "rrc/utils"

// ServingCellConfig-timeDomainHARQ-BundlingType1-r17 ::= ENUMERATED
type ServingcellconfigTimedomainharqBundlingtype1R17 struct {
	Value utils.ENUMERATED
}

const (
	ServingcellconfigTimedomainharqBundlingtype1R17EnumeratedNothing = iota
	ServingcellconfigTimedomainharqBundlingtype1R17EnumeratedEnabled
)

package ies

import "rrc/utils"

// PDSCH-ConfigDedicated-v1530-ce-CQI-AlternativeTableConfig-r15 ::= ENUMERATED
type PdschConfigdedicatedV1530CeCqiAlternativetableconfigR15 struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigdedicatedV1530CeCqiAlternativetableconfigR15EnumeratedNothing = iota
	PdschConfigdedicatedV1530CeCqiAlternativetableconfigR15EnumeratedOn
)

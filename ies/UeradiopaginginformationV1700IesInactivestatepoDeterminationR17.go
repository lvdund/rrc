package ies

import "rrc/utils"

// UERadioPagingInformation-v1700-IEs-inactiveStatePO-Determination-r17 ::= ENUMERATED
type UeradiopaginginformationV1700IesInactivestatepoDeterminationR17 struct {
	Value utils.ENUMERATED
}

const (
	UeradiopaginginformationV1700IesInactivestatepoDeterminationR17EnumeratedNothing = iota
	UeradiopaginginformationV1700IesInactivestatepoDeterminationR17EnumeratedSupported
)

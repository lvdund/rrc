package ies

import "rrc/utils"

// PosSRS-RRC-Inactive-OutsideInitialUL-BWP-r17-srsPosWithoutRestrictionOnBWP-r17 ::= ENUMERATED
type PossrsRrcInactiveOutsideinitialulBwpR17SrsposwithoutrestrictiononbwpR17 struct {
	Value utils.ENUMERATED
}

const (
	PossrsRrcInactiveOutsideinitialulBwpR17SrsposwithoutrestrictiononbwpR17EnumeratedNothing = iota
	PossrsRrcInactiveOutsideinitialulBwpR17SrsposwithoutrestrictiononbwpR17EnumeratedSupported
)

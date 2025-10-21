package ies

import "rrc/utils"

// MIMO-CA-ParametersPerBoBCPerTM-v1430 ::= SEQUENCE
type MimoCaParametersperbobcpertmV1430 struct {
	CsiReportingnpR14       *utils.ENUMERATED
	CsiReportingadvancedR14 *utils.ENUMERATED
}

package ies

import "rrc/utils"

// E-CSFB-r9-mobilityCDMA2000-HRPD-r9 ::= ENUMERATED
type ECsfbR9Mobilitycdma2000HrpdR9 struct {
	Value utils.ENUMERATED
}

const (
	ECsfbR9Mobilitycdma2000HrpdR9EnumeratedNothing = iota
	ECsfbR9Mobilitycdma2000HrpdR9EnumeratedHandover
	ECsfbR9Mobilitycdma2000HrpdR9EnumeratedRedirection
)

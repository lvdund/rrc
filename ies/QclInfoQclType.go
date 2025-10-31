package ies

import "rrc/utils"

// QCL-Info-qcl-Type ::= ENUMERATED
type QclInfoQclType struct {
	Value utils.ENUMERATED
}

const (
	QclInfoQclTypeEnumeratedNothing = iota
	QclInfoQclTypeEnumeratedTypea
	QclInfoQclTypeEnumeratedTypeb
	QclInfoQclTypeEnumeratedTypec
	QclInfoQclTypeEnumeratedTyped
)

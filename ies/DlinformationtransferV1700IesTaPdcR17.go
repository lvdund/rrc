package ies

import "rrc/utils"

// DLInformationTransfer-v1700-IEs-ta-PDC-r17 ::= ENUMERATED
type DlinformationtransferV1700IesTaPdcR17 struct {
	Value utils.ENUMERATED
}

const (
	DlinformationtransferV1700IesTaPdcR17EnumeratedNothing = iota
	DlinformationtransferV1700IesTaPdcR17EnumeratedActivate
	DlinformationtransferV1700IesTaPdcR17EnumeratedDeactivate
)

package ies

import "rrc/utils"

// CSI-RS-ConfigNZP-EMIMO-r13-setup-cdmType-r13 ::= ENUMERATED
type CsiRsConfignzpEmimoR13SetupCdmtypeR13 struct {
	Value utils.ENUMERATED
}

const (
	CsiRsConfignzpEmimoR13SetupCdmtypeR13EnumeratedNothing = iota
	CsiRsConfignzpEmimoR13SetupCdmtypeR13EnumeratedCdm2
	CsiRsConfignzpEmimoR13SetupCdmtypeR13EnumeratedCdm4
)

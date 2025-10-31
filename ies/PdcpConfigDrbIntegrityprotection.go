package ies

import "rrc/utils"

// PDCP-Config-drb-integrityProtection ::= ENUMERATED
type PdcpConfigDrbIntegrityprotection struct {
	Value utils.ENUMERATED
}

const (
	PdcpConfigDrbIntegrityprotectionEnumeratedNothing = iota
	PdcpConfigDrbIntegrityprotectionEnumeratedEnabled
)

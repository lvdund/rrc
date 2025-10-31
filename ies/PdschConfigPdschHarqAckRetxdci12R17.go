package ies

import "rrc/utils"

// PDSCH-Config-pdsch-HARQ-ACK-RetxDCI-1-2-r17 ::= ENUMERATED
type PdschConfigPdschHarqAckRetxdci12R17 struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigPdschHarqAckRetxdci12R17EnumeratedNothing = iota
	PdschConfigPdschHarqAckRetxdci12R17EnumeratedEnabled
)

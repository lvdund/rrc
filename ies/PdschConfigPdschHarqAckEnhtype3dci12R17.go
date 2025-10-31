package ies

import "rrc/utils"

// PDSCH-Config-pdsch-HARQ-ACK-EnhType3DCI-1-2-r17 ::= ENUMERATED
type PdschConfigPdschHarqAckEnhtype3dci12R17 struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigPdschHarqAckEnhtype3dci12R17EnumeratedNothing = iota
	PdschConfigPdschHarqAckEnhtype3dci12R17EnumeratedEnabled
)

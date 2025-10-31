package ies

import "rrc/utils"

// PDSCH-Config-pdsch-HARQ-ACK-EnhType3DCI-Field-1-2-r17 ::= ENUMERATED
type PdschConfigPdschHarqAckEnhtype3dciField12R17 struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigPdschHarqAckEnhtype3dciField12R17EnumeratedNothing = iota
	PdschConfigPdschHarqAckEnhtype3dciField12R17EnumeratedEnabled
)

package ies

import "rrc/utils"

// NRDC-Parameters-v1700-f1c-OverNR-RRC-r17 ::= ENUMERATED
type NrdcParametersV1700F1cOvernrRrcR17 struct {
	Value utils.ENUMERATED
}

const (
	NrdcParametersV1700F1cOvernrRrcR17EnumeratedNothing = iota
	NrdcParametersV1700F1cOvernrRrcR17EnumeratedSupported
)

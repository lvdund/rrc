package ies

// PDCCH-BlindDetectionMixed-r17 ::= SEQUENCE
type PdcchBlinddetectionmixedR17 struct {
	PdcchBlinddetectioncaMixedR17   *PdcchBlinddetectioncaMixedR17
	PdcchBlinddetectioncgUeMixedR17 *PdcchBlinddetectionmixedR17PdcchBlinddetectioncgUeMixedR17
}

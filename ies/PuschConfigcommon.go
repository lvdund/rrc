package ies

// PUSCH-ConfigCommon ::= SEQUENCE
type PuschConfigcommon struct {
	PuschConfigbasic        PuschConfigcommonPuschConfigbasic
	UlReferencesignalspusch UlReferencesignalspusch
}

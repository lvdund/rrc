package ies

// CFR-ConfigMCCH-MTCH-r17 ::= SEQUENCE
type CfrConfigmcchMtchR17 struct {
	LocationandbandwidthbroadcastR17 *LocationandbandwidthbroadcastR17
	PdschConfigmcchR17               *PdschConfigbroadcastR17
	CommoncontrolresourcesetextR17   *Controlresourceset
}

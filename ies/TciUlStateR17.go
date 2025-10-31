package ies

// TCI-UL-State-r17 ::= SEQUENCE
// Extensible
type TciUlStateR17 struct {
	TciUlStateIdR17          TciUlStateIdR17
	ServingcellidR17         *Servcellindex
	BwpIdR17                 *BwpId
	ReferencesignalR17       TciUlStateR17ReferencesignalR17
	AdditionalpciR17         *AdditionalpciindexR17
	UlPowercontrolR17        *UplinkPowercontrolidR17
	PathlossreferencersIdR17 *PathlossreferencersIdR17
}

package ies

// BWP-UplinkDedicated-ul-TCI-StateList-r17-explicitlist ::= SEQUENCE
type BwpUplinkdedicatedUlTciStatelistR17Explicitlist struct {
	UlTciToaddmodlistR17  *[]TciUlStateR17   `lb:1,ub:maxULTciR17`
	UlTciToreleaselistR17 *[]TciUlStateIdR17 `lb:1,ub:maxULTciR17`
}

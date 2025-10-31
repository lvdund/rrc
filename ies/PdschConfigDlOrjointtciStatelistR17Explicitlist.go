package ies

// PDSCH-Config-dl-OrJointTCI-StateList-r17-explicitlist ::= SEQUENCE
type PdschConfigDlOrjointtciStatelistR17Explicitlist struct {
	DlOrjointtciStatetoaddmodlistR17  *[]TciState   `lb:1,ub:maxNrofTCIStates`
	DlOrjointtciStatetoreleaselistR17 *[]TciStateid `lb:1,ub:maxNrofTCIStates`
}

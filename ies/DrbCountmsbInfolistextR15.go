package ies

// DRB-CountMSB-InfoListExt-r15 ::= SEQUENCE OF DRB-CountMSB-Info
// SIZE (1..maxDRBExt-r15)
type DrbCountmsbInfolistextR15 struct {
	Value []DrbCountmsbInfo `lb:1,ub:maxDRBExtR15`
}

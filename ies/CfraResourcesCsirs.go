package ies

// CFRA-resources-csirs ::= SEQUENCE
type CfraResourcesCsirs struct {
	CsirsResourcelist  []CfraCsirsResource `lb:1,ub:maxRACsirsResources`
	RsrpThresholdcsiRs RsrpRange
}

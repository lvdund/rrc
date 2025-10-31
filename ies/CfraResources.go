package ies

// CFRA-resources ::= CHOICE
const (
	CfraResourcesChoiceNothing = iota
	CfraResourcesChoiceSsb
	CfraResourcesChoiceCsirs
)

type CfraResources struct {
	Choice uint64
	Ssb    *CfraResourcesSsb
	Csirs  *CfraResourcesCsirs
}

package ies

// ZP-CSI-RS-Resource ::= SEQUENCE
// Extensible
type ZpCsiRsResource struct {
	ZpCsiRsResourceid    ZpCsiRsResourceid
	Resourcemapping      CsiRsResourcemapping
	Periodicityandoffset *CsiResourceperiodicityandoffset
}

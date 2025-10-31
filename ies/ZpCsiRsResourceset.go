package ies

// ZP-CSI-RS-ResourceSet ::= SEQUENCE
// Extensible
type ZpCsiRsResourceset struct {
	ZpCsiRsResourcesetid  ZpCsiRsResourcesetid
	ZpCsiRsResourceidlist []ZpCsiRsResourceid `lb:1,ub:maxNrofZPCsiRsResourcesperset`
}

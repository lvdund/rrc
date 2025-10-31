package ies

// SPS-Config-v1530 ::= SEQUENCE
type SpsConfigV1530 struct {
	SemipersistschedcRntiR15        *CRnti
	SpsConfigdlR15                  *SpsConfigdl
	SpsConfigulSttiToaddmodlistR15  *SpsConfigulSttiToaddmodlistR15
	SpsConfigulSttiToreleaselistR15 *SpsConfigulSttiToreleaselistR15
	SpsConfigulToaddmodlistR15      *SpsConfigulToaddmodlistR15
	SpsConfigulToreleaselistR15     *SpsConfigulToreleaselistR15
}

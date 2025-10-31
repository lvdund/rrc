package ies

// SPS-Config ::= SEQUENCE
type SpsConfig struct {
	SemipersistschedcRnti *CRnti
	SpsConfigdl           *SpsConfigdl
	SpsConfigul           *SpsConfigul
}

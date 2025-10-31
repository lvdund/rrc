package ies

// SI-RequestConfig ::= SEQUENCE
type SiRequestconfig struct {
	RachOccasionssi    *SiRequestconfigRachOccasionssi
	SiRequestperiod    *SiRequestconfigSiRequestperiod
	SiRequestresources []SiRequestresources `lb:1,ub:maxSIMessage`
}

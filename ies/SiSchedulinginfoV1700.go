package ies

// SI-SchedulingInfo-v1700 ::= SEQUENCE
type SiSchedulinginfoV1700 struct {
	Schedulinginfolist2R17   []Schedulinginfo2R17 `lb:1,ub:maxSIMessage`
	SiRequestconfigredcapR17 *SiRequestconfig
}

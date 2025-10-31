package ies

// PosSIB-ReqInfo-r16 ::= SEQUENCE
// Extensible
type PossibReqinfoR16 struct {
	GnssIdR16     *GnssIdR16
	SbasIdR16     *SbasIdR16
	PossibtypeR16 PossibReqinfoR16PossibtypeR16
}

package ies

// Q-OffsetRangeList ::= SEQUENCE
type QOffsetrangelist struct {
	Rsrpoffsetssb   QOffsetrange
	Rsrqoffsetssb   QOffsetrange
	Sinroffsetssb   QOffsetrange
	RsrpoffsetcsiRs QOffsetrange
	RsrqoffsetcsiRs QOffsetrange
	SinroffsetcsiRs QOffsetrange
}

package ies

// QCL-Info-referenceSignal ::= CHOICE
const (
	QclInfoReferencesignalChoiceNothing = iota
	QclInfoReferencesignalChoiceCsiRs
	QclInfoReferencesignalChoiceSsb
)

type QclInfoReferencesignal struct {
	Choice uint64
	CsiRs  *NzpCsiRsResourceid
	Ssb    *SsbIndex
}

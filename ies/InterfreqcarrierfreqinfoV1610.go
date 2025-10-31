package ies

// InterFreqCarrierFreqInfo-v1610 ::= SEQUENCE
type InterfreqcarrierfreqinfoV1610 struct {
	AltcellreselectionpriorityR16    *Cellreselectionpriority
	AltcellreselectionsubpriorityR16 *CellreselectionsubpriorityR13
	RssConfigcarrierinfoR16          *RssConfigcarrierinfoR16
	InterfreqneighcelllistV1610      *InterfreqneighcelllistV1610
}

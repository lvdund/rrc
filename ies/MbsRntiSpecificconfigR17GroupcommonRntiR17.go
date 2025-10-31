package ies

// MBS-RNTI-SpecificConfig-r17-groupCommon-RNTI-r17 ::= CHOICE
const (
	MbsRntiSpecificconfigR17GroupcommonRntiR17ChoiceNothing = iota
	MbsRntiSpecificconfigR17GroupcommonRntiR17ChoiceGRnti
	MbsRntiSpecificconfigR17GroupcommonRntiR17ChoiceGCsRnti
)

type MbsRntiSpecificconfigR17GroupcommonRntiR17 struct {
	Choice  uint64
	GRnti   *RntiValue
	GCsRnti *RntiValue
}

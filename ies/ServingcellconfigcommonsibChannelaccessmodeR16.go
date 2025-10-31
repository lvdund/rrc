package ies

// ServingCellConfigCommonSIB-channelAccessMode-r16 ::= CHOICE
const (
	ServingcellconfigcommonsibChannelaccessmodeR16ChoiceNothing = iota
	ServingcellconfigcommonsibChannelaccessmodeR16ChoiceDynamic
	ServingcellconfigcommonsibChannelaccessmodeR16ChoiceSemistatic
)

type ServingcellconfigcommonsibChannelaccessmodeR16 struct {
	Choice     uint64
	Dynamic    *struct{}
	Semistatic *SemistaticchannelaccessconfigR16
}

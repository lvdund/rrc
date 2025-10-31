package ies

// ServingCellConfigCommon-channelAccessMode-r16 ::= CHOICE
const (
	ServingcellconfigcommonChannelaccessmodeR16ChoiceNothing = iota
	ServingcellconfigcommonChannelaccessmodeR16ChoiceDynamic
	ServingcellconfigcommonChannelaccessmodeR16ChoiceSemistatic
)

type ServingcellconfigcommonChannelaccessmodeR16 struct {
	Choice     uint64
	Dynamic    *struct{}
	Semistatic *SemistaticchannelaccessconfigR16
}

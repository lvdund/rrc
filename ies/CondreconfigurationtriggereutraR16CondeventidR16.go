package ies

// CondReconfigurationTriggerEUTRA-r16-condEventId-r16 ::= CHOICE
// Extensible
const (
	CondreconfigurationtriggereutraR16CondeventidR16ChoiceNothing = iota
	CondreconfigurationtriggereutraR16CondeventidR16ChoiceCondeventa3R16
	CondreconfigurationtriggereutraR16CondeventidR16ChoiceCondeventa5R16
)

type CondreconfigurationtriggereutraR16CondeventidR16 struct {
	Choice         uint64
	Condeventa3R16 *CondreconfigurationtriggereutraR16CondeventidR16Condeventa3R16
	Condeventa5R16 *CondreconfigurationtriggereutraR16CondeventidR16Condeventa5R16
}

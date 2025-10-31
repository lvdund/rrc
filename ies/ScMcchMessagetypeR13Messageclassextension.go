package ies

// SC-MCCH-MessageType-r13-messageClassExtension ::= CHOICE
const (
	ScMcchMessagetypeR13MessageclassextensionChoiceNothing = iota
	ScMcchMessagetypeR13MessageclassextensionChoiceC2
	ScMcchMessagetypeR13MessageclassextensionChoiceMessageclassextensionfutureR14
)

type ScMcchMessagetypeR13Messageclassextension struct {
	Choice                         uint64
	C2                             *ScMcchMessagetypeR13MessageclassextensionC2
	MessageclassextensionfutureR14 *ScMcchMessagetypeR13MessageclassextensionMessageclassextensionfutureR14
}

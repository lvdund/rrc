package ies

// SCCH-MessageType-messageClassExtension ::= CHOICE
const (
	ScchMessagetypeMessageclassextensionChoiceNothing = iota
	ScchMessagetypeMessageclassextensionChoiceC2
	ScchMessagetypeMessageclassextensionChoiceMessageclassextensionfutureR17
)

type ScchMessagetypeMessageclassextension struct {
	Choice                         uint64
	C2                             *ScchMessagetypeMessageclassextensionC2
	MessageclassextensionfutureR17 *ScchMessagetypeMessageclassextensionMessageclassextensionfutureR17
}

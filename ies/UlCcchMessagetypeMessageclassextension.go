package ies

// UL-CCCH-MessageType-messageClassExtension ::= CHOICE
const (
	UlCcchMessagetypeMessageclassextensionChoiceNothing = iota
	UlCcchMessagetypeMessageclassextensionChoiceC2
	UlCcchMessagetypeMessageclassextensionChoiceMessageclassextensionfutureR13
)

type UlCcchMessagetypeMessageclassextension struct {
	Choice                         uint64
	C2                             *UlCcchMessagetypeMessageclassextensionC2
	MessageclassextensionfutureR13 *UlCcchMessagetypeMessageclassextensionMessageclassextensionfutureR13
}

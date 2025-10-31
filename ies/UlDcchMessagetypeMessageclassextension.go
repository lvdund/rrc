package ies

// UL-DCCH-MessageType-messageClassExtension ::= CHOICE
const (
	UlDcchMessagetypeMessageclassextensionChoiceNothing = iota
	UlDcchMessagetypeMessageclassextensionChoiceC2
	UlDcchMessagetypeMessageclassextensionChoiceMessageclassextensionfutureR11
)

type UlDcchMessagetypeMessageclassextension struct {
	Choice                         uint64
	C2                             *UlDcchMessagetypeMessageclassextensionC2
	MessageclassextensionfutureR11 *UlDcchMessagetypeMessageclassextensionMessageclassextensionfutureR11
}

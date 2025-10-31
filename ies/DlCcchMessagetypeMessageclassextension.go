package ies

// DL-CCCH-MessageType-messageClassExtension ::= CHOICE
const (
	DlCcchMessagetypeMessageclassextensionChoiceNothing = iota
	DlCcchMessagetypeMessageclassextensionChoiceC2
	DlCcchMessagetypeMessageclassextensionChoiceMessageclassextensionfutureR15
)

type DlCcchMessagetypeMessageclassextension struct {
	Choice                         uint64
	C2                             *DlCcchMessagetypeMessageclassextensionC2
	MessageclassextensionfutureR15 *DlCcchMessagetypeMessageclassextensionMessageclassextensionfutureR15
}

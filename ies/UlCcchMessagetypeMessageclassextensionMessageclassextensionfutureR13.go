package ies

// UL-CCCH-MessageType-messageClassExtension-messageClassExtensionFuture-r13 ::= CHOICE
const (
	UlCcchMessagetypeMessageclassextensionMessageclassextensionfutureR13ChoiceNothing = iota
	UlCcchMessagetypeMessageclassextensionMessageclassextensionfutureR13ChoiceC3
	UlCcchMessagetypeMessageclassextensionMessageclassextensionfutureR13ChoiceMessageclassextensionfutureR15
)

type UlCcchMessagetypeMessageclassextensionMessageclassextensionfutureR13 struct {
	Choice                         uint64
	C3                             *UlCcchMessagetypeMessageclassextensionMessageclassextensionfutureR13C3
	MessageclassextensionfutureR15 *UlCcchMessagetypeMessageclassextensionMessageclassextensionfutureR13MessageclassextensionfutureR15
}

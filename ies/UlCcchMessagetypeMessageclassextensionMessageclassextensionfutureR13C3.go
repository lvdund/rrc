package ies

// UL-CCCH-MessageType-messageClassExtension-messageClassExtensionFuture-r13-c3 ::= CHOICE
const (
	UlCcchMessagetypeMessageclassextensionMessageclassextensionfutureR13C3ChoiceNothing = iota
	UlCcchMessagetypeMessageclassextensionMessageclassextensionfutureR13C3ChoiceRrcearlydatarequestR15
	UlCcchMessagetypeMessageclassextensionMessageclassextensionfutureR13C3ChoiceSpare3
	UlCcchMessagetypeMessageclassextensionMessageclassextensionfutureR13C3ChoiceSpare2
	UlCcchMessagetypeMessageclassextensionMessageclassextensionfutureR13C3ChoiceSpare1
)

type UlCcchMessagetypeMessageclassextensionMessageclassextensionfutureR13C3 struct {
	Choice                 uint64
	RrcearlydatarequestR15 *RrcearlydatarequestR15
	Spare3                 *struct{}
	Spare2                 *struct{}
	Spare1                 *struct{}
}

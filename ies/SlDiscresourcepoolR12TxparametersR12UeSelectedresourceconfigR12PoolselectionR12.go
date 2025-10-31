package ies

// SL-DiscResourcePool-r12-txParameters-r12-ue-SelectedResourceConfig-r12-poolSelection-r12 ::= CHOICE
const (
	SlDiscresourcepoolR12TxparametersR12UeSelectedresourceconfigR12PoolselectionR12ChoiceNothing = iota
	SlDiscresourcepoolR12TxparametersR12UeSelectedresourceconfigR12PoolselectionR12ChoiceRsrpbasedR12
	SlDiscresourcepoolR12TxparametersR12UeSelectedresourceconfigR12PoolselectionR12ChoiceRandomR12
)

type SlDiscresourcepoolR12TxparametersR12UeSelectedresourceconfigR12PoolselectionR12 struct {
	Choice       uint64
	RsrpbasedR12 *SlPoolselectionconfigR12
	RandomR12    *struct{}
}

package ies

// SL-CommResourcePool-r12 ::= SEQUENCE
// Extensible
type SlCommresourcepoolR12 struct {
	ScCpLenR12                  SlCpLenR12
	ScPeriodR12                 SlPeriodcommR12
	ScTfResourceconfigR12       SlTfResourceconfigR12
	DataCpLenR12                SlCpLenR12
	DatahoppingconfigR12        SlHoppingconfigcommR12
	UeSelectedresourceconfigR12 *SlCommresourcepoolR12UeSelectedresourceconfigR12
	RxparametersncellR12        *SlCommresourcepoolR12RxparametersncellR12
	TxparametersR12             *SlCommresourcepoolR12TxparametersR12
}

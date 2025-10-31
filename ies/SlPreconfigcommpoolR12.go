package ies

// SL-PreconfigCommPool-r12 ::= SEQUENCE
// Extensible
type SlPreconfigcommpoolR12 struct {
	ScCpLenR12              SlCpLenR12
	ScPeriodR12             SlPeriodcommR12
	ScTfResourceconfigR12   SlTfResourceconfigR12
	ScTxparametersR12       P0SlR12
	DataCpLenR12            SlCpLenR12
	DataTfResourceconfigR12 SlTfResourceconfigR12
	DatahoppingconfigR12    SlHoppingconfigcommR12
	DatatxparametersR12     P0SlR12
	TrptSubsetR12           SlTrptSubsetR12
}

package ies

// ProcessingParameters ::= SEQUENCE
type Processingparameters struct {
	Fallback           ProcessingparametersFallback
	DifferenttbPerslot *ProcessingparametersDifferenttbPerslot
}

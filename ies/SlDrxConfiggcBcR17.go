package ies

// SL-DRX-ConfigGC-BC-r17 ::= SEQUENCE
// Extensible
type SlDrxConfiggcBcR17 struct {
	SlDrxGcBcPerqosListR17 *[]SlDrxGcBcQosR17 `lb:1,ub:maxSLGcBcDrxQosR17`
	SlDrxGcGenericR17      *SlDrxGcGenericR17
	SlDefaultdrxGcBcR17    *SlDrxGcBcQosR17
}

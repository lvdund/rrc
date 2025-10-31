package ies

// SL-DRX-GC-BC-QoS-r17 ::= SEQUENCE
// Extensible
type SlDrxGcBcQosR17 struct {
	SlDrxGcBcMappedqosFlowlistR17 *[]SlQosProfileR16 `lb:1,ub:maxNrofSLQfisR16`
	SlDrxGcBcOndurationtimerR17   SlDrxGcBcQosR17SlDrxGcBcOndurationtimerR17
	SlDrxGcInactivitytimerR17     SlDrxGcBcQosR17SlDrxGcInactivitytimerR17
	SlDrxGcBcCycleR17             SlDrxGcBcQosR17SlDrxGcBcCycleR17
}

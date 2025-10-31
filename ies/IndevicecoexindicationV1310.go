package ies

// InDeviceCoexIndication-v1310-IEs ::= SEQUENCE
type IndevicecoexindicationV1310 struct {
	AffectedcarrierfreqlistV1310   *AffectedcarrierfreqlistV1310
	AffectedcarrierfreqcomblistR13 *AffectedcarrierfreqcomblistR13
	Noncriticalextension           *IndevicecoexindicationV1360
}

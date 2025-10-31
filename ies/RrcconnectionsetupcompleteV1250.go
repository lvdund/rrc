package ies

// RRCConnectionSetupComplete-v1250-IEs ::= SEQUENCE
type RrcconnectionsetupcompleteV1250 struct {
	MobilitystateR12         *RrcconnectionsetupcompleteV1250IesMobilitystateR12
	MobilityhistoryavailR12  *bool
	LogmeasavailablembsfnR12 *bool
	Noncriticalextension     *RrcconnectionsetupcompleteV1320
}

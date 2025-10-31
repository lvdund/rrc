package ies

// LWIP-Config-r13 ::= SEQUENCE
// Extensible
type LwipConfigR13 struct {
	LwipMobilityconfigR13 *WlanMobilityconfigR13
	TunnelconfiglwipR13   *TunnelconfiglwipR13
}

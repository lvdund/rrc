package ies

// DAPS-UplinkPowerConfig-r16 ::= SEQUENCE
type DapsUplinkpowerconfigR16 struct {
	PDapsSourceR16                PMax
	PDapsTargetR16                PMax
	UplinkpowersharingdapsModeR16 DapsUplinkpowerconfigR16UplinkpowersharingdapsModeR16
}

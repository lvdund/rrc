package ies

// ConfigRestrictInfoDAPS-r16-powerCoordination-r16 ::= SEQUENCE
type ConfigrestrictinfodapsR16PowercoordinationR16 struct {
	PDapsSourceR16                PMax
	PDapsTargetR16                PMax
	UplinkpowersharingdapsModeR16 ConfigrestrictinfodapsR16PowercoordinationR16UplinkpowersharingdapsModeR16
}

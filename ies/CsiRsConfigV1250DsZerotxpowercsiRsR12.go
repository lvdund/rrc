package ies

// CSI-RS-Config-v1250-ds-ZeroTxPowerCSI-RS-r12 ::= CHOICE
const (
	CsiRsConfigV1250DsZerotxpowercsiRsR12ChoiceNothing = iota
	CsiRsConfigV1250DsZerotxpowercsiRsR12ChoiceRelease
	CsiRsConfigV1250DsZerotxpowercsiRsR12ChoiceSetup
)

type CsiRsConfigV1250DsZerotxpowercsiRsR12 struct {
	Choice  uint64
	Release *struct{}
	Setup   *CsiRsConfigV1250DsZerotxpowercsiRsR12Setup
}

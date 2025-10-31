package ies

// ZeroTxPowerCSI-RS-Conf-r12 ::= CHOICE
const (
	ZerotxpowercsiRsConfR12ChoiceNothing = iota
	ZerotxpowercsiRsConfR12ChoiceRelease
	ZerotxpowercsiRsConfR12ChoiceSetup
)

type ZerotxpowercsiRsConfR12 struct {
	Choice  uint64
	Release *struct{}
	Setup   *ZerotxpowercsiRsR12
}

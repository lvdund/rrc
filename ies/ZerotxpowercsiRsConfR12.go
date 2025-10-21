package ies

import "rrc/utils"

// ZeroTxPowerCSI-RS-Conf-r12 ::= CHOICE
type ZerotxpowercsiRsConfR12 interface {
	isZerotxpowercsiRsConfR12()
}

type ZerotxpowercsiRsConfR12Release struct {
	Value struct{}
}

func (*ZerotxpowercsiRsConfR12Release) isZerotxpowercsiRsConfR12() {}

type ZerotxpowercsiRsConfR12Setup struct {
	Value ZerotxpowercsiRsR12
}

func (*ZerotxpowercsiRsConfR12Setup) isZerotxpowercsiRsConfR12() {}

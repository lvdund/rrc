package ies

import "rrc/utils"

// BandSidelink-r16-sl-Reception-r16-extendedCP-RxSidelink-r16 ::= ENUMERATED
type BandsidelinkR16SlReceptionR16ExtendedcpRxsidelinkR16 struct {
	Value utils.ENUMERATED
}

const (
	BandsidelinkR16SlReceptionR16ExtendedcpRxsidelinkR16EnumeratedNothing = iota
	BandsidelinkR16SlReceptionR16ExtendedcpRxsidelinkR16EnumeratedSupported
)

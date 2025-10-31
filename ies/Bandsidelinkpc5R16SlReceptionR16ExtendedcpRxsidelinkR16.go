package ies

import "rrc/utils"

// BandSidelinkPC5-r16-sl-Reception-r16-extendedCP-RxSidelink-r16 ::= ENUMERATED
type Bandsidelinkpc5R16SlReceptionR16ExtendedcpRxsidelinkR16 struct {
	Value utils.ENUMERATED
}

const (
	Bandsidelinkpc5R16SlReceptionR16ExtendedcpRxsidelinkR16EnumeratedNothing = iota
	Bandsidelinkpc5R16SlReceptionR16ExtendedcpRxsidelinkR16EnumeratedSupported
)

package ies

import "rrc/utils"

// BandSidelink-r16-sl-Reception-r16-pscch-RxSidelink-r16 ::= ENUMERATED
type BandsidelinkR16SlReceptionR16PscchRxsidelinkR16 struct {
	Value utils.ENUMERATED
}

const (
	BandsidelinkR16SlReceptionR16PscchRxsidelinkR16EnumeratedNothing = iota
	BandsidelinkR16SlReceptionR16PscchRxsidelinkR16EnumeratedValue1
	BandsidelinkR16SlReceptionR16PscchRxsidelinkR16EnumeratedValue2
)

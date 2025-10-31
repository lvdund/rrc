package ies

import "rrc/utils"

// BandSidelinkPC5-r16-sl-Reception-r16-pscch-RxSidelink-r16 ::= ENUMERATED
type Bandsidelinkpc5R16SlReceptionR16PscchRxsidelinkR16 struct {
	Value utils.ENUMERATED
}

const (
	Bandsidelinkpc5R16SlReceptionR16PscchRxsidelinkR16EnumeratedNothing = iota
	Bandsidelinkpc5R16SlReceptionR16PscchRxsidelinkR16EnumeratedValue1
	Bandsidelinkpc5R16SlReceptionR16PscchRxsidelinkR16EnumeratedValue2
)

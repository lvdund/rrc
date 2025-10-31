package ies

// BandSidelinkPC5-r16-sl-Reception-r16 ::= SEQUENCE
type Bandsidelinkpc5R16SlReceptionR16 struct {
	HarqRxprocesssidelinkR16  Bandsidelinkpc5R16SlReceptionR16HarqRxprocesssidelinkR16
	PscchRxsidelinkR16        Bandsidelinkpc5R16SlReceptionR16PscchRxsidelinkR16
	ScsCpPatternrxsidelinkR16 *Bandsidelinkpc5R16SlReceptionR16ScsCpPatternrxsidelinkR16
	ExtendedcpRxsidelinkR16   *Bandsidelinkpc5R16SlReceptionR16ExtendedcpRxsidelinkR16
}

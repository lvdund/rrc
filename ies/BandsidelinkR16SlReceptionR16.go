package ies

// BandSidelink-r16-sl-Reception-r16 ::= SEQUENCE
type BandsidelinkR16SlReceptionR16 struct {
	HarqRxprocesssidelinkR16  BandsidelinkR16SlReceptionR16HarqRxprocesssidelinkR16
	PscchRxsidelinkR16        BandsidelinkR16SlReceptionR16PscchRxsidelinkR16
	ScsCpPatternrxsidelinkR16 *BandsidelinkR16SlReceptionR16ScsCpPatternrxsidelinkR16
	ExtendedcpRxsidelinkR16   *BandsidelinkR16SlReceptionR16ExtendedcpRxsidelinkR16
}

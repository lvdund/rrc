package ies

// BandSidelink-r16-sl-TransmissionMode1-r16 ::= SEQUENCE
type BandsidelinkR16SlTransmissionmode1R16 struct {
	HarqTxprocessmodeonesidelinkR16  BandsidelinkR16SlTransmissionmode1R16HarqTxprocessmodeonesidelinkR16
	ScsCpPatterntxsidelinkmodeoneR16 *BandsidelinkR16SlTransmissionmode1R16ScsCpPatterntxsidelinkmodeoneR16
	ExtendedcpTxsidelinkR16          *BandsidelinkR16SlTransmissionmode1R16ExtendedcpTxsidelinkR16
	HarqReportonpucchR16             *BandsidelinkR16SlTransmissionmode1R16HarqReportonpucchR16
}

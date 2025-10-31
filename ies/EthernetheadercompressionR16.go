package ies

// EthernetHeaderCompression-r16 ::= SEQUENCE
// Extensible
type EthernetheadercompressionR16 struct {
	EhcCommonR16   EthernetheadercompressionR16EhcCommonR16
	EhcDownlinkR16 *EthernetheadercompressionR16EhcDownlinkR16
	EhcUplinkR16   *EthernetheadercompressionR16EhcUplinkR16
}

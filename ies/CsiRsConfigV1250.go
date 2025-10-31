package ies

// CSI-RS-Config-v1250 ::= SEQUENCE
type CsiRsConfigV1250 struct {
	ZerotxpowercsiRs2R12  *ZerotxpowercsiRsConfR12
	DsZerotxpowercsiRsR12 *CsiRsConfigV1250DsZerotxpowercsiRsR12
}

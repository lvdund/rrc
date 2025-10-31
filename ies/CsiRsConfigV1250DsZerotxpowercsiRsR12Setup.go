package ies

// CSI-RS-Config-v1250-ds-ZeroTxPowerCSI-RS-r12-setup ::= SEQUENCE
type CsiRsConfigV1250DsZerotxpowercsiRsR12Setup struct {
	ZerotxpowercsiRsListR12 []ZerotxpowercsiRsR12 `lb:1,ub:maxDSZtpCsiRsR12`
}

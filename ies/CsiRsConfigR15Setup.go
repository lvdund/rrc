package ies

// CSI-RS-Config-r15-setup ::= SEQUENCE
type CsiRsConfigR15Setup struct {
	CsiRsConfigR10   *CsiRsConfigR10
	CsiRsConfigV1250 *CsiRsConfigV1250
	CsiRsConfigV1310 *CsiRsConfigV1310
	CsiRsConfigV1430 *CsiRsConfigV1430
}

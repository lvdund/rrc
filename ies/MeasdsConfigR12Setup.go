package ies

// MeasDS-Config-r12-setup ::= SEQUENCE
// Extensible
type MeasdsConfigR12Setup struct {
	DmtcPeriodoffsetR12      MeasdsConfigR12SetupDmtcPeriodoffsetR12
	DsOccasiondurationR12    MeasdsConfigR12SetupDsOccasiondurationR12
	MeascsiRsToremovelistR12 *MeascsiRsToremovelistR12
	MeascsiRsToaddmodlistR12 *MeascsiRsToaddmodlistR12
}

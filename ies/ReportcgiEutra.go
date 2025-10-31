package ies

// ReportCGI-EUTRA ::= SEQUENCE
// Extensible
type ReportcgiEutra struct {
	Cellforwhichtoreportcgi EutraPhyscellid
	UseautonomousgapsR16    *ReportcgiEutraUseautonomousgapsR16 `ext`
}

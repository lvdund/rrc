package ies

// ReportCGI ::= SEQUENCE
// Extensible
type Reportcgi struct {
	Cellforwhichtoreportcgi Physcellid
	UseautonomousgapsR16    *ReportcgiUseautonomousgapsR16 `ext`
}

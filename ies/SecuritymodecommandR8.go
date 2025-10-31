package ies

// SecurityModeCommand-r8-IEs ::= SEQUENCE
type SecuritymodecommandR8 struct {
	Securityconfigsmc    Securityconfigsmc
	Noncriticalextension *SecuritymodecommandV8a0
}

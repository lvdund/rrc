package ies

// WLANConnectionStatusReport-r13-criticalExtensions-c1 ::= CHOICE
const (
	WlanconnectionstatusreportR13CriticalextensionsC1ChoiceNothing = iota
	WlanconnectionstatusreportR13CriticalextensionsC1ChoiceWlanconnectionstatusreportR13
	WlanconnectionstatusreportR13CriticalextensionsC1ChoiceSpare3
	WlanconnectionstatusreportR13CriticalextensionsC1ChoiceSpare2
	WlanconnectionstatusreportR13CriticalextensionsC1ChoiceSpare1
)

type WlanconnectionstatusreportR13CriticalextensionsC1 struct {
	Choice                        uint64
	WlanconnectionstatusreportR13 *WlanconnectionstatusreportR13
	Spare3                        *struct{}
	Spare2                        *struct{}
	Spare1                        *struct{}
}

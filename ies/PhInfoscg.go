package ies

// PH-InfoSCG ::= SEQUENCE
// Extensible
type PhInfoscg struct {
	Servcellindex            Servcellindex
	PhUplink                 PhUplinkcarrierscg
	PhSupplementaryuplink    *PhUplinkcarrierscg
	TwosrsPuschRepetitionR17 *PhInfoscgTwosrsPuschRepetitionR17 `ext`
}

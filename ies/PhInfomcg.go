package ies

// PH-InfoMCG ::= SEQUENCE
// Extensible
type PhInfomcg struct {
	Servcellindex            Servcellindex
	PhUplink                 PhUplinkcarriermcg
	PhSupplementaryuplink    *PhUplinkcarriermcg
	TwosrsPuschRepetitionR17 *PhInfomcgTwosrsPuschRepetitionR17 `ext`
}

package ies

// SystemInformation-r8-IEs-sib-TypeAndInfo-Item ::= CHOICE
// Extensible
const (
	SysteminformationR8IesSibTypeandinfoItemChoiceNothing = iota
	SysteminformationR8IesSibTypeandinfoItemChoiceSib2
	SysteminformationR8IesSibTypeandinfoItemChoiceSib3
	SysteminformationR8IesSibTypeandinfoItemChoiceSib4
	SysteminformationR8IesSibTypeandinfoItemChoiceSib5
	SysteminformationR8IesSibTypeandinfoItemChoiceSib6
	SysteminformationR8IesSibTypeandinfoItemChoiceSib7
	SysteminformationR8IesSibTypeandinfoItemChoiceSib8
	SysteminformationR8IesSibTypeandinfoItemChoiceSib9
	SysteminformationR8IesSibTypeandinfoItemChoiceSib10
	SysteminformationR8IesSibTypeandinfoItemChoiceSib11
	SysteminformationR8IesSibTypeandinfoItemChoiceSib12V920
	SysteminformationR8IesSibTypeandinfoItemChoiceSib13V920
	SysteminformationR8IesSibTypeandinfoItemChoiceSib14V1130
	SysteminformationR8IesSibTypeandinfoItemChoiceSib15V1130
	SysteminformationR8IesSibTypeandinfoItemChoiceSib16V1130
	SysteminformationR8IesSibTypeandinfoItemChoiceSib17V1250
	SysteminformationR8IesSibTypeandinfoItemChoiceSib18V1250
	SysteminformationR8IesSibTypeandinfoItemChoiceSib19V1250
	SysteminformationR8IesSibTypeandinfoItemChoiceSib20V1310
	SysteminformationR8IesSibTypeandinfoItemChoiceSib21V1430
	SysteminformationR8IesSibTypeandinfoItemChoiceSib24V1530
	SysteminformationR8IesSibTypeandinfoItemChoiceSib25V1530
	SysteminformationR8IesSibTypeandinfoItemChoiceSib26V1530
	SysteminformationR8IesSibTypeandinfoItemChoiceSib26aV1610
	SysteminformationR8IesSibTypeandinfoItemChoiceSib27V1610
	SysteminformationR8IesSibTypeandinfoItemChoiceSib28V1610
	SysteminformationR8IesSibTypeandinfoItemChoiceSib29V1610
)

type SysteminformationR8IesSibTypeandinfoItem struct {
	Choice      uint64
	Sib2        *Systeminformationblocktype2
	Sib3        *Systeminformationblocktype3
	Sib4        *Systeminformationblocktype4
	Sib5        *Systeminformationblocktype5
	Sib6        *Systeminformationblocktype6
	Sib7        *Systeminformationblocktype7
	Sib8        *Systeminformationblocktype8
	Sib9        *Systeminformationblocktype9
	Sib10       *Systeminformationblocktype10
	Sib11       *Systeminformationblocktype11
	Sib12V920   *Systeminformationblocktype12R9
	Sib13V920   *Systeminformationblocktype13R9
	Sib14V1130  *Systeminformationblocktype14R11
	Sib15V1130  *Systeminformationblocktype15R11
	Sib16V1130  *Systeminformationblocktype16R11
	Sib17V1250  *Systeminformationblocktype17R12
	Sib18V1250  *Systeminformationblocktype18R12
	Sib19V1250  *Systeminformationblocktype19R12
	Sib20V1310  *Systeminformationblocktype20R13
	Sib21V1430  *Systeminformationblocktype21R14
	Sib24V1530  *Systeminformationblocktype24R15
	Sib25V1530  *Systeminformationblocktype25R15
	Sib26V1530  *Systeminformationblocktype26R15
	Sib26aV1610 *Systeminformationblocktype26aR16
	Sib27V1610  *Systeminformationblocktype27R16
	Sib28V1610  *Systeminformationblocktype28R16
	Sib29V1610  *Systeminformationblocktype29R16
}

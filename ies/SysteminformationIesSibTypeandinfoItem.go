package ies

// SystemInformation-IEs-sib-TypeAndInfo-Item ::= CHOICE
// Extensible
const (
	SysteminformationIesSibTypeandinfoItemChoiceNothing = iota
	SysteminformationIesSibTypeandinfoItemChoiceSib2
	SysteminformationIesSibTypeandinfoItemChoiceSib3
	SysteminformationIesSibTypeandinfoItemChoiceSib4
	SysteminformationIesSibTypeandinfoItemChoiceSib5
	SysteminformationIesSibTypeandinfoItemChoiceSib6
	SysteminformationIesSibTypeandinfoItemChoiceSib7
	SysteminformationIesSibTypeandinfoItemChoiceSib8
	SysteminformationIesSibTypeandinfoItemChoiceSib9
	SysteminformationIesSibTypeandinfoItemChoiceSib10V1610
	SysteminformationIesSibTypeandinfoItemChoiceSib11V1610
	SysteminformationIesSibTypeandinfoItemChoiceSib12V1610
	SysteminformationIesSibTypeandinfoItemChoiceSib13V1610
	SysteminformationIesSibTypeandinfoItemChoiceSib14V1610
	SysteminformationIesSibTypeandinfoItemChoiceSib15V1700
	SysteminformationIesSibTypeandinfoItemChoiceSib16V1700
	SysteminformationIesSibTypeandinfoItemChoiceSib17V1700
	SysteminformationIesSibTypeandinfoItemChoiceSib18V1700
	SysteminformationIesSibTypeandinfoItemChoiceSib19V1700
	SysteminformationIesSibTypeandinfoItemChoiceSib20V1700
	SysteminformationIesSibTypeandinfoItemChoiceSib21V1700
)

type SysteminformationIesSibTypeandinfoItem struct {
	Choice     uint64
	Sib2       *Sib2
	Sib3       *Sib3
	Sib4       *Sib4
	Sib5       *Sib5
	Sib6       *Sib6
	Sib7       *Sib7
	Sib8       *Sib8
	Sib9       *Sib9
	Sib10V1610 *Sib10R16
	Sib11V1610 *Sib11R16
	Sib12V1610 *Sib12R16
	Sib13V1610 *Sib13R16
	Sib14V1610 *Sib14R16
	Sib15V1700 *Sib15R17
	Sib16V1700 *Sib16R17
	Sib17V1700 *Sib17R17
	Sib18V1700 *Sib18R17
	Sib19V1700 *Sib19R17
	Sib20V1700 *Sib20R17
	Sib21V1700 *Sib21R17
}

package ies

// CondTriggerConfig-r16-condEventId-condEventA5 ::= SEQUENCE
type CondtriggerconfigR16CondeventidCondeventa5 struct {
	A5Threshold1  Meastriggerquantity
	A5Threshold2  Meastriggerquantity
	Hysteresis    Hysteresis
	Timetotrigger Timetotrigger
}

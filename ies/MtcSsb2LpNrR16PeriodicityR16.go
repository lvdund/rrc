package ies

import "rrc/utils"

// MTC-SSB2-LP-NR-r16-periodicity-r16 ::= ENUMERATED
type MtcSsb2LpNrR16PeriodicityR16 struct {
	Value utils.ENUMERATED
}

const (
	MtcSsb2LpNrR16PeriodicityR16EnumeratedNothing = iota
	MtcSsb2LpNrR16PeriodicityR16EnumeratedSf10
	MtcSsb2LpNrR16PeriodicityR16EnumeratedSf20
	MtcSsb2LpNrR16PeriodicityR16EnumeratedSf40
	MtcSsb2LpNrR16PeriodicityR16EnumeratedSf80
	MtcSsb2LpNrR16PeriodicityR16EnumeratedSf160
	MtcSsb2LpNrR16PeriodicityR16EnumeratedSpare3
	MtcSsb2LpNrR16PeriodicityR16EnumeratedSpare2
	MtcSsb2LpNrR16PeriodicityR16EnumeratedSpare1
)

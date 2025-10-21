package ies

import "rrc/utils"

// PollByte-r14 ::= ENUMERATED
type PollbyteR14 struct {
	Value utils.ENUMERATED
}

const (
	PollbyteR14Kb1     = 0
	PollbyteR14Kb2     = 1
	PollbyteR14Kb5     = 2
	PollbyteR14Kb8     = 3
	PollbyteR14Kb10    = 4
	PollbyteR14Kb15    = 5
	PollbyteR14Kb3500  = 6
	PollbyteR14Kb4000  = 7
	PollbyteR14Kb4500  = 8
	PollbyteR14Kb5000  = 9
	PollbyteR14Kb5500  = 10
	PollbyteR14Kb6000  = 11
	PollbyteR14Kb6500  = 12
	PollbyteR14Kb7000  = 13
	PollbyteR14Kb7500  = 14
	PollbyteR14Kb8000  = 15
	PollbyteR14Kb9000  = 16
	PollbyteR14Kb10000 = 17
	PollbyteR14Kb11000 = 18
	PollbyteR14Kb12000 = 19
	PollbyteR14Kb13000 = 20
	PollbyteR14Kb14000 = 21
	PollbyteR14Kb15000 = 22
	PollbyteR14Kb16000 = 23
	PollbyteR14Kb17000 = 24
	PollbyteR14Kb18000 = 25
	PollbyteR14Kb19000 = 26
	PollbyteR14Kb20000 = 27
	PollbyteR14Kb25000 = 28
	PollbyteR14Kb30000 = 29
	PollbyteR14Kb35000 = 30
	PollbyteR14Kb40000 = 31
)

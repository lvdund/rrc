// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PollByte (line 14108).

const (
	PollByte_KB1      = 0
	PollByte_KB2      = 1
	PollByte_KB5      = 2
	PollByte_KB8      = 3
	PollByte_KB10     = 4
	PollByte_KB15     = 5
	PollByte_KB25     = 6
	PollByte_KB50     = 7
	PollByte_KB75     = 8
	PollByte_KB100    = 9
	PollByte_KB125    = 10
	PollByte_KB250    = 11
	PollByte_KB375    = 12
	PollByte_KB500    = 13
	PollByte_KB750    = 14
	PollByte_KB1000   = 15
	PollByte_KB1250   = 16
	PollByte_KB1500   = 17
	PollByte_KB2000   = 18
	PollByte_KB3000   = 19
	PollByte_KB4000   = 20
	PollByte_KB4500   = 21
	PollByte_KB5000   = 22
	PollByte_KB5500   = 23
	PollByte_KB6000   = 24
	PollByte_KB6500   = 25
	PollByte_KB7000   = 26
	PollByte_KB7500   = 27
	PollByte_MB8      = 28
	PollByte_MB9      = 29
	PollByte_MB10     = 30
	PollByte_MB11     = 31
	PollByte_MB12     = 32
	PollByte_MB13     = 33
	PollByte_MB14     = 34
	PollByte_MB15     = 35
	PollByte_MB16     = 36
	PollByte_MB17     = 37
	PollByte_MB18     = 38
	PollByte_MB20     = 39
	PollByte_MB25     = 40
	PollByte_MB30     = 41
	PollByte_MB40     = 42
	PollByte_Infinity = 43
	PollByte_Spare20  = 44
	PollByte_Spare19  = 45
	PollByte_Spare18  = 46
	PollByte_Spare17  = 47
	PollByte_Spare16  = 48
	PollByte_Spare15  = 49
	PollByte_Spare14  = 50
	PollByte_Spare13  = 51
	PollByte_Spare12  = 52
	PollByte_Spare11  = 53
	PollByte_Spare10  = 54
	PollByte_Spare9   = 55
	PollByte_Spare8   = 56
	PollByte_Spare7   = 57
	PollByte_Spare6   = 58
	PollByte_Spare5   = 59
	PollByte_Spare4   = 60
	PollByte_Spare3   = 61
	PollByte_Spare2   = 62
	PollByte_Spare1   = 63
)

var pollByteConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PollByte_KB1, PollByte_KB2, PollByte_KB5, PollByte_KB8, PollByte_KB10, PollByte_KB15, PollByte_KB25, PollByte_KB50, PollByte_KB75, PollByte_KB100, PollByte_KB125, PollByte_KB250, PollByte_KB375, PollByte_KB500, PollByte_KB750, PollByte_KB1000, PollByte_KB1250, PollByte_KB1500, PollByte_KB2000, PollByte_KB3000, PollByte_KB4000, PollByte_KB4500, PollByte_KB5000, PollByte_KB5500, PollByte_KB6000, PollByte_KB6500, PollByte_KB7000, PollByte_KB7500, PollByte_MB8, PollByte_MB9, PollByte_MB10, PollByte_MB11, PollByte_MB12, PollByte_MB13, PollByte_MB14, PollByte_MB15, PollByte_MB16, PollByte_MB17, PollByte_MB18, PollByte_MB20, PollByte_MB25, PollByte_MB30, PollByte_MB40, PollByte_Infinity, PollByte_Spare20, PollByte_Spare19, PollByte_Spare18, PollByte_Spare17, PollByte_Spare16, PollByte_Spare15, PollByte_Spare14, PollByte_Spare13, PollByte_Spare12, PollByte_Spare11, PollByte_Spare10, PollByte_Spare9, PollByte_Spare8, PollByte_Spare7, PollByte_Spare6, PollByte_Spare5, PollByte_Spare4, PollByte_Spare3, PollByte_Spare2, PollByte_Spare1},
}

type PollByte struct {
	Value int64
}

func (v *PollByte) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, pollByteConstraints)
}

func (v *PollByte) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(pollByteConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}

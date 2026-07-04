// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: T-PollRetransmit (line 14089).

const (
	T_PollRetransmit_Ms5       = 0
	T_PollRetransmit_Ms10      = 1
	T_PollRetransmit_Ms15      = 2
	T_PollRetransmit_Ms20      = 3
	T_PollRetransmit_Ms25      = 4
	T_PollRetransmit_Ms30      = 5
	T_PollRetransmit_Ms35      = 6
	T_PollRetransmit_Ms40      = 7
	T_PollRetransmit_Ms45      = 8
	T_PollRetransmit_Ms50      = 9
	T_PollRetransmit_Ms55      = 10
	T_PollRetransmit_Ms60      = 11
	T_PollRetransmit_Ms65      = 12
	T_PollRetransmit_Ms70      = 13
	T_PollRetransmit_Ms75      = 14
	T_PollRetransmit_Ms80      = 15
	T_PollRetransmit_Ms85      = 16
	T_PollRetransmit_Ms90      = 17
	T_PollRetransmit_Ms95      = 18
	T_PollRetransmit_Ms100     = 19
	T_PollRetransmit_Ms105     = 20
	T_PollRetransmit_Ms110     = 21
	T_PollRetransmit_Ms115     = 22
	T_PollRetransmit_Ms120     = 23
	T_PollRetransmit_Ms125     = 24
	T_PollRetransmit_Ms130     = 25
	T_PollRetransmit_Ms135     = 26
	T_PollRetransmit_Ms140     = 27
	T_PollRetransmit_Ms145     = 28
	T_PollRetransmit_Ms150     = 29
	T_PollRetransmit_Ms155     = 30
	T_PollRetransmit_Ms160     = 31
	T_PollRetransmit_Ms165     = 32
	T_PollRetransmit_Ms170     = 33
	T_PollRetransmit_Ms175     = 34
	T_PollRetransmit_Ms180     = 35
	T_PollRetransmit_Ms185     = 36
	T_PollRetransmit_Ms190     = 37
	T_PollRetransmit_Ms195     = 38
	T_PollRetransmit_Ms200     = 39
	T_PollRetransmit_Ms205     = 40
	T_PollRetransmit_Ms210     = 41
	T_PollRetransmit_Ms215     = 42
	T_PollRetransmit_Ms220     = 43
	T_PollRetransmit_Ms225     = 44
	T_PollRetransmit_Ms230     = 45
	T_PollRetransmit_Ms235     = 46
	T_PollRetransmit_Ms240     = 47
	T_PollRetransmit_Ms245     = 48
	T_PollRetransmit_Ms250     = 49
	T_PollRetransmit_Ms300     = 50
	T_PollRetransmit_Ms350     = 51
	T_PollRetransmit_Ms400     = 52
	T_PollRetransmit_Ms450     = 53
	T_PollRetransmit_Ms500     = 54
	T_PollRetransmit_Ms800     = 55
	T_PollRetransmit_Ms1000    = 56
	T_PollRetransmit_Ms2000    = 57
	T_PollRetransmit_Ms4000    = 58
	T_PollRetransmit_Ms1_v1610 = 59
	T_PollRetransmit_Ms2_v1610 = 60
	T_PollRetransmit_Ms3_v1610 = 61
	T_PollRetransmit_Ms4_v1610 = 62
	T_PollRetransmit_Spare1    = 63
)

var tPollRetransmitConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{T_PollRetransmit_Ms5, T_PollRetransmit_Ms10, T_PollRetransmit_Ms15, T_PollRetransmit_Ms20, T_PollRetransmit_Ms25, T_PollRetransmit_Ms30, T_PollRetransmit_Ms35, T_PollRetransmit_Ms40, T_PollRetransmit_Ms45, T_PollRetransmit_Ms50, T_PollRetransmit_Ms55, T_PollRetransmit_Ms60, T_PollRetransmit_Ms65, T_PollRetransmit_Ms70, T_PollRetransmit_Ms75, T_PollRetransmit_Ms80, T_PollRetransmit_Ms85, T_PollRetransmit_Ms90, T_PollRetransmit_Ms95, T_PollRetransmit_Ms100, T_PollRetransmit_Ms105, T_PollRetransmit_Ms110, T_PollRetransmit_Ms115, T_PollRetransmit_Ms120, T_PollRetransmit_Ms125, T_PollRetransmit_Ms130, T_PollRetransmit_Ms135, T_PollRetransmit_Ms140, T_PollRetransmit_Ms145, T_PollRetransmit_Ms150, T_PollRetransmit_Ms155, T_PollRetransmit_Ms160, T_PollRetransmit_Ms165, T_PollRetransmit_Ms170, T_PollRetransmit_Ms175, T_PollRetransmit_Ms180, T_PollRetransmit_Ms185, T_PollRetransmit_Ms190, T_PollRetransmit_Ms195, T_PollRetransmit_Ms200, T_PollRetransmit_Ms205, T_PollRetransmit_Ms210, T_PollRetransmit_Ms215, T_PollRetransmit_Ms220, T_PollRetransmit_Ms225, T_PollRetransmit_Ms230, T_PollRetransmit_Ms235, T_PollRetransmit_Ms240, T_PollRetransmit_Ms245, T_PollRetransmit_Ms250, T_PollRetransmit_Ms300, T_PollRetransmit_Ms350, T_PollRetransmit_Ms400, T_PollRetransmit_Ms450, T_PollRetransmit_Ms500, T_PollRetransmit_Ms800, T_PollRetransmit_Ms1000, T_PollRetransmit_Ms2000, T_PollRetransmit_Ms4000, T_PollRetransmit_Ms1_v1610, T_PollRetransmit_Ms2_v1610, T_PollRetransmit_Ms3_v1610, T_PollRetransmit_Ms4_v1610, T_PollRetransmit_Spare1},
}

type T_PollRetransmit struct {
	Value int64
}

func (v *T_PollRetransmit) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, tPollRetransmitConstraints)
}

func (v *T_PollRetransmit) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(tPollRetransmitConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}

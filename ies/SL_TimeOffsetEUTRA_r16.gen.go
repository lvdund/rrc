// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-TimeOffsetEUTRA-r16 (line 1119).

const (
	SL_TimeOffsetEUTRA_r16_Ms0       = 0
	SL_TimeOffsetEUTRA_r16_Ms0dot25  = 1
	SL_TimeOffsetEUTRA_r16_Ms0dot5   = 2
	SL_TimeOffsetEUTRA_r16_Ms0dot625 = 3
	SL_TimeOffsetEUTRA_r16_Ms0dot75  = 4
	SL_TimeOffsetEUTRA_r16_Ms1       = 5
	SL_TimeOffsetEUTRA_r16_Ms1dot25  = 6
	SL_TimeOffsetEUTRA_r16_Ms1dot5   = 7
	SL_TimeOffsetEUTRA_r16_Ms1dot75  = 8
	SL_TimeOffsetEUTRA_r16_Ms2       = 9
	SL_TimeOffsetEUTRA_r16_Ms2dot5   = 10
	SL_TimeOffsetEUTRA_r16_Ms3       = 11
	SL_TimeOffsetEUTRA_r16_Ms4       = 12
	SL_TimeOffsetEUTRA_r16_Ms5       = 13
	SL_TimeOffsetEUTRA_r16_Ms6       = 14
	SL_TimeOffsetEUTRA_r16_Ms8       = 15
	SL_TimeOffsetEUTRA_r16_Ms10      = 16
	SL_TimeOffsetEUTRA_r16_Ms20      = 17
)

var sLTimeOffsetEUTRAR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_TimeOffsetEUTRA_r16_Ms0, SL_TimeOffsetEUTRA_r16_Ms0dot25, SL_TimeOffsetEUTRA_r16_Ms0dot5, SL_TimeOffsetEUTRA_r16_Ms0dot625, SL_TimeOffsetEUTRA_r16_Ms0dot75, SL_TimeOffsetEUTRA_r16_Ms1, SL_TimeOffsetEUTRA_r16_Ms1dot25, SL_TimeOffsetEUTRA_r16_Ms1dot5, SL_TimeOffsetEUTRA_r16_Ms1dot75, SL_TimeOffsetEUTRA_r16_Ms2, SL_TimeOffsetEUTRA_r16_Ms2dot5, SL_TimeOffsetEUTRA_r16_Ms3, SL_TimeOffsetEUTRA_r16_Ms4, SL_TimeOffsetEUTRA_r16_Ms5, SL_TimeOffsetEUTRA_r16_Ms6, SL_TimeOffsetEUTRA_r16_Ms8, SL_TimeOffsetEUTRA_r16_Ms10, SL_TimeOffsetEUTRA_r16_Ms20},
}

type SL_TimeOffsetEUTRA_r16 struct {
	Value int64
}

func (v *SL_TimeOffsetEUTRA_r16) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, sLTimeOffsetEUTRAR16Constraints)
}

func (v *SL_TimeOffsetEUTRA_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(sLTimeOffsetEUTRAR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}

// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-SSB-TimeAllocation-r16 (line 28329).

var sLSSBTimeAllocationR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-NumSSB-WithinPeriod-r16", Optional: true},
		{Name: "sl-TimeOffsetSSB-r16", Optional: true},
		{Name: "sl-TimeInterval-r16", Optional: true},
	},
}

const (
	SL_SSB_TimeAllocation_r16_Sl_NumSSB_WithinPeriod_r16_N1  = 0
	SL_SSB_TimeAllocation_r16_Sl_NumSSB_WithinPeriod_r16_N2  = 1
	SL_SSB_TimeAllocation_r16_Sl_NumSSB_WithinPeriod_r16_N4  = 2
	SL_SSB_TimeAllocation_r16_Sl_NumSSB_WithinPeriod_r16_N8  = 3
	SL_SSB_TimeAllocation_r16_Sl_NumSSB_WithinPeriod_r16_N16 = 4
	SL_SSB_TimeAllocation_r16_Sl_NumSSB_WithinPeriod_r16_N32 = 5
	SL_SSB_TimeAllocation_r16_Sl_NumSSB_WithinPeriod_r16_N64 = 6
)

var sLSSBTimeAllocationR16SlNumSSBWithinPeriodR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_SSB_TimeAllocation_r16_Sl_NumSSB_WithinPeriod_r16_N1, SL_SSB_TimeAllocation_r16_Sl_NumSSB_WithinPeriod_r16_N2, SL_SSB_TimeAllocation_r16_Sl_NumSSB_WithinPeriod_r16_N4, SL_SSB_TimeAllocation_r16_Sl_NumSSB_WithinPeriod_r16_N8, SL_SSB_TimeAllocation_r16_Sl_NumSSB_WithinPeriod_r16_N16, SL_SSB_TimeAllocation_r16_Sl_NumSSB_WithinPeriod_r16_N32, SL_SSB_TimeAllocation_r16_Sl_NumSSB_WithinPeriod_r16_N64},
}

var sLSSBTimeAllocationR16SlTimeOffsetSSBR16Constraints = per.Constrained(0, 1279)

var sLSSBTimeAllocationR16SlTimeIntervalR16Constraints = per.Constrained(0, 639)

type SL_SSB_TimeAllocation_r16 struct {
	Sl_NumSSB_WithinPeriod_r16 *int64
	Sl_TimeOffsetSSB_r16       *int64
	Sl_TimeInterval_r16        *int64
}

func (ie *SL_SSB_TimeAllocation_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLSSBTimeAllocationR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_NumSSB_WithinPeriod_r16 != nil, ie.Sl_TimeOffsetSSB_r16 != nil, ie.Sl_TimeInterval_r16 != nil}); err != nil {
		return err
	}

	// 2. sl-NumSSB-WithinPeriod-r16: enumerated
	{
		if ie.Sl_NumSSB_WithinPeriod_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_NumSSB_WithinPeriod_r16, sLSSBTimeAllocationR16SlNumSSBWithinPeriodR16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. sl-TimeOffsetSSB-r16: integer
	{
		if ie.Sl_TimeOffsetSSB_r16 != nil {
			if err := e.EncodeInteger(*ie.Sl_TimeOffsetSSB_r16, sLSSBTimeAllocationR16SlTimeOffsetSSBR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. sl-TimeInterval-r16: integer
	{
		if ie.Sl_TimeInterval_r16 != nil {
			if err := e.EncodeInteger(*ie.Sl_TimeInterval_r16, sLSSBTimeAllocationR16SlTimeIntervalR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_SSB_TimeAllocation_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLSSBTimeAllocationR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-NumSSB-WithinPeriod-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(sLSSBTimeAllocationR16SlNumSSBWithinPeriodR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_NumSSB_WithinPeriod_r16 = &idx
		}
	}

	// 3. sl-TimeOffsetSSB-r16: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(sLSSBTimeAllocationR16SlTimeOffsetSSBR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_TimeOffsetSSB_r16 = &v
		}
	}

	// 4. sl-TimeInterval-r16: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(sLSSBTimeAllocationR16SlTimeIntervalR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_TimeInterval_r16 = &v
		}
	}

	return nil
}

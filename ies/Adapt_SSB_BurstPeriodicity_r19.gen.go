// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: Adapt-SSB-BurstPeriodicity-r19 (line 5863).

var adaptSSBBurstPeriodicityR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ssb-Periodicity-r19", Optional: true},
		{Name: "ssb-Offset-r19", Optional: true},
		{Name: "ssb-HalfFrameIndex-r19", Optional: true},
	},
}

const (
	Adapt_SSB_BurstPeriodicity_r19_Ssb_Periodicity_r19_Ms5    = 0
	Adapt_SSB_BurstPeriodicity_r19_Ssb_Periodicity_r19_Ms10   = 1
	Adapt_SSB_BurstPeriodicity_r19_Ssb_Periodicity_r19_Ms20   = 2
	Adapt_SSB_BurstPeriodicity_r19_Ssb_Periodicity_r19_Ms40   = 3
	Adapt_SSB_BurstPeriodicity_r19_Ssb_Periodicity_r19_Ms80   = 4
	Adapt_SSB_BurstPeriodicity_r19_Ssb_Periodicity_r19_Ms160  = 5
	Adapt_SSB_BurstPeriodicity_r19_Ssb_Periodicity_r19_Spare2 = 6
	Adapt_SSB_BurstPeriodicity_r19_Ssb_Periodicity_r19_Spare1 = 7
)

var adaptSSBBurstPeriodicityR19SsbPeriodicityR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Adapt_SSB_BurstPeriodicity_r19_Ssb_Periodicity_r19_Ms5, Adapt_SSB_BurstPeriodicity_r19_Ssb_Periodicity_r19_Ms10, Adapt_SSB_BurstPeriodicity_r19_Ssb_Periodicity_r19_Ms20, Adapt_SSB_BurstPeriodicity_r19_Ssb_Periodicity_r19_Ms40, Adapt_SSB_BurstPeriodicity_r19_Ssb_Periodicity_r19_Ms80, Adapt_SSB_BurstPeriodicity_r19_Ssb_Periodicity_r19_Ms160, Adapt_SSB_BurstPeriodicity_r19_Ssb_Periodicity_r19_Spare2, Adapt_SSB_BurstPeriodicity_r19_Ssb_Periodicity_r19_Spare1},
}

var adaptSSBBurstPeriodicityR19SsbOffsetR19Constraints = per.Constrained(0, 15)

const (
	Adapt_SSB_BurstPeriodicity_r19_Ssb_HalfFrameIndex_r19_Zero = 0
	Adapt_SSB_BurstPeriodicity_r19_Ssb_HalfFrameIndex_r19_One  = 1
)

var adaptSSBBurstPeriodicityR19SsbHalfFrameIndexR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Adapt_SSB_BurstPeriodicity_r19_Ssb_HalfFrameIndex_r19_Zero, Adapt_SSB_BurstPeriodicity_r19_Ssb_HalfFrameIndex_r19_One},
}

type Adapt_SSB_BurstPeriodicity_r19 struct {
	Ssb_Periodicity_r19    *int64
	Ssb_Offset_r19         *int64
	Ssb_HalfFrameIndex_r19 *int64
}

func (ie *Adapt_SSB_BurstPeriodicity_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(adaptSSBBurstPeriodicityR19Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ssb_Periodicity_r19 != nil, ie.Ssb_Offset_r19 != nil, ie.Ssb_HalfFrameIndex_r19 != nil}); err != nil {
		return err
	}

	// 2. ssb-Periodicity-r19: enumerated
	{
		if ie.Ssb_Periodicity_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Ssb_Periodicity_r19, adaptSSBBurstPeriodicityR19SsbPeriodicityR19Constraints); err != nil {
				return err
			}
		}
	}

	// 3. ssb-Offset-r19: integer
	{
		if ie.Ssb_Offset_r19 != nil {
			if err := e.EncodeInteger(*ie.Ssb_Offset_r19, adaptSSBBurstPeriodicityR19SsbOffsetR19Constraints); err != nil {
				return err
			}
		}
	}

	// 4. ssb-HalfFrameIndex-r19: enumerated
	{
		if ie.Ssb_HalfFrameIndex_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Ssb_HalfFrameIndex_r19, adaptSSBBurstPeriodicityR19SsbHalfFrameIndexR19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *Adapt_SSB_BurstPeriodicity_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(adaptSSBBurstPeriodicityR19Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ssb-Periodicity-r19: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(adaptSSBBurstPeriodicityR19SsbPeriodicityR19Constraints)
			if err != nil {
				return err
			}
			ie.Ssb_Periodicity_r19 = &idx
		}
	}

	// 3. ssb-Offset-r19: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(adaptSSBBurstPeriodicityR19SsbOffsetR19Constraints)
			if err != nil {
				return err
			}
			ie.Ssb_Offset_r19 = &v
		}
	}

	// 4. ssb-HalfFrameIndex-r19: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(adaptSSBBurstPeriodicityR19SsbHalfFrameIndexR19Constraints)
			if err != nil {
				return err
			}
			ie.Ssb_HalfFrameIndex_r19 = &idx
		}
	}

	return nil
}

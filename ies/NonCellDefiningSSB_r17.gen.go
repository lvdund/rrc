// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NonCellDefiningSSB-r17 (line 10550).

var nonCellDefiningSSBR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "absoluteFrequencySSB-r17"},
		{Name: "ssb-Periodicity-r17", Optional: true},
		{Name: "ssb-TimeOffset-r17", Optional: true},
	},
}

const (
	NonCellDefiningSSB_r17_Ssb_Periodicity_r17_Ms5    = 0
	NonCellDefiningSSB_r17_Ssb_Periodicity_r17_Ms10   = 1
	NonCellDefiningSSB_r17_Ssb_Periodicity_r17_Ms20   = 2
	NonCellDefiningSSB_r17_Ssb_Periodicity_r17_Ms40   = 3
	NonCellDefiningSSB_r17_Ssb_Periodicity_r17_Ms80   = 4
	NonCellDefiningSSB_r17_Ssb_Periodicity_r17_Ms160  = 5
	NonCellDefiningSSB_r17_Ssb_Periodicity_r17_Spare2 = 6
	NonCellDefiningSSB_r17_Ssb_Periodicity_r17_Spare1 = 7
)

var nonCellDefiningSSBR17SsbPeriodicityR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{NonCellDefiningSSB_r17_Ssb_Periodicity_r17_Ms5, NonCellDefiningSSB_r17_Ssb_Periodicity_r17_Ms10, NonCellDefiningSSB_r17_Ssb_Periodicity_r17_Ms20, NonCellDefiningSSB_r17_Ssb_Periodicity_r17_Ms40, NonCellDefiningSSB_r17_Ssb_Periodicity_r17_Ms80, NonCellDefiningSSB_r17_Ssb_Periodicity_r17_Ms160, NonCellDefiningSSB_r17_Ssb_Periodicity_r17_Spare2, NonCellDefiningSSB_r17_Ssb_Periodicity_r17_Spare1},
}

const (
	NonCellDefiningSSB_r17_Ssb_TimeOffset_r17_Ms5    = 0
	NonCellDefiningSSB_r17_Ssb_TimeOffset_r17_Ms10   = 1
	NonCellDefiningSSB_r17_Ssb_TimeOffset_r17_Ms15   = 2
	NonCellDefiningSSB_r17_Ssb_TimeOffset_r17_Ms20   = 3
	NonCellDefiningSSB_r17_Ssb_TimeOffset_r17_Ms40   = 4
	NonCellDefiningSSB_r17_Ssb_TimeOffset_r17_Ms80   = 5
	NonCellDefiningSSB_r17_Ssb_TimeOffset_r17_Spare2 = 6
	NonCellDefiningSSB_r17_Ssb_TimeOffset_r17_Spare1 = 7
)

var nonCellDefiningSSBR17SsbTimeOffsetR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{NonCellDefiningSSB_r17_Ssb_TimeOffset_r17_Ms5, NonCellDefiningSSB_r17_Ssb_TimeOffset_r17_Ms10, NonCellDefiningSSB_r17_Ssb_TimeOffset_r17_Ms15, NonCellDefiningSSB_r17_Ssb_TimeOffset_r17_Ms20, NonCellDefiningSSB_r17_Ssb_TimeOffset_r17_Ms40, NonCellDefiningSSB_r17_Ssb_TimeOffset_r17_Ms80, NonCellDefiningSSB_r17_Ssb_TimeOffset_r17_Spare2, NonCellDefiningSSB_r17_Ssb_TimeOffset_r17_Spare1},
}

type NonCellDefiningSSB_r17 struct {
	AbsoluteFrequencySSB_r17 ARFCN_ValueNR
	Ssb_Periodicity_r17      *int64
	Ssb_TimeOffset_r17       *int64
}

func (ie *NonCellDefiningSSB_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nonCellDefiningSSBR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ssb_Periodicity_r17 != nil, ie.Ssb_TimeOffset_r17 != nil}); err != nil {
		return err
	}

	// 3. absoluteFrequencySSB-r17: ref
	{
		if err := ie.AbsoluteFrequencySSB_r17.Encode(e); err != nil {
			return err
		}
	}

	// 4. ssb-Periodicity-r17: enumerated
	{
		if ie.Ssb_Periodicity_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Ssb_Periodicity_r17, nonCellDefiningSSBR17SsbPeriodicityR17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. ssb-TimeOffset-r17: enumerated
	{
		if ie.Ssb_TimeOffset_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Ssb_TimeOffset_r17, nonCellDefiningSSBR17SsbTimeOffsetR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *NonCellDefiningSSB_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nonCellDefiningSSBR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. absoluteFrequencySSB-r17: ref
	{
		if err := ie.AbsoluteFrequencySSB_r17.Decode(d); err != nil {
			return err
		}
	}

	// 4. ssb-Periodicity-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(nonCellDefiningSSBR17SsbPeriodicityR17Constraints)
			if err != nil {
				return err
			}
			ie.Ssb_Periodicity_r17 = &idx
		}
	}

	// 5. ssb-TimeOffset-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(nonCellDefiningSSBR17SsbTimeOffsetR17Constraints)
			if err != nil {
				return err
			}
			ie.Ssb_TimeOffset_r17 = &idx
		}
	}

	return nil
}

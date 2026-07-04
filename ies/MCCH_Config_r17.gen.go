// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MCCH-Config-r17 (line 4568).

var mCCHConfigR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "mcch-RepetitionPeriodAndOffset-r17"},
		{Name: "mcch-WindowStartSlot-r17"},
		{Name: "mcch-WindowDuration-r17", Optional: true},
		{Name: "mcch-ModificationPeriod-r17"},
	},
}

var mCCHConfigR17McchWindowStartSlotR17Constraints = per.Constrained(0, 79)

const (
	MCCH_Config_r17_Mcch_WindowDuration_r17_Sl2   = 0
	MCCH_Config_r17_Mcch_WindowDuration_r17_Sl4   = 1
	MCCH_Config_r17_Mcch_WindowDuration_r17_Sl8   = 2
	MCCH_Config_r17_Mcch_WindowDuration_r17_Sl10  = 3
	MCCH_Config_r17_Mcch_WindowDuration_r17_Sl20  = 4
	MCCH_Config_r17_Mcch_WindowDuration_r17_Sl40  = 5
	MCCH_Config_r17_Mcch_WindowDuration_r17_Sl80  = 6
	MCCH_Config_r17_Mcch_WindowDuration_r17_Sl160 = 7
)

var mCCHConfigR17McchWindowDurationR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MCCH_Config_r17_Mcch_WindowDuration_r17_Sl2, MCCH_Config_r17_Mcch_WindowDuration_r17_Sl4, MCCH_Config_r17_Mcch_WindowDuration_r17_Sl8, MCCH_Config_r17_Mcch_WindowDuration_r17_Sl10, MCCH_Config_r17_Mcch_WindowDuration_r17_Sl20, MCCH_Config_r17_Mcch_WindowDuration_r17_Sl40, MCCH_Config_r17_Mcch_WindowDuration_r17_Sl80, MCCH_Config_r17_Mcch_WindowDuration_r17_Sl160},
}

const (
	MCCH_Config_r17_Mcch_ModificationPeriod_r17_Rf2     = 0
	MCCH_Config_r17_Mcch_ModificationPeriod_r17_Rf4     = 1
	MCCH_Config_r17_Mcch_ModificationPeriod_r17_Rf8     = 2
	MCCH_Config_r17_Mcch_ModificationPeriod_r17_Rf16    = 3
	MCCH_Config_r17_Mcch_ModificationPeriod_r17_Rf32    = 4
	MCCH_Config_r17_Mcch_ModificationPeriod_r17_Rf64    = 5
	MCCH_Config_r17_Mcch_ModificationPeriod_r17_Rf128   = 6
	MCCH_Config_r17_Mcch_ModificationPeriod_r17_Rf256   = 7
	MCCH_Config_r17_Mcch_ModificationPeriod_r17_Rf512   = 8
	MCCH_Config_r17_Mcch_ModificationPeriod_r17_Rf1024  = 9
	MCCH_Config_r17_Mcch_ModificationPeriod_r17_Rf2048  = 10
	MCCH_Config_r17_Mcch_ModificationPeriod_r17_Rf4096  = 11
	MCCH_Config_r17_Mcch_ModificationPeriod_r17_Rf8192  = 12
	MCCH_Config_r17_Mcch_ModificationPeriod_r17_Rf16384 = 13
	MCCH_Config_r17_Mcch_ModificationPeriod_r17_Rf32768 = 14
	MCCH_Config_r17_Mcch_ModificationPeriod_r17_Rf65536 = 15
)

var mCCHConfigR17McchModificationPeriodR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MCCH_Config_r17_Mcch_ModificationPeriod_r17_Rf2, MCCH_Config_r17_Mcch_ModificationPeriod_r17_Rf4, MCCH_Config_r17_Mcch_ModificationPeriod_r17_Rf8, MCCH_Config_r17_Mcch_ModificationPeriod_r17_Rf16, MCCH_Config_r17_Mcch_ModificationPeriod_r17_Rf32, MCCH_Config_r17_Mcch_ModificationPeriod_r17_Rf64, MCCH_Config_r17_Mcch_ModificationPeriod_r17_Rf128, MCCH_Config_r17_Mcch_ModificationPeriod_r17_Rf256, MCCH_Config_r17_Mcch_ModificationPeriod_r17_Rf512, MCCH_Config_r17_Mcch_ModificationPeriod_r17_Rf1024, MCCH_Config_r17_Mcch_ModificationPeriod_r17_Rf2048, MCCH_Config_r17_Mcch_ModificationPeriod_r17_Rf4096, MCCH_Config_r17_Mcch_ModificationPeriod_r17_Rf8192, MCCH_Config_r17_Mcch_ModificationPeriod_r17_Rf16384, MCCH_Config_r17_Mcch_ModificationPeriod_r17_Rf32768, MCCH_Config_r17_Mcch_ModificationPeriod_r17_Rf65536},
}

type MCCH_Config_r17 struct {
	Mcch_RepetitionPeriodAndOffset_r17 MCCH_RepetitionPeriodAndOffset_r17
	Mcch_WindowStartSlot_r17           int64
	Mcch_WindowDuration_r17            *int64
	Mcch_ModificationPeriod_r17        int64
}

func (ie *MCCH_Config_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mCCHConfigR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Mcch_WindowDuration_r17 != nil}); err != nil {
		return err
	}

	// 2. mcch-RepetitionPeriodAndOffset-r17: ref
	{
		if err := ie.Mcch_RepetitionPeriodAndOffset_r17.Encode(e); err != nil {
			return err
		}
	}

	// 3. mcch-WindowStartSlot-r17: integer
	{
		if err := e.EncodeInteger(ie.Mcch_WindowStartSlot_r17, mCCHConfigR17McchWindowStartSlotR17Constraints); err != nil {
			return err
		}
	}

	// 4. mcch-WindowDuration-r17: enumerated
	{
		if ie.Mcch_WindowDuration_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Mcch_WindowDuration_r17, mCCHConfigR17McchWindowDurationR17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. mcch-ModificationPeriod-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.Mcch_ModificationPeriod_r17, mCCHConfigR17McchModificationPeriodR17Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MCCH_Config_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mCCHConfigR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. mcch-RepetitionPeriodAndOffset-r17: ref
	{
		if err := ie.Mcch_RepetitionPeriodAndOffset_r17.Decode(d); err != nil {
			return err
		}
	}

	// 3. mcch-WindowStartSlot-r17: integer
	{
		v1, err := d.DecodeInteger(mCCHConfigR17McchWindowStartSlotR17Constraints)
		if err != nil {
			return err
		}
		ie.Mcch_WindowStartSlot_r17 = v1
	}

	// 4. mcch-WindowDuration-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(mCCHConfigR17McchWindowDurationR17Constraints)
			if err != nil {
				return err
			}
			ie.Mcch_WindowDuration_r17 = &idx
		}
	}

	// 5. mcch-ModificationPeriod-r17: enumerated
	{
		v3, err := d.DecodeEnumerated(mCCHConfigR17McchModificationPeriodR17Constraints)
		if err != nil {
			return err
		}
		ie.Mcch_ModificationPeriod_r17 = v3
	}

	return nil
}

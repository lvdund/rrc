// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SSB-Configuration-r16 (line 15627).

var sSBConfigurationR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ssb-Freq-r16"},
		{Name: "halfFrameIndex-r16"},
		{Name: "ssbSubcarrierSpacing-r16"},
		{Name: "ssb-Periodicity-r16", Optional: true},
		{Name: "sfn0-Offset-r16", Optional: true},
		{Name: "sfn-SSB-Offset-r16"},
		{Name: "ss-PBCH-BlockPower-r16", Optional: true},
	},
}

const (
	SSB_Configuration_r16_HalfFrameIndex_r16_Zero = 0
	SSB_Configuration_r16_HalfFrameIndex_r16_One  = 1
)

var sSBConfigurationR16HalfFrameIndexR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SSB_Configuration_r16_HalfFrameIndex_r16_Zero, SSB_Configuration_r16_HalfFrameIndex_r16_One},
}

const (
	SSB_Configuration_r16_Ssb_Periodicity_r16_Ms5    = 0
	SSB_Configuration_r16_Ssb_Periodicity_r16_Ms10   = 1
	SSB_Configuration_r16_Ssb_Periodicity_r16_Ms20   = 2
	SSB_Configuration_r16_Ssb_Periodicity_r16_Ms40   = 3
	SSB_Configuration_r16_Ssb_Periodicity_r16_Ms80   = 4
	SSB_Configuration_r16_Ssb_Periodicity_r16_Ms160  = 5
	SSB_Configuration_r16_Ssb_Periodicity_r16_Spare2 = 6
	SSB_Configuration_r16_Ssb_Periodicity_r16_Spare1 = 7
)

var sSBConfigurationR16SsbPeriodicityR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SSB_Configuration_r16_Ssb_Periodicity_r16_Ms5, SSB_Configuration_r16_Ssb_Periodicity_r16_Ms10, SSB_Configuration_r16_Ssb_Periodicity_r16_Ms20, SSB_Configuration_r16_Ssb_Periodicity_r16_Ms40, SSB_Configuration_r16_Ssb_Periodicity_r16_Ms80, SSB_Configuration_r16_Ssb_Periodicity_r16_Ms160, SSB_Configuration_r16_Ssb_Periodicity_r16_Spare2, SSB_Configuration_r16_Ssb_Periodicity_r16_Spare1},
}

var sSBConfigurationR16SfnSSBOffsetR16Constraints = per.Constrained(0, 15)

var sSBConfigurationR16SsPBCHBlockPowerR16Constraints = per.Constrained(-60, 50)

var sSBConfigurationR16Sfn0OffsetR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sfn-Offset-r16"},
		{Name: "integerSubframeOffset-r16", Optional: true},
	},
}

type SSB_Configuration_r16 struct {
	Ssb_Freq_r16             ARFCN_ValueNR
	HalfFrameIndex_r16       int64
	SsbSubcarrierSpacing_r16 SubcarrierSpacing
	Ssb_Periodicity_r16      *int64
	Sfn0_Offset_r16          *struct {
		Sfn_Offset_r16            int64
		IntegerSubframeOffset_r16 *int64
	}
	Sfn_SSB_Offset_r16     int64
	Ss_PBCH_BlockPower_r16 *int64
}

func (ie *SSB_Configuration_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sSBConfigurationR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ssb_Periodicity_r16 != nil, ie.Sfn0_Offset_r16 != nil, ie.Ss_PBCH_BlockPower_r16 != nil}); err != nil {
		return err
	}

	// 2. ssb-Freq-r16: ref
	{
		if err := ie.Ssb_Freq_r16.Encode(e); err != nil {
			return err
		}
	}

	// 3. halfFrameIndex-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.HalfFrameIndex_r16, sSBConfigurationR16HalfFrameIndexR16Constraints); err != nil {
			return err
		}
	}

	// 4. ssbSubcarrierSpacing-r16: ref
	{
		if err := ie.SsbSubcarrierSpacing_r16.Encode(e); err != nil {
			return err
		}
	}

	// 5. ssb-Periodicity-r16: enumerated
	{
		if ie.Ssb_Periodicity_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Ssb_Periodicity_r16, sSBConfigurationR16SsbPeriodicityR16Constraints); err != nil {
				return err
			}
		}
	}

	// 6. sfn0-Offset-r16: sequence
	{
		if ie.Sfn0_Offset_r16 != nil {
			c := ie.Sfn0_Offset_r16
			sSBConfigurationR16Sfn0OffsetR16Seq := e.NewSequenceEncoder(sSBConfigurationR16Sfn0OffsetR16Constraints)
			if err := sSBConfigurationR16Sfn0OffsetR16Seq.EncodePreamble([]bool{c.IntegerSubframeOffset_r16 != nil}); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.Sfn_Offset_r16, per.Constrained(0, 1023)); err != nil {
				return err
			}
			if c.IntegerSubframeOffset_r16 != nil {
				if err := e.EncodeInteger((*c.IntegerSubframeOffset_r16), per.Constrained(0, 9)); err != nil {
					return err
				}
			}
		}
	}

	// 7. sfn-SSB-Offset-r16: integer
	{
		if err := e.EncodeInteger(ie.Sfn_SSB_Offset_r16, sSBConfigurationR16SfnSSBOffsetR16Constraints); err != nil {
			return err
		}
	}

	// 8. ss-PBCH-BlockPower-r16: integer
	{
		if ie.Ss_PBCH_BlockPower_r16 != nil {
			if err := e.EncodeInteger(*ie.Ss_PBCH_BlockPower_r16, sSBConfigurationR16SsPBCHBlockPowerR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SSB_Configuration_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sSBConfigurationR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ssb-Freq-r16: ref
	{
		if err := ie.Ssb_Freq_r16.Decode(d); err != nil {
			return err
		}
	}

	// 3. halfFrameIndex-r16: enumerated
	{
		v1, err := d.DecodeEnumerated(sSBConfigurationR16HalfFrameIndexR16Constraints)
		if err != nil {
			return err
		}
		ie.HalfFrameIndex_r16 = v1
	}

	// 4. ssbSubcarrierSpacing-r16: ref
	{
		if err := ie.SsbSubcarrierSpacing_r16.Decode(d); err != nil {
			return err
		}
	}

	// 5. ssb-Periodicity-r16: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(sSBConfigurationR16SsbPeriodicityR16Constraints)
			if err != nil {
				return err
			}
			ie.Ssb_Periodicity_r16 = &idx
		}
	}

	// 6. sfn0-Offset-r16: sequence
	{
		if seq.IsComponentPresent(4) {
			ie.Sfn0_Offset_r16 = &struct {
				Sfn_Offset_r16            int64
				IntegerSubframeOffset_r16 *int64
			}{}
			c := ie.Sfn0_Offset_r16
			sSBConfigurationR16Sfn0OffsetR16Seq := d.NewSequenceDecoder(sSBConfigurationR16Sfn0OffsetR16Constraints)
			if err := sSBConfigurationR16Sfn0OffsetR16Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				v, err := d.DecodeInteger(per.Constrained(0, 1023))
				if err != nil {
					return err
				}
				c.Sfn_Offset_r16 = v
			}
			if sSBConfigurationR16Sfn0OffsetR16Seq.IsComponentPresent(1) {
				c.IntegerSubframeOffset_r16 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 9))
				if err != nil {
					return err
				}
				(*c.IntegerSubframeOffset_r16) = v
			}
		}
	}

	// 7. sfn-SSB-Offset-r16: integer
	{
		v5, err := d.DecodeInteger(sSBConfigurationR16SfnSSBOffsetR16Constraints)
		if err != nil {
			return err
		}
		ie.Sfn_SSB_Offset_r16 = v5
	}

	// 8. ss-PBCH-BlockPower-r16: integer
	{
		if seq.IsComponentPresent(6) {
			v, err := d.DecodeInteger(sSBConfigurationR16SsPBCHBlockPowerR16Constraints)
			if err != nil {
				return err
			}
			ie.Ss_PBCH_BlockPower_r16 = &v
		}
	}

	return nil
}

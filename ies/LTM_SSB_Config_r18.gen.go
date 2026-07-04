// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: LTM-SSB-Config-r18 (line 8728).

var lTMSSBConfigR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ssb-Frequency-r18"},
		{Name: "subcarrierSpacing-r18"},
		{Name: "ssb-Periodicity-r18", Optional: true},
		{Name: "ssb-PositionsInBurst-r18", Optional: true},
		{Name: "ss-PBCH-BlockPower-r18", Optional: true},
	},
}

const (
	LTM_SSB_Config_r18_Ssb_Periodicity_r18_Ms5    = 0
	LTM_SSB_Config_r18_Ssb_Periodicity_r18_Ms10   = 1
	LTM_SSB_Config_r18_Ssb_Periodicity_r18_Ms20   = 2
	LTM_SSB_Config_r18_Ssb_Periodicity_r18_Ms40   = 3
	LTM_SSB_Config_r18_Ssb_Periodicity_r18_Ms80   = 4
	LTM_SSB_Config_r18_Ssb_Periodicity_r18_Ms160  = 5
	LTM_SSB_Config_r18_Ssb_Periodicity_r18_Spare2 = 6
	LTM_SSB_Config_r18_Ssb_Periodicity_r18_Spare1 = 7
)

var lTMSSBConfigR18SsbPeriodicityR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LTM_SSB_Config_r18_Ssb_Periodicity_r18_Ms5, LTM_SSB_Config_r18_Ssb_Periodicity_r18_Ms10, LTM_SSB_Config_r18_Ssb_Periodicity_r18_Ms20, LTM_SSB_Config_r18_Ssb_Periodicity_r18_Ms40, LTM_SSB_Config_r18_Ssb_Periodicity_r18_Ms80, LTM_SSB_Config_r18_Ssb_Periodicity_r18_Ms160, LTM_SSB_Config_r18_Ssb_Periodicity_r18_Spare2, LTM_SSB_Config_r18_Ssb_Periodicity_r18_Spare1},
}

var lTM_SSB_Config_r18SsbPositionsInBurstR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "shortBitmap"},
		{Name: "mediumBitmap"},
		{Name: "longBitmap"},
	},
}

const (
	LTM_SSB_Config_r18_Ssb_PositionsInBurst_r18_ShortBitmap  = 0
	LTM_SSB_Config_r18_Ssb_PositionsInBurst_r18_MediumBitmap = 1
	LTM_SSB_Config_r18_Ssb_PositionsInBurst_r18_LongBitmap   = 2
)

var lTMSSBConfigR18SsPBCHBlockPowerR18Constraints = per.Constrained(-60, 50)

type LTM_SSB_Config_r18 struct {
	Ssb_Frequency_r18        ARFCN_ValueNR
	SubcarrierSpacing_r18    SubcarrierSpacing
	Ssb_Periodicity_r18      *int64
	Ssb_PositionsInBurst_r18 *struct {
		Choice       int
		ShortBitmap  *per.BitString
		MediumBitmap *per.BitString
		LongBitmap   *per.BitString
	}
	Ss_PBCH_BlockPower_r18 *int64
}

func (ie *LTM_SSB_Config_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(lTMSSBConfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ssb_Periodicity_r18 != nil, ie.Ssb_PositionsInBurst_r18 != nil, ie.Ss_PBCH_BlockPower_r18 != nil}); err != nil {
		return err
	}

	// 3. ssb-Frequency-r18: ref
	{
		if err := ie.Ssb_Frequency_r18.Encode(e); err != nil {
			return err
		}
	}

	// 4. subcarrierSpacing-r18: ref
	{
		if err := ie.SubcarrierSpacing_r18.Encode(e); err != nil {
			return err
		}
	}

	// 5. ssb-Periodicity-r18: enumerated
	{
		if ie.Ssb_Periodicity_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Ssb_Periodicity_r18, lTMSSBConfigR18SsbPeriodicityR18Constraints); err != nil {
				return err
			}
		}
	}

	// 6. ssb-PositionsInBurst-r18: choice
	{
		if ie.Ssb_PositionsInBurst_r18 != nil {
			choiceEnc := e.NewChoiceEncoder(lTM_SSB_Config_r18SsbPositionsInBurstR18Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Ssb_PositionsInBurst_r18).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Ssb_PositionsInBurst_r18).Choice {
			case LTM_SSB_Config_r18_Ssb_PositionsInBurst_r18_ShortBitmap:
				if err := e.EncodeBitString((*(*ie.Ssb_PositionsInBurst_r18).ShortBitmap), per.FixedSize(4)); err != nil {
					return err
				}
			case LTM_SSB_Config_r18_Ssb_PositionsInBurst_r18_MediumBitmap:
				if err := e.EncodeBitString((*(*ie.Ssb_PositionsInBurst_r18).MediumBitmap), per.FixedSize(8)); err != nil {
					return err
				}
			case LTM_SSB_Config_r18_Ssb_PositionsInBurst_r18_LongBitmap:
				if err := e.EncodeBitString((*(*ie.Ssb_PositionsInBurst_r18).LongBitmap), per.FixedSize(64)); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Ssb_PositionsInBurst_r18).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 7. ss-PBCH-BlockPower-r18: integer
	{
		if ie.Ss_PBCH_BlockPower_r18 != nil {
			if err := e.EncodeInteger(*ie.Ss_PBCH_BlockPower_r18, lTMSSBConfigR18SsPBCHBlockPowerR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *LTM_SSB_Config_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(lTMSSBConfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. ssb-Frequency-r18: ref
	{
		if err := ie.Ssb_Frequency_r18.Decode(d); err != nil {
			return err
		}
	}

	// 4. subcarrierSpacing-r18: ref
	{
		if err := ie.SubcarrierSpacing_r18.Decode(d); err != nil {
			return err
		}
	}

	// 5. ssb-Periodicity-r18: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(lTMSSBConfigR18SsbPeriodicityR18Constraints)
			if err != nil {
				return err
			}
			ie.Ssb_Periodicity_r18 = &idx
		}
	}

	// 6. ssb-PositionsInBurst-r18: choice
	{
		if seq.IsComponentPresent(3) {
			ie.Ssb_PositionsInBurst_r18 = &struct {
				Choice       int
				ShortBitmap  *per.BitString
				MediumBitmap *per.BitString
				LongBitmap   *per.BitString
			}{}
			choiceDec := d.NewChoiceDecoder(lTM_SSB_Config_r18SsbPositionsInBurstR18Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Ssb_PositionsInBurst_r18).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case LTM_SSB_Config_r18_Ssb_PositionsInBurst_r18_ShortBitmap:
				(*ie.Ssb_PositionsInBurst_r18).ShortBitmap = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(4))
				if err != nil {
					return err
				}
				(*(*ie.Ssb_PositionsInBurst_r18).ShortBitmap) = v
			case LTM_SSB_Config_r18_Ssb_PositionsInBurst_r18_MediumBitmap:
				(*ie.Ssb_PositionsInBurst_r18).MediumBitmap = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(8))
				if err != nil {
					return err
				}
				(*(*ie.Ssb_PositionsInBurst_r18).MediumBitmap) = v
			case LTM_SSB_Config_r18_Ssb_PositionsInBurst_r18_LongBitmap:
				(*ie.Ssb_PositionsInBurst_r18).LongBitmap = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(64))
				if err != nil {
					return err
				}
				(*(*ie.Ssb_PositionsInBurst_r18).LongBitmap) = v
			}
		}
	}

	// 7. ss-PBCH-BlockPower-r18: integer
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeInteger(lTMSSBConfigR18SsPBCHBlockPowerR18Constraints)
			if err != nil {
				return err
			}
			ie.Ss_PBCH_BlockPower_r18 = &v
		}
	}

	return nil
}

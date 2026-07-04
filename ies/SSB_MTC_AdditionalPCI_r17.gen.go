// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SSB-MTC-AdditionalPCI-r17 (line 15901).

var sSBMTCAdditionalPCIR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "additionalPCIIndex-r17"},
		{Name: "additionalPCI-r17"},
		{Name: "periodicity-r17"},
		{Name: "ssb-PositionsInBurst-r17"},
		{Name: "ss-PBCH-BlockPower-r17"},
	},
}

const (
	SSB_MTC_AdditionalPCI_r17_Periodicity_r17_Ms5    = 0
	SSB_MTC_AdditionalPCI_r17_Periodicity_r17_Ms10   = 1
	SSB_MTC_AdditionalPCI_r17_Periodicity_r17_Ms20   = 2
	SSB_MTC_AdditionalPCI_r17_Periodicity_r17_Ms40   = 3
	SSB_MTC_AdditionalPCI_r17_Periodicity_r17_Ms80   = 4
	SSB_MTC_AdditionalPCI_r17_Periodicity_r17_Ms160  = 5
	SSB_MTC_AdditionalPCI_r17_Periodicity_r17_Spare2 = 6
	SSB_MTC_AdditionalPCI_r17_Periodicity_r17_Spare1 = 7
)

var sSBMTCAdditionalPCIR17PeriodicityR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SSB_MTC_AdditionalPCI_r17_Periodicity_r17_Ms5, SSB_MTC_AdditionalPCI_r17_Periodicity_r17_Ms10, SSB_MTC_AdditionalPCI_r17_Periodicity_r17_Ms20, SSB_MTC_AdditionalPCI_r17_Periodicity_r17_Ms40, SSB_MTC_AdditionalPCI_r17_Periodicity_r17_Ms80, SSB_MTC_AdditionalPCI_r17_Periodicity_r17_Ms160, SSB_MTC_AdditionalPCI_r17_Periodicity_r17_Spare2, SSB_MTC_AdditionalPCI_r17_Periodicity_r17_Spare1},
}

var sSB_MTC_AdditionalPCI_r17SsbPositionsInBurstR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "shortBitmap"},
		{Name: "mediumBitmap"},
		{Name: "longBitmap"},
	},
}

const (
	SSB_MTC_AdditionalPCI_r17_Ssb_PositionsInBurst_r17_ShortBitmap  = 0
	SSB_MTC_AdditionalPCI_r17_Ssb_PositionsInBurst_r17_MediumBitmap = 1
	SSB_MTC_AdditionalPCI_r17_Ssb_PositionsInBurst_r17_LongBitmap   = 2
)

var sSBMTCAdditionalPCIR17SsPBCHBlockPowerR17Constraints = per.Constrained(-60, 50)

type SSB_MTC_AdditionalPCI_r17 struct {
	AdditionalPCIIndex_r17   AdditionalPCIIndex_r17
	AdditionalPCI_r17        PhysCellId
	Periodicity_r17          int64
	Ssb_PositionsInBurst_r17 struct {
		Choice       int
		ShortBitmap  *per.BitString
		MediumBitmap *per.BitString
		LongBitmap   *per.BitString
	}
	Ss_PBCH_BlockPower_r17 int64
}

func (ie *SSB_MTC_AdditionalPCI_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sSBMTCAdditionalPCIR17Constraints)
	_ = seq

	// 1. additionalPCIIndex-r17: ref
	{
		if err := ie.AdditionalPCIIndex_r17.Encode(e); err != nil {
			return err
		}
	}

	// 2. additionalPCI-r17: ref
	{
		if err := ie.AdditionalPCI_r17.Encode(e); err != nil {
			return err
		}
	}

	// 3. periodicity-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.Periodicity_r17, sSBMTCAdditionalPCIR17PeriodicityR17Constraints); err != nil {
			return err
		}
	}

	// 4. ssb-PositionsInBurst-r17: choice
	{
		choiceEnc := e.NewChoiceEncoder(sSB_MTC_AdditionalPCI_r17SsbPositionsInBurstR17Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.Ssb_PositionsInBurst_r17.Choice), false, nil); err != nil {
			return err
		}
		switch ie.Ssb_PositionsInBurst_r17.Choice {
		case SSB_MTC_AdditionalPCI_r17_Ssb_PositionsInBurst_r17_ShortBitmap:
			if err := e.EncodeBitString((*ie.Ssb_PositionsInBurst_r17.ShortBitmap), per.FixedSize(4)); err != nil {
				return err
			}
		case SSB_MTC_AdditionalPCI_r17_Ssb_PositionsInBurst_r17_MediumBitmap:
			if err := e.EncodeBitString((*ie.Ssb_PositionsInBurst_r17.MediumBitmap), per.FixedSize(8)); err != nil {
				return err
			}
		case SSB_MTC_AdditionalPCI_r17_Ssb_PositionsInBurst_r17_LongBitmap:
			if err := e.EncodeBitString((*ie.Ssb_PositionsInBurst_r17.LongBitmap), per.FixedSize(64)); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.Ssb_PositionsInBurst_r17.Choice), Constraint: "index out of range"}
		}
	}

	// 5. ss-PBCH-BlockPower-r17: integer
	{
		if err := e.EncodeInteger(ie.Ss_PBCH_BlockPower_r17, sSBMTCAdditionalPCIR17SsPBCHBlockPowerR17Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SSB_MTC_AdditionalPCI_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sSBMTCAdditionalPCIR17Constraints)
	_ = seq

	// 1. additionalPCIIndex-r17: ref
	{
		if err := ie.AdditionalPCIIndex_r17.Decode(d); err != nil {
			return err
		}
	}

	// 2. additionalPCI-r17: ref
	{
		if err := ie.AdditionalPCI_r17.Decode(d); err != nil {
			return err
		}
	}

	// 3. periodicity-r17: enumerated
	{
		v2, err := d.DecodeEnumerated(sSBMTCAdditionalPCIR17PeriodicityR17Constraints)
		if err != nil {
			return err
		}
		ie.Periodicity_r17 = v2
	}

	// 4. ssb-PositionsInBurst-r17: choice
	{
		choiceDec := d.NewChoiceDecoder(sSB_MTC_AdditionalPCI_r17SsbPositionsInBurstR17Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.Ssb_PositionsInBurst_r17.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case SSB_MTC_AdditionalPCI_r17_Ssb_PositionsInBurst_r17_ShortBitmap:
			ie.Ssb_PositionsInBurst_r17.ShortBitmap = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(4))
			if err != nil {
				return err
			}
			(*ie.Ssb_PositionsInBurst_r17.ShortBitmap) = v
		case SSB_MTC_AdditionalPCI_r17_Ssb_PositionsInBurst_r17_MediumBitmap:
			ie.Ssb_PositionsInBurst_r17.MediumBitmap = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(8))
			if err != nil {
				return err
			}
			(*ie.Ssb_PositionsInBurst_r17.MediumBitmap) = v
		case SSB_MTC_AdditionalPCI_r17_Ssb_PositionsInBurst_r17_LongBitmap:
			ie.Ssb_PositionsInBurst_r17.LongBitmap = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(64))
			if err != nil {
				return err
			}
			(*ie.Ssb_PositionsInBurst_r17.LongBitmap) = v
		}
	}

	// 5. ss-PBCH-BlockPower-r17: integer
	{
		v4, err := d.DecodeInteger(sSBMTCAdditionalPCIR17SsPBCHBlockPowerR17Constraints)
		if err != nil {
			return err
		}
		ie.Ss_PBCH_BlockPower_r17 = v4
	}

	return nil
}

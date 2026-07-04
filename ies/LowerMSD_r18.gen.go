// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: LowerMSD-r18 (line 24672).

var lowerMSDR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "aggressorband1-r18"},
		{Name: "aggressorband2-r18", Optional: true},
		{Name: "msd-Information-r18"},
	},
}

var lowerMSD_r18Aggressorband1R18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "nr"},
		{Name: "eutra"},
	},
}

const (
	LowerMSD_r18_Aggressorband1_r18_Nr    = 0
	LowerMSD_r18_Aggressorband1_r18_Eutra = 1
)

var lowerMSDR18MsdInformationR18Constraints = per.SizeRange(1, common.MaxLowerMSDInfo_r18)

type LowerMSD_r18 struct {
	Aggressorband1_r18 struct {
		Choice int
		Nr     *FreqBandIndicatorNR
		Eutra  *FreqBandIndicatorEUTRA
	}
	Aggressorband2_r18  *FreqBandIndicatorNR
	Msd_Information_r18 []MSD_Information_r18
}

func (ie *LowerMSD_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(lowerMSDR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Aggressorband2_r18 != nil}); err != nil {
		return err
	}

	// 2. aggressorband1-r18: choice
	{
		choiceEnc := e.NewChoiceEncoder(lowerMSD_r18Aggressorband1R18Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.Aggressorband1_r18.Choice), false, nil); err != nil {
			return err
		}
		switch ie.Aggressorband1_r18.Choice {
		case LowerMSD_r18_Aggressorband1_r18_Nr:
			if err := ie.Aggressorband1_r18.Nr.Encode(e); err != nil {
				return err
			}
		case LowerMSD_r18_Aggressorband1_r18_Eutra:
			if err := ie.Aggressorband1_r18.Eutra.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.Aggressorband1_r18.Choice), Constraint: "index out of range"}
		}
	}

	// 3. aggressorband2-r18: ref
	{
		if ie.Aggressorband2_r18 != nil {
			if err := ie.Aggressorband2_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. msd-Information-r18: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(lowerMSDR18MsdInformationR18Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Msd_Information_r18))); err != nil {
			return err
		}
		for i := range ie.Msd_Information_r18 {
			if err := ie.Msd_Information_r18[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *LowerMSD_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(lowerMSDR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. aggressorband1-r18: choice
	{
		choiceDec := d.NewChoiceDecoder(lowerMSD_r18Aggressorband1R18Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.Aggressorband1_r18.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case LowerMSD_r18_Aggressorband1_r18_Nr:
			ie.Aggressorband1_r18.Nr = new(FreqBandIndicatorNR)
			if err := ie.Aggressorband1_r18.Nr.Decode(d); err != nil {
				return err
			}
		case LowerMSD_r18_Aggressorband1_r18_Eutra:
			ie.Aggressorband1_r18.Eutra = new(FreqBandIndicatorEUTRA)
			if err := ie.Aggressorband1_r18.Eutra.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. aggressorband2-r18: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Aggressorband2_r18 = new(FreqBandIndicatorNR)
			if err := ie.Aggressorband2_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. msd-Information-r18: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(lowerMSDR18MsdInformationR18Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Msd_Information_r18 = make([]MSD_Information_r18, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Msd_Information_r18[i].Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}

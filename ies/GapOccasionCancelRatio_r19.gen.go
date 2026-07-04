// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: GapOccasionCancelRatio-r19 (line 2806).

var gapOccasionCancelRatioR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "gapOccasionCancelRatioPerFR-PerUE-r19", Optional: true},
		{Name: "gapConfigRatioList-r19", Optional: true},
	},
}

var gapOccasionCancelRatio_r19GapOccasionCancelRatioPerFRPerUER19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "perUE-r19"},
		{Name: "perFR-r19"},
	},
}

const (
	GapOccasionCancelRatio_r19_GapOccasionCancelRatioPerFR_PerUE_r19_PerUE_r19 = 0
	GapOccasionCancelRatio_r19_GapOccasionCancelRatioPerFR_PerUE_r19_PerFR_r19 = 1
)

var gapOccasionCancelRatioR19GapConfigRatioListR19Constraints = per.SizeRange(1, common.MaxNrofGapId_r17)

var gapOccasionCancelRatioR19GapOccasionCancelRatioPerFRPerUER19PerFRR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "fr1-r19", Optional: true},
		{Name: "fr2-r19", Optional: true},
	},
}

type GapOccasionCancelRatio_r19 struct {
	GapOccasionCancelRatioPerFR_PerUE_r19 *struct {
		Choice    int
		PerUE_r19 *GapOccasionRatio_r19
		PerFR_r19 *struct {
			Fr1_r19 *GapOccasionRatio_r19
			Fr2_r19 *GapOccasionRatio_r19
		}
	}
	GapConfigRatioList_r19 []GapOccasionRatioPerGapConfig_r19
}

func (ie *GapOccasionCancelRatio_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(gapOccasionCancelRatioR19Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.GapOccasionCancelRatioPerFR_PerUE_r19 != nil, ie.GapConfigRatioList_r19 != nil}); err != nil {
		return err
	}

	// 2. gapOccasionCancelRatioPerFR-PerUE-r19: choice
	{
		if ie.GapOccasionCancelRatioPerFR_PerUE_r19 != nil {
			choiceEnc := e.NewChoiceEncoder(gapOccasionCancelRatio_r19GapOccasionCancelRatioPerFRPerUER19Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.GapOccasionCancelRatioPerFR_PerUE_r19).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.GapOccasionCancelRatioPerFR_PerUE_r19).Choice {
			case GapOccasionCancelRatio_r19_GapOccasionCancelRatioPerFR_PerUE_r19_PerUE_r19:
				if err := (*ie.GapOccasionCancelRatioPerFR_PerUE_r19).PerUE_r19.Encode(e); err != nil {
					return err
				}
			case GapOccasionCancelRatio_r19_GapOccasionCancelRatioPerFR_PerUE_r19_PerFR_r19:
				c := (*(*ie.GapOccasionCancelRatioPerFR_PerUE_r19).PerFR_r19)
				gapOccasionCancelRatioR19GapOccasionCancelRatioPerFRPerUER19PerFRR19Seq := e.NewSequenceEncoder(gapOccasionCancelRatioR19GapOccasionCancelRatioPerFRPerUER19PerFRR19Constraints)
				if err := gapOccasionCancelRatioR19GapOccasionCancelRatioPerFRPerUER19PerFRR19Seq.EncodePreamble([]bool{c.Fr1_r19 != nil, c.Fr2_r19 != nil}); err != nil {
					return err
				}
				if c.Fr1_r19 != nil {
					if err := c.Fr1_r19.Encode(e); err != nil {
						return err
					}
				}
				if c.Fr2_r19 != nil {
					if err := c.Fr2_r19.Encode(e); err != nil {
						return err
					}
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.GapOccasionCancelRatioPerFR_PerUE_r19).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 3. gapConfigRatioList-r19: sequence-of
	{
		if ie.GapConfigRatioList_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(gapOccasionCancelRatioR19GapConfigRatioListR19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.GapConfigRatioList_r19))); err != nil {
				return err
			}
			for i := range ie.GapConfigRatioList_r19 {
				if err := ie.GapConfigRatioList_r19[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *GapOccasionCancelRatio_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(gapOccasionCancelRatioR19Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. gapOccasionCancelRatioPerFR-PerUE-r19: choice
	{
		if seq.IsComponentPresent(0) {
			ie.GapOccasionCancelRatioPerFR_PerUE_r19 = &struct {
				Choice    int
				PerUE_r19 *GapOccasionRatio_r19
				PerFR_r19 *struct {
					Fr1_r19 *GapOccasionRatio_r19
					Fr2_r19 *GapOccasionRatio_r19
				}
			}{}
			choiceDec := d.NewChoiceDecoder(gapOccasionCancelRatio_r19GapOccasionCancelRatioPerFRPerUER19Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.GapOccasionCancelRatioPerFR_PerUE_r19).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case GapOccasionCancelRatio_r19_GapOccasionCancelRatioPerFR_PerUE_r19_PerUE_r19:
				(*ie.GapOccasionCancelRatioPerFR_PerUE_r19).PerUE_r19 = new(GapOccasionRatio_r19)
				if err := (*ie.GapOccasionCancelRatioPerFR_PerUE_r19).PerUE_r19.Decode(d); err != nil {
					return err
				}
			case GapOccasionCancelRatio_r19_GapOccasionCancelRatioPerFR_PerUE_r19_PerFR_r19:
				(*ie.GapOccasionCancelRatioPerFR_PerUE_r19).PerFR_r19 = &struct {
					Fr1_r19 *GapOccasionRatio_r19
					Fr2_r19 *GapOccasionRatio_r19
				}{}
				c := (*(*ie.GapOccasionCancelRatioPerFR_PerUE_r19).PerFR_r19)
				gapOccasionCancelRatioR19GapOccasionCancelRatioPerFRPerUER19PerFRR19Seq := d.NewSequenceDecoder(gapOccasionCancelRatioR19GapOccasionCancelRatioPerFRPerUER19PerFRR19Constraints)
				if err := gapOccasionCancelRatioR19GapOccasionCancelRatioPerFRPerUER19PerFRR19Seq.DecodePreamble(); err != nil {
					return err
				}
				if gapOccasionCancelRatioR19GapOccasionCancelRatioPerFRPerUER19PerFRR19Seq.IsComponentPresent(0) {
					c.Fr1_r19 = new(GapOccasionRatio_r19)
					if err := (*c.Fr1_r19).Decode(d); err != nil {
						return err
					}
				}
				if gapOccasionCancelRatioR19GapOccasionCancelRatioPerFRPerUER19PerFRR19Seq.IsComponentPresent(1) {
					c.Fr2_r19 = new(GapOccasionRatio_r19)
					if err := (*c.Fr2_r19).Decode(d); err != nil {
						return err
					}
				}
			}
		}
	}

	// 3. gapConfigRatioList-r19: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(gapOccasionCancelRatioR19GapConfigRatioListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.GapConfigRatioList_r19 = make([]GapOccasionRatioPerGapConfig_r19, n)
			for i := int64(0); i < n; i++ {
				if err := ie.GapConfigRatioList_r19[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

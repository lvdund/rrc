// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SI-RequestConfigRepetition-r18 (line 15033).

var sIRequestConfigRepetitionR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "rach-OccasionsSI-r18", Optional: true},
		{Name: "si-RequestResourcesRepetitionNum2-r18", Optional: true},
		{Name: "si-RequestResourcesRepetitionNum4-r18", Optional: true},
		{Name: "si-RequestResourcesRepetitionNum8-r18", Optional: true},
	},
}

var sIRequestConfigRepetitionR18SiRequestResourcesRepetitionNum2R18Constraints = per.SizeRange(1, common.MaxSI_Message)

var sIRequestConfigRepetitionR18SiRequestResourcesRepetitionNum4R18Constraints = per.SizeRange(1, common.MaxSI_Message)

var sIRequestConfigRepetitionR18SiRequestResourcesRepetitionNum8R18Constraints = per.SizeRange(1, common.MaxSI_Message)

const (
	SI_RequestConfigRepetition_r18_Rach_OccasionsSI_r18_Ssb_PerRACH_Occasion_r18_OneEighth = 0
	SI_RequestConfigRepetition_r18_Rach_OccasionsSI_r18_Ssb_PerRACH_Occasion_r18_OneFourth = 1
	SI_RequestConfigRepetition_r18_Rach_OccasionsSI_r18_Ssb_PerRACH_Occasion_r18_OneHalf   = 2
	SI_RequestConfigRepetition_r18_Rach_OccasionsSI_r18_Ssb_PerRACH_Occasion_r18_One       = 3
	SI_RequestConfigRepetition_r18_Rach_OccasionsSI_r18_Ssb_PerRACH_Occasion_r18_Two       = 4
	SI_RequestConfigRepetition_r18_Rach_OccasionsSI_r18_Ssb_PerRACH_Occasion_r18_Four      = 5
	SI_RequestConfigRepetition_r18_Rach_OccasionsSI_r18_Ssb_PerRACH_Occasion_r18_Eight     = 6
	SI_RequestConfigRepetition_r18_Rach_OccasionsSI_r18_Ssb_PerRACH_Occasion_r18_Sixteen   = 7
)

var sIRequestConfigRepetitionR18RachOccasionsSIR18SsbPerRACHOccasionR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SI_RequestConfigRepetition_r18_Rach_OccasionsSI_r18_Ssb_PerRACH_Occasion_r18_OneEighth, SI_RequestConfigRepetition_r18_Rach_OccasionsSI_r18_Ssb_PerRACH_Occasion_r18_OneFourth, SI_RequestConfigRepetition_r18_Rach_OccasionsSI_r18_Ssb_PerRACH_Occasion_r18_OneHalf, SI_RequestConfigRepetition_r18_Rach_OccasionsSI_r18_Ssb_PerRACH_Occasion_r18_One, SI_RequestConfigRepetition_r18_Rach_OccasionsSI_r18_Ssb_PerRACH_Occasion_r18_Two, SI_RequestConfigRepetition_r18_Rach_OccasionsSI_r18_Ssb_PerRACH_Occasion_r18_Four, SI_RequestConfigRepetition_r18_Rach_OccasionsSI_r18_Ssb_PerRACH_Occasion_r18_Eight, SI_RequestConfigRepetition_r18_Rach_OccasionsSI_r18_Ssb_PerRACH_Occasion_r18_Sixteen},
}

type SI_RequestConfigRepetition_r18 struct {
	Rach_OccasionsSI_r18 *struct {
		Rach_ConfigSI_r18        RACH_ConfigGeneric
		Ssb_PerRACH_Occasion_r18 int64
	}
	Si_RequestResourcesRepetitionNum2_r18 []SI_RequestResourcesRepetition_r18
	Si_RequestResourcesRepetitionNum4_r18 []SI_RequestResourcesRepetition_r18
	Si_RequestResourcesRepetitionNum8_r18 []SI_RequestResourcesRepetition_r18
}

func (ie *SI_RequestConfigRepetition_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIRequestConfigRepetitionR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Rach_OccasionsSI_r18 != nil, ie.Si_RequestResourcesRepetitionNum2_r18 != nil, ie.Si_RequestResourcesRepetitionNum4_r18 != nil, ie.Si_RequestResourcesRepetitionNum8_r18 != nil}); err != nil {
		return err
	}

	// 3. rach-OccasionsSI-r18: sequence
	{
		if ie.Rach_OccasionsSI_r18 != nil {
			c := ie.Rach_OccasionsSI_r18
			if err := c.Rach_ConfigSI_r18.Encode(e); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.Ssb_PerRACH_Occasion_r18, sIRequestConfigRepetitionR18RachOccasionsSIR18SsbPerRACHOccasionR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. si-RequestResourcesRepetitionNum2-r18: sequence-of
	{
		if ie.Si_RequestResourcesRepetitionNum2_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(sIRequestConfigRepetitionR18SiRequestResourcesRepetitionNum2R18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Si_RequestResourcesRepetitionNum2_r18))); err != nil {
				return err
			}
			for i := range ie.Si_RequestResourcesRepetitionNum2_r18 {
				if err := ie.Si_RequestResourcesRepetitionNum2_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. si-RequestResourcesRepetitionNum4-r18: sequence-of
	{
		if ie.Si_RequestResourcesRepetitionNum4_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(sIRequestConfigRepetitionR18SiRequestResourcesRepetitionNum4R18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Si_RequestResourcesRepetitionNum4_r18))); err != nil {
				return err
			}
			for i := range ie.Si_RequestResourcesRepetitionNum4_r18 {
				if err := ie.Si_RequestResourcesRepetitionNum4_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 6. si-RequestResourcesRepetitionNum8-r18: sequence-of
	{
		if ie.Si_RequestResourcesRepetitionNum8_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(sIRequestConfigRepetitionR18SiRequestResourcesRepetitionNum8R18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Si_RequestResourcesRepetitionNum8_r18))); err != nil {
				return err
			}
			for i := range ie.Si_RequestResourcesRepetitionNum8_r18 {
				if err := ie.Si_RequestResourcesRepetitionNum8_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *SI_RequestConfigRepetition_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIRequestConfigRepetitionR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. rach-OccasionsSI-r18: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.Rach_OccasionsSI_r18 = &struct {
				Rach_ConfigSI_r18        RACH_ConfigGeneric
				Ssb_PerRACH_Occasion_r18 int64
			}{}
			c := ie.Rach_OccasionsSI_r18
			{
				if err := c.Rach_ConfigSI_r18.Decode(d); err != nil {
					return err
				}
			}
			{
				v, err := d.DecodeEnumerated(sIRequestConfigRepetitionR18RachOccasionsSIR18SsbPerRACHOccasionR18Constraints)
				if err != nil {
					return err
				}
				c.Ssb_PerRACH_Occasion_r18 = v
			}
		}
	}

	// 4. si-RequestResourcesRepetitionNum2-r18: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(sIRequestConfigRepetitionR18SiRequestResourcesRepetitionNum2R18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Si_RequestResourcesRepetitionNum2_r18 = make([]SI_RequestResourcesRepetition_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Si_RequestResourcesRepetitionNum2_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. si-RequestResourcesRepetitionNum4-r18: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(sIRequestConfigRepetitionR18SiRequestResourcesRepetitionNum4R18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Si_RequestResourcesRepetitionNum4_r18 = make([]SI_RequestResourcesRepetition_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Si_RequestResourcesRepetitionNum4_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. si-RequestResourcesRepetitionNum8-r18: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(sIRequestConfigRepetitionR18SiRequestResourcesRepetitionNum8R18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Si_RequestResourcesRepetitionNum8_r18 = make([]SI_RequestResourcesRepetition_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Si_RequestResourcesRepetitionNum8_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

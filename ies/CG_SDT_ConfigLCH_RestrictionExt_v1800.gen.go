// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CG-SDT-ConfigLCH-RestrictionExt-v1800 (line 1425).

var cGSDTConfigLCHRestrictionExtV1800Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "cg-SDT-MaxDurationToNextCG-Occasion-r18", Optional: true},
	},
}

const (
	CG_SDT_ConfigLCH_RestrictionExt_v1800_Cg_SDT_MaxDurationToNextCG_Occasion_r18_Ms10    = 0
	CG_SDT_ConfigLCH_RestrictionExt_v1800_Cg_SDT_MaxDurationToNextCG_Occasion_r18_Ms100   = 1
	CG_SDT_ConfigLCH_RestrictionExt_v1800_Cg_SDT_MaxDurationToNextCG_Occasion_r18_Sec1    = 2
	CG_SDT_ConfigLCH_RestrictionExt_v1800_Cg_SDT_MaxDurationToNextCG_Occasion_r18_Sec10   = 3
	CG_SDT_ConfigLCH_RestrictionExt_v1800_Cg_SDT_MaxDurationToNextCG_Occasion_r18_Sec60   = 4
	CG_SDT_ConfigLCH_RestrictionExt_v1800_Cg_SDT_MaxDurationToNextCG_Occasion_r18_Sec100  = 5
	CG_SDT_ConfigLCH_RestrictionExt_v1800_Cg_SDT_MaxDurationToNextCG_Occasion_r18_Sec300  = 6
	CG_SDT_ConfigLCH_RestrictionExt_v1800_Cg_SDT_MaxDurationToNextCG_Occasion_r18_Sec600  = 7
	CG_SDT_ConfigLCH_RestrictionExt_v1800_Cg_SDT_MaxDurationToNextCG_Occasion_r18_Sec1200 = 8
	CG_SDT_ConfigLCH_RestrictionExt_v1800_Cg_SDT_MaxDurationToNextCG_Occasion_r18_Sec1800 = 9
	CG_SDT_ConfigLCH_RestrictionExt_v1800_Cg_SDT_MaxDurationToNextCG_Occasion_r18_Sec3600 = 10
	CG_SDT_ConfigLCH_RestrictionExt_v1800_Cg_SDT_MaxDurationToNextCG_Occasion_r18_Spare5  = 11
	CG_SDT_ConfigLCH_RestrictionExt_v1800_Cg_SDT_MaxDurationToNextCG_Occasion_r18_Spare4  = 12
	CG_SDT_ConfigLCH_RestrictionExt_v1800_Cg_SDT_MaxDurationToNextCG_Occasion_r18_Spare3  = 13
	CG_SDT_ConfigLCH_RestrictionExt_v1800_Cg_SDT_MaxDurationToNextCG_Occasion_r18_Spare2  = 14
	CG_SDT_ConfigLCH_RestrictionExt_v1800_Cg_SDT_MaxDurationToNextCG_Occasion_r18_Spare1  = 15
)

var cGSDTConfigLCHRestrictionExtV1800CgSDTMaxDurationToNextCGOccasionR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CG_SDT_ConfigLCH_RestrictionExt_v1800_Cg_SDT_MaxDurationToNextCG_Occasion_r18_Ms10, CG_SDT_ConfigLCH_RestrictionExt_v1800_Cg_SDT_MaxDurationToNextCG_Occasion_r18_Ms100, CG_SDT_ConfigLCH_RestrictionExt_v1800_Cg_SDT_MaxDurationToNextCG_Occasion_r18_Sec1, CG_SDT_ConfigLCH_RestrictionExt_v1800_Cg_SDT_MaxDurationToNextCG_Occasion_r18_Sec10, CG_SDT_ConfigLCH_RestrictionExt_v1800_Cg_SDT_MaxDurationToNextCG_Occasion_r18_Sec60, CG_SDT_ConfigLCH_RestrictionExt_v1800_Cg_SDT_MaxDurationToNextCG_Occasion_r18_Sec100, CG_SDT_ConfigLCH_RestrictionExt_v1800_Cg_SDT_MaxDurationToNextCG_Occasion_r18_Sec300, CG_SDT_ConfigLCH_RestrictionExt_v1800_Cg_SDT_MaxDurationToNextCG_Occasion_r18_Sec600, CG_SDT_ConfigLCH_RestrictionExt_v1800_Cg_SDT_MaxDurationToNextCG_Occasion_r18_Sec1200, CG_SDT_ConfigLCH_RestrictionExt_v1800_Cg_SDT_MaxDurationToNextCG_Occasion_r18_Sec1800, CG_SDT_ConfigLCH_RestrictionExt_v1800_Cg_SDT_MaxDurationToNextCG_Occasion_r18_Sec3600, CG_SDT_ConfigLCH_RestrictionExt_v1800_Cg_SDT_MaxDurationToNextCG_Occasion_r18_Spare5, CG_SDT_ConfigLCH_RestrictionExt_v1800_Cg_SDT_MaxDurationToNextCG_Occasion_r18_Spare4, CG_SDT_ConfigLCH_RestrictionExt_v1800_Cg_SDT_MaxDurationToNextCG_Occasion_r18_Spare3, CG_SDT_ConfigLCH_RestrictionExt_v1800_Cg_SDT_MaxDurationToNextCG_Occasion_r18_Spare2, CG_SDT_ConfigLCH_RestrictionExt_v1800_Cg_SDT_MaxDurationToNextCG_Occasion_r18_Spare1},
}

type CG_SDT_ConfigLCH_RestrictionExt_v1800 struct {
	Cg_SDT_MaxDurationToNextCG_Occasion_r18 *int64
}

func (ie *CG_SDT_ConfigLCH_RestrictionExt_v1800) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cGSDTConfigLCHRestrictionExtV1800Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Cg_SDT_MaxDurationToNextCG_Occasion_r18 != nil}); err != nil {
		return err
	}

	// 2. cg-SDT-MaxDurationToNextCG-Occasion-r18: enumerated
	{
		if ie.Cg_SDT_MaxDurationToNextCG_Occasion_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Cg_SDT_MaxDurationToNextCG_Occasion_r18, cGSDTConfigLCHRestrictionExtV1800CgSDTMaxDurationToNextCGOccasionR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CG_SDT_ConfigLCH_RestrictionExt_v1800) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cGSDTConfigLCHRestrictionExtV1800Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. cg-SDT-MaxDurationToNextCG-Occasion-r18: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(cGSDTConfigLCHRestrictionExtV1800CgSDTMaxDurationToNextCGOccasionR18Constraints)
			if err != nil {
				return err
			}
			ie.Cg_SDT_MaxDurationToNextCG_Occasion_r18 = &idx
		}
	}

	return nil
}

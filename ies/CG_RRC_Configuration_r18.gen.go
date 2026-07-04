// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CG-RRC-Configuration-r18 (line 6721).

var cGRRCConfigurationR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "cg-RRC-RetransmissionTimer-r18", Optional: true},
		{Name: "cg-RRC-RSRP-ThresholdSSB-r18", Optional: true},
		{Name: "rrc-SSB-Subset-r18", Optional: true},
		{Name: "rrc-SSB-PerCG-PUSCH-r18", Optional: true},
		{Name: "rrc-P0-PUSCH-r18", Optional: true},
		{Name: "rrc-Alpha-r18", Optional: true},
		{Name: "rrc-DMRS-Ports-r18", Optional: true},
		{Name: "rrc-NrofDMRS-Sequences-r18", Optional: true},
	},
}

var cGRRCConfigurationR18CgRRCRetransmissionTimerR18Constraints = per.Constrained(1, 288)

var cG_RRC_Configuration_r18RrcSSBSubsetR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "shortBitmap-r18"},
		{Name: "mediumBitmap-r18"},
		{Name: "longBitmap-r18"},
	},
}

const (
	CG_RRC_Configuration_r18_Rrc_SSB_Subset_r18_ShortBitmap_r18  = 0
	CG_RRC_Configuration_r18_Rrc_SSB_Subset_r18_MediumBitmap_r18 = 1
	CG_RRC_Configuration_r18_Rrc_SSB_Subset_r18_LongBitmap_r18   = 2
)

const (
	CG_RRC_Configuration_r18_Rrc_SSB_PerCG_PUSCH_r18_OneEighth = 0
	CG_RRC_Configuration_r18_Rrc_SSB_PerCG_PUSCH_r18_OneFourth = 1
	CG_RRC_Configuration_r18_Rrc_SSB_PerCG_PUSCH_r18_Half      = 2
	CG_RRC_Configuration_r18_Rrc_SSB_PerCG_PUSCH_r18_One       = 3
	CG_RRC_Configuration_r18_Rrc_SSB_PerCG_PUSCH_r18_Two       = 4
	CG_RRC_Configuration_r18_Rrc_SSB_PerCG_PUSCH_r18_Four      = 5
	CG_RRC_Configuration_r18_Rrc_SSB_PerCG_PUSCH_r18_Eight     = 6
	CG_RRC_Configuration_r18_Rrc_SSB_PerCG_PUSCH_r18_Sixteen   = 7
)

var cGRRCConfigurationR18RrcSSBPerCGPUSCHR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CG_RRC_Configuration_r18_Rrc_SSB_PerCG_PUSCH_r18_OneEighth, CG_RRC_Configuration_r18_Rrc_SSB_PerCG_PUSCH_r18_OneFourth, CG_RRC_Configuration_r18_Rrc_SSB_PerCG_PUSCH_r18_Half, CG_RRC_Configuration_r18_Rrc_SSB_PerCG_PUSCH_r18_One, CG_RRC_Configuration_r18_Rrc_SSB_PerCG_PUSCH_r18_Two, CG_RRC_Configuration_r18_Rrc_SSB_PerCG_PUSCH_r18_Four, CG_RRC_Configuration_r18_Rrc_SSB_PerCG_PUSCH_r18_Eight, CG_RRC_Configuration_r18_Rrc_SSB_PerCG_PUSCH_r18_Sixteen},
}

var cGRRCConfigurationR18RrcP0PUSCHR18Constraints = per.Constrained(-16, 15)

const (
	CG_RRC_Configuration_r18_Rrc_Alpha_r18_Alpha0  = 0
	CG_RRC_Configuration_r18_Rrc_Alpha_r18_Alpha04 = 1
	CG_RRC_Configuration_r18_Rrc_Alpha_r18_Alpha05 = 2
	CG_RRC_Configuration_r18_Rrc_Alpha_r18_Alpha06 = 3
	CG_RRC_Configuration_r18_Rrc_Alpha_r18_Alpha07 = 4
	CG_RRC_Configuration_r18_Rrc_Alpha_r18_Alpha08 = 5
	CG_RRC_Configuration_r18_Rrc_Alpha_r18_Alpha09 = 6
	CG_RRC_Configuration_r18_Rrc_Alpha_r18_Alpha1  = 7
)

var cGRRCConfigurationR18RrcAlphaR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CG_RRC_Configuration_r18_Rrc_Alpha_r18_Alpha0, CG_RRC_Configuration_r18_Rrc_Alpha_r18_Alpha04, CG_RRC_Configuration_r18_Rrc_Alpha_r18_Alpha05, CG_RRC_Configuration_r18_Rrc_Alpha_r18_Alpha06, CG_RRC_Configuration_r18_Rrc_Alpha_r18_Alpha07, CG_RRC_Configuration_r18_Rrc_Alpha_r18_Alpha08, CG_RRC_Configuration_r18_Rrc_Alpha_r18_Alpha09, CG_RRC_Configuration_r18_Rrc_Alpha_r18_Alpha1},
}

var cG_RRC_Configuration_r18RrcDMRSPortsR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "dmrsType1-r18"},
		{Name: "dmrsType2-r18"},
	},
}

const (
	CG_RRC_Configuration_r18_Rrc_DMRS_Ports_r18_DmrsType1_r18 = 0
	CG_RRC_Configuration_r18_Rrc_DMRS_Ports_r18_DmrsType2_r18 = 1
)

var cGRRCConfigurationR18RrcNrofDMRSSequencesR18Constraints = per.Constrained(1, 2)

type CG_RRC_Configuration_r18 struct {
	Cg_RRC_RetransmissionTimer_r18 *int64
	Cg_RRC_RSRP_ThresholdSSB_r18   *RSRP_Range
	Rrc_SSB_Subset_r18             *struct {
		Choice           int
		ShortBitmap_r18  *per.BitString
		MediumBitmap_r18 *per.BitString
		LongBitmap_r18   *per.BitString
	}
	Rrc_SSB_PerCG_PUSCH_r18 *int64
	Rrc_P0_PUSCH_r18        *int64
	Rrc_Alpha_r18           *int64
	Rrc_DMRS_Ports_r18      *struct {
		Choice        int
		DmrsType1_r18 *per.BitString
		DmrsType2_r18 *per.BitString
	}
	Rrc_NrofDMRS_Sequences_r18 *int64
}

func (ie *CG_RRC_Configuration_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cGRRCConfigurationR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Cg_RRC_RetransmissionTimer_r18 != nil, ie.Cg_RRC_RSRP_ThresholdSSB_r18 != nil, ie.Rrc_SSB_Subset_r18 != nil, ie.Rrc_SSB_PerCG_PUSCH_r18 != nil, ie.Rrc_P0_PUSCH_r18 != nil, ie.Rrc_Alpha_r18 != nil, ie.Rrc_DMRS_Ports_r18 != nil, ie.Rrc_NrofDMRS_Sequences_r18 != nil}); err != nil {
		return err
	}

	// 3. cg-RRC-RetransmissionTimer-r18: integer
	{
		if ie.Cg_RRC_RetransmissionTimer_r18 != nil {
			if err := e.EncodeInteger(*ie.Cg_RRC_RetransmissionTimer_r18, cGRRCConfigurationR18CgRRCRetransmissionTimerR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. cg-RRC-RSRP-ThresholdSSB-r18: ref
	{
		if ie.Cg_RRC_RSRP_ThresholdSSB_r18 != nil {
			if err := ie.Cg_RRC_RSRP_ThresholdSSB_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. rrc-SSB-Subset-r18: choice
	{
		if ie.Rrc_SSB_Subset_r18 != nil {
			choiceEnc := e.NewChoiceEncoder(cG_RRC_Configuration_r18RrcSSBSubsetR18Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Rrc_SSB_Subset_r18).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Rrc_SSB_Subset_r18).Choice {
			case CG_RRC_Configuration_r18_Rrc_SSB_Subset_r18_ShortBitmap_r18:
				if err := e.EncodeBitString((*(*ie.Rrc_SSB_Subset_r18).ShortBitmap_r18), per.FixedSize(4)); err != nil {
					return err
				}
			case CG_RRC_Configuration_r18_Rrc_SSB_Subset_r18_MediumBitmap_r18:
				if err := e.EncodeBitString((*(*ie.Rrc_SSB_Subset_r18).MediumBitmap_r18), per.FixedSize(8)); err != nil {
					return err
				}
			case CG_RRC_Configuration_r18_Rrc_SSB_Subset_r18_LongBitmap_r18:
				if err := e.EncodeBitString((*(*ie.Rrc_SSB_Subset_r18).LongBitmap_r18), per.FixedSize(64)); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Rrc_SSB_Subset_r18).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 6. rrc-SSB-PerCG-PUSCH-r18: enumerated
	{
		if ie.Rrc_SSB_PerCG_PUSCH_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Rrc_SSB_PerCG_PUSCH_r18, cGRRCConfigurationR18RrcSSBPerCGPUSCHR18Constraints); err != nil {
				return err
			}
		}
	}

	// 7. rrc-P0-PUSCH-r18: integer
	{
		if ie.Rrc_P0_PUSCH_r18 != nil {
			if err := e.EncodeInteger(*ie.Rrc_P0_PUSCH_r18, cGRRCConfigurationR18RrcP0PUSCHR18Constraints); err != nil {
				return err
			}
		}
	}

	// 8. rrc-Alpha-r18: enumerated
	{
		if ie.Rrc_Alpha_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Rrc_Alpha_r18, cGRRCConfigurationR18RrcAlphaR18Constraints); err != nil {
				return err
			}
		}
	}

	// 9. rrc-DMRS-Ports-r18: choice
	{
		if ie.Rrc_DMRS_Ports_r18 != nil {
			choiceEnc := e.NewChoiceEncoder(cG_RRC_Configuration_r18RrcDMRSPortsR18Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Rrc_DMRS_Ports_r18).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Rrc_DMRS_Ports_r18).Choice {
			case CG_RRC_Configuration_r18_Rrc_DMRS_Ports_r18_DmrsType1_r18:
				if err := e.EncodeBitString((*(*ie.Rrc_DMRS_Ports_r18).DmrsType1_r18), per.FixedSize(8)); err != nil {
					return err
				}
			case CG_RRC_Configuration_r18_Rrc_DMRS_Ports_r18_DmrsType2_r18:
				if err := e.EncodeBitString((*(*ie.Rrc_DMRS_Ports_r18).DmrsType2_r18), per.FixedSize(12)); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Rrc_DMRS_Ports_r18).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 10. rrc-NrofDMRS-Sequences-r18: integer
	{
		if ie.Rrc_NrofDMRS_Sequences_r18 != nil {
			if err := e.EncodeInteger(*ie.Rrc_NrofDMRS_Sequences_r18, cGRRCConfigurationR18RrcNrofDMRSSequencesR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CG_RRC_Configuration_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cGRRCConfigurationR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. cg-RRC-RetransmissionTimer-r18: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(cGRRCConfigurationR18CgRRCRetransmissionTimerR18Constraints)
			if err != nil {
				return err
			}
			ie.Cg_RRC_RetransmissionTimer_r18 = &v
		}
	}

	// 4. cg-RRC-RSRP-ThresholdSSB-r18: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Cg_RRC_RSRP_ThresholdSSB_r18 = new(RSRP_Range)
			if err := ie.Cg_RRC_RSRP_ThresholdSSB_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. rrc-SSB-Subset-r18: choice
	{
		if seq.IsComponentPresent(2) {
			ie.Rrc_SSB_Subset_r18 = &struct {
				Choice           int
				ShortBitmap_r18  *per.BitString
				MediumBitmap_r18 *per.BitString
				LongBitmap_r18   *per.BitString
			}{}
			choiceDec := d.NewChoiceDecoder(cG_RRC_Configuration_r18RrcSSBSubsetR18Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Rrc_SSB_Subset_r18).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case CG_RRC_Configuration_r18_Rrc_SSB_Subset_r18_ShortBitmap_r18:
				(*ie.Rrc_SSB_Subset_r18).ShortBitmap_r18 = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(4))
				if err != nil {
					return err
				}
				(*(*ie.Rrc_SSB_Subset_r18).ShortBitmap_r18) = v
			case CG_RRC_Configuration_r18_Rrc_SSB_Subset_r18_MediumBitmap_r18:
				(*ie.Rrc_SSB_Subset_r18).MediumBitmap_r18 = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(8))
				if err != nil {
					return err
				}
				(*(*ie.Rrc_SSB_Subset_r18).MediumBitmap_r18) = v
			case CG_RRC_Configuration_r18_Rrc_SSB_Subset_r18_LongBitmap_r18:
				(*ie.Rrc_SSB_Subset_r18).LongBitmap_r18 = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(64))
				if err != nil {
					return err
				}
				(*(*ie.Rrc_SSB_Subset_r18).LongBitmap_r18) = v
			}
		}
	}

	// 6. rrc-SSB-PerCG-PUSCH-r18: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(cGRRCConfigurationR18RrcSSBPerCGPUSCHR18Constraints)
			if err != nil {
				return err
			}
			ie.Rrc_SSB_PerCG_PUSCH_r18 = &idx
		}
	}

	// 7. rrc-P0-PUSCH-r18: integer
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeInteger(cGRRCConfigurationR18RrcP0PUSCHR18Constraints)
			if err != nil {
				return err
			}
			ie.Rrc_P0_PUSCH_r18 = &v
		}
	}

	// 8. rrc-Alpha-r18: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(cGRRCConfigurationR18RrcAlphaR18Constraints)
			if err != nil {
				return err
			}
			ie.Rrc_Alpha_r18 = &idx
		}
	}

	// 9. rrc-DMRS-Ports-r18: choice
	{
		if seq.IsComponentPresent(6) {
			ie.Rrc_DMRS_Ports_r18 = &struct {
				Choice        int
				DmrsType1_r18 *per.BitString
				DmrsType2_r18 *per.BitString
			}{}
			choiceDec := d.NewChoiceDecoder(cG_RRC_Configuration_r18RrcDMRSPortsR18Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Rrc_DMRS_Ports_r18).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case CG_RRC_Configuration_r18_Rrc_DMRS_Ports_r18_DmrsType1_r18:
				(*ie.Rrc_DMRS_Ports_r18).DmrsType1_r18 = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(8))
				if err != nil {
					return err
				}
				(*(*ie.Rrc_DMRS_Ports_r18).DmrsType1_r18) = v
			case CG_RRC_Configuration_r18_Rrc_DMRS_Ports_r18_DmrsType2_r18:
				(*ie.Rrc_DMRS_Ports_r18).DmrsType2_r18 = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(12))
				if err != nil {
					return err
				}
				(*(*ie.Rrc_DMRS_Ports_r18).DmrsType2_r18) = v
			}
		}
	}

	// 10. rrc-NrofDMRS-Sequences-r18: integer
	{
		if seq.IsComponentPresent(7) {
			v, err := d.DecodeInteger(cGRRCConfigurationR18RrcNrofDMRSSequencesR18Constraints)
			if err != nil {
				return err
			}
			ie.Rrc_NrofDMRS_Sequences_r18 = &v
		}
	}

	return nil
}

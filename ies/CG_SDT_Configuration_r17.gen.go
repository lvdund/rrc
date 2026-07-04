// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CG-SDT-Configuration-r17 (line 6704).

var cGSDTConfigurationR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "cg-SDT-RetransmissionTimer", Optional: true},
		{Name: "sdt-SSB-Subset-r17", Optional: true},
		{Name: "sdt-SSB-PerCG-PUSCH-r17", Optional: true},
		{Name: "sdt-P0-PUSCH-r17", Optional: true},
		{Name: "sdt-Alpha-r17", Optional: true},
		{Name: "sdt-DMRS-Ports-r17", Optional: true},
		{Name: "sdt-NrofDMRS-Sequences-r17", Optional: true},
	},
}

var cGSDTConfigurationR17CgSDTRetransmissionTimerConstraints = per.Constrained(1, 64)

var cG_SDT_Configuration_r17SdtSSBSubsetR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "shortBitmap-r17"},
		{Name: "mediumBitmap-r17"},
		{Name: "longBitmap-r17"},
	},
}

const (
	CG_SDT_Configuration_r17_Sdt_SSB_Subset_r17_ShortBitmap_r17  = 0
	CG_SDT_Configuration_r17_Sdt_SSB_Subset_r17_MediumBitmap_r17 = 1
	CG_SDT_Configuration_r17_Sdt_SSB_Subset_r17_LongBitmap_r17   = 2
)

const (
	CG_SDT_Configuration_r17_Sdt_SSB_PerCG_PUSCH_r17_OneEighth = 0
	CG_SDT_Configuration_r17_Sdt_SSB_PerCG_PUSCH_r17_OneFourth = 1
	CG_SDT_Configuration_r17_Sdt_SSB_PerCG_PUSCH_r17_Half      = 2
	CG_SDT_Configuration_r17_Sdt_SSB_PerCG_PUSCH_r17_One       = 3
	CG_SDT_Configuration_r17_Sdt_SSB_PerCG_PUSCH_r17_Two       = 4
	CG_SDT_Configuration_r17_Sdt_SSB_PerCG_PUSCH_r17_Four      = 5
	CG_SDT_Configuration_r17_Sdt_SSB_PerCG_PUSCH_r17_Eight     = 6
	CG_SDT_Configuration_r17_Sdt_SSB_PerCG_PUSCH_r17_Sixteen   = 7
)

var cGSDTConfigurationR17SdtSSBPerCGPUSCHR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CG_SDT_Configuration_r17_Sdt_SSB_PerCG_PUSCH_r17_OneEighth, CG_SDT_Configuration_r17_Sdt_SSB_PerCG_PUSCH_r17_OneFourth, CG_SDT_Configuration_r17_Sdt_SSB_PerCG_PUSCH_r17_Half, CG_SDT_Configuration_r17_Sdt_SSB_PerCG_PUSCH_r17_One, CG_SDT_Configuration_r17_Sdt_SSB_PerCG_PUSCH_r17_Two, CG_SDT_Configuration_r17_Sdt_SSB_PerCG_PUSCH_r17_Four, CG_SDT_Configuration_r17_Sdt_SSB_PerCG_PUSCH_r17_Eight, CG_SDT_Configuration_r17_Sdt_SSB_PerCG_PUSCH_r17_Sixteen},
}

var cGSDTConfigurationR17SdtP0PUSCHR17Constraints = per.Constrained(-16, 15)

const (
	CG_SDT_Configuration_r17_Sdt_Alpha_r17_Alpha0  = 0
	CG_SDT_Configuration_r17_Sdt_Alpha_r17_Alpha04 = 1
	CG_SDT_Configuration_r17_Sdt_Alpha_r17_Alpha05 = 2
	CG_SDT_Configuration_r17_Sdt_Alpha_r17_Alpha06 = 3
	CG_SDT_Configuration_r17_Sdt_Alpha_r17_Alpha07 = 4
	CG_SDT_Configuration_r17_Sdt_Alpha_r17_Alpha08 = 5
	CG_SDT_Configuration_r17_Sdt_Alpha_r17_Alpha09 = 6
	CG_SDT_Configuration_r17_Sdt_Alpha_r17_Alpha1  = 7
)

var cGSDTConfigurationR17SdtAlphaR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CG_SDT_Configuration_r17_Sdt_Alpha_r17_Alpha0, CG_SDT_Configuration_r17_Sdt_Alpha_r17_Alpha04, CG_SDT_Configuration_r17_Sdt_Alpha_r17_Alpha05, CG_SDT_Configuration_r17_Sdt_Alpha_r17_Alpha06, CG_SDT_Configuration_r17_Sdt_Alpha_r17_Alpha07, CG_SDT_Configuration_r17_Sdt_Alpha_r17_Alpha08, CG_SDT_Configuration_r17_Sdt_Alpha_r17_Alpha09, CG_SDT_Configuration_r17_Sdt_Alpha_r17_Alpha1},
}

var cG_SDT_Configuration_r17SdtDMRSPortsR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "dmrsType1-r17"},
		{Name: "dmrsType2-r17"},
	},
}

const (
	CG_SDT_Configuration_r17_Sdt_DMRS_Ports_r17_DmrsType1_r17 = 0
	CG_SDT_Configuration_r17_Sdt_DMRS_Ports_r17_DmrsType2_r17 = 1
)

var cGSDTConfigurationR17SdtNrofDMRSSequencesR17Constraints = per.Constrained(1, 2)

type CG_SDT_Configuration_r17 struct {
	Cg_SDT_RetransmissionTimer *int64
	Sdt_SSB_Subset_r17         *struct {
		Choice           int
		ShortBitmap_r17  *per.BitString
		MediumBitmap_r17 *per.BitString
		LongBitmap_r17   *per.BitString
	}
	Sdt_SSB_PerCG_PUSCH_r17 *int64
	Sdt_P0_PUSCH_r17        *int64
	Sdt_Alpha_r17           *int64
	Sdt_DMRS_Ports_r17      *struct {
		Choice        int
		DmrsType1_r17 *per.BitString
		DmrsType2_r17 *per.BitString
	}
	Sdt_NrofDMRS_Sequences_r17 *int64
}

func (ie *CG_SDT_Configuration_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cGSDTConfigurationR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Cg_SDT_RetransmissionTimer != nil, ie.Sdt_SSB_Subset_r17 != nil, ie.Sdt_SSB_PerCG_PUSCH_r17 != nil, ie.Sdt_P0_PUSCH_r17 != nil, ie.Sdt_Alpha_r17 != nil, ie.Sdt_DMRS_Ports_r17 != nil, ie.Sdt_NrofDMRS_Sequences_r17 != nil}); err != nil {
		return err
	}

	// 2. cg-SDT-RetransmissionTimer: integer
	{
		if ie.Cg_SDT_RetransmissionTimer != nil {
			if err := e.EncodeInteger(*ie.Cg_SDT_RetransmissionTimer, cGSDTConfigurationR17CgSDTRetransmissionTimerConstraints); err != nil {
				return err
			}
		}
	}

	// 3. sdt-SSB-Subset-r17: choice
	{
		if ie.Sdt_SSB_Subset_r17 != nil {
			choiceEnc := e.NewChoiceEncoder(cG_SDT_Configuration_r17SdtSSBSubsetR17Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Sdt_SSB_Subset_r17).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Sdt_SSB_Subset_r17).Choice {
			case CG_SDT_Configuration_r17_Sdt_SSB_Subset_r17_ShortBitmap_r17:
				if err := e.EncodeBitString((*(*ie.Sdt_SSB_Subset_r17).ShortBitmap_r17), per.FixedSize(4)); err != nil {
					return err
				}
			case CG_SDT_Configuration_r17_Sdt_SSB_Subset_r17_MediumBitmap_r17:
				if err := e.EncodeBitString((*(*ie.Sdt_SSB_Subset_r17).MediumBitmap_r17), per.FixedSize(8)); err != nil {
					return err
				}
			case CG_SDT_Configuration_r17_Sdt_SSB_Subset_r17_LongBitmap_r17:
				if err := e.EncodeBitString((*(*ie.Sdt_SSB_Subset_r17).LongBitmap_r17), per.FixedSize(64)); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Sdt_SSB_Subset_r17).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 4. sdt-SSB-PerCG-PUSCH-r17: enumerated
	{
		if ie.Sdt_SSB_PerCG_PUSCH_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Sdt_SSB_PerCG_PUSCH_r17, cGSDTConfigurationR17SdtSSBPerCGPUSCHR17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. sdt-P0-PUSCH-r17: integer
	{
		if ie.Sdt_P0_PUSCH_r17 != nil {
			if err := e.EncodeInteger(*ie.Sdt_P0_PUSCH_r17, cGSDTConfigurationR17SdtP0PUSCHR17Constraints); err != nil {
				return err
			}
		}
	}

	// 6. sdt-Alpha-r17: enumerated
	{
		if ie.Sdt_Alpha_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Sdt_Alpha_r17, cGSDTConfigurationR17SdtAlphaR17Constraints); err != nil {
				return err
			}
		}
	}

	// 7. sdt-DMRS-Ports-r17: choice
	{
		if ie.Sdt_DMRS_Ports_r17 != nil {
			choiceEnc := e.NewChoiceEncoder(cG_SDT_Configuration_r17SdtDMRSPortsR17Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Sdt_DMRS_Ports_r17).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Sdt_DMRS_Ports_r17).Choice {
			case CG_SDT_Configuration_r17_Sdt_DMRS_Ports_r17_DmrsType1_r17:
				if err := e.EncodeBitString((*(*ie.Sdt_DMRS_Ports_r17).DmrsType1_r17), per.FixedSize(8)); err != nil {
					return err
				}
			case CG_SDT_Configuration_r17_Sdt_DMRS_Ports_r17_DmrsType2_r17:
				if err := e.EncodeBitString((*(*ie.Sdt_DMRS_Ports_r17).DmrsType2_r17), per.FixedSize(12)); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Sdt_DMRS_Ports_r17).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 8. sdt-NrofDMRS-Sequences-r17: integer
	{
		if ie.Sdt_NrofDMRS_Sequences_r17 != nil {
			if err := e.EncodeInteger(*ie.Sdt_NrofDMRS_Sequences_r17, cGSDTConfigurationR17SdtNrofDMRSSequencesR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CG_SDT_Configuration_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cGSDTConfigurationR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. cg-SDT-RetransmissionTimer: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(cGSDTConfigurationR17CgSDTRetransmissionTimerConstraints)
			if err != nil {
				return err
			}
			ie.Cg_SDT_RetransmissionTimer = &v
		}
	}

	// 3. sdt-SSB-Subset-r17: choice
	{
		if seq.IsComponentPresent(1) {
			ie.Sdt_SSB_Subset_r17 = &struct {
				Choice           int
				ShortBitmap_r17  *per.BitString
				MediumBitmap_r17 *per.BitString
				LongBitmap_r17   *per.BitString
			}{}
			choiceDec := d.NewChoiceDecoder(cG_SDT_Configuration_r17SdtSSBSubsetR17Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sdt_SSB_Subset_r17).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case CG_SDT_Configuration_r17_Sdt_SSB_Subset_r17_ShortBitmap_r17:
				(*ie.Sdt_SSB_Subset_r17).ShortBitmap_r17 = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(4))
				if err != nil {
					return err
				}
				(*(*ie.Sdt_SSB_Subset_r17).ShortBitmap_r17) = v
			case CG_SDT_Configuration_r17_Sdt_SSB_Subset_r17_MediumBitmap_r17:
				(*ie.Sdt_SSB_Subset_r17).MediumBitmap_r17 = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(8))
				if err != nil {
					return err
				}
				(*(*ie.Sdt_SSB_Subset_r17).MediumBitmap_r17) = v
			case CG_SDT_Configuration_r17_Sdt_SSB_Subset_r17_LongBitmap_r17:
				(*ie.Sdt_SSB_Subset_r17).LongBitmap_r17 = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(64))
				if err != nil {
					return err
				}
				(*(*ie.Sdt_SSB_Subset_r17).LongBitmap_r17) = v
			}
		}
	}

	// 4. sdt-SSB-PerCG-PUSCH-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(cGSDTConfigurationR17SdtSSBPerCGPUSCHR17Constraints)
			if err != nil {
				return err
			}
			ie.Sdt_SSB_PerCG_PUSCH_r17 = &idx
		}
	}

	// 5. sdt-P0-PUSCH-r17: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(cGSDTConfigurationR17SdtP0PUSCHR17Constraints)
			if err != nil {
				return err
			}
			ie.Sdt_P0_PUSCH_r17 = &v
		}
	}

	// 6. sdt-Alpha-r17: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(cGSDTConfigurationR17SdtAlphaR17Constraints)
			if err != nil {
				return err
			}
			ie.Sdt_Alpha_r17 = &idx
		}
	}

	// 7. sdt-DMRS-Ports-r17: choice
	{
		if seq.IsComponentPresent(5) {
			ie.Sdt_DMRS_Ports_r17 = &struct {
				Choice        int
				DmrsType1_r17 *per.BitString
				DmrsType2_r17 *per.BitString
			}{}
			choiceDec := d.NewChoiceDecoder(cG_SDT_Configuration_r17SdtDMRSPortsR17Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sdt_DMRS_Ports_r17).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case CG_SDT_Configuration_r17_Sdt_DMRS_Ports_r17_DmrsType1_r17:
				(*ie.Sdt_DMRS_Ports_r17).DmrsType1_r17 = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(8))
				if err != nil {
					return err
				}
				(*(*ie.Sdt_DMRS_Ports_r17).DmrsType1_r17) = v
			case CG_SDT_Configuration_r17_Sdt_DMRS_Ports_r17_DmrsType2_r17:
				(*ie.Sdt_DMRS_Ports_r17).DmrsType2_r17 = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(12))
				if err != nil {
					return err
				}
				(*(*ie.Sdt_DMRS_Ports_r17).DmrsType2_r17) = v
			}
		}
	}

	// 8. sdt-NrofDMRS-Sequences-r17: integer
	{
		if seq.IsComponentPresent(6) {
			v, err := d.DecodeInteger(cGSDTConfigurationR17SdtNrofDMRSSequencesR17Constraints)
			if err != nil {
				return err
			}
			ie.Sdt_NrofDMRS_Sequences_r17 = &v
		}
	}

	return nil
}

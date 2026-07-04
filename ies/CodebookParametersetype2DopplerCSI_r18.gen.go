// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CodebookParametersetype2DopplerCSI-r18 (line 18875).

var codebookParametersetype2DopplerCSIR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "eType2Doppler-r18"},
		{Name: "eType2DopplerN4-r18", Optional: true},
		{Name: "ddUnitSize-A-CSI-RS-CMR-r18", Optional: true},
		{Name: "maxNumberAperiodicCSI-RS-Resource-r18", Optional: true},
		{Name: "eType2DopplerR2-r18", Optional: true},
		{Name: "eType2DopplerX1-r18", Optional: true},
		{Name: "eType2DopplerX2-r18", Optional: true},
		{Name: "eType2DopplerL-N4D1-r18", Optional: true},
		{Name: "eType2DopplerL6-r18", Optional: true},
		{Name: "eType2DopplerR3R4-r18", Optional: true},
	},
}

const (
	CodebookParametersetype2DopplerCSI_r18_DdUnitSize_A_CSI_RS_CMR_r18_Supported = 0
)

var codebookParametersetype2DopplerCSIR18DdUnitSizeACSIRSCMRR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersetype2DopplerCSI_r18_DdUnitSize_A_CSI_RS_CMR_r18_Supported},
}

const (
	CodebookParametersetype2DopplerCSI_r18_MaxNumberAperiodicCSI_RS_Resource_r18_N4  = 0
	CodebookParametersetype2DopplerCSI_r18_MaxNumberAperiodicCSI_RS_Resource_r18_N8  = 1
	CodebookParametersetype2DopplerCSI_r18_MaxNumberAperiodicCSI_RS_Resource_r18_N12 = 2
)

var codebookParametersetype2DopplerCSIR18MaxNumberAperiodicCSIRSResourceR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersetype2DopplerCSI_r18_MaxNumberAperiodicCSI_RS_Resource_r18_N4, CodebookParametersetype2DopplerCSI_r18_MaxNumberAperiodicCSI_RS_Resource_r18_N8, CodebookParametersetype2DopplerCSI_r18_MaxNumberAperiodicCSI_RS_Resource_r18_N12},
}

var codebookParametersetype2DopplerCSIR18EType2DopplerR2R18Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

const (
	CodebookParametersetype2DopplerCSI_r18_EType2DopplerX1_r18_Supported = 0
)

var codebookParametersetype2DopplerCSIR18EType2DopplerX1R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersetype2DopplerCSI_r18_EType2DopplerX1_r18_Supported},
}

const (
	CodebookParametersetype2DopplerCSI_r18_EType2DopplerX2_r18_Supported = 0
)

var codebookParametersetype2DopplerCSIR18EType2DopplerX2R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersetype2DopplerCSI_r18_EType2DopplerX2_r18_Supported},
}

const (
	CodebookParametersetype2DopplerCSI_r18_EType2DopplerL_N4D1_r18_Supported = 0
)

var codebookParametersetype2DopplerCSIR18EType2DopplerLN4D1R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersetype2DopplerCSI_r18_EType2DopplerL_N4D1_r18_Supported},
}

const (
	CodebookParametersetype2DopplerCSI_r18_EType2DopplerL6_r18_Supported = 0
)

var codebookParametersetype2DopplerCSIR18EType2DopplerL6R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersetype2DopplerCSI_r18_EType2DopplerL6_r18_Supported},
}

const (
	CodebookParametersetype2DopplerCSI_r18_EType2DopplerR3R4_r18_Supported = 0
)

var codebookParametersetype2DopplerCSIR18EType2DopplerR3R4R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersetype2DopplerCSI_r18_EType2DopplerR3R4_r18_Supported},
}

var codebookParametersetype2DopplerCSIR18EType2DopplerR18SupportedCSIRSResourceListR18Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

const (
	CodebookParametersetype2DopplerCSI_r18_EType2Doppler_r18_Scalingfactor_r18_N1 = 0
	CodebookParametersetype2DopplerCSI_r18_EType2Doppler_r18_Scalingfactor_r18_N2 = 1
	CodebookParametersetype2DopplerCSI_r18_EType2Doppler_r18_Scalingfactor_r18_N4 = 2
)

var codebookParametersetype2DopplerCSIR18EType2DopplerR18ScalingfactorR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersetype2DopplerCSI_r18_EType2Doppler_r18_Scalingfactor_r18_N1, CodebookParametersetype2DopplerCSI_r18_EType2Doppler_r18_Scalingfactor_r18_N2, CodebookParametersetype2DopplerCSI_r18_EType2Doppler_r18_Scalingfactor_r18_N4},
}

var codebookParametersetype2DopplerCSIR18EType2DopplerN4R18SupportedCSIRSReportSettingList1R18Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookParametersetype2DopplerCSIR18EType2DopplerN4R18SupportedCSIRSReportSettingList2R18Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

type CodebookParametersetype2DopplerCSI_r18 struct {
	EType2Doppler_r18 struct {
		SupportedCSI_RS_ResourceList_r18 []int64
		ValueY_P_SP_CSI_RS_r18           int64
		ValueY_A_CSI_RS_r18              int64
		Scalingfactor_r18                int64
	}
	EType2DopplerN4_r18 *struct {
		SupportedCSI_RS_ReportSettingList1_r18 []SupportedCSI_RS_ReportSetting_r18
		SupportedCSI_RS_ReportSettingList2_r18 []SupportedCSI_RS_ReportSetting_r18
	}
	DdUnitSize_A_CSI_RS_CMR_r18           *int64
	MaxNumberAperiodicCSI_RS_Resource_r18 *int64
	EType2DopplerR2_r18                   []int64
	EType2DopplerX1_r18                   *int64
	EType2DopplerX2_r18                   *int64
	EType2DopplerL_N4D1_r18               *int64
	EType2DopplerL6_r18                   *int64
	EType2DopplerR3R4_r18                 *int64
}

func (ie *CodebookParametersetype2DopplerCSI_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(codebookParametersetype2DopplerCSIR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.EType2DopplerN4_r18 != nil, ie.DdUnitSize_A_CSI_RS_CMR_r18 != nil, ie.MaxNumberAperiodicCSI_RS_Resource_r18 != nil, ie.EType2DopplerR2_r18 != nil, ie.EType2DopplerX1_r18 != nil, ie.EType2DopplerX2_r18 != nil, ie.EType2DopplerL_N4D1_r18 != nil, ie.EType2DopplerL6_r18 != nil, ie.EType2DopplerR3R4_r18 != nil}); err != nil {
		return err
	}

	// 2. eType2Doppler-r18: sequence
	{
		{
			c := &ie.EType2Doppler_r18
			{
				seqOf := e.NewSequenceOfEncoder(codebookParametersetype2DopplerCSIR18EType2DopplerR18SupportedCSIRSResourceListR18Constraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ResourceList_r18))); err != nil {
					return err
				}
				for i := range c.SupportedCSI_RS_ResourceList_r18 {
					if err := e.EncodeInteger(c.SupportedCSI_RS_ResourceList_r18[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
			if err := e.EncodeInteger(c.ValueY_P_SP_CSI_RS_r18, per.Constrained(1, 3)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.ValueY_A_CSI_RS_r18, per.Constrained(1, 3)); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.Scalingfactor_r18, codebookParametersetype2DopplerCSIR18EType2DopplerR18ScalingfactorR18Constraints); err != nil {
				return err
			}
		}
	}

	// 3. eType2DopplerN4-r18: sequence
	{
		if ie.EType2DopplerN4_r18 != nil {
			c := ie.EType2DopplerN4_r18
			{
				seqOf := e.NewSequenceOfEncoder(codebookParametersetype2DopplerCSIR18EType2DopplerN4R18SupportedCSIRSReportSettingList1R18Constraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ReportSettingList1_r18))); err != nil {
					return err
				}
				for i := range c.SupportedCSI_RS_ReportSettingList1_r18 {
					if err := c.SupportedCSI_RS_ReportSettingList1_r18[i].Encode(e); err != nil {
						return err
					}
				}
			}
			{
				seqOf := e.NewSequenceOfEncoder(codebookParametersetype2DopplerCSIR18EType2DopplerN4R18SupportedCSIRSReportSettingList2R18Constraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ReportSettingList2_r18))); err != nil {
					return err
				}
				for i := range c.SupportedCSI_RS_ReportSettingList2_r18 {
					if err := c.SupportedCSI_RS_ReportSettingList2_r18[i].Encode(e); err != nil {
						return err
					}
				}
			}
		}
	}

	// 4. ddUnitSize-A-CSI-RS-CMR-r18: enumerated
	{
		if ie.DdUnitSize_A_CSI_RS_CMR_r18 != nil {
			if err := e.EncodeEnumerated(*ie.DdUnitSize_A_CSI_RS_CMR_r18, codebookParametersetype2DopplerCSIR18DdUnitSizeACSIRSCMRR18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. maxNumberAperiodicCSI-RS-Resource-r18: enumerated
	{
		if ie.MaxNumberAperiodicCSI_RS_Resource_r18 != nil {
			if err := e.EncodeEnumerated(*ie.MaxNumberAperiodicCSI_RS_Resource_r18, codebookParametersetype2DopplerCSIR18MaxNumberAperiodicCSIRSResourceR18Constraints); err != nil {
				return err
			}
		}
	}

	// 6. eType2DopplerR2-r18: sequence-of
	{
		if ie.EType2DopplerR2_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookParametersetype2DopplerCSIR18EType2DopplerR2R18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.EType2DopplerR2_r18))); err != nil {
				return err
			}
			for i := range ie.EType2DopplerR2_r18 {
				if err := e.EncodeInteger(ie.EType2DopplerR2_r18[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 7. eType2DopplerX1-r18: enumerated
	{
		if ie.EType2DopplerX1_r18 != nil {
			if err := e.EncodeEnumerated(*ie.EType2DopplerX1_r18, codebookParametersetype2DopplerCSIR18EType2DopplerX1R18Constraints); err != nil {
				return err
			}
		}
	}

	// 8. eType2DopplerX2-r18: enumerated
	{
		if ie.EType2DopplerX2_r18 != nil {
			if err := e.EncodeEnumerated(*ie.EType2DopplerX2_r18, codebookParametersetype2DopplerCSIR18EType2DopplerX2R18Constraints); err != nil {
				return err
			}
		}
	}

	// 9. eType2DopplerL-N4D1-r18: enumerated
	{
		if ie.EType2DopplerL_N4D1_r18 != nil {
			if err := e.EncodeEnumerated(*ie.EType2DopplerL_N4D1_r18, codebookParametersetype2DopplerCSIR18EType2DopplerLN4D1R18Constraints); err != nil {
				return err
			}
		}
	}

	// 10. eType2DopplerL6-r18: enumerated
	{
		if ie.EType2DopplerL6_r18 != nil {
			if err := e.EncodeEnumerated(*ie.EType2DopplerL6_r18, codebookParametersetype2DopplerCSIR18EType2DopplerL6R18Constraints); err != nil {
				return err
			}
		}
	}

	// 11. eType2DopplerR3R4-r18: enumerated
	{
		if ie.EType2DopplerR3R4_r18 != nil {
			if err := e.EncodeEnumerated(*ie.EType2DopplerR3R4_r18, codebookParametersetype2DopplerCSIR18EType2DopplerR3R4R18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CodebookParametersetype2DopplerCSI_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(codebookParametersetype2DopplerCSIR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. eType2Doppler-r18: sequence
	{
		{
			c := &ie.EType2Doppler_r18
			{
				seqOf := d.NewSequenceOfDecoder(codebookParametersetype2DopplerCSIR18EType2DopplerR18SupportedCSIRSResourceListR18Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SupportedCSI_RS_ResourceList_r18 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
					if err != nil {
						return err
					}
					c.SupportedCSI_RS_ResourceList_r18[i] = v
				}
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 3))
				if err != nil {
					return err
				}
				c.ValueY_P_SP_CSI_RS_r18 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 3))
				if err != nil {
					return err
				}
				c.ValueY_A_CSI_RS_r18 = v
			}
			{
				v, err := d.DecodeEnumerated(codebookParametersetype2DopplerCSIR18EType2DopplerR18ScalingfactorR18Constraints)
				if err != nil {
					return err
				}
				c.Scalingfactor_r18 = v
			}
		}
	}

	// 3. eType2DopplerN4-r18: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.EType2DopplerN4_r18 = &struct {
				SupportedCSI_RS_ReportSettingList1_r18 []SupportedCSI_RS_ReportSetting_r18
				SupportedCSI_RS_ReportSettingList2_r18 []SupportedCSI_RS_ReportSetting_r18
			}{}
			c := ie.EType2DopplerN4_r18
			{
				seqOf := d.NewSequenceOfDecoder(codebookParametersetype2DopplerCSIR18EType2DopplerN4R18SupportedCSIRSReportSettingList1R18Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SupportedCSI_RS_ReportSettingList1_r18 = make([]SupportedCSI_RS_ReportSetting_r18, n)
				for i := int64(0); i < n; i++ {
					if err := c.SupportedCSI_RS_ReportSettingList1_r18[i].Decode(d); err != nil {
						return err
					}
				}
			}
			{
				seqOf := d.NewSequenceOfDecoder(codebookParametersetype2DopplerCSIR18EType2DopplerN4R18SupportedCSIRSReportSettingList2R18Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SupportedCSI_RS_ReportSettingList2_r18 = make([]SupportedCSI_RS_ReportSetting_r18, n)
				for i := int64(0); i < n; i++ {
					if err := c.SupportedCSI_RS_ReportSettingList2_r18[i].Decode(d); err != nil {
						return err
					}
				}
			}
		}
	}

	// 4. ddUnitSize-A-CSI-RS-CMR-r18: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(codebookParametersetype2DopplerCSIR18DdUnitSizeACSIRSCMRR18Constraints)
			if err != nil {
				return err
			}
			ie.DdUnitSize_A_CSI_RS_CMR_r18 = &idx
		}
	}

	// 5. maxNumberAperiodicCSI-RS-Resource-r18: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(codebookParametersetype2DopplerCSIR18MaxNumberAperiodicCSIRSResourceR18Constraints)
			if err != nil {
				return err
			}
			ie.MaxNumberAperiodicCSI_RS_Resource_r18 = &idx
		}
	}

	// 6. eType2DopplerR2-r18: sequence-of
	{
		if seq.IsComponentPresent(4) {
			seqOf := d.NewSequenceOfDecoder(codebookParametersetype2DopplerCSIR18EType2DopplerR2R18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.EType2DopplerR2_r18 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.EType2DopplerR2_r18[i] = v
			}
		}
	}

	// 7. eType2DopplerX1-r18: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(codebookParametersetype2DopplerCSIR18EType2DopplerX1R18Constraints)
			if err != nil {
				return err
			}
			ie.EType2DopplerX1_r18 = &idx
		}
	}

	// 8. eType2DopplerX2-r18: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(codebookParametersetype2DopplerCSIR18EType2DopplerX2R18Constraints)
			if err != nil {
				return err
			}
			ie.EType2DopplerX2_r18 = &idx
		}
	}

	// 9. eType2DopplerL-N4D1-r18: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(codebookParametersetype2DopplerCSIR18EType2DopplerLN4D1R18Constraints)
			if err != nil {
				return err
			}
			ie.EType2DopplerL_N4D1_r18 = &idx
		}
	}

	// 10. eType2DopplerL6-r18: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(codebookParametersetype2DopplerCSIR18EType2DopplerL6R18Constraints)
			if err != nil {
				return err
			}
			ie.EType2DopplerL6_r18 = &idx
		}
	}

	// 11. eType2DopplerR3R4-r18: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(codebookParametersetype2DopplerCSIR18EType2DopplerR3R4R18Constraints)
			if err != nil {
				return err
			}
			ie.EType2DopplerR3R4_r18 = &idx
		}
	}

	return nil
}

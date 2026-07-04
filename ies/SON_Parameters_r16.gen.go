// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SON-Parameters-r16 (line 25328).

var sONParametersR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "rach-Report-r16", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
	},
}

const (
	SON_Parameters_r16_Rach_Report_r16_Supported = 0
)

var sONParametersR16RachReportR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SON_Parameters_r16_Rach_Report_r16_Supported},
}

const (
	SON_Parameters_r16_Ext_RlfReportCHO_r17_Supported = 0
)

var sONParametersR16ExtRlfReportCHOR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SON_Parameters_r16_Ext_RlfReportCHO_r17_Supported},
}

const (
	SON_Parameters_r16_Ext_RlfReportDAPS_r17_Supported = 0
)

var sONParametersR16ExtRlfReportDAPSR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SON_Parameters_r16_Ext_RlfReportDAPS_r17_Supported},
}

const (
	SON_Parameters_r16_Ext_Success_HO_Report_r17_Supported = 0
)

var sONParametersR16ExtSuccessHOReportR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SON_Parameters_r16_Ext_Success_HO_Report_r17_Supported},
}

const (
	SON_Parameters_r16_Ext_TwoStepRACH_Report_r17_Supported = 0
)

var sONParametersR16ExtTwoStepRACHReportR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SON_Parameters_r16_Ext_TwoStepRACH_Report_r17_Supported},
}

const (
	SON_Parameters_r16_Ext_Pscell_MHI_Report_r17_Supported = 0
)

var sONParametersR16ExtPscellMHIReportR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SON_Parameters_r16_Ext_Pscell_MHI_Report_r17_Supported},
}

const (
	SON_Parameters_r16_Ext_OnDemandSI_Report_r17_Supported = 0
)

var sONParametersR16ExtOnDemandSIReportR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SON_Parameters_r16_Ext_OnDemandSI_Report_r17_Supported},
}

const (
	SON_Parameters_r16_Ext_Cef_ReportRedCap_r17_Supported = 0
)

var sONParametersR16ExtCefReportRedCapR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SON_Parameters_r16_Ext_Cef_ReportRedCap_r17_Supported},
}

const (
	SON_Parameters_r16_Ext_Rlf_ReportRedCap_r17_Supported = 0
)

var sONParametersR16ExtRlfReportRedCapR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SON_Parameters_r16_Ext_Rlf_ReportRedCap_r17_Supported},
}

const (
	SON_Parameters_r16_Ext_Spr_Report_r18_Supported = 0
)

var sONParametersR16ExtSprReportR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SON_Parameters_r16_Ext_Spr_Report_r18_Supported},
}

const (
	SON_Parameters_r16_Ext_SuccessIRAT_HO_Report_r18_Supported = 0
)

var sONParametersR16ExtSuccessIRATHOReportR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SON_Parameters_r16_Ext_SuccessIRAT_HO_Report_r18_Supported},
}

const (
	SON_Parameters_r16_Ext_RlfReportLTM_r19_Supported = 0
)

var sONParametersR16ExtRlfReportLTMR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SON_Parameters_r16_Ext_RlfReportLTM_r19_Supported},
}

const (
	SON_Parameters_r16_Ext_RlfReportCondHandoverWithCandSCG_r19_Supported = 0
)

var sONParametersR16ExtRlfReportCondHandoverWithCandSCGR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SON_Parameters_r16_Ext_RlfReportCondHandoverWithCandSCG_r19_Supported},
}

type SON_Parameters_r16 struct {
	Rach_Report_r16                      *int64
	RlfReportCHO_r17                     *int64
	RlfReportDAPS_r17                    *int64
	Success_HO_Report_r17                *int64
	TwoStepRACH_Report_r17               *int64
	Pscell_MHI_Report_r17                *int64
	OnDemandSI_Report_r17                *int64
	Cef_ReportRedCap_r17                 *int64
	Rlf_ReportRedCap_r17                 *int64
	Spr_Report_r18                       *int64
	SuccessIRAT_HO_Report_r18            *int64
	RlfReportLTM_r19                     *int64
	RlfReportCondHandoverWithCandSCG_r19 *int64
}

func (ie *SON_Parameters_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sONParametersR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.RlfReportCHO_r17 != nil || ie.RlfReportDAPS_r17 != nil || ie.Success_HO_Report_r17 != nil || ie.TwoStepRACH_Report_r17 != nil || ie.Pscell_MHI_Report_r17 != nil || ie.OnDemandSI_Report_r17 != nil
	hasExtGroup1 := ie.Cef_ReportRedCap_r17 != nil || ie.Rlf_ReportRedCap_r17 != nil
	hasExtGroup2 := ie.Spr_Report_r18 != nil || ie.SuccessIRAT_HO_Report_r18 != nil
	hasExtGroup3 := ie.RlfReportLTM_r19 != nil || ie.RlfReportCondHandoverWithCandSCG_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Rach_Report_r16 != nil}); err != nil {
		return err
	}

	// 3. rach-Report-r16: enumerated
	{
		if ie.Rach_Report_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Rach_Report_r16, sONParametersR16RachReportR16Constraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2, hasExtGroup3}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "rlfReportCHO-r17", Optional: true},
					{Name: "rlfReportDAPS-r17", Optional: true},
					{Name: "success-HO-Report-r17", Optional: true},
					{Name: "twoStepRACH-Report-r17", Optional: true},
					{Name: "pscell-MHI-Report-r17", Optional: true},
					{Name: "onDemandSI-Report-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.RlfReportCHO_r17 != nil, ie.RlfReportDAPS_r17 != nil, ie.Success_HO_Report_r17 != nil, ie.TwoStepRACH_Report_r17 != nil, ie.Pscell_MHI_Report_r17 != nil, ie.OnDemandSI_Report_r17 != nil}); err != nil {
				return err
			}

			if ie.RlfReportCHO_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.RlfReportCHO_r17, sONParametersR16ExtRlfReportCHOR17Constraints); err != nil {
					return err
				}
			}

			if ie.RlfReportDAPS_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.RlfReportDAPS_r17, sONParametersR16ExtRlfReportDAPSR17Constraints); err != nil {
					return err
				}
			}

			if ie.Success_HO_Report_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Success_HO_Report_r17, sONParametersR16ExtSuccessHOReportR17Constraints); err != nil {
					return err
				}
			}

			if ie.TwoStepRACH_Report_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.TwoStepRACH_Report_r17, sONParametersR16ExtTwoStepRACHReportR17Constraints); err != nil {
					return err
				}
			}

			if ie.Pscell_MHI_Report_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Pscell_MHI_Report_r17, sONParametersR16ExtPscellMHIReportR17Constraints); err != nil {
					return err
				}
			}

			if ie.OnDemandSI_Report_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.OnDemandSI_Report_r17, sONParametersR16ExtOnDemandSIReportR17Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 1:
		if hasExtGroup1 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "cef-ReportRedCap-r17", Optional: true},
					{Name: "rlf-ReportRedCap-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Cef_ReportRedCap_r17 != nil, ie.Rlf_ReportRedCap_r17 != nil}); err != nil {
				return err
			}

			if ie.Cef_ReportRedCap_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Cef_ReportRedCap_r17, sONParametersR16ExtCefReportRedCapR17Constraints); err != nil {
					return err
				}
			}

			if ie.Rlf_ReportRedCap_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Rlf_ReportRedCap_r17, sONParametersR16ExtRlfReportRedCapR17Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 2:
		if hasExtGroup2 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "spr-Report-r18", Optional: true},
					{Name: "successIRAT-HO-Report-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Spr_Report_r18 != nil, ie.SuccessIRAT_HO_Report_r18 != nil}); err != nil {
				return err
			}

			if ie.Spr_Report_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Spr_Report_r18, sONParametersR16ExtSprReportR18Constraints); err != nil {
					return err
				}
			}

			if ie.SuccessIRAT_HO_Report_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.SuccessIRAT_HO_Report_r18, sONParametersR16ExtSuccessIRATHOReportR18Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 3:
		if hasExtGroup3 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "rlfReportLTM-r19", Optional: true},
					{Name: "rlfReportCondHandoverWithCandSCG-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.RlfReportLTM_r19 != nil, ie.RlfReportCondHandoverWithCandSCG_r19 != nil}); err != nil {
				return err
			}

			if ie.RlfReportLTM_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.RlfReportLTM_r19, sONParametersR16ExtRlfReportLTMR19Constraints); err != nil {
					return err
				}
			}

			if ie.RlfReportCondHandoverWithCandSCG_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.RlfReportCondHandoverWithCandSCG_r19, sONParametersR16ExtRlfReportCondHandoverWithCandSCGR19Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SON_Parameters_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sONParametersR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. rach-Report-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(sONParametersR16RachReportR16Constraints)
			if err != nil {
				return err
			}
			ie.Rach_Report_r16 = &idx
		}
	}

	// Extension additions as open types.
	extValues, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	extValueIndex := 0

	// Extension group 0:
	if seq.IsExtensionPresent(0) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "rlfReportCHO-r17", Optional: true},
				{Name: "rlfReportDAPS-r17", Optional: true},
				{Name: "success-HO-Report-r17", Optional: true},
				{Name: "twoStepRACH-Report-r17", Optional: true},
				{Name: "pscell-MHI-Report-r17", Optional: true},
				{Name: "onDemandSI-Report-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(sONParametersR16ExtRlfReportCHOR17Constraints)
			if err != nil {
				return err
			}
			ie.RlfReportCHO_r17 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(sONParametersR16ExtRlfReportDAPSR17Constraints)
			if err != nil {
				return err
			}
			ie.RlfReportDAPS_r17 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(sONParametersR16ExtSuccessHOReportR17Constraints)
			if err != nil {
				return err
			}
			ie.Success_HO_Report_r17 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(sONParametersR16ExtTwoStepRACHReportR17Constraints)
			if err != nil {
				return err
			}
			ie.TwoStepRACH_Report_r17 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(sONParametersR16ExtPscellMHIReportR17Constraints)
			if err != nil {
				return err
			}
			ie.Pscell_MHI_Report_r17 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(sONParametersR16ExtOnDemandSIReportR17Constraints)
			if err != nil {
				return err
			}
			ie.OnDemandSI_Report_r17 = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "cef-ReportRedCap-r17", Optional: true},
				{Name: "rlf-ReportRedCap-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(sONParametersR16ExtCefReportRedCapR17Constraints)
			if err != nil {
				return err
			}
			ie.Cef_ReportRedCap_r17 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(sONParametersR16ExtRlfReportRedCapR17Constraints)
			if err != nil {
				return err
			}
			ie.Rlf_ReportRedCap_r17 = &v
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "spr-Report-r18", Optional: true},
				{Name: "successIRAT-HO-Report-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(sONParametersR16ExtSprReportR18Constraints)
			if err != nil {
				return err
			}
			ie.Spr_Report_r18 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(sONParametersR16ExtSuccessIRATHOReportR18Constraints)
			if err != nil {
				return err
			}
			ie.SuccessIRAT_HO_Report_r18 = &v
		}
	}

	// Extension group 3:
	if seq.IsExtensionPresent(3) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "rlfReportLTM-r19", Optional: true},
				{Name: "rlfReportCondHandoverWithCandSCG-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(sONParametersR16ExtRlfReportLTMR19Constraints)
			if err != nil {
				return err
			}
			ie.RlfReportLTM_r19 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(sONParametersR16ExtRlfReportCondHandoverWithCandSCGR19Constraints)
			if err != nil {
				return err
			}
			ie.RlfReportCondHandoverWithCandSCG_r19 = &v
		}
	}

	return nil
}

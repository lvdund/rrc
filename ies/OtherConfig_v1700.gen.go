// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: OtherConfig-v1700 (line 26330).

var otherConfigV1700Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ul-GapFR2-PreferenceConfig-r17", Optional: true},
		{Name: "musim-GapAssistanceConfig-r17", Optional: true},
		{Name: "musim-LeaveAssistanceConfig-r17", Optional: true},
		{Name: "successHO-Config-r17", Optional: true},
		{Name: "maxBW-PreferenceConfigFR2-2-r17", Optional: true},
		{Name: "maxMIMO-LayerPreferenceConfigFR2-2-r17", Optional: true},
		{Name: "minSchedulingOffsetPreferenceConfigExt-r17", Optional: true},
		{Name: "rlm-RelaxationReportingConfig-r17", Optional: true},
		{Name: "bfd-RelaxationReportingConfig-r17", Optional: true},
		{Name: "scg-DeactivationPreferenceConfig-r17", Optional: true},
		{Name: "rrm-MeasRelaxationReportingConfig-r17", Optional: true},
		{Name: "propDelayDiffReportConfig-r17", Optional: true},
	},
}

const (
	OtherConfig_v1700_Ul_GapFR2_PreferenceConfig_r17_True = 0
)

var otherConfigV1700UlGapFR2PreferenceConfigR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{OtherConfig_v1700_Ul_GapFR2_PreferenceConfig_r17_True},
}

var otherConfig_v1700MusimGapAssistanceConfigR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	OtherConfig_v1700_Musim_GapAssistanceConfig_r17_Release = 0
	OtherConfig_v1700_Musim_GapAssistanceConfig_r17_Setup   = 1
)

var otherConfig_v1700MusimLeaveAssistanceConfigR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	OtherConfig_v1700_Musim_LeaveAssistanceConfig_r17_Release = 0
	OtherConfig_v1700_Musim_LeaveAssistanceConfig_r17_Setup   = 1
)

var otherConfig_v1700SuccessHOConfigR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	OtherConfig_v1700_SuccessHO_Config_r17_Release = 0
	OtherConfig_v1700_SuccessHO_Config_r17_Setup   = 1
)

const (
	OtherConfig_v1700_MaxBW_PreferenceConfigFR2_2_r17_True = 0
)

var otherConfigV1700MaxBWPreferenceConfigFR22R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{OtherConfig_v1700_MaxBW_PreferenceConfigFR2_2_r17_True},
}

const (
	OtherConfig_v1700_MaxMIMO_LayerPreferenceConfigFR2_2_r17_True = 0
)

var otherConfigV1700MaxMIMOLayerPreferenceConfigFR22R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{OtherConfig_v1700_MaxMIMO_LayerPreferenceConfigFR2_2_r17_True},
}

const (
	OtherConfig_v1700_MinSchedulingOffsetPreferenceConfigExt_r17_True = 0
)

var otherConfigV1700MinSchedulingOffsetPreferenceConfigExtR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{OtherConfig_v1700_MinSchedulingOffsetPreferenceConfigExt_r17_True},
}

var otherConfig_v1700RlmRelaxationReportingConfigR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	OtherConfig_v1700_Rlm_RelaxationReportingConfig_r17_Release = 0
	OtherConfig_v1700_Rlm_RelaxationReportingConfig_r17_Setup   = 1
)

var otherConfig_v1700BfdRelaxationReportingConfigR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	OtherConfig_v1700_Bfd_RelaxationReportingConfig_r17_Release = 0
	OtherConfig_v1700_Bfd_RelaxationReportingConfig_r17_Setup   = 1
)

var otherConfig_v1700ScgDeactivationPreferenceConfigR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	OtherConfig_v1700_Scg_DeactivationPreferenceConfig_r17_Release = 0
	OtherConfig_v1700_Scg_DeactivationPreferenceConfig_r17_Setup   = 1
)

var otherConfig_v1700RrmMeasRelaxationReportingConfigR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	OtherConfig_v1700_Rrm_MeasRelaxationReportingConfig_r17_Release = 0
	OtherConfig_v1700_Rrm_MeasRelaxationReportingConfig_r17_Setup   = 1
)

var otherConfig_v1700PropDelayDiffReportConfigR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	OtherConfig_v1700_PropDelayDiffReportConfig_r17_Release = 0
	OtherConfig_v1700_PropDelayDiffReportConfig_r17_Setup   = 1
)

type OtherConfig_v1700 struct {
	Ul_GapFR2_PreferenceConfig_r17 *int64
	Musim_GapAssistanceConfig_r17  *struct {
		Choice  int
		Release *struct{}
		Setup   *MUSIM_GapAssistanceConfig_r17
	}
	Musim_LeaveAssistanceConfig_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *MUSIM_LeaveAssistanceConfig_r17
	}
	SuccessHO_Config_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *SuccessHO_Config_r17
	}
	MaxBW_PreferenceConfigFR2_2_r17            *int64
	MaxMIMO_LayerPreferenceConfigFR2_2_r17     *int64
	MinSchedulingOffsetPreferenceConfigExt_r17 *int64
	Rlm_RelaxationReportingConfig_r17          *struct {
		Choice  int
		Release *struct{}
		Setup   *RLM_RelaxationReportingConfig_r17
	}
	Bfd_RelaxationReportingConfig_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *BFD_RelaxationReportingConfig_r17
	}
	Scg_DeactivationPreferenceConfig_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *SCG_DeactivationPreferenceConfig_r17
	}
	Rrm_MeasRelaxationReportingConfig_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *RRM_MeasRelaxationReportingConfig_r17
	}
	PropDelayDiffReportConfig_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *PropDelayDiffReportConfig_r17
	}
}

func (ie *OtherConfig_v1700) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(otherConfigV1700Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ul_GapFR2_PreferenceConfig_r17 != nil, ie.Musim_GapAssistanceConfig_r17 != nil, ie.Musim_LeaveAssistanceConfig_r17 != nil, ie.SuccessHO_Config_r17 != nil, ie.MaxBW_PreferenceConfigFR2_2_r17 != nil, ie.MaxMIMO_LayerPreferenceConfigFR2_2_r17 != nil, ie.MinSchedulingOffsetPreferenceConfigExt_r17 != nil, ie.Rlm_RelaxationReportingConfig_r17 != nil, ie.Bfd_RelaxationReportingConfig_r17 != nil, ie.Scg_DeactivationPreferenceConfig_r17 != nil, ie.Rrm_MeasRelaxationReportingConfig_r17 != nil, ie.PropDelayDiffReportConfig_r17 != nil}); err != nil {
		return err
	}

	// 2. ul-GapFR2-PreferenceConfig-r17: enumerated
	{
		if ie.Ul_GapFR2_PreferenceConfig_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Ul_GapFR2_PreferenceConfig_r17, otherConfigV1700UlGapFR2PreferenceConfigR17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. musim-GapAssistanceConfig-r17: choice
	{
		if ie.Musim_GapAssistanceConfig_r17 != nil {
			choiceEnc := e.NewChoiceEncoder(otherConfig_v1700MusimGapAssistanceConfigR17Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Musim_GapAssistanceConfig_r17).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Musim_GapAssistanceConfig_r17).Choice {
			case OtherConfig_v1700_Musim_GapAssistanceConfig_r17_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1700_Musim_GapAssistanceConfig_r17_Setup:
				if err := (*ie.Musim_GapAssistanceConfig_r17).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Musim_GapAssistanceConfig_r17).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 4. musim-LeaveAssistanceConfig-r17: choice
	{
		if ie.Musim_LeaveAssistanceConfig_r17 != nil {
			choiceEnc := e.NewChoiceEncoder(otherConfig_v1700MusimLeaveAssistanceConfigR17Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Musim_LeaveAssistanceConfig_r17).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Musim_LeaveAssistanceConfig_r17).Choice {
			case OtherConfig_v1700_Musim_LeaveAssistanceConfig_r17_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1700_Musim_LeaveAssistanceConfig_r17_Setup:
				if err := (*ie.Musim_LeaveAssistanceConfig_r17).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Musim_LeaveAssistanceConfig_r17).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 5. successHO-Config-r17: choice
	{
		if ie.SuccessHO_Config_r17 != nil {
			choiceEnc := e.NewChoiceEncoder(otherConfig_v1700SuccessHOConfigR17Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.SuccessHO_Config_r17).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.SuccessHO_Config_r17).Choice {
			case OtherConfig_v1700_SuccessHO_Config_r17_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1700_SuccessHO_Config_r17_Setup:
				if err := (*ie.SuccessHO_Config_r17).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.SuccessHO_Config_r17).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 6. maxBW-PreferenceConfigFR2-2-r17: enumerated
	{
		if ie.MaxBW_PreferenceConfigFR2_2_r17 != nil {
			if err := e.EncodeEnumerated(*ie.MaxBW_PreferenceConfigFR2_2_r17, otherConfigV1700MaxBWPreferenceConfigFR22R17Constraints); err != nil {
				return err
			}
		}
	}

	// 7. maxMIMO-LayerPreferenceConfigFR2-2-r17: enumerated
	{
		if ie.MaxMIMO_LayerPreferenceConfigFR2_2_r17 != nil {
			if err := e.EncodeEnumerated(*ie.MaxMIMO_LayerPreferenceConfigFR2_2_r17, otherConfigV1700MaxMIMOLayerPreferenceConfigFR22R17Constraints); err != nil {
				return err
			}
		}
	}

	// 8. minSchedulingOffsetPreferenceConfigExt-r17: enumerated
	{
		if ie.MinSchedulingOffsetPreferenceConfigExt_r17 != nil {
			if err := e.EncodeEnumerated(*ie.MinSchedulingOffsetPreferenceConfigExt_r17, otherConfigV1700MinSchedulingOffsetPreferenceConfigExtR17Constraints); err != nil {
				return err
			}
		}
	}

	// 9. rlm-RelaxationReportingConfig-r17: choice
	{
		if ie.Rlm_RelaxationReportingConfig_r17 != nil {
			choiceEnc := e.NewChoiceEncoder(otherConfig_v1700RlmRelaxationReportingConfigR17Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Rlm_RelaxationReportingConfig_r17).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Rlm_RelaxationReportingConfig_r17).Choice {
			case OtherConfig_v1700_Rlm_RelaxationReportingConfig_r17_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1700_Rlm_RelaxationReportingConfig_r17_Setup:
				if err := (*ie.Rlm_RelaxationReportingConfig_r17).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Rlm_RelaxationReportingConfig_r17).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 10. bfd-RelaxationReportingConfig-r17: choice
	{
		if ie.Bfd_RelaxationReportingConfig_r17 != nil {
			choiceEnc := e.NewChoiceEncoder(otherConfig_v1700BfdRelaxationReportingConfigR17Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Bfd_RelaxationReportingConfig_r17).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Bfd_RelaxationReportingConfig_r17).Choice {
			case OtherConfig_v1700_Bfd_RelaxationReportingConfig_r17_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1700_Bfd_RelaxationReportingConfig_r17_Setup:
				if err := (*ie.Bfd_RelaxationReportingConfig_r17).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Bfd_RelaxationReportingConfig_r17).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 11. scg-DeactivationPreferenceConfig-r17: choice
	{
		if ie.Scg_DeactivationPreferenceConfig_r17 != nil {
			choiceEnc := e.NewChoiceEncoder(otherConfig_v1700ScgDeactivationPreferenceConfigR17Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Scg_DeactivationPreferenceConfig_r17).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Scg_DeactivationPreferenceConfig_r17).Choice {
			case OtherConfig_v1700_Scg_DeactivationPreferenceConfig_r17_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1700_Scg_DeactivationPreferenceConfig_r17_Setup:
				if err := (*ie.Scg_DeactivationPreferenceConfig_r17).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Scg_DeactivationPreferenceConfig_r17).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 12. rrm-MeasRelaxationReportingConfig-r17: choice
	{
		if ie.Rrm_MeasRelaxationReportingConfig_r17 != nil {
			choiceEnc := e.NewChoiceEncoder(otherConfig_v1700RrmMeasRelaxationReportingConfigR17Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Rrm_MeasRelaxationReportingConfig_r17).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Rrm_MeasRelaxationReportingConfig_r17).Choice {
			case OtherConfig_v1700_Rrm_MeasRelaxationReportingConfig_r17_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1700_Rrm_MeasRelaxationReportingConfig_r17_Setup:
				if err := (*ie.Rrm_MeasRelaxationReportingConfig_r17).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Rrm_MeasRelaxationReportingConfig_r17).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 13. propDelayDiffReportConfig-r17: choice
	{
		if ie.PropDelayDiffReportConfig_r17 != nil {
			choiceEnc := e.NewChoiceEncoder(otherConfig_v1700PropDelayDiffReportConfigR17Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.PropDelayDiffReportConfig_r17).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.PropDelayDiffReportConfig_r17).Choice {
			case OtherConfig_v1700_PropDelayDiffReportConfig_r17_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1700_PropDelayDiffReportConfig_r17_Setup:
				if err := (*ie.PropDelayDiffReportConfig_r17).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.PropDelayDiffReportConfig_r17).Choice), Constraint: "index out of range"}
			}
		}
	}

	return nil
}

func (ie *OtherConfig_v1700) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(otherConfigV1700Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ul-GapFR2-PreferenceConfig-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(otherConfigV1700UlGapFR2PreferenceConfigR17Constraints)
			if err != nil {
				return err
			}
			ie.Ul_GapFR2_PreferenceConfig_r17 = &idx
		}
	}

	// 3. musim-GapAssistanceConfig-r17: choice
	{
		if seq.IsComponentPresent(1) {
			ie.Musim_GapAssistanceConfig_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *MUSIM_GapAssistanceConfig_r17
			}{}
			choiceDec := d.NewChoiceDecoder(otherConfig_v1700MusimGapAssistanceConfigR17Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Musim_GapAssistanceConfig_r17).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case OtherConfig_v1700_Musim_GapAssistanceConfig_r17_Release:
				(*ie.Musim_GapAssistanceConfig_r17).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1700_Musim_GapAssistanceConfig_r17_Setup:
				(*ie.Musim_GapAssistanceConfig_r17).Setup = new(MUSIM_GapAssistanceConfig_r17)
				if err := (*ie.Musim_GapAssistanceConfig_r17).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. musim-LeaveAssistanceConfig-r17: choice
	{
		if seq.IsComponentPresent(2) {
			ie.Musim_LeaveAssistanceConfig_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *MUSIM_LeaveAssistanceConfig_r17
			}{}
			choiceDec := d.NewChoiceDecoder(otherConfig_v1700MusimLeaveAssistanceConfigR17Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Musim_LeaveAssistanceConfig_r17).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case OtherConfig_v1700_Musim_LeaveAssistanceConfig_r17_Release:
				(*ie.Musim_LeaveAssistanceConfig_r17).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1700_Musim_LeaveAssistanceConfig_r17_Setup:
				(*ie.Musim_LeaveAssistanceConfig_r17).Setup = new(MUSIM_LeaveAssistanceConfig_r17)
				if err := (*ie.Musim_LeaveAssistanceConfig_r17).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. successHO-Config-r17: choice
	{
		if seq.IsComponentPresent(3) {
			ie.SuccessHO_Config_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SuccessHO_Config_r17
			}{}
			choiceDec := d.NewChoiceDecoder(otherConfig_v1700SuccessHOConfigR17Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.SuccessHO_Config_r17).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case OtherConfig_v1700_SuccessHO_Config_r17_Release:
				(*ie.SuccessHO_Config_r17).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1700_SuccessHO_Config_r17_Setup:
				(*ie.SuccessHO_Config_r17).Setup = new(SuccessHO_Config_r17)
				if err := (*ie.SuccessHO_Config_r17).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. maxBW-PreferenceConfigFR2-2-r17: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(otherConfigV1700MaxBWPreferenceConfigFR22R17Constraints)
			if err != nil {
				return err
			}
			ie.MaxBW_PreferenceConfigFR2_2_r17 = &idx
		}
	}

	// 7. maxMIMO-LayerPreferenceConfigFR2-2-r17: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(otherConfigV1700MaxMIMOLayerPreferenceConfigFR22R17Constraints)
			if err != nil {
				return err
			}
			ie.MaxMIMO_LayerPreferenceConfigFR2_2_r17 = &idx
		}
	}

	// 8. minSchedulingOffsetPreferenceConfigExt-r17: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(otherConfigV1700MinSchedulingOffsetPreferenceConfigExtR17Constraints)
			if err != nil {
				return err
			}
			ie.MinSchedulingOffsetPreferenceConfigExt_r17 = &idx
		}
	}

	// 9. rlm-RelaxationReportingConfig-r17: choice
	{
		if seq.IsComponentPresent(7) {
			ie.Rlm_RelaxationReportingConfig_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *RLM_RelaxationReportingConfig_r17
			}{}
			choiceDec := d.NewChoiceDecoder(otherConfig_v1700RlmRelaxationReportingConfigR17Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Rlm_RelaxationReportingConfig_r17).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case OtherConfig_v1700_Rlm_RelaxationReportingConfig_r17_Release:
				(*ie.Rlm_RelaxationReportingConfig_r17).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1700_Rlm_RelaxationReportingConfig_r17_Setup:
				(*ie.Rlm_RelaxationReportingConfig_r17).Setup = new(RLM_RelaxationReportingConfig_r17)
				if err := (*ie.Rlm_RelaxationReportingConfig_r17).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 10. bfd-RelaxationReportingConfig-r17: choice
	{
		if seq.IsComponentPresent(8) {
			ie.Bfd_RelaxationReportingConfig_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *BFD_RelaxationReportingConfig_r17
			}{}
			choiceDec := d.NewChoiceDecoder(otherConfig_v1700BfdRelaxationReportingConfigR17Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Bfd_RelaxationReportingConfig_r17).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case OtherConfig_v1700_Bfd_RelaxationReportingConfig_r17_Release:
				(*ie.Bfd_RelaxationReportingConfig_r17).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1700_Bfd_RelaxationReportingConfig_r17_Setup:
				(*ie.Bfd_RelaxationReportingConfig_r17).Setup = new(BFD_RelaxationReportingConfig_r17)
				if err := (*ie.Bfd_RelaxationReportingConfig_r17).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 11. scg-DeactivationPreferenceConfig-r17: choice
	{
		if seq.IsComponentPresent(9) {
			ie.Scg_DeactivationPreferenceConfig_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SCG_DeactivationPreferenceConfig_r17
			}{}
			choiceDec := d.NewChoiceDecoder(otherConfig_v1700ScgDeactivationPreferenceConfigR17Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Scg_DeactivationPreferenceConfig_r17).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case OtherConfig_v1700_Scg_DeactivationPreferenceConfig_r17_Release:
				(*ie.Scg_DeactivationPreferenceConfig_r17).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1700_Scg_DeactivationPreferenceConfig_r17_Setup:
				(*ie.Scg_DeactivationPreferenceConfig_r17).Setup = new(SCG_DeactivationPreferenceConfig_r17)
				if err := (*ie.Scg_DeactivationPreferenceConfig_r17).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 12. rrm-MeasRelaxationReportingConfig-r17: choice
	{
		if seq.IsComponentPresent(10) {
			ie.Rrm_MeasRelaxationReportingConfig_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *RRM_MeasRelaxationReportingConfig_r17
			}{}
			choiceDec := d.NewChoiceDecoder(otherConfig_v1700RrmMeasRelaxationReportingConfigR17Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Rrm_MeasRelaxationReportingConfig_r17).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case OtherConfig_v1700_Rrm_MeasRelaxationReportingConfig_r17_Release:
				(*ie.Rrm_MeasRelaxationReportingConfig_r17).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1700_Rrm_MeasRelaxationReportingConfig_r17_Setup:
				(*ie.Rrm_MeasRelaxationReportingConfig_r17).Setup = new(RRM_MeasRelaxationReportingConfig_r17)
				if err := (*ie.Rrm_MeasRelaxationReportingConfig_r17).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 13. propDelayDiffReportConfig-r17: choice
	{
		if seq.IsComponentPresent(11) {
			ie.PropDelayDiffReportConfig_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PropDelayDiffReportConfig_r17
			}{}
			choiceDec := d.NewChoiceDecoder(otherConfig_v1700PropDelayDiffReportConfigR17Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.PropDelayDiffReportConfig_r17).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case OtherConfig_v1700_PropDelayDiffReportConfig_r17_Release:
				(*ie.PropDelayDiffReportConfig_r17).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1700_PropDelayDiffReportConfig_r17_Setup:
				(*ie.PropDelayDiffReportConfig_r17).Setup = new(PropDelayDiffReportConfig_r17)
				if err := (*ie.PropDelayDiffReportConfig_r17).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

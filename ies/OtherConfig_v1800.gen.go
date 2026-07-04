// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: OtherConfig-v1800 (line 26345).

var otherConfigV1800Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "idc-AssistanceConfig-v1800", Optional: true},
		{Name: "multiRx-PreferenceReportingConfigFR2-r18", Optional: true},
		{Name: "aerial-FlightPathAvailabilityConfig-r18", Optional: true},
		{Name: "ul-TrafficInfoReportingConfig-r18", Optional: true},
		{Name: "n3c-RelayUE-InfoReportConfig-r18", Optional: true},
		{Name: "successPSCell-Config-r18", Optional: true},
		{Name: "sn-InitiatedPSCellChange-r18", Optional: true},
		{Name: "musim-GapPriorityAssistanceConfig-r18", Optional: true},
		{Name: "musim-CapabilityRestrictionConfig-r18", Optional: true},
	},
}

var otherConfig_v1800IdcAssistanceConfigV1800Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	OtherConfig_v1800_Idc_AssistanceConfig_v1800_Release = 0
	OtherConfig_v1800_Idc_AssistanceConfig_v1800_Setup   = 1
)

var otherConfig_v1800MultiRxPreferenceReportingConfigFR2R18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	OtherConfig_v1800_MultiRx_PreferenceReportingConfigFR2_r18_Release = 0
	OtherConfig_v1800_MultiRx_PreferenceReportingConfigFR2_r18_Setup   = 1
)

const (
	OtherConfig_v1800_Aerial_FlightPathAvailabilityConfig_r18_True = 0
)

var otherConfigV1800AerialFlightPathAvailabilityConfigR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{OtherConfig_v1800_Aerial_FlightPathAvailabilityConfig_r18_True},
}

var otherConfig_v1800UlTrafficInfoReportingConfigR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	OtherConfig_v1800_Ul_TrafficInfoReportingConfig_r18_Release = 0
	OtherConfig_v1800_Ul_TrafficInfoReportingConfig_r18_Setup   = 1
)

const (
	OtherConfig_v1800_N3c_RelayUE_InfoReportConfig_r18_True = 0
)

var otherConfigV1800N3cRelayUEInfoReportConfigR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{OtherConfig_v1800_N3c_RelayUE_InfoReportConfig_r18_True},
}

var otherConfig_v1800SuccessPSCellConfigR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	OtherConfig_v1800_SuccessPSCell_Config_r18_Release = 0
	OtherConfig_v1800_SuccessPSCell_Config_r18_Setup   = 1
)

const (
	OtherConfig_v1800_Sn_InitiatedPSCellChange_r18_True = 0
)

var otherConfigV1800SnInitiatedPSCellChangeR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{OtherConfig_v1800_Sn_InitiatedPSCellChange_r18_True},
}

const (
	OtherConfig_v1800_Musim_GapPriorityAssistanceConfig_r18_True = 0
)

var otherConfigV1800MusimGapPriorityAssistanceConfigR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{OtherConfig_v1800_Musim_GapPriorityAssistanceConfig_r18_True},
}

var otherConfig_v1800MusimCapabilityRestrictionConfigR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	OtherConfig_v1800_Musim_CapabilityRestrictionConfig_r18_Release = 0
	OtherConfig_v1800_Musim_CapabilityRestrictionConfig_r18_Setup   = 1
)

type OtherConfig_v1800 struct {
	Idc_AssistanceConfig_v1800 *struct {
		Choice  int
		Release *struct{}
		Setup   *IDC_AssistanceConfig_v1800
	}
	MultiRx_PreferenceReportingConfigFR2_r18 *struct {
		Choice  int
		Release *struct{}
		Setup   *MultiRx_PreferenceReportingConfigFR2_r18
	}
	Aerial_FlightPathAvailabilityConfig_r18 *int64
	Ul_TrafficInfoReportingConfig_r18       *struct {
		Choice  int
		Release *struct{}
		Setup   *UL_TrafficInfoReportingConfig_r18
	}
	N3c_RelayUE_InfoReportConfig_r18 *int64
	SuccessPSCell_Config_r18         *struct {
		Choice  int
		Release *struct{}
		Setup   *SuccessPSCell_Config_r18
	}
	Sn_InitiatedPSCellChange_r18          *int64
	Musim_GapPriorityAssistanceConfig_r18 *int64
	Musim_CapabilityRestrictionConfig_r18 *struct {
		Choice  int
		Release *struct{}
		Setup   *MUSIM_CapabilityRestrictionConfig_r18
	}
}

func (ie *OtherConfig_v1800) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(otherConfigV1800Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Idc_AssistanceConfig_v1800 != nil, ie.MultiRx_PreferenceReportingConfigFR2_r18 != nil, ie.Aerial_FlightPathAvailabilityConfig_r18 != nil, ie.Ul_TrafficInfoReportingConfig_r18 != nil, ie.N3c_RelayUE_InfoReportConfig_r18 != nil, ie.SuccessPSCell_Config_r18 != nil, ie.Sn_InitiatedPSCellChange_r18 != nil, ie.Musim_GapPriorityAssistanceConfig_r18 != nil, ie.Musim_CapabilityRestrictionConfig_r18 != nil}); err != nil {
		return err
	}

	// 2. idc-AssistanceConfig-v1800: choice
	{
		if ie.Idc_AssistanceConfig_v1800 != nil {
			choiceEnc := e.NewChoiceEncoder(otherConfig_v1800IdcAssistanceConfigV1800Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Idc_AssistanceConfig_v1800).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Idc_AssistanceConfig_v1800).Choice {
			case OtherConfig_v1800_Idc_AssistanceConfig_v1800_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1800_Idc_AssistanceConfig_v1800_Setup:
				if err := (*ie.Idc_AssistanceConfig_v1800).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Idc_AssistanceConfig_v1800).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 3. multiRx-PreferenceReportingConfigFR2-r18: choice
	{
		if ie.MultiRx_PreferenceReportingConfigFR2_r18 != nil {
			choiceEnc := e.NewChoiceEncoder(otherConfig_v1800MultiRxPreferenceReportingConfigFR2R18Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.MultiRx_PreferenceReportingConfigFR2_r18).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.MultiRx_PreferenceReportingConfigFR2_r18).Choice {
			case OtherConfig_v1800_MultiRx_PreferenceReportingConfigFR2_r18_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1800_MultiRx_PreferenceReportingConfigFR2_r18_Setup:
				if err := (*ie.MultiRx_PreferenceReportingConfigFR2_r18).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.MultiRx_PreferenceReportingConfigFR2_r18).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 4. aerial-FlightPathAvailabilityConfig-r18: enumerated
	{
		if ie.Aerial_FlightPathAvailabilityConfig_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Aerial_FlightPathAvailabilityConfig_r18, otherConfigV1800AerialFlightPathAvailabilityConfigR18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. ul-TrafficInfoReportingConfig-r18: choice
	{
		if ie.Ul_TrafficInfoReportingConfig_r18 != nil {
			choiceEnc := e.NewChoiceEncoder(otherConfig_v1800UlTrafficInfoReportingConfigR18Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Ul_TrafficInfoReportingConfig_r18).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Ul_TrafficInfoReportingConfig_r18).Choice {
			case OtherConfig_v1800_Ul_TrafficInfoReportingConfig_r18_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1800_Ul_TrafficInfoReportingConfig_r18_Setup:
				if err := (*ie.Ul_TrafficInfoReportingConfig_r18).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Ul_TrafficInfoReportingConfig_r18).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 6. n3c-RelayUE-InfoReportConfig-r18: enumerated
	{
		if ie.N3c_RelayUE_InfoReportConfig_r18 != nil {
			if err := e.EncodeEnumerated(*ie.N3c_RelayUE_InfoReportConfig_r18, otherConfigV1800N3cRelayUEInfoReportConfigR18Constraints); err != nil {
				return err
			}
		}
	}

	// 7. successPSCell-Config-r18: choice
	{
		if ie.SuccessPSCell_Config_r18 != nil {
			choiceEnc := e.NewChoiceEncoder(otherConfig_v1800SuccessPSCellConfigR18Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.SuccessPSCell_Config_r18).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.SuccessPSCell_Config_r18).Choice {
			case OtherConfig_v1800_SuccessPSCell_Config_r18_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1800_SuccessPSCell_Config_r18_Setup:
				if err := (*ie.SuccessPSCell_Config_r18).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.SuccessPSCell_Config_r18).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 8. sn-InitiatedPSCellChange-r18: enumerated
	{
		if ie.Sn_InitiatedPSCellChange_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Sn_InitiatedPSCellChange_r18, otherConfigV1800SnInitiatedPSCellChangeR18Constraints); err != nil {
				return err
			}
		}
	}

	// 9. musim-GapPriorityAssistanceConfig-r18: enumerated
	{
		if ie.Musim_GapPriorityAssistanceConfig_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Musim_GapPriorityAssistanceConfig_r18, otherConfigV1800MusimGapPriorityAssistanceConfigR18Constraints); err != nil {
				return err
			}
		}
	}

	// 10. musim-CapabilityRestrictionConfig-r18: choice
	{
		if ie.Musim_CapabilityRestrictionConfig_r18 != nil {
			choiceEnc := e.NewChoiceEncoder(otherConfig_v1800MusimCapabilityRestrictionConfigR18Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Musim_CapabilityRestrictionConfig_r18).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Musim_CapabilityRestrictionConfig_r18).Choice {
			case OtherConfig_v1800_Musim_CapabilityRestrictionConfig_r18_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1800_Musim_CapabilityRestrictionConfig_r18_Setup:
				if err := (*ie.Musim_CapabilityRestrictionConfig_r18).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Musim_CapabilityRestrictionConfig_r18).Choice), Constraint: "index out of range"}
			}
		}
	}

	return nil
}

func (ie *OtherConfig_v1800) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(otherConfigV1800Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. idc-AssistanceConfig-v1800: choice
	{
		if seq.IsComponentPresent(0) {
			ie.Idc_AssistanceConfig_v1800 = &struct {
				Choice  int
				Release *struct{}
				Setup   *IDC_AssistanceConfig_v1800
			}{}
			choiceDec := d.NewChoiceDecoder(otherConfig_v1800IdcAssistanceConfigV1800Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Idc_AssistanceConfig_v1800).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case OtherConfig_v1800_Idc_AssistanceConfig_v1800_Release:
				(*ie.Idc_AssistanceConfig_v1800).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1800_Idc_AssistanceConfig_v1800_Setup:
				(*ie.Idc_AssistanceConfig_v1800).Setup = new(IDC_AssistanceConfig_v1800)
				if err := (*ie.Idc_AssistanceConfig_v1800).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. multiRx-PreferenceReportingConfigFR2-r18: choice
	{
		if seq.IsComponentPresent(1) {
			ie.MultiRx_PreferenceReportingConfigFR2_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *MultiRx_PreferenceReportingConfigFR2_r18
			}{}
			choiceDec := d.NewChoiceDecoder(otherConfig_v1800MultiRxPreferenceReportingConfigFR2R18Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.MultiRx_PreferenceReportingConfigFR2_r18).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case OtherConfig_v1800_MultiRx_PreferenceReportingConfigFR2_r18_Release:
				(*ie.MultiRx_PreferenceReportingConfigFR2_r18).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1800_MultiRx_PreferenceReportingConfigFR2_r18_Setup:
				(*ie.MultiRx_PreferenceReportingConfigFR2_r18).Setup = new(MultiRx_PreferenceReportingConfigFR2_r18)
				if err := (*ie.MultiRx_PreferenceReportingConfigFR2_r18).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. aerial-FlightPathAvailabilityConfig-r18: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(otherConfigV1800AerialFlightPathAvailabilityConfigR18Constraints)
			if err != nil {
				return err
			}
			ie.Aerial_FlightPathAvailabilityConfig_r18 = &idx
		}
	}

	// 5. ul-TrafficInfoReportingConfig-r18: choice
	{
		if seq.IsComponentPresent(3) {
			ie.Ul_TrafficInfoReportingConfig_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *UL_TrafficInfoReportingConfig_r18
			}{}
			choiceDec := d.NewChoiceDecoder(otherConfig_v1800UlTrafficInfoReportingConfigR18Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Ul_TrafficInfoReportingConfig_r18).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case OtherConfig_v1800_Ul_TrafficInfoReportingConfig_r18_Release:
				(*ie.Ul_TrafficInfoReportingConfig_r18).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1800_Ul_TrafficInfoReportingConfig_r18_Setup:
				(*ie.Ul_TrafficInfoReportingConfig_r18).Setup = new(UL_TrafficInfoReportingConfig_r18)
				if err := (*ie.Ul_TrafficInfoReportingConfig_r18).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. n3c-RelayUE-InfoReportConfig-r18: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(otherConfigV1800N3cRelayUEInfoReportConfigR18Constraints)
			if err != nil {
				return err
			}
			ie.N3c_RelayUE_InfoReportConfig_r18 = &idx
		}
	}

	// 7. successPSCell-Config-r18: choice
	{
		if seq.IsComponentPresent(5) {
			ie.SuccessPSCell_Config_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SuccessPSCell_Config_r18
			}{}
			choiceDec := d.NewChoiceDecoder(otherConfig_v1800SuccessPSCellConfigR18Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.SuccessPSCell_Config_r18).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case OtherConfig_v1800_SuccessPSCell_Config_r18_Release:
				(*ie.SuccessPSCell_Config_r18).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1800_SuccessPSCell_Config_r18_Setup:
				(*ie.SuccessPSCell_Config_r18).Setup = new(SuccessPSCell_Config_r18)
				if err := (*ie.SuccessPSCell_Config_r18).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 8. sn-InitiatedPSCellChange-r18: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(otherConfigV1800SnInitiatedPSCellChangeR18Constraints)
			if err != nil {
				return err
			}
			ie.Sn_InitiatedPSCellChange_r18 = &idx
		}
	}

	// 9. musim-GapPriorityAssistanceConfig-r18: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(otherConfigV1800MusimGapPriorityAssistanceConfigR18Constraints)
			if err != nil {
				return err
			}
			ie.Musim_GapPriorityAssistanceConfig_r18 = &idx
		}
	}

	// 10. musim-CapabilityRestrictionConfig-r18: choice
	{
		if seq.IsComponentPresent(8) {
			ie.Musim_CapabilityRestrictionConfig_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *MUSIM_CapabilityRestrictionConfig_r18
			}{}
			choiceDec := d.NewChoiceDecoder(otherConfig_v1800MusimCapabilityRestrictionConfigR18Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Musim_CapabilityRestrictionConfig_r18).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case OtherConfig_v1800_Musim_CapabilityRestrictionConfig_r18_Release:
				(*ie.Musim_CapabilityRestrictionConfig_r18).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1800_Musim_CapabilityRestrictionConfig_r18_Setup:
				(*ie.Musim_CapabilityRestrictionConfig_r18).Setup = new(MUSIM_CapabilityRestrictionConfig_r18)
				if err := (*ie.Musim_CapabilityRestrictionConfig_r18).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

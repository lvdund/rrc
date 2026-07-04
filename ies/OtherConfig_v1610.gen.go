// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: OtherConfig-v1610 (line 26314).

var otherConfigV1610Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "idc-AssistanceConfig-r16", Optional: true},
		{Name: "drx-PreferenceConfig-r16", Optional: true},
		{Name: "maxBW-PreferenceConfig-r16", Optional: true},
		{Name: "maxCC-PreferenceConfig-r16", Optional: true},
		{Name: "maxMIMO-LayerPreferenceConfig-r16", Optional: true},
		{Name: "minSchedulingOffsetPreferenceConfig-r16", Optional: true},
		{Name: "releasePreferenceConfig-r16", Optional: true},
		{Name: "referenceTimePreferenceReporting-r16", Optional: true},
		{Name: "btNameList-r16", Optional: true},
		{Name: "wlanNameList-r16", Optional: true},
		{Name: "sensorNameList-r16", Optional: true},
		{Name: "obtainCommonLocation-r16", Optional: true},
		{Name: "sl-AssistanceConfigNR-r16", Optional: true},
	},
}

var otherConfig_v1610IdcAssistanceConfigR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	OtherConfig_v1610_Idc_AssistanceConfig_r16_Release = 0
	OtherConfig_v1610_Idc_AssistanceConfig_r16_Setup   = 1
)

var otherConfig_v1610DrxPreferenceConfigR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	OtherConfig_v1610_Drx_PreferenceConfig_r16_Release = 0
	OtherConfig_v1610_Drx_PreferenceConfig_r16_Setup   = 1
)

var otherConfig_v1610MaxBWPreferenceConfigR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	OtherConfig_v1610_MaxBW_PreferenceConfig_r16_Release = 0
	OtherConfig_v1610_MaxBW_PreferenceConfig_r16_Setup   = 1
)

var otherConfig_v1610MaxCCPreferenceConfigR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	OtherConfig_v1610_MaxCC_PreferenceConfig_r16_Release = 0
	OtherConfig_v1610_MaxCC_PreferenceConfig_r16_Setup   = 1
)

var otherConfig_v1610MaxMIMOLayerPreferenceConfigR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	OtherConfig_v1610_MaxMIMO_LayerPreferenceConfig_r16_Release = 0
	OtherConfig_v1610_MaxMIMO_LayerPreferenceConfig_r16_Setup   = 1
)

var otherConfig_v1610MinSchedulingOffsetPreferenceConfigR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	OtherConfig_v1610_MinSchedulingOffsetPreferenceConfig_r16_Release = 0
	OtherConfig_v1610_MinSchedulingOffsetPreferenceConfig_r16_Setup   = 1
)

var otherConfig_v1610ReleasePreferenceConfigR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	OtherConfig_v1610_ReleasePreferenceConfig_r16_Release = 0
	OtherConfig_v1610_ReleasePreferenceConfig_r16_Setup   = 1
)

const (
	OtherConfig_v1610_ReferenceTimePreferenceReporting_r16_True = 0
)

var otherConfigV1610ReferenceTimePreferenceReportingR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{OtherConfig_v1610_ReferenceTimePreferenceReporting_r16_True},
}

var otherConfig_v1610BtNameListR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	OtherConfig_v1610_BtNameList_r16_Release = 0
	OtherConfig_v1610_BtNameList_r16_Setup   = 1
)

var otherConfig_v1610WlanNameListR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	OtherConfig_v1610_WlanNameList_r16_Release = 0
	OtherConfig_v1610_WlanNameList_r16_Setup   = 1
)

var otherConfig_v1610SensorNameListR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	OtherConfig_v1610_SensorNameList_r16_Release = 0
	OtherConfig_v1610_SensorNameList_r16_Setup   = 1
)

const (
	OtherConfig_v1610_ObtainCommonLocation_r16_True = 0
)

var otherConfigV1610ObtainCommonLocationR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{OtherConfig_v1610_ObtainCommonLocation_r16_True},
}

const (
	OtherConfig_v1610_Sl_AssistanceConfigNR_r16_True = 0
)

var otherConfigV1610SlAssistanceConfigNRR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{OtherConfig_v1610_Sl_AssistanceConfigNR_r16_True},
}

type OtherConfig_v1610 struct {
	Idc_AssistanceConfig_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *IDC_AssistanceConfig_r16
	}
	Drx_PreferenceConfig_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *DRX_PreferenceConfig_r16
	}
	MaxBW_PreferenceConfig_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *MaxBW_PreferenceConfig_r16
	}
	MaxCC_PreferenceConfig_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *MaxCC_PreferenceConfig_r16
	}
	MaxMIMO_LayerPreferenceConfig_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *MaxMIMO_LayerPreferenceConfig_r16
	}
	MinSchedulingOffsetPreferenceConfig_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *MinSchedulingOffsetPreferenceConfig_r16
	}
	ReleasePreferenceConfig_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *ReleasePreferenceConfig_r16
	}
	ReferenceTimePreferenceReporting_r16 *int64
	BtNameList_r16                       *struct {
		Choice  int
		Release *struct{}
		Setup   *BT_NameList_r16
	}
	WlanNameList_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *WLAN_NameList_r16
	}
	SensorNameList_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *Sensor_NameList_r16
	}
	ObtainCommonLocation_r16  *int64
	Sl_AssistanceConfigNR_r16 *int64
}

func (ie *OtherConfig_v1610) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(otherConfigV1610Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Idc_AssistanceConfig_r16 != nil, ie.Drx_PreferenceConfig_r16 != nil, ie.MaxBW_PreferenceConfig_r16 != nil, ie.MaxCC_PreferenceConfig_r16 != nil, ie.MaxMIMO_LayerPreferenceConfig_r16 != nil, ie.MinSchedulingOffsetPreferenceConfig_r16 != nil, ie.ReleasePreferenceConfig_r16 != nil, ie.ReferenceTimePreferenceReporting_r16 != nil, ie.BtNameList_r16 != nil, ie.WlanNameList_r16 != nil, ie.SensorNameList_r16 != nil, ie.ObtainCommonLocation_r16 != nil, ie.Sl_AssistanceConfigNR_r16 != nil}); err != nil {
		return err
	}

	// 2. idc-AssistanceConfig-r16: choice
	{
		if ie.Idc_AssistanceConfig_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(otherConfig_v1610IdcAssistanceConfigR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Idc_AssistanceConfig_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Idc_AssistanceConfig_r16).Choice {
			case OtherConfig_v1610_Idc_AssistanceConfig_r16_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1610_Idc_AssistanceConfig_r16_Setup:
				if err := (*ie.Idc_AssistanceConfig_r16).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Idc_AssistanceConfig_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 3. drx-PreferenceConfig-r16: choice
	{
		if ie.Drx_PreferenceConfig_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(otherConfig_v1610DrxPreferenceConfigR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Drx_PreferenceConfig_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Drx_PreferenceConfig_r16).Choice {
			case OtherConfig_v1610_Drx_PreferenceConfig_r16_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1610_Drx_PreferenceConfig_r16_Setup:
				if err := (*ie.Drx_PreferenceConfig_r16).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Drx_PreferenceConfig_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 4. maxBW-PreferenceConfig-r16: choice
	{
		if ie.MaxBW_PreferenceConfig_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(otherConfig_v1610MaxBWPreferenceConfigR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.MaxBW_PreferenceConfig_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.MaxBW_PreferenceConfig_r16).Choice {
			case OtherConfig_v1610_MaxBW_PreferenceConfig_r16_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1610_MaxBW_PreferenceConfig_r16_Setup:
				if err := (*ie.MaxBW_PreferenceConfig_r16).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.MaxBW_PreferenceConfig_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 5. maxCC-PreferenceConfig-r16: choice
	{
		if ie.MaxCC_PreferenceConfig_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(otherConfig_v1610MaxCCPreferenceConfigR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.MaxCC_PreferenceConfig_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.MaxCC_PreferenceConfig_r16).Choice {
			case OtherConfig_v1610_MaxCC_PreferenceConfig_r16_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1610_MaxCC_PreferenceConfig_r16_Setup:
				if err := (*ie.MaxCC_PreferenceConfig_r16).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.MaxCC_PreferenceConfig_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 6. maxMIMO-LayerPreferenceConfig-r16: choice
	{
		if ie.MaxMIMO_LayerPreferenceConfig_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(otherConfig_v1610MaxMIMOLayerPreferenceConfigR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.MaxMIMO_LayerPreferenceConfig_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.MaxMIMO_LayerPreferenceConfig_r16).Choice {
			case OtherConfig_v1610_MaxMIMO_LayerPreferenceConfig_r16_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1610_MaxMIMO_LayerPreferenceConfig_r16_Setup:
				if err := (*ie.MaxMIMO_LayerPreferenceConfig_r16).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.MaxMIMO_LayerPreferenceConfig_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 7. minSchedulingOffsetPreferenceConfig-r16: choice
	{
		if ie.MinSchedulingOffsetPreferenceConfig_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(otherConfig_v1610MinSchedulingOffsetPreferenceConfigR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.MinSchedulingOffsetPreferenceConfig_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.MinSchedulingOffsetPreferenceConfig_r16).Choice {
			case OtherConfig_v1610_MinSchedulingOffsetPreferenceConfig_r16_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1610_MinSchedulingOffsetPreferenceConfig_r16_Setup:
				if err := (*ie.MinSchedulingOffsetPreferenceConfig_r16).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.MinSchedulingOffsetPreferenceConfig_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 8. releasePreferenceConfig-r16: choice
	{
		if ie.ReleasePreferenceConfig_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(otherConfig_v1610ReleasePreferenceConfigR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.ReleasePreferenceConfig_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.ReleasePreferenceConfig_r16).Choice {
			case OtherConfig_v1610_ReleasePreferenceConfig_r16_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1610_ReleasePreferenceConfig_r16_Setup:
				if err := (*ie.ReleasePreferenceConfig_r16).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.ReleasePreferenceConfig_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 9. referenceTimePreferenceReporting-r16: enumerated
	{
		if ie.ReferenceTimePreferenceReporting_r16 != nil {
			if err := e.EncodeEnumerated(*ie.ReferenceTimePreferenceReporting_r16, otherConfigV1610ReferenceTimePreferenceReportingR16Constraints); err != nil {
				return err
			}
		}
	}

	// 10. btNameList-r16: choice
	{
		if ie.BtNameList_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(otherConfig_v1610BtNameListR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.BtNameList_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.BtNameList_r16).Choice {
			case OtherConfig_v1610_BtNameList_r16_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1610_BtNameList_r16_Setup:
				if err := (*ie.BtNameList_r16).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.BtNameList_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 11. wlanNameList-r16: choice
	{
		if ie.WlanNameList_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(otherConfig_v1610WlanNameListR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.WlanNameList_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.WlanNameList_r16).Choice {
			case OtherConfig_v1610_WlanNameList_r16_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1610_WlanNameList_r16_Setup:
				if err := (*ie.WlanNameList_r16).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.WlanNameList_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 12. sensorNameList-r16: choice
	{
		if ie.SensorNameList_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(otherConfig_v1610SensorNameListR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.SensorNameList_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.SensorNameList_r16).Choice {
			case OtherConfig_v1610_SensorNameList_r16_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1610_SensorNameList_r16_Setup:
				if err := (*ie.SensorNameList_r16).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.SensorNameList_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 13. obtainCommonLocation-r16: enumerated
	{
		if ie.ObtainCommonLocation_r16 != nil {
			if err := e.EncodeEnumerated(*ie.ObtainCommonLocation_r16, otherConfigV1610ObtainCommonLocationR16Constraints); err != nil {
				return err
			}
		}
	}

	// 14. sl-AssistanceConfigNR-r16: enumerated
	{
		if ie.Sl_AssistanceConfigNR_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_AssistanceConfigNR_r16, otherConfigV1610SlAssistanceConfigNRR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *OtherConfig_v1610) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(otherConfigV1610Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. idc-AssistanceConfig-r16: choice
	{
		if seq.IsComponentPresent(0) {
			ie.Idc_AssistanceConfig_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *IDC_AssistanceConfig_r16
			}{}
			choiceDec := d.NewChoiceDecoder(otherConfig_v1610IdcAssistanceConfigR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Idc_AssistanceConfig_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case OtherConfig_v1610_Idc_AssistanceConfig_r16_Release:
				(*ie.Idc_AssistanceConfig_r16).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1610_Idc_AssistanceConfig_r16_Setup:
				(*ie.Idc_AssistanceConfig_r16).Setup = new(IDC_AssistanceConfig_r16)
				if err := (*ie.Idc_AssistanceConfig_r16).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. drx-PreferenceConfig-r16: choice
	{
		if seq.IsComponentPresent(1) {
			ie.Drx_PreferenceConfig_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *DRX_PreferenceConfig_r16
			}{}
			choiceDec := d.NewChoiceDecoder(otherConfig_v1610DrxPreferenceConfigR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Drx_PreferenceConfig_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case OtherConfig_v1610_Drx_PreferenceConfig_r16_Release:
				(*ie.Drx_PreferenceConfig_r16).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1610_Drx_PreferenceConfig_r16_Setup:
				(*ie.Drx_PreferenceConfig_r16).Setup = new(DRX_PreferenceConfig_r16)
				if err := (*ie.Drx_PreferenceConfig_r16).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. maxBW-PreferenceConfig-r16: choice
	{
		if seq.IsComponentPresent(2) {
			ie.MaxBW_PreferenceConfig_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *MaxBW_PreferenceConfig_r16
			}{}
			choiceDec := d.NewChoiceDecoder(otherConfig_v1610MaxBWPreferenceConfigR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.MaxBW_PreferenceConfig_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case OtherConfig_v1610_MaxBW_PreferenceConfig_r16_Release:
				(*ie.MaxBW_PreferenceConfig_r16).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1610_MaxBW_PreferenceConfig_r16_Setup:
				(*ie.MaxBW_PreferenceConfig_r16).Setup = new(MaxBW_PreferenceConfig_r16)
				if err := (*ie.MaxBW_PreferenceConfig_r16).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. maxCC-PreferenceConfig-r16: choice
	{
		if seq.IsComponentPresent(3) {
			ie.MaxCC_PreferenceConfig_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *MaxCC_PreferenceConfig_r16
			}{}
			choiceDec := d.NewChoiceDecoder(otherConfig_v1610MaxCCPreferenceConfigR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.MaxCC_PreferenceConfig_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case OtherConfig_v1610_MaxCC_PreferenceConfig_r16_Release:
				(*ie.MaxCC_PreferenceConfig_r16).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1610_MaxCC_PreferenceConfig_r16_Setup:
				(*ie.MaxCC_PreferenceConfig_r16).Setup = new(MaxCC_PreferenceConfig_r16)
				if err := (*ie.MaxCC_PreferenceConfig_r16).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. maxMIMO-LayerPreferenceConfig-r16: choice
	{
		if seq.IsComponentPresent(4) {
			ie.MaxMIMO_LayerPreferenceConfig_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *MaxMIMO_LayerPreferenceConfig_r16
			}{}
			choiceDec := d.NewChoiceDecoder(otherConfig_v1610MaxMIMOLayerPreferenceConfigR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.MaxMIMO_LayerPreferenceConfig_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case OtherConfig_v1610_MaxMIMO_LayerPreferenceConfig_r16_Release:
				(*ie.MaxMIMO_LayerPreferenceConfig_r16).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1610_MaxMIMO_LayerPreferenceConfig_r16_Setup:
				(*ie.MaxMIMO_LayerPreferenceConfig_r16).Setup = new(MaxMIMO_LayerPreferenceConfig_r16)
				if err := (*ie.MaxMIMO_LayerPreferenceConfig_r16).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 7. minSchedulingOffsetPreferenceConfig-r16: choice
	{
		if seq.IsComponentPresent(5) {
			ie.MinSchedulingOffsetPreferenceConfig_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *MinSchedulingOffsetPreferenceConfig_r16
			}{}
			choiceDec := d.NewChoiceDecoder(otherConfig_v1610MinSchedulingOffsetPreferenceConfigR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.MinSchedulingOffsetPreferenceConfig_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case OtherConfig_v1610_MinSchedulingOffsetPreferenceConfig_r16_Release:
				(*ie.MinSchedulingOffsetPreferenceConfig_r16).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1610_MinSchedulingOffsetPreferenceConfig_r16_Setup:
				(*ie.MinSchedulingOffsetPreferenceConfig_r16).Setup = new(MinSchedulingOffsetPreferenceConfig_r16)
				if err := (*ie.MinSchedulingOffsetPreferenceConfig_r16).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 8. releasePreferenceConfig-r16: choice
	{
		if seq.IsComponentPresent(6) {
			ie.ReleasePreferenceConfig_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *ReleasePreferenceConfig_r16
			}{}
			choiceDec := d.NewChoiceDecoder(otherConfig_v1610ReleasePreferenceConfigR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.ReleasePreferenceConfig_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case OtherConfig_v1610_ReleasePreferenceConfig_r16_Release:
				(*ie.ReleasePreferenceConfig_r16).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1610_ReleasePreferenceConfig_r16_Setup:
				(*ie.ReleasePreferenceConfig_r16).Setup = new(ReleasePreferenceConfig_r16)
				if err := (*ie.ReleasePreferenceConfig_r16).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 9. referenceTimePreferenceReporting-r16: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(otherConfigV1610ReferenceTimePreferenceReportingR16Constraints)
			if err != nil {
				return err
			}
			ie.ReferenceTimePreferenceReporting_r16 = &idx
		}
	}

	// 10. btNameList-r16: choice
	{
		if seq.IsComponentPresent(8) {
			ie.BtNameList_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *BT_NameList_r16
			}{}
			choiceDec := d.NewChoiceDecoder(otherConfig_v1610BtNameListR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.BtNameList_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case OtherConfig_v1610_BtNameList_r16_Release:
				(*ie.BtNameList_r16).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1610_BtNameList_r16_Setup:
				(*ie.BtNameList_r16).Setup = new(BT_NameList_r16)
				if err := (*ie.BtNameList_r16).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 11. wlanNameList-r16: choice
	{
		if seq.IsComponentPresent(9) {
			ie.WlanNameList_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *WLAN_NameList_r16
			}{}
			choiceDec := d.NewChoiceDecoder(otherConfig_v1610WlanNameListR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.WlanNameList_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case OtherConfig_v1610_WlanNameList_r16_Release:
				(*ie.WlanNameList_r16).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1610_WlanNameList_r16_Setup:
				(*ie.WlanNameList_r16).Setup = new(WLAN_NameList_r16)
				if err := (*ie.WlanNameList_r16).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 12. sensorNameList-r16: choice
	{
		if seq.IsComponentPresent(10) {
			ie.SensorNameList_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *Sensor_NameList_r16
			}{}
			choiceDec := d.NewChoiceDecoder(otherConfig_v1610SensorNameListR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.SensorNameList_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case OtherConfig_v1610_SensorNameList_r16_Release:
				(*ie.SensorNameList_r16).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1610_SensorNameList_r16_Setup:
				(*ie.SensorNameList_r16).Setup = new(Sensor_NameList_r16)
				if err := (*ie.SensorNameList_r16).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 13. obtainCommonLocation-r16: enumerated
	{
		if seq.IsComponentPresent(11) {
			idx, err := d.DecodeEnumerated(otherConfigV1610ObtainCommonLocationR16Constraints)
			if err != nil {
				return err
			}
			ie.ObtainCommonLocation_r16 = &idx
		}
	}

	// 14. sl-AssistanceConfigNR-r16: enumerated
	{
		if seq.IsComponentPresent(12) {
			idx, err := d.DecodeEnumerated(otherConfigV1610SlAssistanceConfigNRR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_AssistanceConfigNR_r16 = &idx
		}
	}

	return nil
}

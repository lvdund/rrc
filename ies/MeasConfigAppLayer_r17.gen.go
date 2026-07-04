// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasConfigAppLayer-r17 (line 26021).

var measConfigAppLayerR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "measConfigAppLayerId-r17"},
		{Name: "measConfigAppLayerContainer-r17", Optional: true},
		{Name: "serviceType-r17", Optional: true},
		{Name: "pauseReporting-r17", Optional: true},
		{Name: "transmissionOfSessionStartStop-r17", Optional: true},
		{Name: "ran-VisibleParameters-r17", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var measConfigAppLayerR17MeasConfigAppLayerContainerR17Constraints = per.SizeRange(1, 8000)

const (
	MeasConfigAppLayer_r17_ServiceType_r17_Streaming = 0
	MeasConfigAppLayer_r17_ServiceType_r17_Mtsi      = 1
	MeasConfigAppLayer_r17_ServiceType_r17_Vr        = 2
	MeasConfigAppLayer_r17_ServiceType_r17_Spare5    = 3
	MeasConfigAppLayer_r17_ServiceType_r17_Spare4    = 4
	MeasConfigAppLayer_r17_ServiceType_r17_Spare3    = 5
	MeasConfigAppLayer_r17_ServiceType_r17_Spare2    = 6
	MeasConfigAppLayer_r17_ServiceType_r17_Spare1    = 7
)

var measConfigAppLayerR17ServiceTypeR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasConfigAppLayer_r17_ServiceType_r17_Streaming, MeasConfigAppLayer_r17_ServiceType_r17_Mtsi, MeasConfigAppLayer_r17_ServiceType_r17_Vr, MeasConfigAppLayer_r17_ServiceType_r17_Spare5, MeasConfigAppLayer_r17_ServiceType_r17_Spare4, MeasConfigAppLayer_r17_ServiceType_r17_Spare3, MeasConfigAppLayer_r17_ServiceType_r17_Spare2, MeasConfigAppLayer_r17_ServiceType_r17_Spare1},
}

var measConfigAppLayer_r17RanVisibleParametersR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	MeasConfigAppLayer_r17_Ran_VisibleParameters_r17_Release = 0
	MeasConfigAppLayer_r17_Ran_VisibleParameters_r17_Setup   = 1
)

const (
	MeasConfigAppLayer_r17_Ext_ReportingSRB_r18_Srb4 = 0
	MeasConfigAppLayer_r17_Ext_ReportingSRB_r18_Srb5 = 1
)

var measConfigAppLayerR17ExtReportingSRBR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasConfigAppLayer_r17_Ext_ReportingSRB_r18_Srb4, MeasConfigAppLayer_r17_Ext_ReportingSRB_r18_Srb5},
}

var measConfigAppLayerR17AppLayerMeasPriorityR18Constraints = per.Constrained(1, 16)

var measConfigAppLayerR17ExtAppLayerIdleInactiveConfigR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	MeasConfigAppLayer_r17_Ext_AppLayerIdleInactiveConfig_r18_Release = 0
	MeasConfigAppLayer_r17_Ext_AppLayerIdleInactiveConfig_r18_Setup   = 1
)

type MeasConfigAppLayer_r17 struct {
	MeasConfigAppLayerId_r17           MeasConfigAppLayerId_r17
	MeasConfigAppLayerContainer_r17    []byte
	ServiceType_r17                    *int64
	PauseReporting_r17                 *bool
	TransmissionOfSessionStartStop_r17 *bool
	Ran_VisibleParameters_r17          *struct {
		Choice  int
		Release *struct{}
		Setup   *RAN_VisibleParameters_r17
	}
	ReportingSRB_r18               *int64
	AppLayerMeasPriority_r18       *int64
	AppLayerIdleInactiveConfig_r18 *struct {
		Choice  int
		Release *struct{}
		Setup   *AppLayerIdleInactiveConfig_r18
	}
}

func (ie *MeasConfigAppLayer_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measConfigAppLayerR17Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.ReportingSRB_r18 != nil || ie.AppLayerMeasPriority_r18 != nil || ie.AppLayerIdleInactiveConfig_r18 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MeasConfigAppLayerContainer_r17 != nil, ie.ServiceType_r17 != nil, ie.PauseReporting_r17 != nil, ie.TransmissionOfSessionStartStop_r17 != nil, ie.Ran_VisibleParameters_r17 != nil}); err != nil {
		return err
	}

	// 3. measConfigAppLayerId-r17: ref
	{
		if err := ie.MeasConfigAppLayerId_r17.Encode(e); err != nil {
			return err
		}
	}

	// 4. measConfigAppLayerContainer-r17: octet-string
	{
		if ie.MeasConfigAppLayerContainer_r17 != nil {
			if err := e.EncodeOctetString(ie.MeasConfigAppLayerContainer_r17, measConfigAppLayerR17MeasConfigAppLayerContainerR17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. serviceType-r17: enumerated
	{
		if ie.ServiceType_r17 != nil {
			if err := e.EncodeEnumerated(*ie.ServiceType_r17, measConfigAppLayerR17ServiceTypeR17Constraints); err != nil {
				return err
			}
		}
	}

	// 6. pauseReporting-r17: boolean
	{
		if ie.PauseReporting_r17 != nil {
			if err := e.EncodeBoolean(*ie.PauseReporting_r17); err != nil {
				return err
			}
		}
	}

	// 7. transmissionOfSessionStartStop-r17: boolean
	{
		if ie.TransmissionOfSessionStartStop_r17 != nil {
			if err := e.EncodeBoolean(*ie.TransmissionOfSessionStartStop_r17); err != nil {
				return err
			}
		}
	}

	// 8. ran-VisibleParameters-r17: choice
	{
		if ie.Ran_VisibleParameters_r17 != nil {
			choiceEnc := e.NewChoiceEncoder(measConfigAppLayer_r17RanVisibleParametersR17Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Ran_VisibleParameters_r17).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Ran_VisibleParameters_r17).Choice {
			case MeasConfigAppLayer_r17_Ran_VisibleParameters_r17_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case MeasConfigAppLayer_r17_Ran_VisibleParameters_r17_Setup:
				if err := (*ie.Ran_VisibleParameters_r17).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Ran_VisibleParameters_r17).Choice), Constraint: "index out of range"}
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "reportingSRB-r18", Optional: true},
					{Name: "appLayerMeasPriority-r18", Optional: true},
					{Name: "appLayerIdleInactiveConfig-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ReportingSRB_r18 != nil, ie.AppLayerMeasPriority_r18 != nil, ie.AppLayerIdleInactiveConfig_r18 != nil}); err != nil {
				return err
			}

			if ie.ReportingSRB_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.ReportingSRB_r18, measConfigAppLayerR17ExtReportingSRBR18Constraints); err != nil {
					return err
				}
			}

			if ie.AppLayerMeasPriority_r18 != nil {
				if err := ex.EncodeInteger(*ie.AppLayerMeasPriority_r18, measConfigAppLayerR17AppLayerMeasPriorityR18Constraints); err != nil {
					return err
				}
			}

			if ie.AppLayerIdleInactiveConfig_r18 != nil {
				choiceEnc := ex.NewChoiceEncoder(measConfigAppLayerR17ExtAppLayerIdleInactiveConfigR18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.AppLayerIdleInactiveConfig_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.AppLayerIdleInactiveConfig_r18).Choice {
				case MeasConfigAppLayer_r17_Ext_AppLayerIdleInactiveConfig_r18_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case MeasConfigAppLayer_r17_Ext_AppLayerIdleInactiveConfig_r18_Setup:
					if err := (*ie.AppLayerIdleInactiveConfig_r18).Setup.Encode(ex); err != nil {
						return err
					}
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

func (ie *MeasConfigAppLayer_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measConfigAppLayerR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. measConfigAppLayerId-r17: ref
	{
		if err := ie.MeasConfigAppLayerId_r17.Decode(d); err != nil {
			return err
		}
	}

	// 4. measConfigAppLayerContainer-r17: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(measConfigAppLayerR17MeasConfigAppLayerContainerR17Constraints)
			if err != nil {
				return err
			}
			ie.MeasConfigAppLayerContainer_r17 = v
		}
	}

	// 5. serviceType-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(measConfigAppLayerR17ServiceTypeR17Constraints)
			if err != nil {
				return err
			}
			ie.ServiceType_r17 = &idx
		}
	}

	// 6. pauseReporting-r17: boolean
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeBoolean()
			if err != nil {
				return err
			}
			ie.PauseReporting_r17 = &v
		}
	}

	// 7. transmissionOfSessionStartStop-r17: boolean
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeBoolean()
			if err != nil {
				return err
			}
			ie.TransmissionOfSessionStartStop_r17 = &v
		}
	}

	// 8. ran-VisibleParameters-r17: choice
	{
		if seq.IsComponentPresent(5) {
			ie.Ran_VisibleParameters_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *RAN_VisibleParameters_r17
			}{}
			choiceDec := d.NewChoiceDecoder(measConfigAppLayer_r17RanVisibleParametersR17Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Ran_VisibleParameters_r17).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case MeasConfigAppLayer_r17_Ran_VisibleParameters_r17_Release:
				(*ie.Ran_VisibleParameters_r17).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case MeasConfigAppLayer_r17_Ran_VisibleParameters_r17_Setup:
				(*ie.Ran_VisibleParameters_r17).Setup = new(RAN_VisibleParameters_r17)
				if err := (*ie.Ran_VisibleParameters_r17).Setup.Decode(d); err != nil {
					return err
				}
			}
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
				{Name: "reportingSRB-r18", Optional: true},
				{Name: "appLayerMeasPriority-r18", Optional: true},
				{Name: "appLayerIdleInactiveConfig-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(measConfigAppLayerR17ExtReportingSRBR18Constraints)
			if err != nil {
				return err
			}
			ie.ReportingSRB_r18 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeInteger(measConfigAppLayerR17AppLayerMeasPriorityR18Constraints)
			if err != nil {
				return err
			}
			ie.AppLayerMeasPriority_r18 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			ie.AppLayerIdleInactiveConfig_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *AppLayerIdleInactiveConfig_r18
			}{}
			choiceDec := dx.NewChoiceDecoder(measConfigAppLayerR17ExtAppLayerIdleInactiveConfigR18Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.AppLayerIdleInactiveConfig_r18).Choice = int(index)
			switch index {
			case MeasConfigAppLayer_r17_Ext_AppLayerIdleInactiveConfig_r18_Release:
				(*ie.AppLayerIdleInactiveConfig_r18).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case MeasConfigAppLayer_r17_Ext_AppLayerIdleInactiveConfig_r18_Setup:
				(*ie.AppLayerIdleInactiveConfig_r18).Setup = new(AppLayerIdleInactiveConfig_r18)
				if err := (*ie.AppLayerIdleInactiveConfig_r18).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

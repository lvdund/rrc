// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: LoggedMeasurementConfiguration-r16-IEs (line 541).

var loggedMeasurementConfigurationR16IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "traceReference-r16"},
		{Name: "traceRecordingSessionRef-r16"},
		{Name: "tce-Id-r16"},
		{Name: "absoluteTimeInfo-r16"},
		{Name: "areaConfiguration-r16", Optional: true},
		{Name: "plmn-IdentityList-r16", Optional: true},
		{Name: "bt-NameList-r16", Optional: true},
		{Name: "wlan-NameList-r16", Optional: true},
		{Name: "sensor-NameList-r16", Optional: true},
		{Name: "loggingDuration-r16"},
		{Name: "reportType"},
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var loggedMeasurementConfigurationR16IEsTraceRecordingSessionRefR16Constraints = per.FixedSize(2)

var loggedMeasurementConfigurationR16IEsTceIdR16Constraints = per.FixedSize(1)

var loggedMeasurementConfiguration_r16_IEsBtNameListR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	LoggedMeasurementConfiguration_r16_IEs_Bt_NameList_r16_Release = 0
	LoggedMeasurementConfiguration_r16_IEs_Bt_NameList_r16_Setup   = 1
)

var loggedMeasurementConfiguration_r16_IEsWlanNameListR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	LoggedMeasurementConfiguration_r16_IEs_Wlan_NameList_r16_Release = 0
	LoggedMeasurementConfiguration_r16_IEs_Wlan_NameList_r16_Setup   = 1
)

var loggedMeasurementConfiguration_r16_IEsSensorNameListR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	LoggedMeasurementConfiguration_r16_IEs_Sensor_NameList_r16_Release = 0
	LoggedMeasurementConfiguration_r16_IEs_Sensor_NameList_r16_Setup   = 1
)

var loggedMeasurementConfiguration_r16_IEsReportTypeConstraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "periodical"},
		{Name: "eventTriggered"},
	},
}

const (
	LoggedMeasurementConfiguration_r16_IEs_ReportType_Periodical     = 0
	LoggedMeasurementConfiguration_r16_IEs_ReportType_EventTriggered = 1
)

var loggedMeasurementConfigurationR16IEsLateNonCriticalExtensionConstraints = per.SizeConstraints{}

type LoggedMeasurementConfiguration_r16_IEs struct {
	TraceReference_r16           TraceReference_r16
	TraceRecordingSessionRef_r16 []byte
	Tce_Id_r16                   []byte
	AbsoluteTimeInfo_r16         AbsoluteTimeInfo_r16
	AreaConfiguration_r16        *AreaConfiguration_r16
	Plmn_IdentityList_r16        *PLMN_IdentityList2_r16
	Bt_NameList_r16              *struct {
		Choice  int
		Release *struct{}
		Setup   *BT_NameList_r16
	}
	Wlan_NameList_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *WLAN_NameList_r16
	}
	Sensor_NameList_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *Sensor_NameList_r16
	}
	LoggingDuration_r16 LoggingDuration_r16
	ReportType          struct {
		Choice         int
		Periodical     *LoggedPeriodicalReportConfig_r16
		EventTriggered *LoggedEventTriggerConfig_r16
	}
	LateNonCriticalExtension []byte
	NonCriticalExtension     *LoggedMeasurementConfiguration_v1700_IEs
}

func (ie *LoggedMeasurementConfiguration_r16_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(loggedMeasurementConfigurationR16IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.AreaConfiguration_r16 != nil, ie.Plmn_IdentityList_r16 != nil, ie.Bt_NameList_r16 != nil, ie.Wlan_NameList_r16 != nil, ie.Sensor_NameList_r16 != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. traceReference-r16: ref
	{
		if err := ie.TraceReference_r16.Encode(e); err != nil {
			return err
		}
	}

	// 3. traceRecordingSessionRef-r16: octet-string
	{
		if err := e.EncodeOctetString(ie.TraceRecordingSessionRef_r16, loggedMeasurementConfigurationR16IEsTraceRecordingSessionRefR16Constraints); err != nil {
			return err
		}
	}

	// 4. tce-Id-r16: octet-string
	{
		if err := e.EncodeOctetString(ie.Tce_Id_r16, loggedMeasurementConfigurationR16IEsTceIdR16Constraints); err != nil {
			return err
		}
	}

	// 5. absoluteTimeInfo-r16: ref
	{
		if err := ie.AbsoluteTimeInfo_r16.Encode(e); err != nil {
			return err
		}
	}

	// 6. areaConfiguration-r16: ref
	{
		if ie.AreaConfiguration_r16 != nil {
			if err := ie.AreaConfiguration_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. plmn-IdentityList-r16: ref
	{
		if ie.Plmn_IdentityList_r16 != nil {
			if err := ie.Plmn_IdentityList_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. bt-NameList-r16: choice
	{
		if ie.Bt_NameList_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(loggedMeasurementConfiguration_r16_IEsBtNameListR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Bt_NameList_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Bt_NameList_r16).Choice {
			case LoggedMeasurementConfiguration_r16_IEs_Bt_NameList_r16_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case LoggedMeasurementConfiguration_r16_IEs_Bt_NameList_r16_Setup:
				if err := (*ie.Bt_NameList_r16).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Bt_NameList_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 9. wlan-NameList-r16: choice
	{
		if ie.Wlan_NameList_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(loggedMeasurementConfiguration_r16_IEsWlanNameListR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Wlan_NameList_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Wlan_NameList_r16).Choice {
			case LoggedMeasurementConfiguration_r16_IEs_Wlan_NameList_r16_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case LoggedMeasurementConfiguration_r16_IEs_Wlan_NameList_r16_Setup:
				if err := (*ie.Wlan_NameList_r16).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Wlan_NameList_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 10. sensor-NameList-r16: choice
	{
		if ie.Sensor_NameList_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(loggedMeasurementConfiguration_r16_IEsSensorNameListR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Sensor_NameList_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Sensor_NameList_r16).Choice {
			case LoggedMeasurementConfiguration_r16_IEs_Sensor_NameList_r16_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case LoggedMeasurementConfiguration_r16_IEs_Sensor_NameList_r16_Setup:
				if err := (*ie.Sensor_NameList_r16).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Sensor_NameList_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 11. loggingDuration-r16: ref
	{
		if err := ie.LoggingDuration_r16.Encode(e); err != nil {
			return err
		}
	}

	// 12. reportType: choice
	{
		choiceEnc := e.NewChoiceEncoder(loggedMeasurementConfiguration_r16_IEsReportTypeConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.ReportType.Choice), false, nil); err != nil {
			return err
		}
		switch ie.ReportType.Choice {
		case LoggedMeasurementConfiguration_r16_IEs_ReportType_Periodical:
			if err := ie.ReportType.Periodical.Encode(e); err != nil {
				return err
			}
		case LoggedMeasurementConfiguration_r16_IEs_ReportType_EventTriggered:
			if err := ie.ReportType.EventTriggered.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.ReportType.Choice), Constraint: "index out of range"}
		}
	}

	// 13. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, loggedMeasurementConfigurationR16IEsLateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	// 14. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *LoggedMeasurementConfiguration_r16_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(loggedMeasurementConfigurationR16IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. traceReference-r16: ref
	{
		if err := ie.TraceReference_r16.Decode(d); err != nil {
			return err
		}
	}

	// 3. traceRecordingSessionRef-r16: octet-string
	{
		v1, err := d.DecodeOctetString(loggedMeasurementConfigurationR16IEsTraceRecordingSessionRefR16Constraints)
		if err != nil {
			return err
		}
		ie.TraceRecordingSessionRef_r16 = v1
	}

	// 4. tce-Id-r16: octet-string
	{
		v2, err := d.DecodeOctetString(loggedMeasurementConfigurationR16IEsTceIdR16Constraints)
		if err != nil {
			return err
		}
		ie.Tce_Id_r16 = v2
	}

	// 5. absoluteTimeInfo-r16: ref
	{
		if err := ie.AbsoluteTimeInfo_r16.Decode(d); err != nil {
			return err
		}
	}

	// 6. areaConfiguration-r16: ref
	{
		if seq.IsComponentPresent(4) {
			ie.AreaConfiguration_r16 = new(AreaConfiguration_r16)
			if err := ie.AreaConfiguration_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. plmn-IdentityList-r16: ref
	{
		if seq.IsComponentPresent(5) {
			ie.Plmn_IdentityList_r16 = new(PLMN_IdentityList2_r16)
			if err := ie.Plmn_IdentityList_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. bt-NameList-r16: choice
	{
		if seq.IsComponentPresent(6) {
			ie.Bt_NameList_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *BT_NameList_r16
			}{}
			choiceDec := d.NewChoiceDecoder(loggedMeasurementConfiguration_r16_IEsBtNameListR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Bt_NameList_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case LoggedMeasurementConfiguration_r16_IEs_Bt_NameList_r16_Release:
				(*ie.Bt_NameList_r16).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case LoggedMeasurementConfiguration_r16_IEs_Bt_NameList_r16_Setup:
				(*ie.Bt_NameList_r16).Setup = new(BT_NameList_r16)
				if err := (*ie.Bt_NameList_r16).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 9. wlan-NameList-r16: choice
	{
		if seq.IsComponentPresent(7) {
			ie.Wlan_NameList_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *WLAN_NameList_r16
			}{}
			choiceDec := d.NewChoiceDecoder(loggedMeasurementConfiguration_r16_IEsWlanNameListR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Wlan_NameList_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case LoggedMeasurementConfiguration_r16_IEs_Wlan_NameList_r16_Release:
				(*ie.Wlan_NameList_r16).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case LoggedMeasurementConfiguration_r16_IEs_Wlan_NameList_r16_Setup:
				(*ie.Wlan_NameList_r16).Setup = new(WLAN_NameList_r16)
				if err := (*ie.Wlan_NameList_r16).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 10. sensor-NameList-r16: choice
	{
		if seq.IsComponentPresent(8) {
			ie.Sensor_NameList_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *Sensor_NameList_r16
			}{}
			choiceDec := d.NewChoiceDecoder(loggedMeasurementConfiguration_r16_IEsSensorNameListR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sensor_NameList_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case LoggedMeasurementConfiguration_r16_IEs_Sensor_NameList_r16_Release:
				(*ie.Sensor_NameList_r16).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case LoggedMeasurementConfiguration_r16_IEs_Sensor_NameList_r16_Setup:
				(*ie.Sensor_NameList_r16).Setup = new(Sensor_NameList_r16)
				if err := (*ie.Sensor_NameList_r16).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 11. loggingDuration-r16: ref
	{
		if err := ie.LoggingDuration_r16.Decode(d); err != nil {
			return err
		}
	}

	// 12. reportType: choice
	{
		choiceDec := d.NewChoiceDecoder(loggedMeasurementConfiguration_r16_IEsReportTypeConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.ReportType.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case LoggedMeasurementConfiguration_r16_IEs_ReportType_Periodical:
			ie.ReportType.Periodical = new(LoggedPeriodicalReportConfig_r16)
			if err := ie.ReportType.Periodical.Decode(d); err != nil {
				return err
			}
		case LoggedMeasurementConfiguration_r16_IEs_ReportType_EventTriggered:
			ie.ReportType.EventTriggered = new(LoggedEventTriggerConfig_r16)
			if err := ie.ReportType.EventTriggered.Decode(d); err != nil {
				return err
			}
		}
	}

	// 13. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(11) {
			v, err := d.DecodeOctetString(loggedMeasurementConfigurationR16IEsLateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	// 14. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(12) {
			ie.NonCriticalExtension = new(LoggedMeasurementConfiguration_v1700_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}

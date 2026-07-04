// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: OtherConfig-v1900 (line 26361).

var otherConfigV1900Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "gapOccasionCancelRatioReportConfig-r19", Optional: true},
		{Name: "lpwus-OffsetPreferenceConfig-r19", Optional: true},
		{Name: "applicabilityReportConfig-r19", Optional: true},
		{Name: "dataCollectionPreferenceConfig-r19", Optional: true},
		{Name: "loggedDataCollectionAssistanceConfig-r19", Optional: true},
		{Name: "assisted-SSB-MTC-MG-Config-r19", Optional: true},
		{Name: "fbs-PreferenceReportingConfig-r19", Optional: true},
	},
}

var otherConfig_v1900GapOccasionCancelRatioReportConfigR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	OtherConfig_v1900_GapOccasionCancelRatioReportConfig_r19_Release = 0
	OtherConfig_v1900_GapOccasionCancelRatioReportConfig_r19_Setup   = 1
)

var otherConfig_v1900LpwusOffsetPreferenceConfigR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	OtherConfig_v1900_Lpwus_OffsetPreferenceConfig_r19_Release = 0
	OtherConfig_v1900_Lpwus_OffsetPreferenceConfig_r19_Setup   = 1
)

var otherConfig_v1900ApplicabilityReportConfigR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	OtherConfig_v1900_ApplicabilityReportConfig_r19_Release = 0
	OtherConfig_v1900_ApplicabilityReportConfig_r19_Setup   = 1
)

var otherConfig_v1900DataCollectionPreferenceConfigR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	OtherConfig_v1900_DataCollectionPreferenceConfig_r19_Release = 0
	OtherConfig_v1900_DataCollectionPreferenceConfig_r19_Setup   = 1
)

var otherConfig_v1900LoggedDataCollectionAssistanceConfigR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	OtherConfig_v1900_LoggedDataCollectionAssistanceConfig_r19_Release = 0
	OtherConfig_v1900_LoggedDataCollectionAssistanceConfig_r19_Setup   = 1
)

var otherConfig_v1900AssistedSSBMTCMGConfigR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	OtherConfig_v1900_Assisted_SSB_MTC_MG_Config_r19_Release = 0
	OtherConfig_v1900_Assisted_SSB_MTC_MG_Config_r19_Setup   = 1
)

var otherConfig_v1900FbsPreferenceReportingConfigR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	OtherConfig_v1900_Fbs_PreferenceReportingConfig_r19_Release = 0
	OtherConfig_v1900_Fbs_PreferenceReportingConfig_r19_Setup   = 1
)

type OtherConfig_v1900 struct {
	GapOccasionCancelRatioReportConfig_r19 *struct {
		Choice  int
		Release *struct{}
		Setup   *GapOccasionCancelRatioReportConfig_r19
	}
	Lpwus_OffsetPreferenceConfig_r19 *struct {
		Choice  int
		Release *struct{}
		Setup   *LPWUS_OffsetPreferenceConfig_r19
	}
	ApplicabilityReportConfig_r19 *struct {
		Choice  int
		Release *struct{}
		Setup   *ApplicabilityReportConfig_r19
	}
	DataCollectionPreferenceConfig_r19 *struct {
		Choice  int
		Release *struct{}
		Setup   *DataCollectionPreferenceConfig_r19
	}
	LoggedDataCollectionAssistanceConfig_r19 *struct {
		Choice  int
		Release *struct{}
		Setup   *LoggedDataCollectionAssistanceConfig_r19
	}
	Assisted_SSB_MTC_MG_Config_r19 *struct {
		Choice  int
		Release *struct{}
		Setup   *Assisted_SSB_MTC_MG_Config_r19
	}
	Fbs_PreferenceReportingConfig_r19 *struct {
		Choice  int
		Release *struct{}
		Setup   *FBS_PreferenceReportingConfig_r19
	}
}

func (ie *OtherConfig_v1900) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(otherConfigV1900Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.GapOccasionCancelRatioReportConfig_r19 != nil, ie.Lpwus_OffsetPreferenceConfig_r19 != nil, ie.ApplicabilityReportConfig_r19 != nil, ie.DataCollectionPreferenceConfig_r19 != nil, ie.LoggedDataCollectionAssistanceConfig_r19 != nil, ie.Assisted_SSB_MTC_MG_Config_r19 != nil, ie.Fbs_PreferenceReportingConfig_r19 != nil}); err != nil {
		return err
	}

	// 2. gapOccasionCancelRatioReportConfig-r19: choice
	{
		if ie.GapOccasionCancelRatioReportConfig_r19 != nil {
			choiceEnc := e.NewChoiceEncoder(otherConfig_v1900GapOccasionCancelRatioReportConfigR19Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.GapOccasionCancelRatioReportConfig_r19).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.GapOccasionCancelRatioReportConfig_r19).Choice {
			case OtherConfig_v1900_GapOccasionCancelRatioReportConfig_r19_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1900_GapOccasionCancelRatioReportConfig_r19_Setup:
				if err := (*ie.GapOccasionCancelRatioReportConfig_r19).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.GapOccasionCancelRatioReportConfig_r19).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 3. lpwus-OffsetPreferenceConfig-r19: choice
	{
		if ie.Lpwus_OffsetPreferenceConfig_r19 != nil {
			choiceEnc := e.NewChoiceEncoder(otherConfig_v1900LpwusOffsetPreferenceConfigR19Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Lpwus_OffsetPreferenceConfig_r19).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Lpwus_OffsetPreferenceConfig_r19).Choice {
			case OtherConfig_v1900_Lpwus_OffsetPreferenceConfig_r19_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1900_Lpwus_OffsetPreferenceConfig_r19_Setup:
				if err := (*ie.Lpwus_OffsetPreferenceConfig_r19).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Lpwus_OffsetPreferenceConfig_r19).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 4. applicabilityReportConfig-r19: choice
	{
		if ie.ApplicabilityReportConfig_r19 != nil {
			choiceEnc := e.NewChoiceEncoder(otherConfig_v1900ApplicabilityReportConfigR19Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.ApplicabilityReportConfig_r19).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.ApplicabilityReportConfig_r19).Choice {
			case OtherConfig_v1900_ApplicabilityReportConfig_r19_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1900_ApplicabilityReportConfig_r19_Setup:
				if err := (*ie.ApplicabilityReportConfig_r19).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.ApplicabilityReportConfig_r19).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 5. dataCollectionPreferenceConfig-r19: choice
	{
		if ie.DataCollectionPreferenceConfig_r19 != nil {
			choiceEnc := e.NewChoiceEncoder(otherConfig_v1900DataCollectionPreferenceConfigR19Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.DataCollectionPreferenceConfig_r19).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.DataCollectionPreferenceConfig_r19).Choice {
			case OtherConfig_v1900_DataCollectionPreferenceConfig_r19_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1900_DataCollectionPreferenceConfig_r19_Setup:
				if err := (*ie.DataCollectionPreferenceConfig_r19).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.DataCollectionPreferenceConfig_r19).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 6. loggedDataCollectionAssistanceConfig-r19: choice
	{
		if ie.LoggedDataCollectionAssistanceConfig_r19 != nil {
			choiceEnc := e.NewChoiceEncoder(otherConfig_v1900LoggedDataCollectionAssistanceConfigR19Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.LoggedDataCollectionAssistanceConfig_r19).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.LoggedDataCollectionAssistanceConfig_r19).Choice {
			case OtherConfig_v1900_LoggedDataCollectionAssistanceConfig_r19_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1900_LoggedDataCollectionAssistanceConfig_r19_Setup:
				if err := (*ie.LoggedDataCollectionAssistanceConfig_r19).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.LoggedDataCollectionAssistanceConfig_r19).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 7. assisted-SSB-MTC-MG-Config-r19: choice
	{
		if ie.Assisted_SSB_MTC_MG_Config_r19 != nil {
			choiceEnc := e.NewChoiceEncoder(otherConfig_v1900AssistedSSBMTCMGConfigR19Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Assisted_SSB_MTC_MG_Config_r19).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Assisted_SSB_MTC_MG_Config_r19).Choice {
			case OtherConfig_v1900_Assisted_SSB_MTC_MG_Config_r19_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1900_Assisted_SSB_MTC_MG_Config_r19_Setup:
				if err := (*ie.Assisted_SSB_MTC_MG_Config_r19).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Assisted_SSB_MTC_MG_Config_r19).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 8. fbs-PreferenceReportingConfig-r19: choice
	{
		if ie.Fbs_PreferenceReportingConfig_r19 != nil {
			choiceEnc := e.NewChoiceEncoder(otherConfig_v1900FbsPreferenceReportingConfigR19Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Fbs_PreferenceReportingConfig_r19).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Fbs_PreferenceReportingConfig_r19).Choice {
			case OtherConfig_v1900_Fbs_PreferenceReportingConfig_r19_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1900_Fbs_PreferenceReportingConfig_r19_Setup:
				if err := (*ie.Fbs_PreferenceReportingConfig_r19).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Fbs_PreferenceReportingConfig_r19).Choice), Constraint: "index out of range"}
			}
		}
	}

	return nil
}

func (ie *OtherConfig_v1900) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(otherConfigV1900Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. gapOccasionCancelRatioReportConfig-r19: choice
	{
		if seq.IsComponentPresent(0) {
			ie.GapOccasionCancelRatioReportConfig_r19 = &struct {
				Choice  int
				Release *struct{}
				Setup   *GapOccasionCancelRatioReportConfig_r19
			}{}
			choiceDec := d.NewChoiceDecoder(otherConfig_v1900GapOccasionCancelRatioReportConfigR19Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.GapOccasionCancelRatioReportConfig_r19).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case OtherConfig_v1900_GapOccasionCancelRatioReportConfig_r19_Release:
				(*ie.GapOccasionCancelRatioReportConfig_r19).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1900_GapOccasionCancelRatioReportConfig_r19_Setup:
				(*ie.GapOccasionCancelRatioReportConfig_r19).Setup = new(GapOccasionCancelRatioReportConfig_r19)
				if err := (*ie.GapOccasionCancelRatioReportConfig_r19).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. lpwus-OffsetPreferenceConfig-r19: choice
	{
		if seq.IsComponentPresent(1) {
			ie.Lpwus_OffsetPreferenceConfig_r19 = &struct {
				Choice  int
				Release *struct{}
				Setup   *LPWUS_OffsetPreferenceConfig_r19
			}{}
			choiceDec := d.NewChoiceDecoder(otherConfig_v1900LpwusOffsetPreferenceConfigR19Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Lpwus_OffsetPreferenceConfig_r19).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case OtherConfig_v1900_Lpwus_OffsetPreferenceConfig_r19_Release:
				(*ie.Lpwus_OffsetPreferenceConfig_r19).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1900_Lpwus_OffsetPreferenceConfig_r19_Setup:
				(*ie.Lpwus_OffsetPreferenceConfig_r19).Setup = new(LPWUS_OffsetPreferenceConfig_r19)
				if err := (*ie.Lpwus_OffsetPreferenceConfig_r19).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. applicabilityReportConfig-r19: choice
	{
		if seq.IsComponentPresent(2) {
			ie.ApplicabilityReportConfig_r19 = &struct {
				Choice  int
				Release *struct{}
				Setup   *ApplicabilityReportConfig_r19
			}{}
			choiceDec := d.NewChoiceDecoder(otherConfig_v1900ApplicabilityReportConfigR19Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.ApplicabilityReportConfig_r19).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case OtherConfig_v1900_ApplicabilityReportConfig_r19_Release:
				(*ie.ApplicabilityReportConfig_r19).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1900_ApplicabilityReportConfig_r19_Setup:
				(*ie.ApplicabilityReportConfig_r19).Setup = new(ApplicabilityReportConfig_r19)
				if err := (*ie.ApplicabilityReportConfig_r19).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. dataCollectionPreferenceConfig-r19: choice
	{
		if seq.IsComponentPresent(3) {
			ie.DataCollectionPreferenceConfig_r19 = &struct {
				Choice  int
				Release *struct{}
				Setup   *DataCollectionPreferenceConfig_r19
			}{}
			choiceDec := d.NewChoiceDecoder(otherConfig_v1900DataCollectionPreferenceConfigR19Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.DataCollectionPreferenceConfig_r19).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case OtherConfig_v1900_DataCollectionPreferenceConfig_r19_Release:
				(*ie.DataCollectionPreferenceConfig_r19).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1900_DataCollectionPreferenceConfig_r19_Setup:
				(*ie.DataCollectionPreferenceConfig_r19).Setup = new(DataCollectionPreferenceConfig_r19)
				if err := (*ie.DataCollectionPreferenceConfig_r19).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. loggedDataCollectionAssistanceConfig-r19: choice
	{
		if seq.IsComponentPresent(4) {
			ie.LoggedDataCollectionAssistanceConfig_r19 = &struct {
				Choice  int
				Release *struct{}
				Setup   *LoggedDataCollectionAssistanceConfig_r19
			}{}
			choiceDec := d.NewChoiceDecoder(otherConfig_v1900LoggedDataCollectionAssistanceConfigR19Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.LoggedDataCollectionAssistanceConfig_r19).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case OtherConfig_v1900_LoggedDataCollectionAssistanceConfig_r19_Release:
				(*ie.LoggedDataCollectionAssistanceConfig_r19).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1900_LoggedDataCollectionAssistanceConfig_r19_Setup:
				(*ie.LoggedDataCollectionAssistanceConfig_r19).Setup = new(LoggedDataCollectionAssistanceConfig_r19)
				if err := (*ie.LoggedDataCollectionAssistanceConfig_r19).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 7. assisted-SSB-MTC-MG-Config-r19: choice
	{
		if seq.IsComponentPresent(5) {
			ie.Assisted_SSB_MTC_MG_Config_r19 = &struct {
				Choice  int
				Release *struct{}
				Setup   *Assisted_SSB_MTC_MG_Config_r19
			}{}
			choiceDec := d.NewChoiceDecoder(otherConfig_v1900AssistedSSBMTCMGConfigR19Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Assisted_SSB_MTC_MG_Config_r19).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case OtherConfig_v1900_Assisted_SSB_MTC_MG_Config_r19_Release:
				(*ie.Assisted_SSB_MTC_MG_Config_r19).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1900_Assisted_SSB_MTC_MG_Config_r19_Setup:
				(*ie.Assisted_SSB_MTC_MG_Config_r19).Setup = new(Assisted_SSB_MTC_MG_Config_r19)
				if err := (*ie.Assisted_SSB_MTC_MG_Config_r19).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 8. fbs-PreferenceReportingConfig-r19: choice
	{
		if seq.IsComponentPresent(6) {
			ie.Fbs_PreferenceReportingConfig_r19 = &struct {
				Choice  int
				Release *struct{}
				Setup   *FBS_PreferenceReportingConfig_r19
			}{}
			choiceDec := d.NewChoiceDecoder(otherConfig_v1900FbsPreferenceReportingConfigR19Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Fbs_PreferenceReportingConfig_r19).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case OtherConfig_v1900_Fbs_PreferenceReportingConfig_r19_Release:
				(*ie.Fbs_PreferenceReportingConfig_r19).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1900_Fbs_PreferenceReportingConfig_r19_Setup:
				(*ie.Fbs_PreferenceReportingConfig_r19).Setup = new(FBS_PreferenceReportingConfig_r19)
				if err := (*ie.Fbs_PreferenceReportingConfig_r19).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

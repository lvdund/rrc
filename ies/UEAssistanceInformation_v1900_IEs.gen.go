// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UEAssistanceInformation-v1900-IEs (line 2487).

var uEAssistanceInformationV1900IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "gapOccasionCancelRatio-r19", Optional: true},
		{Name: "lpwus-OffsetPreference-r19", Optional: true},
		{Name: "applicabilityReportList-r19", Optional: true},
		{Name: "dataCollectionPreference-r19", Optional: true},
		{Name: "loggedDataCollectionAssistance-r19", Optional: true},
		{Name: "assisted-SSB-MTC-MG-Report-r19", Optional: true},
		{Name: "fbs-Preference-r19", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var uEAssistanceInformationV1900IEsAssistedSSBMTCMGReportR19Constraints = per.FixedSize(7)

const (
	UEAssistanceInformation_v1900_IEs_Fbs_Preference_r19_Preferred    = 0
	UEAssistanceInformation_v1900_IEs_Fbs_Preference_r19_NotPreferred = 1
)

var uEAssistanceInformationV1900IEsFbsPreferenceR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UEAssistanceInformation_v1900_IEs_Fbs_Preference_r19_Preferred, UEAssistanceInformation_v1900_IEs_Fbs_Preference_r19_NotPreferred},
}

type UEAssistanceInformation_v1900_IEs struct {
	GapOccasionCancelRatio_r19         *GapOccasionCancelRatio_r19
	Lpwus_OffsetPreference_r19         *LPWUS_OffsetPreference_r19
	ApplicabilityReportList_r19        *ApplicabilityReportList_r19
	DataCollectionPreference_r19       *DataCollectionPreference_r19
	LoggedDataCollectionAssistance_r19 *LoggedDataCollectionAssistance_r19
	Assisted_SSB_MTC_MG_Report_r19     *per.BitString
	Fbs_Preference_r19                 *int64
	NonCriticalExtension               *struct{}
}

func (ie *UEAssistanceInformation_v1900_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEAssistanceInformationV1900IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.GapOccasionCancelRatio_r19 != nil, ie.Lpwus_OffsetPreference_r19 != nil, ie.ApplicabilityReportList_r19 != nil, ie.DataCollectionPreference_r19 != nil, ie.LoggedDataCollectionAssistance_r19 != nil, ie.Assisted_SSB_MTC_MG_Report_r19 != nil, ie.Fbs_Preference_r19 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. gapOccasionCancelRatio-r19: ref
	{
		if ie.GapOccasionCancelRatio_r19 != nil {
			if err := ie.GapOccasionCancelRatio_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. lpwus-OffsetPreference-r19: ref
	{
		if ie.Lpwus_OffsetPreference_r19 != nil {
			if err := ie.Lpwus_OffsetPreference_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. applicabilityReportList-r19: ref
	{
		if ie.ApplicabilityReportList_r19 != nil {
			if err := ie.ApplicabilityReportList_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. dataCollectionPreference-r19: ref
	{
		if ie.DataCollectionPreference_r19 != nil {
			if err := ie.DataCollectionPreference_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. loggedDataCollectionAssistance-r19: ref
	{
		if ie.LoggedDataCollectionAssistance_r19 != nil {
			if err := ie.LoggedDataCollectionAssistance_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. assisted-SSB-MTC-MG-Report-r19: bit-string
	{
		if ie.Assisted_SSB_MTC_MG_Report_r19 != nil {
			if err := e.EncodeBitString(*ie.Assisted_SSB_MTC_MG_Report_r19, uEAssistanceInformationV1900IEsAssistedSSBMTCMGReportR19Constraints); err != nil {
				return err
			}
		}
	}

	// 8. fbs-Preference-r19: enumerated
	{
		if ie.Fbs_Preference_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Fbs_Preference_r19, uEAssistanceInformationV1900IEsFbsPreferenceR19Constraints); err != nil {
				return err
			}
		}
	}

	// 9. nonCriticalExtension: sequence
	{
		if ie.NonCriticalExtension != nil {
			// empty SEQUENCE {}: no content to encode.
		}
	}

	return nil
}

func (ie *UEAssistanceInformation_v1900_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEAssistanceInformationV1900IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. gapOccasionCancelRatio-r19: ref
	{
		if seq.IsComponentPresent(0) {
			ie.GapOccasionCancelRatio_r19 = new(GapOccasionCancelRatio_r19)
			if err := ie.GapOccasionCancelRatio_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. lpwus-OffsetPreference-r19: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Lpwus_OffsetPreference_r19 = new(LPWUS_OffsetPreference_r19)
			if err := ie.Lpwus_OffsetPreference_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. applicabilityReportList-r19: ref
	{
		if seq.IsComponentPresent(2) {
			ie.ApplicabilityReportList_r19 = new(ApplicabilityReportList_r19)
			if err := ie.ApplicabilityReportList_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. dataCollectionPreference-r19: ref
	{
		if seq.IsComponentPresent(3) {
			ie.DataCollectionPreference_r19 = new(DataCollectionPreference_r19)
			if err := ie.DataCollectionPreference_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. loggedDataCollectionAssistance-r19: ref
	{
		if seq.IsComponentPresent(4) {
			ie.LoggedDataCollectionAssistance_r19 = new(LoggedDataCollectionAssistance_r19)
			if err := ie.LoggedDataCollectionAssistance_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. assisted-SSB-MTC-MG-Report-r19: bit-string
	{
		if seq.IsComponentPresent(5) {
			v, err := d.DecodeBitString(uEAssistanceInformationV1900IEsAssistedSSBMTCMGReportR19Constraints)
			if err != nil {
				return err
			}
			ie.Assisted_SSB_MTC_MG_Report_r19 = &v
		}
	}

	// 8. fbs-Preference-r19: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(uEAssistanceInformationV1900IEsFbsPreferenceR19Constraints)
			if err != nil {
				return err
			}
			ie.Fbs_Preference_r19 = &idx
		}
	}

	// 9. nonCriticalExtension: sequence
	{
		if seq.IsComponentPresent(7) {
			ie.NonCriticalExtension = &struct{}{}
		}
	}

	return nil
}

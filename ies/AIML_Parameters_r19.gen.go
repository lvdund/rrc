// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: AIML-Parameters-r19 (line 16511).

var aIMLParametersR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "applicabilityReportingCSI-r19", Optional: true},
		{Name: "applicabilityReportingOther-r19", Optional: true},
		{Name: "loggedDataCollection-r19", Optional: true},
		{Name: "eventBasedLoggedDataCollection-r19", Optional: true},
		{Name: "dataThresholdAvailabilityIndication-r19", Optional: true},
	},
}

const (
	AIML_Parameters_r19_ApplicabilityReportingCSI_r19_Supported = 0
)

var aIMLParametersR19ApplicabilityReportingCSIR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AIML_Parameters_r19_ApplicabilityReportingCSI_r19_Supported},
}

const (
	AIML_Parameters_r19_ApplicabilityReportingOther_r19_Supported = 0
)

var aIMLParametersR19ApplicabilityReportingOtherR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AIML_Parameters_r19_ApplicabilityReportingOther_r19_Supported},
}

const (
	AIML_Parameters_r19_LoggedDataCollection_r19_Supported = 0
)

var aIMLParametersR19LoggedDataCollectionR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AIML_Parameters_r19_LoggedDataCollection_r19_Supported},
}

const (
	AIML_Parameters_r19_EventBasedLoggedDataCollection_r19_Supported = 0
)

var aIMLParametersR19EventBasedLoggedDataCollectionR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AIML_Parameters_r19_EventBasedLoggedDataCollection_r19_Supported},
}

const (
	AIML_Parameters_r19_DataThresholdAvailabilityIndication_r19_Supported = 0
)

var aIMLParametersR19DataThresholdAvailabilityIndicationR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AIML_Parameters_r19_DataThresholdAvailabilityIndication_r19_Supported},
}

type AIML_Parameters_r19 struct {
	ApplicabilityReportingCSI_r19           *int64
	ApplicabilityReportingOther_r19         *int64
	LoggedDataCollection_r19                *int64
	EventBasedLoggedDataCollection_r19      *int64
	DataThresholdAvailabilityIndication_r19 *int64
}

func (ie *AIML_Parameters_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(aIMLParametersR19Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ApplicabilityReportingCSI_r19 != nil, ie.ApplicabilityReportingOther_r19 != nil, ie.LoggedDataCollection_r19 != nil, ie.EventBasedLoggedDataCollection_r19 != nil, ie.DataThresholdAvailabilityIndication_r19 != nil}); err != nil {
		return err
	}

	// 2. applicabilityReportingCSI-r19: enumerated
	{
		if ie.ApplicabilityReportingCSI_r19 != nil {
			if err := e.EncodeEnumerated(*ie.ApplicabilityReportingCSI_r19, aIMLParametersR19ApplicabilityReportingCSIR19Constraints); err != nil {
				return err
			}
		}
	}

	// 3. applicabilityReportingOther-r19: enumerated
	{
		if ie.ApplicabilityReportingOther_r19 != nil {
			if err := e.EncodeEnumerated(*ie.ApplicabilityReportingOther_r19, aIMLParametersR19ApplicabilityReportingOtherR19Constraints); err != nil {
				return err
			}
		}
	}

	// 4. loggedDataCollection-r19: enumerated
	{
		if ie.LoggedDataCollection_r19 != nil {
			if err := e.EncodeEnumerated(*ie.LoggedDataCollection_r19, aIMLParametersR19LoggedDataCollectionR19Constraints); err != nil {
				return err
			}
		}
	}

	// 5. eventBasedLoggedDataCollection-r19: enumerated
	{
		if ie.EventBasedLoggedDataCollection_r19 != nil {
			if err := e.EncodeEnumerated(*ie.EventBasedLoggedDataCollection_r19, aIMLParametersR19EventBasedLoggedDataCollectionR19Constraints); err != nil {
				return err
			}
		}
	}

	// 6. dataThresholdAvailabilityIndication-r19: enumerated
	{
		if ie.DataThresholdAvailabilityIndication_r19 != nil {
			if err := e.EncodeEnumerated(*ie.DataThresholdAvailabilityIndication_r19, aIMLParametersR19DataThresholdAvailabilityIndicationR19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *AIML_Parameters_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(aIMLParametersR19Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. applicabilityReportingCSI-r19: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(aIMLParametersR19ApplicabilityReportingCSIR19Constraints)
			if err != nil {
				return err
			}
			ie.ApplicabilityReportingCSI_r19 = &idx
		}
	}

	// 3. applicabilityReportingOther-r19: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(aIMLParametersR19ApplicabilityReportingOtherR19Constraints)
			if err != nil {
				return err
			}
			ie.ApplicabilityReportingOther_r19 = &idx
		}
	}

	// 4. loggedDataCollection-r19: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(aIMLParametersR19LoggedDataCollectionR19Constraints)
			if err != nil {
				return err
			}
			ie.LoggedDataCollection_r19 = &idx
		}
	}

	// 5. eventBasedLoggedDataCollection-r19: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(aIMLParametersR19EventBasedLoggedDataCollectionR19Constraints)
			if err != nil {
				return err
			}
			ie.EventBasedLoggedDataCollection_r19 = &idx
		}
	}

	// 6. dataThresholdAvailabilityIndication-r19: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(aIMLParametersR19DataThresholdAvailabilityIndicationR19Constraints)
			if err != nil {
				return err
			}
			ie.DataThresholdAvailabilityIndication_r19 = &idx
		}
	}

	return nil
}

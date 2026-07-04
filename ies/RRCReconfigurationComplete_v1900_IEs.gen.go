// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCReconfigurationComplete-v1900-IEs (line 1195).

var rRCReconfigurationCompleteV1900IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "applicabilityReportList-r19", Optional: true},
		{Name: "csi-LogMeasAvailable-r19", Optional: true},
		{Name: "assisted-SSB-MTC-MG-Report-r19", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	RRCReconfigurationComplete_v1900_IEs_Csi_LogMeasAvailable_r19_True = 0
)

var rRCReconfigurationCompleteV1900IEsCsiLogMeasAvailableR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCReconfigurationComplete_v1900_IEs_Csi_LogMeasAvailable_r19_True},
}

var rRCReconfigurationCompleteV1900IEsAssistedSSBMTCMGReportR19Constraints = per.FixedSize(7)

type RRCReconfigurationComplete_v1900_IEs struct {
	ApplicabilityReportList_r19    *ApplicabilityReportList_r19
	Csi_LogMeasAvailable_r19       *int64
	Assisted_SSB_MTC_MG_Report_r19 *per.BitString
	NonCriticalExtension           *struct{}
}

func (ie *RRCReconfigurationComplete_v1900_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCReconfigurationCompleteV1900IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ApplicabilityReportList_r19 != nil, ie.Csi_LogMeasAvailable_r19 != nil, ie.Assisted_SSB_MTC_MG_Report_r19 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. applicabilityReportList-r19: ref
	{
		if ie.ApplicabilityReportList_r19 != nil {
			if err := ie.ApplicabilityReportList_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. csi-LogMeasAvailable-r19: enumerated
	{
		if ie.Csi_LogMeasAvailable_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Csi_LogMeasAvailable_r19, rRCReconfigurationCompleteV1900IEsCsiLogMeasAvailableR19Constraints); err != nil {
				return err
			}
		}
	}

	// 4. assisted-SSB-MTC-MG-Report-r19: bit-string
	{
		if ie.Assisted_SSB_MTC_MG_Report_r19 != nil {
			if err := e.EncodeBitString(*ie.Assisted_SSB_MTC_MG_Report_r19, rRCReconfigurationCompleteV1900IEsAssistedSSBMTCMGReportR19Constraints); err != nil {
				return err
			}
		}
	}

	// 5. nonCriticalExtension: sequence
	{
		if ie.NonCriticalExtension != nil {
			// empty SEQUENCE {}: no content to encode.
		}
	}

	return nil
}

func (ie *RRCReconfigurationComplete_v1900_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCReconfigurationCompleteV1900IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. applicabilityReportList-r19: ref
	{
		if seq.IsComponentPresent(0) {
			ie.ApplicabilityReportList_r19 = new(ApplicabilityReportList_r19)
			if err := ie.ApplicabilityReportList_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. csi-LogMeasAvailable-r19: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(rRCReconfigurationCompleteV1900IEsCsiLogMeasAvailableR19Constraints)
			if err != nil {
				return err
			}
			ie.Csi_LogMeasAvailable_r19 = &idx
		}
	}

	// 4. assisted-SSB-MTC-MG-Report-r19: bit-string
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeBitString(rRCReconfigurationCompleteV1900IEsAssistedSSBMTCMGReportR19Constraints)
			if err != nil {
				return err
			}
			ie.Assisted_SSB_MTC_MG_Report_r19 = &v
		}
	}

	// 5. nonCriticalExtension: sequence
	{
		if seq.IsComponentPresent(3) {
			ie.NonCriticalExtension = &struct{}{}
		}
	}

	return nil
}

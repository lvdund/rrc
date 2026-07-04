// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DLInformationTransfer-v1800-IEs (line 365).

var dLInformationTransferV1800IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "eventID-TSS-r18", Optional: true},
		{Name: "clockQualityDetailsLevel-r18", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var dLInformationTransferV1800IEsEventIDTSSR18Constraints = per.Constrained(0, 63)

var dLInformationTransfer_v1800_IEsClockQualityDetailsLevelR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "clockQualityMetrics-r18"},
		{Name: "clockQualityAcceptanceStatus-r18"},
	},
}

const (
	DLInformationTransfer_v1800_IEs_ClockQualityDetailsLevel_r18_ClockQualityMetrics_r18          = 0
	DLInformationTransfer_v1800_IEs_ClockQualityDetailsLevel_r18_ClockQualityAcceptanceStatus_r18 = 1
)

const (
	DLInformationTransfer_v1800_IEs_ClockQualityDetailsLevel_r18_ClockQualityAcceptanceStatus_r18_Acceptable    = 0
	DLInformationTransfer_v1800_IEs_ClockQualityDetailsLevel_r18_ClockQualityAcceptanceStatus_r18_NotAcceptable = 1
)

var dLInformationTransferV1800IEsClockQualityDetailsLevelR18ClockQualityAcceptanceStatusR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DLInformationTransfer_v1800_IEs_ClockQualityDetailsLevel_r18_ClockQualityAcceptanceStatus_r18_Acceptable, DLInformationTransfer_v1800_IEs_ClockQualityDetailsLevel_r18_ClockQualityAcceptanceStatus_r18_NotAcceptable},
}

type DLInformationTransfer_v1800_IEs struct {
	EventID_TSS_r18              *int64
	ClockQualityDetailsLevel_r18 *struct {
		Choice                           int
		ClockQualityMetrics_r18          *ClockQualityMetrics_r18
		ClockQualityAcceptanceStatus_r18 *int64
	}
	NonCriticalExtension *struct{}
}

func (ie *DLInformationTransfer_v1800_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dLInformationTransferV1800IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.EventID_TSS_r18 != nil, ie.ClockQualityDetailsLevel_r18 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. eventID-TSS-r18: integer
	{
		if ie.EventID_TSS_r18 != nil {
			if err := e.EncodeInteger(*ie.EventID_TSS_r18, dLInformationTransferV1800IEsEventIDTSSR18Constraints); err != nil {
				return err
			}
		}
	}

	// 3. clockQualityDetailsLevel-r18: choice
	{
		if ie.ClockQualityDetailsLevel_r18 != nil {
			choiceEnc := e.NewChoiceEncoder(dLInformationTransfer_v1800_IEsClockQualityDetailsLevelR18Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.ClockQualityDetailsLevel_r18).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.ClockQualityDetailsLevel_r18).Choice {
			case DLInformationTransfer_v1800_IEs_ClockQualityDetailsLevel_r18_ClockQualityMetrics_r18:
				if err := (*ie.ClockQualityDetailsLevel_r18).ClockQualityMetrics_r18.Encode(e); err != nil {
					return err
				}
			case DLInformationTransfer_v1800_IEs_ClockQualityDetailsLevel_r18_ClockQualityAcceptanceStatus_r18:
				if err := e.EncodeEnumerated((*(*ie.ClockQualityDetailsLevel_r18).ClockQualityAcceptanceStatus_r18), dLInformationTransferV1800IEsClockQualityDetailsLevelR18ClockQualityAcceptanceStatusR18Constraints); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.ClockQualityDetailsLevel_r18).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 4. nonCriticalExtension: sequence
	{
		if ie.NonCriticalExtension != nil {
			// empty SEQUENCE {}: no content to encode.
		}
	}

	return nil
}

func (ie *DLInformationTransfer_v1800_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dLInformationTransferV1800IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. eventID-TSS-r18: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(dLInformationTransferV1800IEsEventIDTSSR18Constraints)
			if err != nil {
				return err
			}
			ie.EventID_TSS_r18 = &v
		}
	}

	// 3. clockQualityDetailsLevel-r18: choice
	{
		if seq.IsComponentPresent(1) {
			ie.ClockQualityDetailsLevel_r18 = &struct {
				Choice                           int
				ClockQualityMetrics_r18          *ClockQualityMetrics_r18
				ClockQualityAcceptanceStatus_r18 *int64
			}{}
			choiceDec := d.NewChoiceDecoder(dLInformationTransfer_v1800_IEsClockQualityDetailsLevelR18Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.ClockQualityDetailsLevel_r18).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case DLInformationTransfer_v1800_IEs_ClockQualityDetailsLevel_r18_ClockQualityMetrics_r18:
				(*ie.ClockQualityDetailsLevel_r18).ClockQualityMetrics_r18 = new(ClockQualityMetrics_r18)
				if err := (*ie.ClockQualityDetailsLevel_r18).ClockQualityMetrics_r18.Decode(d); err != nil {
					return err
				}
			case DLInformationTransfer_v1800_IEs_ClockQualityDetailsLevel_r18_ClockQualityAcceptanceStatus_r18:
				(*ie.ClockQualityDetailsLevel_r18).ClockQualityAcceptanceStatus_r18 = new(int64)
				v, err := d.DecodeEnumerated(dLInformationTransferV1800IEsClockQualityDetailsLevelR18ClockQualityAcceptanceStatusR18Constraints)
				if err != nil {
					return err
				}
				(*(*ie.ClockQualityDetailsLevel_r18).ClockQualityAcceptanceStatus_r18) = v
			}
		}
	}

	// 4. nonCriticalExtension: sequence
	{
		if seq.IsComponentPresent(2) {
			ie.NonCriticalExtension = &struct{}{}
		}
	}

	return nil
}

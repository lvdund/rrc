// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: LogMeasReport-r16 (line 2989).

var logMeasReportR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "absoluteTimeStamp-r16"},
		{Name: "traceReference-r16"},
		{Name: "traceRecordingSessionRef-r16"},
		{Name: "tce-Id-r16"},
		{Name: "logMeasInfoList-r16"},
		{Name: "logMeasAvailable-r16", Optional: true},
		{Name: "logMeasAvailableBT-r16", Optional: true},
		{Name: "logMeasAvailableWLAN-r16", Optional: true},
	},
}

var logMeasReportR16TraceRecordingSessionRefR16Constraints = per.FixedSize(2)

var logMeasReportR16TceIdR16Constraints = per.FixedSize(1)

const (
	LogMeasReport_r16_LogMeasAvailable_r16_True = 0
)

var logMeasReportR16LogMeasAvailableR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LogMeasReport_r16_LogMeasAvailable_r16_True},
}

const (
	LogMeasReport_r16_LogMeasAvailableBT_r16_True = 0
)

var logMeasReportR16LogMeasAvailableBTR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LogMeasReport_r16_LogMeasAvailableBT_r16_True},
}

const (
	LogMeasReport_r16_LogMeasAvailableWLAN_r16_True = 0
)

var logMeasReportR16LogMeasAvailableWLANR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LogMeasReport_r16_LogMeasAvailableWLAN_r16_True},
}

type LogMeasReport_r16 struct {
	AbsoluteTimeStamp_r16        AbsoluteTimeInfo_r16
	TraceReference_r16           TraceReference_r16
	TraceRecordingSessionRef_r16 []byte
	Tce_Id_r16                   []byte
	LogMeasInfoList_r16          LogMeasInfoList_r16
	LogMeasAvailable_r16         *int64
	LogMeasAvailableBT_r16       *int64
	LogMeasAvailableWLAN_r16     *int64
}

func (ie *LogMeasReport_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(logMeasReportR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.LogMeasAvailable_r16 != nil, ie.LogMeasAvailableBT_r16 != nil, ie.LogMeasAvailableWLAN_r16 != nil}); err != nil {
		return err
	}

	// 3. absoluteTimeStamp-r16: ref
	{
		if err := ie.AbsoluteTimeStamp_r16.Encode(e); err != nil {
			return err
		}
	}

	// 4. traceReference-r16: ref
	{
		if err := ie.TraceReference_r16.Encode(e); err != nil {
			return err
		}
	}

	// 5. traceRecordingSessionRef-r16: octet-string
	{
		if err := e.EncodeOctetString(ie.TraceRecordingSessionRef_r16, logMeasReportR16TraceRecordingSessionRefR16Constraints); err != nil {
			return err
		}
	}

	// 6. tce-Id-r16: octet-string
	{
		if err := e.EncodeOctetString(ie.Tce_Id_r16, logMeasReportR16TceIdR16Constraints); err != nil {
			return err
		}
	}

	// 7. logMeasInfoList-r16: ref
	{
		if err := ie.LogMeasInfoList_r16.Encode(e); err != nil {
			return err
		}
	}

	// 8. logMeasAvailable-r16: enumerated
	{
		if ie.LogMeasAvailable_r16 != nil {
			if err := e.EncodeEnumerated(*ie.LogMeasAvailable_r16, logMeasReportR16LogMeasAvailableR16Constraints); err != nil {
				return err
			}
		}
	}

	// 9. logMeasAvailableBT-r16: enumerated
	{
		if ie.LogMeasAvailableBT_r16 != nil {
			if err := e.EncodeEnumerated(*ie.LogMeasAvailableBT_r16, logMeasReportR16LogMeasAvailableBTR16Constraints); err != nil {
				return err
			}
		}
	}

	// 10. logMeasAvailableWLAN-r16: enumerated
	{
		if ie.LogMeasAvailableWLAN_r16 != nil {
			if err := e.EncodeEnumerated(*ie.LogMeasAvailableWLAN_r16, logMeasReportR16LogMeasAvailableWLANR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *LogMeasReport_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(logMeasReportR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. absoluteTimeStamp-r16: ref
	{
		if err := ie.AbsoluteTimeStamp_r16.Decode(d); err != nil {
			return err
		}
	}

	// 4. traceReference-r16: ref
	{
		if err := ie.TraceReference_r16.Decode(d); err != nil {
			return err
		}
	}

	// 5. traceRecordingSessionRef-r16: octet-string
	{
		v2, err := d.DecodeOctetString(logMeasReportR16TraceRecordingSessionRefR16Constraints)
		if err != nil {
			return err
		}
		ie.TraceRecordingSessionRef_r16 = v2
	}

	// 6. tce-Id-r16: octet-string
	{
		v3, err := d.DecodeOctetString(logMeasReportR16TceIdR16Constraints)
		if err != nil {
			return err
		}
		ie.Tce_Id_r16 = v3
	}

	// 7. logMeasInfoList-r16: ref
	{
		if err := ie.LogMeasInfoList_r16.Decode(d); err != nil {
			return err
		}
	}

	// 8. logMeasAvailable-r16: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(logMeasReportR16LogMeasAvailableR16Constraints)
			if err != nil {
				return err
			}
			ie.LogMeasAvailable_r16 = &idx
		}
	}

	// 9. logMeasAvailableBT-r16: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(logMeasReportR16LogMeasAvailableBTR16Constraints)
			if err != nil {
				return err
			}
			ie.LogMeasAvailableBT_r16 = &idx
		}
	}

	// 10. logMeasAvailableWLAN-r16: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(logMeasReportR16LogMeasAvailableWLANR16Constraints)
			if err != nil {
				return err
			}
			ie.LogMeasAvailableWLAN_r16 = &idx
		}
	}

	return nil
}

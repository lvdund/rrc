// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: AppLayerIdleInactiveConfig-r18 (line 25981).

var appLayerIdleInactiveConfigR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "measConfigAppLayerId-r18"},
		{Name: "serviceType-r18", Optional: true},
		{Name: "appLayerMeasPriority-r18", Optional: true},
		{Name: "qoe-Reference-r18", Optional: true},
		{Name: "qoe-MeasurementType-r18", Optional: true},
		{Name: "qoe-AreaScope-r18", Optional: true},
		{Name: "mce-Id-r18", Optional: true},
		{Name: "availableRAN-VisibleMetrics-r18", Optional: true},
	},
}

const (
	AppLayerIdleInactiveConfig_r18_ServiceType_r18_Streaming = 0
	AppLayerIdleInactiveConfig_r18_ServiceType_r18_Mtsi      = 1
	AppLayerIdleInactiveConfig_r18_ServiceType_r18_Vr        = 2
	AppLayerIdleInactiveConfig_r18_ServiceType_r18_Spare5    = 3
	AppLayerIdleInactiveConfig_r18_ServiceType_r18_Spare4    = 4
	AppLayerIdleInactiveConfig_r18_ServiceType_r18_Spare3    = 5
	AppLayerIdleInactiveConfig_r18_ServiceType_r18_Spare2    = 6
	AppLayerIdleInactiveConfig_r18_ServiceType_r18_Spare1    = 7
)

var appLayerIdleInactiveConfigR18ServiceTypeR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AppLayerIdleInactiveConfig_r18_ServiceType_r18_Streaming, AppLayerIdleInactiveConfig_r18_ServiceType_r18_Mtsi, AppLayerIdleInactiveConfig_r18_ServiceType_r18_Vr, AppLayerIdleInactiveConfig_r18_ServiceType_r18_Spare5, AppLayerIdleInactiveConfig_r18_ServiceType_r18_Spare4, AppLayerIdleInactiveConfig_r18_ServiceType_r18_Spare3, AppLayerIdleInactiveConfig_r18_ServiceType_r18_Spare2, AppLayerIdleInactiveConfig_r18_ServiceType_r18_Spare1},
}

var appLayerIdleInactiveConfigR18AppLayerMeasPriorityR18Constraints = per.Constrained(1, 16)

var appLayerIdleInactiveConfigR18QoeReferenceR18Constraints = per.FixedSize(6)

const (
	AppLayerIdleInactiveConfig_r18_Qoe_MeasurementType_r18_Sbased = 0
	AppLayerIdleInactiveConfig_r18_Qoe_MeasurementType_r18_Mbased = 1
)

var appLayerIdleInactiveConfigR18QoeMeasurementTypeR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AppLayerIdleInactiveConfig_r18_Qoe_MeasurementType_r18_Sbased, AppLayerIdleInactiveConfig_r18_Qoe_MeasurementType_r18_Mbased},
}

var appLayerIdleInactiveConfigR18MceIdR18Constraints = per.FixedSize(1)

type AppLayerIdleInactiveConfig_r18 struct {
	MeasConfigAppLayerId_r18        MeasConfigAppLayerId_r17
	ServiceType_r18                 *int64
	AppLayerMeasPriority_r18        *int64
	Qoe_Reference_r18               []byte
	Qoe_MeasurementType_r18         *int64
	Qoe_AreaScope_r18               *Qoe_AreaScope_r18
	Mce_Id_r18                      []byte
	AvailableRAN_VisibleMetrics_r18 *AvailableRAN_VisibleMetrics_r18
}

func (ie *AppLayerIdleInactiveConfig_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(appLayerIdleInactiveConfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ServiceType_r18 != nil, ie.AppLayerMeasPriority_r18 != nil, ie.Qoe_Reference_r18 != nil, ie.Qoe_MeasurementType_r18 != nil, ie.Qoe_AreaScope_r18 != nil, ie.Mce_Id_r18 != nil, ie.AvailableRAN_VisibleMetrics_r18 != nil}); err != nil {
		return err
	}

	// 3. measConfigAppLayerId-r18: ref
	{
		if err := ie.MeasConfigAppLayerId_r18.Encode(e); err != nil {
			return err
		}
	}

	// 4. serviceType-r18: enumerated
	{
		if ie.ServiceType_r18 != nil {
			if err := e.EncodeEnumerated(*ie.ServiceType_r18, appLayerIdleInactiveConfigR18ServiceTypeR18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. appLayerMeasPriority-r18: integer
	{
		if ie.AppLayerMeasPriority_r18 != nil {
			if err := e.EncodeInteger(*ie.AppLayerMeasPriority_r18, appLayerIdleInactiveConfigR18AppLayerMeasPriorityR18Constraints); err != nil {
				return err
			}
		}
	}

	// 6. qoe-Reference-r18: octet-string
	{
		if ie.Qoe_Reference_r18 != nil {
			if err := e.EncodeOctetString(ie.Qoe_Reference_r18, appLayerIdleInactiveConfigR18QoeReferenceR18Constraints); err != nil {
				return err
			}
		}
	}

	// 7. qoe-MeasurementType-r18: enumerated
	{
		if ie.Qoe_MeasurementType_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Qoe_MeasurementType_r18, appLayerIdleInactiveConfigR18QoeMeasurementTypeR18Constraints); err != nil {
				return err
			}
		}
	}

	// 8. qoe-AreaScope-r18: ref
	{
		if ie.Qoe_AreaScope_r18 != nil {
			if err := ie.Qoe_AreaScope_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 9. mce-Id-r18: octet-string
	{
		if ie.Mce_Id_r18 != nil {
			if err := e.EncodeOctetString(ie.Mce_Id_r18, appLayerIdleInactiveConfigR18MceIdR18Constraints); err != nil {
				return err
			}
		}
	}

	// 10. availableRAN-VisibleMetrics-r18: ref
	{
		if ie.AvailableRAN_VisibleMetrics_r18 != nil {
			if err := ie.AvailableRAN_VisibleMetrics_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *AppLayerIdleInactiveConfig_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(appLayerIdleInactiveConfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. measConfigAppLayerId-r18: ref
	{
		if err := ie.MeasConfigAppLayerId_r18.Decode(d); err != nil {
			return err
		}
	}

	// 4. serviceType-r18: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(appLayerIdleInactiveConfigR18ServiceTypeR18Constraints)
			if err != nil {
				return err
			}
			ie.ServiceType_r18 = &idx
		}
	}

	// 5. appLayerMeasPriority-r18: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(appLayerIdleInactiveConfigR18AppLayerMeasPriorityR18Constraints)
			if err != nil {
				return err
			}
			ie.AppLayerMeasPriority_r18 = &v
		}
	}

	// 6. qoe-Reference-r18: octet-string
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeOctetString(appLayerIdleInactiveConfigR18QoeReferenceR18Constraints)
			if err != nil {
				return err
			}
			ie.Qoe_Reference_r18 = v
		}
	}

	// 7. qoe-MeasurementType-r18: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(appLayerIdleInactiveConfigR18QoeMeasurementTypeR18Constraints)
			if err != nil {
				return err
			}
			ie.Qoe_MeasurementType_r18 = &idx
		}
	}

	// 8. qoe-AreaScope-r18: ref
	{
		if seq.IsComponentPresent(5) {
			ie.Qoe_AreaScope_r18 = new(Qoe_AreaScope_r18)
			if err := ie.Qoe_AreaScope_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 9. mce-Id-r18: octet-string
	{
		if seq.IsComponentPresent(6) {
			v, err := d.DecodeOctetString(appLayerIdleInactiveConfigR18MceIdR18Constraints)
			if err != nil {
				return err
			}
			ie.Mce_Id_r18 = v
		}
	}

	// 10. availableRAN-VisibleMetrics-r18: ref
	{
		if seq.IsComponentPresent(7) {
			ie.AvailableRAN_VisibleMetrics_r18 = new(AvailableRAN_VisibleMetrics_r18)
			if err := ie.AvailableRAN_VisibleMetrics_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}

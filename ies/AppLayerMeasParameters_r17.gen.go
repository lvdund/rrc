// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: AppLayerMeasParameters-r17 (line 16522).

var appLayerMeasParametersR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "qoe-Streaming-MeasReport-r17", Optional: true},
		{Name: "qoe-MTSI-MeasReport-r17", Optional: true},
		{Name: "qoe-VR-MeasReport-r17", Optional: true},
		{Name: "ran-VisibleQoE-Streaming-MeasReport-r17", Optional: true},
		{Name: "ran-VisibleQoE-VR-MeasReport-r17", Optional: true},
		{Name: "ul-MeasurementReportAppLayer-Seg-r17", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

const (
	AppLayerMeasParameters_r17_Qoe_Streaming_MeasReport_r17_Supported = 0
)

var appLayerMeasParametersR17QoeStreamingMeasReportR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AppLayerMeasParameters_r17_Qoe_Streaming_MeasReport_r17_Supported},
}

const (
	AppLayerMeasParameters_r17_Qoe_MTSI_MeasReport_r17_Supported = 0
)

var appLayerMeasParametersR17QoeMTSIMeasReportR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AppLayerMeasParameters_r17_Qoe_MTSI_MeasReport_r17_Supported},
}

const (
	AppLayerMeasParameters_r17_Qoe_VR_MeasReport_r17_Supported = 0
)

var appLayerMeasParametersR17QoeVRMeasReportR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AppLayerMeasParameters_r17_Qoe_VR_MeasReport_r17_Supported},
}

const (
	AppLayerMeasParameters_r17_Ran_VisibleQoE_Streaming_MeasReport_r17_Supported = 0
)

var appLayerMeasParametersR17RanVisibleQoEStreamingMeasReportR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AppLayerMeasParameters_r17_Ran_VisibleQoE_Streaming_MeasReport_r17_Supported},
}

const (
	AppLayerMeasParameters_r17_Ran_VisibleQoE_VR_MeasReport_r17_Supported = 0
)

var appLayerMeasParametersR17RanVisibleQoEVRMeasReportR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AppLayerMeasParameters_r17_Ran_VisibleQoE_VR_MeasReport_r17_Supported},
}

const (
	AppLayerMeasParameters_r17_Ul_MeasurementReportAppLayer_Seg_r17_Supported = 0
)

var appLayerMeasParametersR17UlMeasurementReportAppLayerSegR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AppLayerMeasParameters_r17_Ul_MeasurementReportAppLayer_Seg_r17_Supported},
}

const (
	AppLayerMeasParameters_r17_Ext_Qoe_IdleInactiveMeasReport_r18_Supported = 0
)

var appLayerMeasParametersR17ExtQoeIdleInactiveMeasReportR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AppLayerMeasParameters_r17_Ext_Qoe_IdleInactiveMeasReport_r18_Supported},
}

const (
	AppLayerMeasParameters_r17_Ext_Qoe_NRDC_MeasReport_r18_Supported = 0
)

var appLayerMeasParametersR17ExtQoeNRDCMeasReportR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AppLayerMeasParameters_r17_Ext_Qoe_NRDC_MeasReport_r18_Supported},
}

const (
	AppLayerMeasParameters_r17_Ext_Qoe_AdditionalMemoryMeasReport_r18_KB128  = 0
	AppLayerMeasParameters_r17_Ext_Qoe_AdditionalMemoryMeasReport_r18_KB256  = 1
	AppLayerMeasParameters_r17_Ext_Qoe_AdditionalMemoryMeasReport_r18_KB512  = 2
	AppLayerMeasParameters_r17_Ext_Qoe_AdditionalMemoryMeasReport_r18_KB1024 = 3
)

var appLayerMeasParametersR17ExtQoeAdditionalMemoryMeasReportR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AppLayerMeasParameters_r17_Ext_Qoe_AdditionalMemoryMeasReport_r18_KB128, AppLayerMeasParameters_r17_Ext_Qoe_AdditionalMemoryMeasReport_r18_KB256, AppLayerMeasParameters_r17_Ext_Qoe_AdditionalMemoryMeasReport_r18_KB512, AppLayerMeasParameters_r17_Ext_Qoe_AdditionalMemoryMeasReport_r18_KB1024},
}

const (
	AppLayerMeasParameters_r17_Ext_Qoe_PriorityBasedDiscarding_r18_Supported = 0
)

var appLayerMeasParametersR17ExtQoePriorityBasedDiscardingR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AppLayerMeasParameters_r17_Ext_Qoe_PriorityBasedDiscarding_r18_Supported},
}

const (
	AppLayerMeasParameters_r17_Ext_Srb5_r18_Supported = 0
)

var appLayerMeasParametersR17ExtSrb5R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AppLayerMeasParameters_r17_Ext_Srb5_r18_Supported},
}

type AppLayerMeasParameters_r17 struct {
	Qoe_Streaming_MeasReport_r17            *int64
	Qoe_MTSI_MeasReport_r17                 *int64
	Qoe_VR_MeasReport_r17                   *int64
	Ran_VisibleQoE_Streaming_MeasReport_r17 *int64
	Ran_VisibleQoE_VR_MeasReport_r17        *int64
	Ul_MeasurementReportAppLayer_Seg_r17    *int64
	Qoe_IdleInactiveMeasReport_r18          *int64
	Qoe_NRDC_MeasReport_r18                 *int64
	Qoe_AdditionalMemoryMeasReport_r18      *int64
	Qoe_PriorityBasedDiscarding_r18         *int64
	Srb5_r18                                *int64
}

func (ie *AppLayerMeasParameters_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(appLayerMeasParametersR17Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Qoe_IdleInactiveMeasReport_r18 != nil || ie.Qoe_NRDC_MeasReport_r18 != nil || ie.Qoe_AdditionalMemoryMeasReport_r18 != nil || ie.Qoe_PriorityBasedDiscarding_r18 != nil || ie.Srb5_r18 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Qoe_Streaming_MeasReport_r17 != nil, ie.Qoe_MTSI_MeasReport_r17 != nil, ie.Qoe_VR_MeasReport_r17 != nil, ie.Ran_VisibleQoE_Streaming_MeasReport_r17 != nil, ie.Ran_VisibleQoE_VR_MeasReport_r17 != nil, ie.Ul_MeasurementReportAppLayer_Seg_r17 != nil}); err != nil {
		return err
	}

	// 3. qoe-Streaming-MeasReport-r17: enumerated
	{
		if ie.Qoe_Streaming_MeasReport_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Qoe_Streaming_MeasReport_r17, appLayerMeasParametersR17QoeStreamingMeasReportR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. qoe-MTSI-MeasReport-r17: enumerated
	{
		if ie.Qoe_MTSI_MeasReport_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Qoe_MTSI_MeasReport_r17, appLayerMeasParametersR17QoeMTSIMeasReportR17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. qoe-VR-MeasReport-r17: enumerated
	{
		if ie.Qoe_VR_MeasReport_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Qoe_VR_MeasReport_r17, appLayerMeasParametersR17QoeVRMeasReportR17Constraints); err != nil {
				return err
			}
		}
	}

	// 6. ran-VisibleQoE-Streaming-MeasReport-r17: enumerated
	{
		if ie.Ran_VisibleQoE_Streaming_MeasReport_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Ran_VisibleQoE_Streaming_MeasReport_r17, appLayerMeasParametersR17RanVisibleQoEStreamingMeasReportR17Constraints); err != nil {
				return err
			}
		}
	}

	// 7. ran-VisibleQoE-VR-MeasReport-r17: enumerated
	{
		if ie.Ran_VisibleQoE_VR_MeasReport_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Ran_VisibleQoE_VR_MeasReport_r17, appLayerMeasParametersR17RanVisibleQoEVRMeasReportR17Constraints); err != nil {
				return err
			}
		}
	}

	// 8. ul-MeasurementReportAppLayer-Seg-r17: enumerated
	{
		if ie.Ul_MeasurementReportAppLayer_Seg_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Ul_MeasurementReportAppLayer_Seg_r17, appLayerMeasParametersR17UlMeasurementReportAppLayerSegR17Constraints); err != nil {
				return err
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
					{Name: "qoe-IdleInactiveMeasReport-r18", Optional: true},
					{Name: "qoe-NRDC-MeasReport-r18", Optional: true},
					{Name: "qoe-AdditionalMemoryMeasReport-r18", Optional: true},
					{Name: "qoe-PriorityBasedDiscarding-r18", Optional: true},
					{Name: "srb5-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Qoe_IdleInactiveMeasReport_r18 != nil, ie.Qoe_NRDC_MeasReport_r18 != nil, ie.Qoe_AdditionalMemoryMeasReport_r18 != nil, ie.Qoe_PriorityBasedDiscarding_r18 != nil, ie.Srb5_r18 != nil}); err != nil {
				return err
			}

			if ie.Qoe_IdleInactiveMeasReport_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Qoe_IdleInactiveMeasReport_r18, appLayerMeasParametersR17ExtQoeIdleInactiveMeasReportR18Constraints); err != nil {
					return err
				}
			}

			if ie.Qoe_NRDC_MeasReport_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Qoe_NRDC_MeasReport_r18, appLayerMeasParametersR17ExtQoeNRDCMeasReportR18Constraints); err != nil {
					return err
				}
			}

			if ie.Qoe_AdditionalMemoryMeasReport_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Qoe_AdditionalMemoryMeasReport_r18, appLayerMeasParametersR17ExtQoeAdditionalMemoryMeasReportR18Constraints); err != nil {
					return err
				}
			}

			if ie.Qoe_PriorityBasedDiscarding_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Qoe_PriorityBasedDiscarding_r18, appLayerMeasParametersR17ExtQoePriorityBasedDiscardingR18Constraints); err != nil {
					return err
				}
			}

			if ie.Srb5_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Srb5_r18, appLayerMeasParametersR17ExtSrb5R18Constraints); err != nil {
					return err
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

func (ie *AppLayerMeasParameters_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(appLayerMeasParametersR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. qoe-Streaming-MeasReport-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(appLayerMeasParametersR17QoeStreamingMeasReportR17Constraints)
			if err != nil {
				return err
			}
			ie.Qoe_Streaming_MeasReport_r17 = &idx
		}
	}

	// 4. qoe-MTSI-MeasReport-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(appLayerMeasParametersR17QoeMTSIMeasReportR17Constraints)
			if err != nil {
				return err
			}
			ie.Qoe_MTSI_MeasReport_r17 = &idx
		}
	}

	// 5. qoe-VR-MeasReport-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(appLayerMeasParametersR17QoeVRMeasReportR17Constraints)
			if err != nil {
				return err
			}
			ie.Qoe_VR_MeasReport_r17 = &idx
		}
	}

	// 6. ran-VisibleQoE-Streaming-MeasReport-r17: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(appLayerMeasParametersR17RanVisibleQoEStreamingMeasReportR17Constraints)
			if err != nil {
				return err
			}
			ie.Ran_VisibleQoE_Streaming_MeasReport_r17 = &idx
		}
	}

	// 7. ran-VisibleQoE-VR-MeasReport-r17: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(appLayerMeasParametersR17RanVisibleQoEVRMeasReportR17Constraints)
			if err != nil {
				return err
			}
			ie.Ran_VisibleQoE_VR_MeasReport_r17 = &idx
		}
	}

	// 8. ul-MeasurementReportAppLayer-Seg-r17: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(appLayerMeasParametersR17UlMeasurementReportAppLayerSegR17Constraints)
			if err != nil {
				return err
			}
			ie.Ul_MeasurementReportAppLayer_Seg_r17 = &idx
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
				{Name: "qoe-IdleInactiveMeasReport-r18", Optional: true},
				{Name: "qoe-NRDC-MeasReport-r18", Optional: true},
				{Name: "qoe-AdditionalMemoryMeasReport-r18", Optional: true},
				{Name: "qoe-PriorityBasedDiscarding-r18", Optional: true},
				{Name: "srb5-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(appLayerMeasParametersR17ExtQoeIdleInactiveMeasReportR18Constraints)
			if err != nil {
				return err
			}
			ie.Qoe_IdleInactiveMeasReport_r18 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(appLayerMeasParametersR17ExtQoeNRDCMeasReportR18Constraints)
			if err != nil {
				return err
			}
			ie.Qoe_NRDC_MeasReport_r18 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(appLayerMeasParametersR17ExtQoeAdditionalMemoryMeasReportR18Constraints)
			if err != nil {
				return err
			}
			ie.Qoe_AdditionalMemoryMeasReport_r18 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(appLayerMeasParametersR17ExtQoePriorityBasedDiscardingR18Constraints)
			if err != nil {
				return err
			}
			ie.Qoe_PriorityBasedDiscarding_r18 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(appLayerMeasParametersR17ExtSrb5R18Constraints)
			if err != nil {
				return err
			}
			ie.Srb5_r18 = &v
		}
	}

	return nil
}

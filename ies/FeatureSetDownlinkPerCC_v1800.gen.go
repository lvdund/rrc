// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FeatureSetDownlinkPerCC-v1800 (line 19931).

var featureSetDownlinkPerCCV1800Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "multiDCI-IntraCellMultiTRP-TwoTA-r18", Optional: true},
		{Name: "multiDCI-InterCellMultiTRP-TwoTA-r18", Optional: true},
		{Name: "rxTimingDiff-r18", Optional: true},
		{Name: "multiDCI-MultiTRP-CORESET-Monitoring-r18", Optional: true},
		{Name: "broadcastNonServingCell-r18", Optional: true},
		{Name: "schedulingMeasurementRelaxation-r18", Optional: true},
	},
}

const (
	FeatureSetDownlinkPerCC_v1800_MultiDCI_IntraCellMultiTRP_TwoTA_r18_Supported = 0
)

var featureSetDownlinkPerCCV1800MultiDCIIntraCellMultiTRPTwoTAR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlinkPerCC_v1800_MultiDCI_IntraCellMultiTRP_TwoTA_r18_Supported},
}

var featureSetDownlinkPerCCV1800MultiDCIInterCellMultiTRPTwoTAR18Constraints = per.Constrained(1, 2)

const (
	FeatureSetDownlinkPerCC_v1800_RxTimingDiff_r18_Supported = 0
)

var featureSetDownlinkPerCCV1800RxTimingDiffR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlinkPerCC_v1800_RxTimingDiff_r18_Supported},
}

const (
	FeatureSetDownlinkPerCC_v1800_MultiDCI_MultiTRP_CORESET_Monitoring_r18_Supported = 0
)

var featureSetDownlinkPerCCV1800MultiDCIMultiTRPCORESETMonitoringR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlinkPerCC_v1800_MultiDCI_MultiTRP_CORESET_Monitoring_r18_Supported},
}

const (
	FeatureSetDownlinkPerCC_v1800_BroadcastNonServingCell_r18_Supported = 0
)

var featureSetDownlinkPerCCV1800BroadcastNonServingCellR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlinkPerCC_v1800_BroadcastNonServingCell_r18_Supported},
}

const (
	FeatureSetDownlinkPerCC_v1800_SchedulingMeasurementRelaxation_r18_Supported = 0
)

var featureSetDownlinkPerCCV1800SchedulingMeasurementRelaxationR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlinkPerCC_v1800_SchedulingMeasurementRelaxation_r18_Supported},
}

type FeatureSetDownlinkPerCC_v1800 struct {
	MultiDCI_IntraCellMultiTRP_TwoTA_r18     *int64
	MultiDCI_InterCellMultiTRP_TwoTA_r18     *int64
	RxTimingDiff_r18                         *int64
	MultiDCI_MultiTRP_CORESET_Monitoring_r18 *int64
	BroadcastNonServingCell_r18              *int64
	SchedulingMeasurementRelaxation_r18      *int64
}

func (ie *FeatureSetDownlinkPerCC_v1800) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetDownlinkPerCCV1800Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MultiDCI_IntraCellMultiTRP_TwoTA_r18 != nil, ie.MultiDCI_InterCellMultiTRP_TwoTA_r18 != nil, ie.RxTimingDiff_r18 != nil, ie.MultiDCI_MultiTRP_CORESET_Monitoring_r18 != nil, ie.BroadcastNonServingCell_r18 != nil, ie.SchedulingMeasurementRelaxation_r18 != nil}); err != nil {
		return err
	}

	// 2. multiDCI-IntraCellMultiTRP-TwoTA-r18: enumerated
	{
		if ie.MultiDCI_IntraCellMultiTRP_TwoTA_r18 != nil {
			if err := e.EncodeEnumerated(*ie.MultiDCI_IntraCellMultiTRP_TwoTA_r18, featureSetDownlinkPerCCV1800MultiDCIIntraCellMultiTRPTwoTAR18Constraints); err != nil {
				return err
			}
		}
	}

	// 3. multiDCI-InterCellMultiTRP-TwoTA-r18: integer
	{
		if ie.MultiDCI_InterCellMultiTRP_TwoTA_r18 != nil {
			if err := e.EncodeInteger(*ie.MultiDCI_InterCellMultiTRP_TwoTA_r18, featureSetDownlinkPerCCV1800MultiDCIInterCellMultiTRPTwoTAR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. rxTimingDiff-r18: enumerated
	{
		if ie.RxTimingDiff_r18 != nil {
			if err := e.EncodeEnumerated(*ie.RxTimingDiff_r18, featureSetDownlinkPerCCV1800RxTimingDiffR18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. multiDCI-MultiTRP-CORESET-Monitoring-r18: enumerated
	{
		if ie.MultiDCI_MultiTRP_CORESET_Monitoring_r18 != nil {
			if err := e.EncodeEnumerated(*ie.MultiDCI_MultiTRP_CORESET_Monitoring_r18, featureSetDownlinkPerCCV1800MultiDCIMultiTRPCORESETMonitoringR18Constraints); err != nil {
				return err
			}
		}
	}

	// 6. broadcastNonServingCell-r18: enumerated
	{
		if ie.BroadcastNonServingCell_r18 != nil {
			if err := e.EncodeEnumerated(*ie.BroadcastNonServingCell_r18, featureSetDownlinkPerCCV1800BroadcastNonServingCellR18Constraints); err != nil {
				return err
			}
		}
	}

	// 7. schedulingMeasurementRelaxation-r18: enumerated
	{
		if ie.SchedulingMeasurementRelaxation_r18 != nil {
			if err := e.EncodeEnumerated(*ie.SchedulingMeasurementRelaxation_r18, featureSetDownlinkPerCCV1800SchedulingMeasurementRelaxationR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FeatureSetDownlinkPerCC_v1800) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetDownlinkPerCCV1800Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. multiDCI-IntraCellMultiTRP-TwoTA-r18: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkPerCCV1800MultiDCIIntraCellMultiTRPTwoTAR18Constraints)
			if err != nil {
				return err
			}
			ie.MultiDCI_IntraCellMultiTRP_TwoTA_r18 = &idx
		}
	}

	// 3. multiDCI-InterCellMultiTRP-TwoTA-r18: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(featureSetDownlinkPerCCV1800MultiDCIInterCellMultiTRPTwoTAR18Constraints)
			if err != nil {
				return err
			}
			ie.MultiDCI_InterCellMultiTRP_TwoTA_r18 = &v
		}
	}

	// 4. rxTimingDiff-r18: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkPerCCV1800RxTimingDiffR18Constraints)
			if err != nil {
				return err
			}
			ie.RxTimingDiff_r18 = &idx
		}
	}

	// 5. multiDCI-MultiTRP-CORESET-Monitoring-r18: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkPerCCV1800MultiDCIMultiTRPCORESETMonitoringR18Constraints)
			if err != nil {
				return err
			}
			ie.MultiDCI_MultiTRP_CORESET_Monitoring_r18 = &idx
		}
	}

	// 6. broadcastNonServingCell-r18: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkPerCCV1800BroadcastNonServingCellR18Constraints)
			if err != nil {
				return err
			}
			ie.BroadcastNonServingCell_r18 = &idx
		}
	}

	// 7. schedulingMeasurementRelaxation-r18: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkPerCCV1800SchedulingMeasurementRelaxationR18Constraints)
			if err != nil {
				return err
			}
			ie.SchedulingMeasurementRelaxation_r18 = &idx
		}
	}

	return nil
}

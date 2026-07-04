// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FeatureSetUplink-v1630 (line 20223).

var featureSetUplinkV1630Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "offsetSRS-CB-PUSCH-Ant-Switch-fr1-r16", Optional: true},
		{Name: "offsetSRS-CB-PUSCH-PDCCH-MonitorSingleOcc-fr1-r16", Optional: true},
		{Name: "offsetSRS-CB-PUSCH-PDCCH-MonitorAnyOccWithoutGap-fr1-r16", Optional: true},
		{Name: "offsetSRS-CB-PUSCH-PDCCH-MonitorAnyOccWithGap-fr1-r16", Optional: true},
		{Name: "dummy", Optional: true},
		{Name: "partialCancellationPUCCH-PUSCH-PRACH-TX-r16", Optional: true},
	},
}

const (
	FeatureSetUplink_v1630_OffsetSRS_CB_PUSCH_Ant_Switch_Fr1_r16_Supported = 0
)

var featureSetUplinkV1630OffsetSRSCBPUSCHAntSwitchFr1R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1630_OffsetSRS_CB_PUSCH_Ant_Switch_Fr1_r16_Supported},
}

const (
	FeatureSetUplink_v1630_OffsetSRS_CB_PUSCH_PDCCH_MonitorSingleOcc_Fr1_r16_Supported = 0
)

var featureSetUplinkV1630OffsetSRSCBPUSCHPDCCHMonitorSingleOccFr1R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1630_OffsetSRS_CB_PUSCH_PDCCH_MonitorSingleOcc_Fr1_r16_Supported},
}

const (
	FeatureSetUplink_v1630_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithoutGap_Fr1_r16_Supported = 0
)

var featureSetUplinkV1630OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithoutGapFr1R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1630_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithoutGap_Fr1_r16_Supported},
}

const (
	FeatureSetUplink_v1630_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithGap_Fr1_r16_Supported = 0
)

var featureSetUplinkV1630OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithGapFr1R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1630_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithGap_Fr1_r16_Supported},
}

const (
	FeatureSetUplink_v1630_Dummy_Supported = 0
)

var featureSetUplinkV1630DummyConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1630_Dummy_Supported},
}

const (
	FeatureSetUplink_v1630_PartialCancellationPUCCH_PUSCH_PRACH_TX_r16_Supported = 0
)

var featureSetUplinkV1630PartialCancellationPUCCHPUSCHPRACHTXR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1630_PartialCancellationPUCCH_PUSCH_PRACH_TX_r16_Supported},
}

type FeatureSetUplink_v1630 struct {
	OffsetSRS_CB_PUSCH_Ant_Switch_Fr1_r16                    *int64
	OffsetSRS_CB_PUSCH_PDCCH_MonitorSingleOcc_Fr1_r16        *int64
	OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithoutGap_Fr1_r16 *int64
	OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithGap_Fr1_r16    *int64
	Dummy                                                    *int64
	PartialCancellationPUCCH_PUSCH_PRACH_TX_r16              *int64
}

func (ie *FeatureSetUplink_v1630) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetUplinkV1630Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.OffsetSRS_CB_PUSCH_Ant_Switch_Fr1_r16 != nil, ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorSingleOcc_Fr1_r16 != nil, ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithoutGap_Fr1_r16 != nil, ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithGap_Fr1_r16 != nil, ie.Dummy != nil, ie.PartialCancellationPUCCH_PUSCH_PRACH_TX_r16 != nil}); err != nil {
		return err
	}

	// 2. offsetSRS-CB-PUSCH-Ant-Switch-fr1-r16: enumerated
	{
		if ie.OffsetSRS_CB_PUSCH_Ant_Switch_Fr1_r16 != nil {
			if err := e.EncodeEnumerated(*ie.OffsetSRS_CB_PUSCH_Ant_Switch_Fr1_r16, featureSetUplinkV1630OffsetSRSCBPUSCHAntSwitchFr1R16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. offsetSRS-CB-PUSCH-PDCCH-MonitorSingleOcc-fr1-r16: enumerated
	{
		if ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorSingleOcc_Fr1_r16 != nil {
			if err := e.EncodeEnumerated(*ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorSingleOcc_Fr1_r16, featureSetUplinkV1630OffsetSRSCBPUSCHPDCCHMonitorSingleOccFr1R16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. offsetSRS-CB-PUSCH-PDCCH-MonitorAnyOccWithoutGap-fr1-r16: enumerated
	{
		if ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithoutGap_Fr1_r16 != nil {
			if err := e.EncodeEnumerated(*ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithoutGap_Fr1_r16, featureSetUplinkV1630OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithoutGapFr1R16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. offsetSRS-CB-PUSCH-PDCCH-MonitorAnyOccWithGap-fr1-r16: enumerated
	{
		if ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithGap_Fr1_r16 != nil {
			if err := e.EncodeEnumerated(*ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithGap_Fr1_r16, featureSetUplinkV1630OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithGapFr1R16Constraints); err != nil {
				return err
			}
		}
	}

	// 6. dummy: enumerated
	{
		if ie.Dummy != nil {
			if err := e.EncodeEnumerated(*ie.Dummy, featureSetUplinkV1630DummyConstraints); err != nil {
				return err
			}
		}
	}

	// 7. partialCancellationPUCCH-PUSCH-PRACH-TX-r16: enumerated
	{
		if ie.PartialCancellationPUCCH_PUSCH_PRACH_TX_r16 != nil {
			if err := e.EncodeEnumerated(*ie.PartialCancellationPUCCH_PUSCH_PRACH_TX_r16, featureSetUplinkV1630PartialCancellationPUCCHPUSCHPRACHTXR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FeatureSetUplink_v1630) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetUplinkV1630Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. offsetSRS-CB-PUSCH-Ant-Switch-fr1-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1630OffsetSRSCBPUSCHAntSwitchFr1R16Constraints)
			if err != nil {
				return err
			}
			ie.OffsetSRS_CB_PUSCH_Ant_Switch_Fr1_r16 = &idx
		}
	}

	// 3. offsetSRS-CB-PUSCH-PDCCH-MonitorSingleOcc-fr1-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1630OffsetSRSCBPUSCHPDCCHMonitorSingleOccFr1R16Constraints)
			if err != nil {
				return err
			}
			ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorSingleOcc_Fr1_r16 = &idx
		}
	}

	// 4. offsetSRS-CB-PUSCH-PDCCH-MonitorAnyOccWithoutGap-fr1-r16: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1630OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithoutGapFr1R16Constraints)
			if err != nil {
				return err
			}
			ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithoutGap_Fr1_r16 = &idx
		}
	}

	// 5. offsetSRS-CB-PUSCH-PDCCH-MonitorAnyOccWithGap-fr1-r16: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1630OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithGapFr1R16Constraints)
			if err != nil {
				return err
			}
			ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithGap_Fr1_r16 = &idx
		}
	}

	// 6. dummy: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1630DummyConstraints)
			if err != nil {
				return err
			}
			ie.Dummy = &idx
		}
	}

	// 7. partialCancellationPUCCH-PUSCH-PRACH-TX-r16: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1630PartialCancellationPUCCHPUSCHPRACHTXR16Constraints)
			if err != nil {
				return err
			}
			ie.PartialCancellationPUCCH_PUSCH_PRACH_TX_r16 = &idx
		}
	}

	return nil
}

// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FeatureSetUplink-v1710 (line 20263).

var featureSetUplinkV1710Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "mTRP-PUSCH-TypeA-CB-r17", Optional: true},
		{Name: "mTRP-PUSCH-RepetitionTypeA-r17", Optional: true},
		{Name: "mTRP-PUCCH-IntraSlot-r17", Optional: true},
		{Name: "srs-AntennaSwitching2SP-1Periodic-r17", Optional: true},
		{Name: "srs-ExtensionAperiodicSRS-r17", Optional: true},
		{Name: "srs-OneAP-SRS-r17", Optional: true},
		{Name: "ue-PowerClassPerBandPerBC-r17", Optional: true},
		{Name: "tx-Support-UL-GapFR2-r17", Optional: true},
	},
}

const (
	FeatureSetUplink_v1710_MTRP_PUSCH_TypeA_CB_r17_N1 = 0
	FeatureSetUplink_v1710_MTRP_PUSCH_TypeA_CB_r17_N2 = 1
	FeatureSetUplink_v1710_MTRP_PUSCH_TypeA_CB_r17_N4 = 2
)

var featureSetUplinkV1710MTRPPUSCHTypeACBR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1710_MTRP_PUSCH_TypeA_CB_r17_N1, FeatureSetUplink_v1710_MTRP_PUSCH_TypeA_CB_r17_N2, FeatureSetUplink_v1710_MTRP_PUSCH_TypeA_CB_r17_N4},
}

const (
	FeatureSetUplink_v1710_MTRP_PUSCH_RepetitionTypeA_r17_N1 = 0
	FeatureSetUplink_v1710_MTRP_PUSCH_RepetitionTypeA_r17_N2 = 1
	FeatureSetUplink_v1710_MTRP_PUSCH_RepetitionTypeA_r17_N3 = 2
	FeatureSetUplink_v1710_MTRP_PUSCH_RepetitionTypeA_r17_N4 = 3
)

var featureSetUplinkV1710MTRPPUSCHRepetitionTypeAR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1710_MTRP_PUSCH_RepetitionTypeA_r17_N1, FeatureSetUplink_v1710_MTRP_PUSCH_RepetitionTypeA_r17_N2, FeatureSetUplink_v1710_MTRP_PUSCH_RepetitionTypeA_r17_N3, FeatureSetUplink_v1710_MTRP_PUSCH_RepetitionTypeA_r17_N4},
}

const (
	FeatureSetUplink_v1710_MTRP_PUCCH_IntraSlot_r17_Pf0_2   = 0
	FeatureSetUplink_v1710_MTRP_PUCCH_IntraSlot_r17_Pf1_3_4 = 1
	FeatureSetUplink_v1710_MTRP_PUCCH_IntraSlot_r17_Pf0_4   = 2
)

var featureSetUplinkV1710MTRPPUCCHIntraSlotR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1710_MTRP_PUCCH_IntraSlot_r17_Pf0_2, FeatureSetUplink_v1710_MTRP_PUCCH_IntraSlot_r17_Pf1_3_4, FeatureSetUplink_v1710_MTRP_PUCCH_IntraSlot_r17_Pf0_4},
}

const (
	FeatureSetUplink_v1710_Srs_AntennaSwitching2SP_1Periodic_r17_Supported = 0
)

var featureSetUplinkV1710SrsAntennaSwitching2SP1PeriodicR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1710_Srs_AntennaSwitching2SP_1Periodic_r17_Supported},
}

const (
	FeatureSetUplink_v1710_Srs_ExtensionAperiodicSRS_r17_Supported = 0
)

var featureSetUplinkV1710SrsExtensionAperiodicSRSR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1710_Srs_ExtensionAperiodicSRS_r17_Supported},
}

const (
	FeatureSetUplink_v1710_Srs_OneAP_SRS_r17_Supported = 0
)

var featureSetUplinkV1710SrsOneAPSRSR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1710_Srs_OneAP_SRS_r17_Supported},
}

const (
	FeatureSetUplink_v1710_Ue_PowerClassPerBandPerBC_r17_Pc1dot5 = 0
	FeatureSetUplink_v1710_Ue_PowerClassPerBandPerBC_r17_Pc2     = 1
	FeatureSetUplink_v1710_Ue_PowerClassPerBandPerBC_r17_Pc3     = 2
)

var featureSetUplinkV1710UePowerClassPerBandPerBCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1710_Ue_PowerClassPerBandPerBC_r17_Pc1dot5, FeatureSetUplink_v1710_Ue_PowerClassPerBandPerBC_r17_Pc2, FeatureSetUplink_v1710_Ue_PowerClassPerBandPerBC_r17_Pc3},
}

const (
	FeatureSetUplink_v1710_Tx_Support_UL_GapFR2_r17_Supported = 0
)

var featureSetUplinkV1710TxSupportULGapFR2R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1710_Tx_Support_UL_GapFR2_r17_Supported},
}

type FeatureSetUplink_v1710 struct {
	MTRP_PUSCH_TypeA_CB_r17               *int64
	MTRP_PUSCH_RepetitionTypeA_r17        *int64
	MTRP_PUCCH_IntraSlot_r17              *int64
	Srs_AntennaSwitching2SP_1Periodic_r17 *int64
	Srs_ExtensionAperiodicSRS_r17         *int64
	Srs_OneAP_SRS_r17                     *int64
	Ue_PowerClassPerBandPerBC_r17         *int64
	Tx_Support_UL_GapFR2_r17              *int64
}

func (ie *FeatureSetUplink_v1710) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetUplinkV1710Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MTRP_PUSCH_TypeA_CB_r17 != nil, ie.MTRP_PUSCH_RepetitionTypeA_r17 != nil, ie.MTRP_PUCCH_IntraSlot_r17 != nil, ie.Srs_AntennaSwitching2SP_1Periodic_r17 != nil, ie.Srs_ExtensionAperiodicSRS_r17 != nil, ie.Srs_OneAP_SRS_r17 != nil, ie.Ue_PowerClassPerBandPerBC_r17 != nil, ie.Tx_Support_UL_GapFR2_r17 != nil}); err != nil {
		return err
	}

	// 2. mTRP-PUSCH-TypeA-CB-r17: enumerated
	{
		if ie.MTRP_PUSCH_TypeA_CB_r17 != nil {
			if err := e.EncodeEnumerated(*ie.MTRP_PUSCH_TypeA_CB_r17, featureSetUplinkV1710MTRPPUSCHTypeACBR17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. mTRP-PUSCH-RepetitionTypeA-r17: enumerated
	{
		if ie.MTRP_PUSCH_RepetitionTypeA_r17 != nil {
			if err := e.EncodeEnumerated(*ie.MTRP_PUSCH_RepetitionTypeA_r17, featureSetUplinkV1710MTRPPUSCHRepetitionTypeAR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. mTRP-PUCCH-IntraSlot-r17: enumerated
	{
		if ie.MTRP_PUCCH_IntraSlot_r17 != nil {
			if err := e.EncodeEnumerated(*ie.MTRP_PUCCH_IntraSlot_r17, featureSetUplinkV1710MTRPPUCCHIntraSlotR17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. srs-AntennaSwitching2SP-1Periodic-r17: enumerated
	{
		if ie.Srs_AntennaSwitching2SP_1Periodic_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Srs_AntennaSwitching2SP_1Periodic_r17, featureSetUplinkV1710SrsAntennaSwitching2SP1PeriodicR17Constraints); err != nil {
				return err
			}
		}
	}

	// 6. srs-ExtensionAperiodicSRS-r17: enumerated
	{
		if ie.Srs_ExtensionAperiodicSRS_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Srs_ExtensionAperiodicSRS_r17, featureSetUplinkV1710SrsExtensionAperiodicSRSR17Constraints); err != nil {
				return err
			}
		}
	}

	// 7. srs-OneAP-SRS-r17: enumerated
	{
		if ie.Srs_OneAP_SRS_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Srs_OneAP_SRS_r17, featureSetUplinkV1710SrsOneAPSRSR17Constraints); err != nil {
				return err
			}
		}
	}

	// 8. ue-PowerClassPerBandPerBC-r17: enumerated
	{
		if ie.Ue_PowerClassPerBandPerBC_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Ue_PowerClassPerBandPerBC_r17, featureSetUplinkV1710UePowerClassPerBandPerBCR17Constraints); err != nil {
				return err
			}
		}
	}

	// 9. tx-Support-UL-GapFR2-r17: enumerated
	{
		if ie.Tx_Support_UL_GapFR2_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Tx_Support_UL_GapFR2_r17, featureSetUplinkV1710TxSupportULGapFR2R17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FeatureSetUplink_v1710) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetUplinkV1710Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. mTRP-PUSCH-TypeA-CB-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1710MTRPPUSCHTypeACBR17Constraints)
			if err != nil {
				return err
			}
			ie.MTRP_PUSCH_TypeA_CB_r17 = &idx
		}
	}

	// 3. mTRP-PUSCH-RepetitionTypeA-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1710MTRPPUSCHRepetitionTypeAR17Constraints)
			if err != nil {
				return err
			}
			ie.MTRP_PUSCH_RepetitionTypeA_r17 = &idx
		}
	}

	// 4. mTRP-PUCCH-IntraSlot-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1710MTRPPUCCHIntraSlotR17Constraints)
			if err != nil {
				return err
			}
			ie.MTRP_PUCCH_IntraSlot_r17 = &idx
		}
	}

	// 5. srs-AntennaSwitching2SP-1Periodic-r17: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1710SrsAntennaSwitching2SP1PeriodicR17Constraints)
			if err != nil {
				return err
			}
			ie.Srs_AntennaSwitching2SP_1Periodic_r17 = &idx
		}
	}

	// 6. srs-ExtensionAperiodicSRS-r17: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1710SrsExtensionAperiodicSRSR17Constraints)
			if err != nil {
				return err
			}
			ie.Srs_ExtensionAperiodicSRS_r17 = &idx
		}
	}

	// 7. srs-OneAP-SRS-r17: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1710SrsOneAPSRSR17Constraints)
			if err != nil {
				return err
			}
			ie.Srs_OneAP_SRS_r17 = &idx
		}
	}

	// 8. ue-PowerClassPerBandPerBC-r17: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1710UePowerClassPerBandPerBCR17Constraints)
			if err != nil {
				return err
			}
			ie.Ue_PowerClassPerBandPerBC_r17 = &idx
		}
	}

	// 9. tx-Support-UL-GapFR2-r17: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1710TxSupportULGapFR2R17Constraints)
			if err != nil {
				return err
			}
			ie.Tx_Support_UL_GapFR2_r17 = &idx
		}
	}

	return nil
}

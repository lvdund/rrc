// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CA-ParametersNR-v1640 (line 17390).

var cAParametersNRV1640Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "uplinkTxDC-TwoCarrierReport-r16", Optional: true},
		{Name: "maxUpTo3Diff-NumerologiesConfigSinglePUCCH-grp-r16", Optional: true},
		{Name: "maxUpTo4Diff-NumerologiesConfigSinglePUCCH-grp-r16", Optional: true},
		{Name: "twoPUCCH-Grp-ConfigurationsList-r16", Optional: true},
		{Name: "diffNumerologyAcrossPUCCH-Group-CarrierTypes-r16", Optional: true},
		{Name: "diffNumerologyWithinPUCCH-GroupSmallerSCS-CarrierTypes-r16", Optional: true},
		{Name: "diffNumerologyWithinPUCCH-GroupLargerSCS-CarrierTypes-r16", Optional: true},
		{Name: "pdcch-MonitoringCA-NonAlignedSpan-r16", Optional: true},
		{Name: "pdcch-BlindDetectionCA-Mixed-NonAlignedSpan-r16", Optional: true},
	},
}

const (
	CA_ParametersNR_v1640_UplinkTxDC_TwoCarrierReport_r16_Supported = 0
)

var cAParametersNRV1640UplinkTxDCTwoCarrierReportR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1640_UplinkTxDC_TwoCarrierReport_r16_Supported},
}

var cAParametersNRV1640TwoPUCCHGrpConfigurationsListR16Constraints = per.SizeRange(1, common.MaxTwoPUCCH_Grp_ConfigList_r16)

const (
	CA_ParametersNR_v1640_DiffNumerologyAcrossPUCCH_Group_CarrierTypes_r16_Supported = 0
)

var cAParametersNRV1640DiffNumerologyAcrossPUCCHGroupCarrierTypesR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1640_DiffNumerologyAcrossPUCCH_Group_CarrierTypes_r16_Supported},
}

const (
	CA_ParametersNR_v1640_DiffNumerologyWithinPUCCH_GroupSmallerSCS_CarrierTypes_r16_Supported = 0
)

var cAParametersNRV1640DiffNumerologyWithinPUCCHGroupSmallerSCSCarrierTypesR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1640_DiffNumerologyWithinPUCCH_GroupSmallerSCS_CarrierTypes_r16_Supported},
}

const (
	CA_ParametersNR_v1640_DiffNumerologyWithinPUCCH_GroupLargerSCS_CarrierTypes_r16_Supported = 0
)

var cAParametersNRV1640DiffNumerologyWithinPUCCHGroupLargerSCSCarrierTypesR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1640_DiffNumerologyWithinPUCCH_GroupLargerSCS_CarrierTypes_r16_Supported},
}

var cAParametersNRV1640PdcchMonitoringCANonAlignedSpanR16Constraints = per.Constrained(2, 16)

type CA_ParametersNR_v1640 struct {
	UplinkTxDC_TwoCarrierReport_r16                            *int64
	MaxUpTo3Diff_NumerologiesConfigSinglePUCCH_Grp_r16         *PUCCH_Grp_CarrierTypes_r16
	MaxUpTo4Diff_NumerologiesConfigSinglePUCCH_Grp_r16         *PUCCH_Grp_CarrierTypes_r16
	TwoPUCCH_Grp_ConfigurationsList_r16                        []TwoPUCCH_Grp_Configurations_r16
	DiffNumerologyAcrossPUCCH_Group_CarrierTypes_r16           *int64
	DiffNumerologyWithinPUCCH_GroupSmallerSCS_CarrierTypes_r16 *int64
	DiffNumerologyWithinPUCCH_GroupLargerSCS_CarrierTypes_r16  *int64
	Pdcch_MonitoringCA_NonAlignedSpan_r16                      *int64
	Pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_r16            *struct {
		Pdcch_BlindDetectionCA1_r16 int64
		Pdcch_BlindDetectionCA2_r16 int64
	}
}

func (ie *CA_ParametersNR_v1640) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cAParametersNRV1640Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.UplinkTxDC_TwoCarrierReport_r16 != nil, ie.MaxUpTo3Diff_NumerologiesConfigSinglePUCCH_Grp_r16 != nil, ie.MaxUpTo4Diff_NumerologiesConfigSinglePUCCH_Grp_r16 != nil, ie.TwoPUCCH_Grp_ConfigurationsList_r16 != nil, ie.DiffNumerologyAcrossPUCCH_Group_CarrierTypes_r16 != nil, ie.DiffNumerologyWithinPUCCH_GroupSmallerSCS_CarrierTypes_r16 != nil, ie.DiffNumerologyWithinPUCCH_GroupLargerSCS_CarrierTypes_r16 != nil, ie.Pdcch_MonitoringCA_NonAlignedSpan_r16 != nil, ie.Pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_r16 != nil}); err != nil {
		return err
	}

	// 2. uplinkTxDC-TwoCarrierReport-r16: enumerated
	{
		if ie.UplinkTxDC_TwoCarrierReport_r16 != nil {
			if err := e.EncodeEnumerated(*ie.UplinkTxDC_TwoCarrierReport_r16, cAParametersNRV1640UplinkTxDCTwoCarrierReportR16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. maxUpTo3Diff-NumerologiesConfigSinglePUCCH-grp-r16: ref
	{
		if ie.MaxUpTo3Diff_NumerologiesConfigSinglePUCCH_Grp_r16 != nil {
			if err := ie.MaxUpTo3Diff_NumerologiesConfigSinglePUCCH_Grp_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. maxUpTo4Diff-NumerologiesConfigSinglePUCCH-grp-r16: ref
	{
		if ie.MaxUpTo4Diff_NumerologiesConfigSinglePUCCH_Grp_r16 != nil {
			if err := ie.MaxUpTo4Diff_NumerologiesConfigSinglePUCCH_Grp_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. twoPUCCH-Grp-ConfigurationsList-r16: sequence-of
	{
		if ie.TwoPUCCH_Grp_ConfigurationsList_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(cAParametersNRV1640TwoPUCCHGrpConfigurationsListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.TwoPUCCH_Grp_ConfigurationsList_r16))); err != nil {
				return err
			}
			for i := range ie.TwoPUCCH_Grp_ConfigurationsList_r16 {
				if err := ie.TwoPUCCH_Grp_ConfigurationsList_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 6. diffNumerologyAcrossPUCCH-Group-CarrierTypes-r16: enumerated
	{
		if ie.DiffNumerologyAcrossPUCCH_Group_CarrierTypes_r16 != nil {
			if err := e.EncodeEnumerated(*ie.DiffNumerologyAcrossPUCCH_Group_CarrierTypes_r16, cAParametersNRV1640DiffNumerologyAcrossPUCCHGroupCarrierTypesR16Constraints); err != nil {
				return err
			}
		}
	}

	// 7. diffNumerologyWithinPUCCH-GroupSmallerSCS-CarrierTypes-r16: enumerated
	{
		if ie.DiffNumerologyWithinPUCCH_GroupSmallerSCS_CarrierTypes_r16 != nil {
			if err := e.EncodeEnumerated(*ie.DiffNumerologyWithinPUCCH_GroupSmallerSCS_CarrierTypes_r16, cAParametersNRV1640DiffNumerologyWithinPUCCHGroupSmallerSCSCarrierTypesR16Constraints); err != nil {
				return err
			}
		}
	}

	// 8. diffNumerologyWithinPUCCH-GroupLargerSCS-CarrierTypes-r16: enumerated
	{
		if ie.DiffNumerologyWithinPUCCH_GroupLargerSCS_CarrierTypes_r16 != nil {
			if err := e.EncodeEnumerated(*ie.DiffNumerologyWithinPUCCH_GroupLargerSCS_CarrierTypes_r16, cAParametersNRV1640DiffNumerologyWithinPUCCHGroupLargerSCSCarrierTypesR16Constraints); err != nil {
				return err
			}
		}
	}

	// 9. pdcch-MonitoringCA-NonAlignedSpan-r16: integer
	{
		if ie.Pdcch_MonitoringCA_NonAlignedSpan_r16 != nil {
			if err := e.EncodeInteger(*ie.Pdcch_MonitoringCA_NonAlignedSpan_r16, cAParametersNRV1640PdcchMonitoringCANonAlignedSpanR16Constraints); err != nil {
				return err
			}
		}
	}

	// 10. pdcch-BlindDetectionCA-Mixed-NonAlignedSpan-r16: sequence
	{
		if ie.Pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_r16 != nil {
			c := ie.Pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_r16
			if err := e.EncodeInteger(c.Pdcch_BlindDetectionCA1_r16, per.Constrained(1, 15)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.Pdcch_BlindDetectionCA2_r16, per.Constrained(1, 15)); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CA_ParametersNR_v1640) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cAParametersNRV1640Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. uplinkTxDC-TwoCarrierReport-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1640UplinkTxDCTwoCarrierReportR16Constraints)
			if err != nil {
				return err
			}
			ie.UplinkTxDC_TwoCarrierReport_r16 = &idx
		}
	}

	// 3. maxUpTo3Diff-NumerologiesConfigSinglePUCCH-grp-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.MaxUpTo3Diff_NumerologiesConfigSinglePUCCH_Grp_r16 = new(PUCCH_Grp_CarrierTypes_r16)
			if err := ie.MaxUpTo3Diff_NumerologiesConfigSinglePUCCH_Grp_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. maxUpTo4Diff-NumerologiesConfigSinglePUCCH-grp-r16: ref
	{
		if seq.IsComponentPresent(2) {
			ie.MaxUpTo4Diff_NumerologiesConfigSinglePUCCH_Grp_r16 = new(PUCCH_Grp_CarrierTypes_r16)
			if err := ie.MaxUpTo4Diff_NumerologiesConfigSinglePUCCH_Grp_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. twoPUCCH-Grp-ConfigurationsList-r16: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(cAParametersNRV1640TwoPUCCHGrpConfigurationsListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.TwoPUCCH_Grp_ConfigurationsList_r16 = make([]TwoPUCCH_Grp_Configurations_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.TwoPUCCH_Grp_ConfigurationsList_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. diffNumerologyAcrossPUCCH-Group-CarrierTypes-r16: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1640DiffNumerologyAcrossPUCCHGroupCarrierTypesR16Constraints)
			if err != nil {
				return err
			}
			ie.DiffNumerologyAcrossPUCCH_Group_CarrierTypes_r16 = &idx
		}
	}

	// 7. diffNumerologyWithinPUCCH-GroupSmallerSCS-CarrierTypes-r16: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1640DiffNumerologyWithinPUCCHGroupSmallerSCSCarrierTypesR16Constraints)
			if err != nil {
				return err
			}
			ie.DiffNumerologyWithinPUCCH_GroupSmallerSCS_CarrierTypes_r16 = &idx
		}
	}

	// 8. diffNumerologyWithinPUCCH-GroupLargerSCS-CarrierTypes-r16: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1640DiffNumerologyWithinPUCCHGroupLargerSCSCarrierTypesR16Constraints)
			if err != nil {
				return err
			}
			ie.DiffNumerologyWithinPUCCH_GroupLargerSCS_CarrierTypes_r16 = &idx
		}
	}

	// 9. pdcch-MonitoringCA-NonAlignedSpan-r16: integer
	{
		if seq.IsComponentPresent(7) {
			v, err := d.DecodeInteger(cAParametersNRV1640PdcchMonitoringCANonAlignedSpanR16Constraints)
			if err != nil {
				return err
			}
			ie.Pdcch_MonitoringCA_NonAlignedSpan_r16 = &v
		}
	}

	// 10. pdcch-BlindDetectionCA-Mixed-NonAlignedSpan-r16: sequence
	{
		if seq.IsComponentPresent(8) {
			ie.Pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_r16 = &struct {
				Pdcch_BlindDetectionCA1_r16 int64
				Pdcch_BlindDetectionCA2_r16 int64
			}{}
			c := ie.Pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_r16
			{
				v, err := d.DecodeInteger(per.Constrained(1, 15))
				if err != nil {
					return err
				}
				c.Pdcch_BlindDetectionCA1_r16 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 15))
				if err != nil {
					return err
				}
				c.Pdcch_BlindDetectionCA2_r16 = v
			}
		}
	}

	return nil
}

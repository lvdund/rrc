// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasIdleCarrierNR-r16 (line 9272).

var measIdleCarrierNRR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "carrierFreq-r16"},
		{Name: "ssbSubcarrierSpacing-r16"},
		{Name: "frequencyBandList", Optional: true},
		{Name: "measCellListNR-r16", Optional: true},
		{Name: "reportQuantities-r16"},
		{Name: "qualityThreshold-r16", Optional: true},
		{Name: "ssb-MeasConfig-r16", Optional: true},
		{Name: "beamMeasConfigIdle-r16", Optional: true},
	},
}

const (
	MeasIdleCarrierNR_r16_ReportQuantities_r16_Rsrp = 0
	MeasIdleCarrierNR_r16_ReportQuantities_r16_Rsrq = 1
	MeasIdleCarrierNR_r16_ReportQuantities_r16_Both = 2
)

var measIdleCarrierNRR16ReportQuantitiesR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasIdleCarrierNR_r16_ReportQuantities_r16_Rsrp, MeasIdleCarrierNR_r16_ReportQuantities_r16_Rsrq, MeasIdleCarrierNR_r16_ReportQuantities_r16_Both},
}

var measIdleCarrierNRR16QualityThresholdR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "idleRSRP-Threshold-NR-r16", Optional: true},
		{Name: "idleRSRQ-Threshold-NR-r16", Optional: true},
	},
}

var measIdleCarrierNRR16SsbMeasConfigR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "nrofSS-BlocksToAverage-r16", Optional: true},
		{Name: "absThreshSS-BlocksConsolidation-r16", Optional: true},
		{Name: "smtc-r16", Optional: true},
		{Name: "ssb-ToMeasure-r16", Optional: true},
		{Name: "deriveSSB-IndexFromCell-r16"},
		{Name: "ss-RSSI-Measurement-r16", Optional: true},
	},
}

type MeasIdleCarrierNR_r16 struct {
	CarrierFreq_r16          ARFCN_ValueNR
	SsbSubcarrierSpacing_r16 SubcarrierSpacing
	FrequencyBandList        *MultiFrequencyBandListNR
	MeasCellListNR_r16       *CellListNR_r16
	ReportQuantities_r16     int64
	QualityThreshold_r16     *struct {
		IdleRSRP_Threshold_NR_r16 *RSRP_Range
		IdleRSRQ_Threshold_NR_r16 *RSRQ_Range
	}
	Ssb_MeasConfig_r16 *struct {
		NrofSS_BlocksToAverage_r16          *int64
		AbsThreshSS_BlocksConsolidation_r16 *ThresholdNR
		Smtc_r16                            *SSB_MTC
		Ssb_ToMeasure_r16                   *SSB_ToMeasure
		DeriveSSB_IndexFromCell_r16         bool
		Ss_RSSI_Measurement_r16             *SS_RSSI_Measurement
	}
	BeamMeasConfigIdle_r16 *BeamMeasConfigIdle_NR_r16
}

func (ie *MeasIdleCarrierNR_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measIdleCarrierNRR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FrequencyBandList != nil, ie.MeasCellListNR_r16 != nil, ie.QualityThreshold_r16 != nil, ie.Ssb_MeasConfig_r16 != nil, ie.BeamMeasConfigIdle_r16 != nil}); err != nil {
		return err
	}

	// 3. carrierFreq-r16: ref
	{
		if err := ie.CarrierFreq_r16.Encode(e); err != nil {
			return err
		}
	}

	// 4. ssbSubcarrierSpacing-r16: ref
	{
		if err := ie.SsbSubcarrierSpacing_r16.Encode(e); err != nil {
			return err
		}
	}

	// 5. frequencyBandList: ref
	{
		if ie.FrequencyBandList != nil {
			if err := ie.FrequencyBandList.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. measCellListNR-r16: ref
	{
		if ie.MeasCellListNR_r16 != nil {
			if err := ie.MeasCellListNR_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. reportQuantities-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.ReportQuantities_r16, measIdleCarrierNRR16ReportQuantitiesR16Constraints); err != nil {
			return err
		}
	}

	// 8. qualityThreshold-r16: sequence
	{
		if ie.QualityThreshold_r16 != nil {
			c := ie.QualityThreshold_r16
			measIdleCarrierNRR16QualityThresholdR16Seq := e.NewSequenceEncoder(measIdleCarrierNRR16QualityThresholdR16Constraints)
			if err := measIdleCarrierNRR16QualityThresholdR16Seq.EncodePreamble([]bool{c.IdleRSRP_Threshold_NR_r16 != nil, c.IdleRSRQ_Threshold_NR_r16 != nil}); err != nil {
				return err
			}
			if c.IdleRSRP_Threshold_NR_r16 != nil {
				if err := c.IdleRSRP_Threshold_NR_r16.Encode(e); err != nil {
					return err
				}
			}
			if c.IdleRSRQ_Threshold_NR_r16 != nil {
				if err := c.IdleRSRQ_Threshold_NR_r16.Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 9. ssb-MeasConfig-r16: sequence
	{
		if ie.Ssb_MeasConfig_r16 != nil {
			c := ie.Ssb_MeasConfig_r16
			measIdleCarrierNRR16SsbMeasConfigR16Seq := e.NewSequenceEncoder(measIdleCarrierNRR16SsbMeasConfigR16Constraints)
			if err := measIdleCarrierNRR16SsbMeasConfigR16Seq.EncodePreamble([]bool{c.NrofSS_BlocksToAverage_r16 != nil, c.AbsThreshSS_BlocksConsolidation_r16 != nil, c.Smtc_r16 != nil, c.Ssb_ToMeasure_r16 != nil, c.Ss_RSSI_Measurement_r16 != nil}); err != nil {
				return err
			}
			if c.NrofSS_BlocksToAverage_r16 != nil {
				if err := e.EncodeInteger((*c.NrofSS_BlocksToAverage_r16), per.Constrained(2, common.MaxNrofSS_BlocksToAverage)); err != nil {
					return err
				}
			}
			if c.AbsThreshSS_BlocksConsolidation_r16 != nil {
				if err := c.AbsThreshSS_BlocksConsolidation_r16.Encode(e); err != nil {
					return err
				}
			}
			if c.Smtc_r16 != nil {
				if err := c.Smtc_r16.Encode(e); err != nil {
					return err
				}
			}
			if c.Ssb_ToMeasure_r16 != nil {
				if err := c.Ssb_ToMeasure_r16.Encode(e); err != nil {
					return err
				}
			}
			if err := e.EncodeBoolean(c.DeriveSSB_IndexFromCell_r16); err != nil {
				return err
			}
			if c.Ss_RSSI_Measurement_r16 != nil {
				if err := c.Ss_RSSI_Measurement_r16.Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 10. beamMeasConfigIdle-r16: ref
	{
		if ie.BeamMeasConfigIdle_r16 != nil {
			if err := ie.BeamMeasConfigIdle_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasIdleCarrierNR_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measIdleCarrierNRR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. carrierFreq-r16: ref
	{
		if err := ie.CarrierFreq_r16.Decode(d); err != nil {
			return err
		}
	}

	// 4. ssbSubcarrierSpacing-r16: ref
	{
		if err := ie.SsbSubcarrierSpacing_r16.Decode(d); err != nil {
			return err
		}
	}

	// 5. frequencyBandList: ref
	{
		if seq.IsComponentPresent(2) {
			ie.FrequencyBandList = new(MultiFrequencyBandListNR)
			if err := ie.FrequencyBandList.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. measCellListNR-r16: ref
	{
		if seq.IsComponentPresent(3) {
			ie.MeasCellListNR_r16 = new(CellListNR_r16)
			if err := ie.MeasCellListNR_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. reportQuantities-r16: enumerated
	{
		v4, err := d.DecodeEnumerated(measIdleCarrierNRR16ReportQuantitiesR16Constraints)
		if err != nil {
			return err
		}
		ie.ReportQuantities_r16 = v4
	}

	// 8. qualityThreshold-r16: sequence
	{
		if seq.IsComponentPresent(5) {
			ie.QualityThreshold_r16 = &struct {
				IdleRSRP_Threshold_NR_r16 *RSRP_Range
				IdleRSRQ_Threshold_NR_r16 *RSRQ_Range
			}{}
			c := ie.QualityThreshold_r16
			measIdleCarrierNRR16QualityThresholdR16Seq := d.NewSequenceDecoder(measIdleCarrierNRR16QualityThresholdR16Constraints)
			if err := measIdleCarrierNRR16QualityThresholdR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if measIdleCarrierNRR16QualityThresholdR16Seq.IsComponentPresent(0) {
				c.IdleRSRP_Threshold_NR_r16 = new(RSRP_Range)
				if err := (*c.IdleRSRP_Threshold_NR_r16).Decode(d); err != nil {
					return err
				}
			}
			if measIdleCarrierNRR16QualityThresholdR16Seq.IsComponentPresent(1) {
				c.IdleRSRQ_Threshold_NR_r16 = new(RSRQ_Range)
				if err := (*c.IdleRSRQ_Threshold_NR_r16).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 9. ssb-MeasConfig-r16: sequence
	{
		if seq.IsComponentPresent(6) {
			ie.Ssb_MeasConfig_r16 = &struct {
				NrofSS_BlocksToAverage_r16          *int64
				AbsThreshSS_BlocksConsolidation_r16 *ThresholdNR
				Smtc_r16                            *SSB_MTC
				Ssb_ToMeasure_r16                   *SSB_ToMeasure
				DeriveSSB_IndexFromCell_r16         bool
				Ss_RSSI_Measurement_r16             *SS_RSSI_Measurement
			}{}
			c := ie.Ssb_MeasConfig_r16
			measIdleCarrierNRR16SsbMeasConfigR16Seq := d.NewSequenceDecoder(measIdleCarrierNRR16SsbMeasConfigR16Constraints)
			if err := measIdleCarrierNRR16SsbMeasConfigR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if measIdleCarrierNRR16SsbMeasConfigR16Seq.IsComponentPresent(0) {
				c.NrofSS_BlocksToAverage_r16 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(2, common.MaxNrofSS_BlocksToAverage))
				if err != nil {
					return err
				}
				(*c.NrofSS_BlocksToAverage_r16) = v
			}
			if measIdleCarrierNRR16SsbMeasConfigR16Seq.IsComponentPresent(1) {
				c.AbsThreshSS_BlocksConsolidation_r16 = new(ThresholdNR)
				if err := (*c.AbsThreshSS_BlocksConsolidation_r16).Decode(d); err != nil {
					return err
				}
			}
			if measIdleCarrierNRR16SsbMeasConfigR16Seq.IsComponentPresent(2) {
				c.Smtc_r16 = new(SSB_MTC)
				if err := (*c.Smtc_r16).Decode(d); err != nil {
					return err
				}
			}
			if measIdleCarrierNRR16SsbMeasConfigR16Seq.IsComponentPresent(3) {
				c.Ssb_ToMeasure_r16 = new(SSB_ToMeasure)
				if err := (*c.Ssb_ToMeasure_r16).Decode(d); err != nil {
					return err
				}
			}
			{
				v, err := d.DecodeBoolean()
				if err != nil {
					return err
				}
				c.DeriveSSB_IndexFromCell_r16 = v
			}
			if measIdleCarrierNRR16SsbMeasConfigR16Seq.IsComponentPresent(5) {
				c.Ss_RSSI_Measurement_r16 = new(SS_RSSI_Measurement)
				if err := (*c.Ss_RSSI_Measurement_r16).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 10. beamMeasConfigIdle-r16: ref
	{
		if seq.IsComponentPresent(7) {
			ie.BeamMeasConfigIdle_r16 = new(BeamMeasConfigIdle_NR_r16)
			if err := ie.BeamMeasConfigIdle_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}

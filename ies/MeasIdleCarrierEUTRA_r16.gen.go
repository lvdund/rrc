// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasIdleCarrierEUTRA-r16 (line 9294).

var measIdleCarrierEUTRAR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "carrierFreqEUTRA-r16"},
		{Name: "allowedMeasBandwidth-r16"},
		{Name: "measCellListEUTRA-r16", Optional: true},
		{Name: "reportQuantitiesEUTRA-r16"},
		{Name: "qualityThresholdEUTRA-r16", Optional: true},
	},
}

const (
	MeasIdleCarrierEUTRA_r16_ReportQuantitiesEUTRA_r16_Rsrp = 0
	MeasIdleCarrierEUTRA_r16_ReportQuantitiesEUTRA_r16_Rsrq = 1
	MeasIdleCarrierEUTRA_r16_ReportQuantitiesEUTRA_r16_Both = 2
)

var measIdleCarrierEUTRAR16ReportQuantitiesEUTRAR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasIdleCarrierEUTRA_r16_ReportQuantitiesEUTRA_r16_Rsrp, MeasIdleCarrierEUTRA_r16_ReportQuantitiesEUTRA_r16_Rsrq, MeasIdleCarrierEUTRA_r16_ReportQuantitiesEUTRA_r16_Both},
}

var measIdleCarrierEUTRAR16QualityThresholdEUTRAR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "idleRSRP-Threshold-EUTRA-r16", Optional: true},
		{Name: "idleRSRQ-Threshold-EUTRA-r16", Optional: true},
	},
}

type MeasIdleCarrierEUTRA_r16 struct {
	CarrierFreqEUTRA_r16      ARFCN_ValueEUTRA
	AllowedMeasBandwidth_r16  EUTRA_AllowedMeasBandwidth
	MeasCellListEUTRA_r16     *CellListEUTRA_r16
	ReportQuantitiesEUTRA_r16 int64
	QualityThresholdEUTRA_r16 *struct {
		IdleRSRP_Threshold_EUTRA_r16 *RSRP_RangeEUTRA
		IdleRSRQ_Threshold_EUTRA_r16 *RSRQ_RangeEUTRA_r16
	}
}

func (ie *MeasIdleCarrierEUTRA_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measIdleCarrierEUTRAR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MeasCellListEUTRA_r16 != nil, ie.QualityThresholdEUTRA_r16 != nil}); err != nil {
		return err
	}

	// 3. carrierFreqEUTRA-r16: ref
	{
		if err := ie.CarrierFreqEUTRA_r16.Encode(e); err != nil {
			return err
		}
	}

	// 4. allowedMeasBandwidth-r16: ref
	{
		if err := ie.AllowedMeasBandwidth_r16.Encode(e); err != nil {
			return err
		}
	}

	// 5. measCellListEUTRA-r16: ref
	{
		if ie.MeasCellListEUTRA_r16 != nil {
			if err := ie.MeasCellListEUTRA_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. reportQuantitiesEUTRA-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.ReportQuantitiesEUTRA_r16, measIdleCarrierEUTRAR16ReportQuantitiesEUTRAR16Constraints); err != nil {
			return err
		}
	}

	// 7. qualityThresholdEUTRA-r16: sequence
	{
		if ie.QualityThresholdEUTRA_r16 != nil {
			c := ie.QualityThresholdEUTRA_r16
			measIdleCarrierEUTRAR16QualityThresholdEUTRAR16Seq := e.NewSequenceEncoder(measIdleCarrierEUTRAR16QualityThresholdEUTRAR16Constraints)
			if err := measIdleCarrierEUTRAR16QualityThresholdEUTRAR16Seq.EncodePreamble([]bool{c.IdleRSRP_Threshold_EUTRA_r16 != nil, c.IdleRSRQ_Threshold_EUTRA_r16 != nil}); err != nil {
				return err
			}
			if c.IdleRSRP_Threshold_EUTRA_r16 != nil {
				if err := c.IdleRSRP_Threshold_EUTRA_r16.Encode(e); err != nil {
					return err
				}
			}
			if c.IdleRSRQ_Threshold_EUTRA_r16 != nil {
				if err := c.IdleRSRQ_Threshold_EUTRA_r16.Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *MeasIdleCarrierEUTRA_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measIdleCarrierEUTRAR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. carrierFreqEUTRA-r16: ref
	{
		if err := ie.CarrierFreqEUTRA_r16.Decode(d); err != nil {
			return err
		}
	}

	// 4. allowedMeasBandwidth-r16: ref
	{
		if err := ie.AllowedMeasBandwidth_r16.Decode(d); err != nil {
			return err
		}
	}

	// 5. measCellListEUTRA-r16: ref
	{
		if seq.IsComponentPresent(2) {
			ie.MeasCellListEUTRA_r16 = new(CellListEUTRA_r16)
			if err := ie.MeasCellListEUTRA_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. reportQuantitiesEUTRA-r16: enumerated
	{
		v3, err := d.DecodeEnumerated(measIdleCarrierEUTRAR16ReportQuantitiesEUTRAR16Constraints)
		if err != nil {
			return err
		}
		ie.ReportQuantitiesEUTRA_r16 = v3
	}

	// 7. qualityThresholdEUTRA-r16: sequence
	{
		if seq.IsComponentPresent(4) {
			ie.QualityThresholdEUTRA_r16 = &struct {
				IdleRSRP_Threshold_EUTRA_r16 *RSRP_RangeEUTRA
				IdleRSRQ_Threshold_EUTRA_r16 *RSRQ_RangeEUTRA_r16
			}{}
			c := ie.QualityThresholdEUTRA_r16
			measIdleCarrierEUTRAR16QualityThresholdEUTRAR16Seq := d.NewSequenceDecoder(measIdleCarrierEUTRAR16QualityThresholdEUTRAR16Constraints)
			if err := measIdleCarrierEUTRAR16QualityThresholdEUTRAR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if measIdleCarrierEUTRAR16QualityThresholdEUTRAR16Seq.IsComponentPresent(0) {
				c.IdleRSRP_Threshold_EUTRA_r16 = new(RSRP_RangeEUTRA)
				if err := (*c.IdleRSRP_Threshold_EUTRA_r16).Decode(d); err != nil {
					return err
				}
			}
			if measIdleCarrierEUTRAR16QualityThresholdEUTRAR16Seq.IsComponentPresent(1) {
				c.IdleRSRQ_Threshold_EUTRA_r16 = new(RSRQ_RangeEUTRA_r16)
				if err := (*c.IdleRSRQ_Threshold_EUTRA_r16).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

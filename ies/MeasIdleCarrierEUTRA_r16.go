package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasIdleCarrierEUTRA_r16 struct {
	CarrierFreqEUTRA_r16      ARFCN_ValueEUTRA                                    `madatory`
	AllowedMeasBandwidth_r16  EUTRA_AllowedMeasBandwidth                          `madatory`
	MeasCellListEUTRA_r16     *CellListEUTRA_r16                                  `optional`
	ReportQuantitiesEUTRA_r16 MeasIdleCarrierEUTRA_r16_reportQuantitiesEUTRA_r16  `madatory`
	QualityThresholdEUTRA_r16 *MeasIdleCarrierEUTRA_r16_qualityThresholdEUTRA_r16 `optional`
}

func (ie *MeasIdleCarrierEUTRA_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.MeasCellListEUTRA_r16 != nil, ie.QualityThresholdEUTRA_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.CarrierFreqEUTRA_r16.Encode(w); err != nil {
		return utils.WrapError("Encode CarrierFreqEUTRA_r16", err)
	}
	if err = ie.AllowedMeasBandwidth_r16.Encode(w); err != nil {
		return utils.WrapError("Encode AllowedMeasBandwidth_r16", err)
	}
	if ie.MeasCellListEUTRA_r16 != nil {
		if err = ie.MeasCellListEUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MeasCellListEUTRA_r16", err)
		}
	}
	if err = ie.ReportQuantitiesEUTRA_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ReportQuantitiesEUTRA_r16", err)
	}
	if ie.QualityThresholdEUTRA_r16 != nil {
		if err = ie.QualityThresholdEUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode QualityThresholdEUTRA_r16", err)
		}
	}
	return nil
}

func (ie *MeasIdleCarrierEUTRA_r16) Decode(r *uper.UperReader) error {
	var err error
	var MeasCellListEUTRA_r16Present bool
	if MeasCellListEUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var QualityThresholdEUTRA_r16Present bool
	if QualityThresholdEUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.CarrierFreqEUTRA_r16.Decode(r); err != nil {
		return utils.WrapError("Decode CarrierFreqEUTRA_r16", err)
	}
	if err = ie.AllowedMeasBandwidth_r16.Decode(r); err != nil {
		return utils.WrapError("Decode AllowedMeasBandwidth_r16", err)
	}
	if MeasCellListEUTRA_r16Present {
		ie.MeasCellListEUTRA_r16 = new(CellListEUTRA_r16)
		if err = ie.MeasCellListEUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MeasCellListEUTRA_r16", err)
		}
	}
	if err = ie.ReportQuantitiesEUTRA_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ReportQuantitiesEUTRA_r16", err)
	}
	if QualityThresholdEUTRA_r16Present {
		ie.QualityThresholdEUTRA_r16 = new(MeasIdleCarrierEUTRA_r16_qualityThresholdEUTRA_r16)
		if err = ie.QualityThresholdEUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode QualityThresholdEUTRA_r16", err)
		}
	}
	return nil
}

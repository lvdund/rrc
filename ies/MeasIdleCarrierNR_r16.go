package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasIdleCarrierNR_r16 struct {
	CarrierFreq_r16          ARFCN_ValueNR                               `madatory`
	SsbSubcarrierSpacing_r16 SubcarrierSpacing                           `madatory`
	FrequencyBandList        *MultiFrequencyBandListNR                   `optional`
	MeasCellListNR_r16       *CellListNR_r16                             `optional`
	ReportQuantities_r16     MeasIdleCarrierNR_r16_reportQuantities_r16  `madatory`
	QualityThreshold_r16     *MeasIdleCarrierNR_r16_qualityThreshold_r16 `optional`
	Ssb_MeasConfig_r16       *MeasIdleCarrierNR_r16_ssb_MeasConfig_r16   `lb:2,ub:maxNrofSS_BlocksToAverage,optional`
	BeamMeasConfigIdle_r16   *BeamMeasConfigIdle_NR_r16                  `optional`
}

func (ie *MeasIdleCarrierNR_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.FrequencyBandList != nil, ie.MeasCellListNR_r16 != nil, ie.QualityThreshold_r16 != nil, ie.Ssb_MeasConfig_r16 != nil, ie.BeamMeasConfigIdle_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.CarrierFreq_r16.Encode(w); err != nil {
		return utils.WrapError("Encode CarrierFreq_r16", err)
	}
	if err = ie.SsbSubcarrierSpacing_r16.Encode(w); err != nil {
		return utils.WrapError("Encode SsbSubcarrierSpacing_r16", err)
	}
	if ie.FrequencyBandList != nil {
		if err = ie.FrequencyBandList.Encode(w); err != nil {
			return utils.WrapError("Encode FrequencyBandList", err)
		}
	}
	if ie.MeasCellListNR_r16 != nil {
		if err = ie.MeasCellListNR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MeasCellListNR_r16", err)
		}
	}
	if err = ie.ReportQuantities_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ReportQuantities_r16", err)
	}
	if ie.QualityThreshold_r16 != nil {
		if err = ie.QualityThreshold_r16.Encode(w); err != nil {
			return utils.WrapError("Encode QualityThreshold_r16", err)
		}
	}
	if ie.Ssb_MeasConfig_r16 != nil {
		if err = ie.Ssb_MeasConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ssb_MeasConfig_r16", err)
		}
	}
	if ie.BeamMeasConfigIdle_r16 != nil {
		if err = ie.BeamMeasConfigIdle_r16.Encode(w); err != nil {
			return utils.WrapError("Encode BeamMeasConfigIdle_r16", err)
		}
	}
	return nil
}

func (ie *MeasIdleCarrierNR_r16) Decode(r *aper.AperReader) error {
	var err error
	var FrequencyBandListPresent bool
	if FrequencyBandListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasCellListNR_r16Present bool
	if MeasCellListNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var QualityThreshold_r16Present bool
	if QualityThreshold_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ssb_MeasConfig_r16Present bool
	if Ssb_MeasConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var BeamMeasConfigIdle_r16Present bool
	if BeamMeasConfigIdle_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.CarrierFreq_r16.Decode(r); err != nil {
		return utils.WrapError("Decode CarrierFreq_r16", err)
	}
	if err = ie.SsbSubcarrierSpacing_r16.Decode(r); err != nil {
		return utils.WrapError("Decode SsbSubcarrierSpacing_r16", err)
	}
	if FrequencyBandListPresent {
		ie.FrequencyBandList = new(MultiFrequencyBandListNR)
		if err = ie.FrequencyBandList.Decode(r); err != nil {
			return utils.WrapError("Decode FrequencyBandList", err)
		}
	}
	if MeasCellListNR_r16Present {
		ie.MeasCellListNR_r16 = new(CellListNR_r16)
		if err = ie.MeasCellListNR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MeasCellListNR_r16", err)
		}
	}
	if err = ie.ReportQuantities_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ReportQuantities_r16", err)
	}
	if QualityThreshold_r16Present {
		ie.QualityThreshold_r16 = new(MeasIdleCarrierNR_r16_qualityThreshold_r16)
		if err = ie.QualityThreshold_r16.Decode(r); err != nil {
			return utils.WrapError("Decode QualityThreshold_r16", err)
		}
	}
	if Ssb_MeasConfig_r16Present {
		ie.Ssb_MeasConfig_r16 = new(MeasIdleCarrierNR_r16_ssb_MeasConfig_r16)
		if err = ie.Ssb_MeasConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_MeasConfig_r16", err)
		}
	}
	if BeamMeasConfigIdle_r16Present {
		ie.BeamMeasConfigIdle_r16 = new(BeamMeasConfigIdle_NR_r16)
		if err = ie.BeamMeasConfigIdle_r16.Decode(r); err != nil {
			return utils.WrapError("Decode BeamMeasConfigIdle_r16", err)
		}
	}
	return nil
}

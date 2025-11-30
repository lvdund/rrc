package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type FrequencyConfig_NR_r16 struct {
	FreqBandIndicatorNR_r16  FreqBandIndicatorNR `madatory`
	CarrierCenterFreq_NR_r16 ARFCN_ValueNR       `madatory`
	CarrierBandwidth_NR_r16  int64               `lb:1,ub:maxNrofPhysicalResourceBlocks,madatory`
	SubcarrierSpacing_NR_r16 SubcarrierSpacing   `madatory`
}

func (ie *FrequencyConfig_NR_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.FreqBandIndicatorNR_r16.Encode(w); err != nil {
		return utils.WrapError("Encode FreqBandIndicatorNR_r16", err)
	}
	if err = ie.CarrierCenterFreq_NR_r16.Encode(w); err != nil {
		return utils.WrapError("Encode CarrierCenterFreq_NR_r16", err)
	}
	if err = w.WriteInteger(ie.CarrierBandwidth_NR_r16, &aper.Constraint{Lb: 1, Ub: maxNrofPhysicalResourceBlocks}, false); err != nil {
		return utils.WrapError("WriteInteger CarrierBandwidth_NR_r16", err)
	}
	if err = ie.SubcarrierSpacing_NR_r16.Encode(w); err != nil {
		return utils.WrapError("Encode SubcarrierSpacing_NR_r16", err)
	}
	return nil
}

func (ie *FrequencyConfig_NR_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.FreqBandIndicatorNR_r16.Decode(r); err != nil {
		return utils.WrapError("Decode FreqBandIndicatorNR_r16", err)
	}
	if err = ie.CarrierCenterFreq_NR_r16.Decode(r); err != nil {
		return utils.WrapError("Decode CarrierCenterFreq_NR_r16", err)
	}
	var tmp_int_CarrierBandwidth_NR_r16 int64
	if tmp_int_CarrierBandwidth_NR_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: maxNrofPhysicalResourceBlocks}, false); err != nil {
		return utils.WrapError("ReadInteger CarrierBandwidth_NR_r16", err)
	}
	ie.CarrierBandwidth_NR_r16 = tmp_int_CarrierBandwidth_NR_r16
	if err = ie.SubcarrierSpacing_NR_r16.Decode(r); err != nil {
		return utils.WrapError("Decode SubcarrierSpacing_NR_r16", err)
	}
	return nil
}

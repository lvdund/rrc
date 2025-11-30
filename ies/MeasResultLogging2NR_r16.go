package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultLogging2NR_r16 struct {
	CarrierFreq_r16             ARFCN_ValueNR               `madatory`
	MeasResultListLoggingNR_r16 MeasResultListLoggingNR_r16 `madatory`
}

func (ie *MeasResultLogging2NR_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.CarrierFreq_r16.Encode(w); err != nil {
		return utils.WrapError("Encode CarrierFreq_r16", err)
	}
	if err = ie.MeasResultListLoggingNR_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MeasResultListLoggingNR_r16", err)
	}
	return nil
}

func (ie *MeasResultLogging2NR_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.CarrierFreq_r16.Decode(r); err != nil {
		return utils.WrapError("Decode CarrierFreq_r16", err)
	}
	if err = ie.MeasResultListLoggingNR_r16.Decode(r); err != nil {
		return utils.WrapError("Decode MeasResultListLoggingNR_r16", err)
	}
	return nil
}

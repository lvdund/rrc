package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PSCellId_r17 struct {
	PhysCellId_r17  PhysCellId    `madatory`
	CarrierFreq_r17 ARFCN_ValueNR `madatory`
}

func (ie *PSCellId_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.PhysCellId_r17.Encode(w); err != nil {
		return utils.WrapError("Encode PhysCellId_r17", err)
	}
	if err = ie.CarrierFreq_r17.Encode(w); err != nil {
		return utils.WrapError("Encode CarrierFreq_r17", err)
	}
	return nil
}

func (ie *PSCellId_r17) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.PhysCellId_r17.Decode(r); err != nil {
		return utils.WrapError("Decode PhysCellId_r17", err)
	}
	if err = ie.CarrierFreq_r17.Decode(r); err != nil {
		return utils.WrapError("Decode CarrierFreq_r17", err)
	}
	return nil
}

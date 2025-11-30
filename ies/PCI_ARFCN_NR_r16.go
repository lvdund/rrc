package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PCI_ARFCN_NR_r16 struct {
	PhysCellId_r16  PhysCellId    `madatory`
	CarrierFreq_r16 ARFCN_ValueNR `madatory`
}

func (ie *PCI_ARFCN_NR_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.PhysCellId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode PhysCellId_r16", err)
	}
	if err = ie.CarrierFreq_r16.Encode(w); err != nil {
		return utils.WrapError("Encode CarrierFreq_r16", err)
	}
	return nil
}

func (ie *PCI_ARFCN_NR_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.PhysCellId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode PhysCellId_r16", err)
	}
	if err = ie.CarrierFreq_r16.Decode(r); err != nil {
		return utils.WrapError("Decode CarrierFreq_r16", err)
	}
	return nil
}

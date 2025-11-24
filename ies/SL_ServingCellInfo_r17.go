package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_ServingCellInfo_r17 struct {
	Sl_PhysCellId_r17    PhysCellId    `madatory`
	Sl_CarrierFreqNR_r17 ARFCN_ValueNR `madatory`
}

func (ie *SL_ServingCellInfo_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Sl_PhysCellId_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_PhysCellId_r17", err)
	}
	if err = ie.Sl_CarrierFreqNR_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_CarrierFreqNR_r17", err)
	}
	return nil
}

func (ie *SL_ServingCellInfo_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Sl_PhysCellId_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_PhysCellId_r17", err)
	}
	if err = ie.Sl_CarrierFreqNR_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_CarrierFreqNR_r17", err)
	}
	return nil
}

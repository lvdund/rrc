package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UplinkTxDirectCurrentCarrierInfo_r16 struct {
	ServCellIndex_r16 ServCellIndex                                         `madatory`
	ServCellInfo_r16  UplinkTxDirectCurrentCarrierInfo_r16_servCellInfo_r16 `madatory`
}

func (ie *UplinkTxDirectCurrentCarrierInfo_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.ServCellIndex_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ServCellIndex_r16", err)
	}
	if err = ie.ServCellInfo_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ServCellInfo_r16", err)
	}
	return nil
}

func (ie *UplinkTxDirectCurrentCarrierInfo_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.ServCellIndex_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ServCellIndex_r16", err)
	}
	if err = ie.ServCellInfo_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ServCellInfo_r16", err)
	}
	return nil
}

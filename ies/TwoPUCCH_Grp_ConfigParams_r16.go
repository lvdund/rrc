package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type TwoPUCCH_Grp_ConfigParams_r16 struct {
	Pucch_GroupMapping_r16 PUCCH_Grp_CarrierTypes_r16 `madatory`
	Pucch_TX_r16           PUCCH_Grp_CarrierTypes_r16 `madatory`
}

func (ie *TwoPUCCH_Grp_ConfigParams_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Pucch_GroupMapping_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Pucch_GroupMapping_r16", err)
	}
	if err = ie.Pucch_TX_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Pucch_TX_r16", err)
	}
	return nil
}

func (ie *TwoPUCCH_Grp_ConfigParams_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Pucch_GroupMapping_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Pucch_GroupMapping_r16", err)
	}
	if err = ie.Pucch_TX_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Pucch_TX_r16", err)
	}
	return nil
}

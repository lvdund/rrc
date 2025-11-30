package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MBS_FSAI_InterFreq_r17 struct {
	Dl_CarrierFreq_r17 ARFCN_ValueNR     `madatory`
	Mbs_FSAI_List_r17  MBS_FSAI_List_r17 `madatory`
}

func (ie *MBS_FSAI_InterFreq_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Dl_CarrierFreq_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Dl_CarrierFreq_r17", err)
	}
	if err = ie.Mbs_FSAI_List_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Mbs_FSAI_List_r17", err)
	}
	return nil
}

func (ie *MBS_FSAI_InterFreq_r17) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Dl_CarrierFreq_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Dl_CarrierFreq_r17", err)
	}
	if err = ie.Mbs_FSAI_List_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Mbs_FSAI_List_r17", err)
	}
	return nil
}

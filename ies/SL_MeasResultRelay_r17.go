package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_MeasResultRelay_r17 struct {
	CellIdentity_r17        CellAccessRelatedInfo `madatory`
	Sl_RelayUE_Identity_r17 SL_SourceIdentity_r17 `madatory`
	Sl_MeasResult_r17       SL_MeasResult_r16     `madatory`
}

func (ie *SL_MeasResultRelay_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.CellIdentity_r17.Encode(w); err != nil {
		return utils.WrapError("Encode CellIdentity_r17", err)
	}
	if err = ie.Sl_RelayUE_Identity_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_RelayUE_Identity_r17", err)
	}
	if err = ie.Sl_MeasResult_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_MeasResult_r17", err)
	}
	return nil
}

func (ie *SL_MeasResultRelay_r17) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.CellIdentity_r17.Decode(r); err != nil {
		return utils.WrapError("Decode CellIdentity_r17", err)
	}
	if err = ie.Sl_RelayUE_Identity_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_RelayUE_Identity_r17", err)
	}
	if err = ie.Sl_MeasResult_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_MeasResult_r17", err)
	}
	return nil
}

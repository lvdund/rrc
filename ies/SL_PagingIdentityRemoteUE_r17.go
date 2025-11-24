package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_PagingIdentityRemoteUE_r17 struct {
	Ng_5G_S_TMSI_r17 NG_5G_S_TMSI  `madatory`
	FullI_RNTI_r17   *I_RNTI_Value `optional`
}

func (ie *SL_PagingIdentityRemoteUE_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.FullI_RNTI_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Ng_5G_S_TMSI_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Ng_5G_S_TMSI_r17", err)
	}
	if ie.FullI_RNTI_r17 != nil {
		if err = ie.FullI_RNTI_r17.Encode(w); err != nil {
			return utils.WrapError("Encode FullI_RNTI_r17", err)
		}
	}
	return nil
}

func (ie *SL_PagingIdentityRemoteUE_r17) Decode(r *uper.UperReader) error {
	var err error
	var FullI_RNTI_r17Present bool
	if FullI_RNTI_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Ng_5G_S_TMSI_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Ng_5G_S_TMSI_r17", err)
	}
	if FullI_RNTI_r17Present {
		ie.FullI_RNTI_r17 = new(I_RNTI_Value)
		if err = ie.FullI_RNTI_r17.Decode(r); err != nil {
			return utils.WrapError("Decode FullI_RNTI_r17", err)
		}
	}
	return nil
}

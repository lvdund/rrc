package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_PagingInfo_RemoteUE_r17 struct {
	Sl_PagingIdentityRemoteUE_r17 SL_PagingIdentityRemoteUE_r17 `madatory`
	Sl_PagingCycleRemoteUE_r17    *PagingCycle                  `optional`
}

func (ie *SL_PagingInfo_RemoteUE_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_PagingCycleRemoteUE_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Sl_PagingIdentityRemoteUE_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_PagingIdentityRemoteUE_r17", err)
	}
	if ie.Sl_PagingCycleRemoteUE_r17 != nil {
		if err = ie.Sl_PagingCycleRemoteUE_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_PagingCycleRemoteUE_r17", err)
		}
	}
	return nil
}

func (ie *SL_PagingInfo_RemoteUE_r17) Decode(r *aper.AperReader) error {
	var err error
	var Sl_PagingCycleRemoteUE_r17Present bool
	if Sl_PagingCycleRemoteUE_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Sl_PagingIdentityRemoteUE_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_PagingIdentityRemoteUE_r17", err)
	}
	if Sl_PagingCycleRemoteUE_r17Present {
		ie.Sl_PagingCycleRemoteUE_r17 = new(PagingCycle)
		if err = ie.Sl_PagingCycleRemoteUE_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_PagingCycleRemoteUE_r17", err)
		}
	}
	return nil
}

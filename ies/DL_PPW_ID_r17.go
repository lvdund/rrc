package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type DL_PPW_ID_r17 struct {
	Value uint64 `lb:0,ub:maxNrofPPW_ID_1_r17,madatory`
}

func (ie *DL_PPW_ID_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 0, Ub: maxNrofPPW_ID_1_r17}, false); err != nil {
		return utils.WrapError("Encode DL_PPW_ID_r17", err)
	}
	return nil
}

func (ie *DL_PPW_ID_r17) Decode(r *aper.AperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: maxNrofPPW_ID_1_r17}, false); err != nil {
		return utils.WrapError("Decode DL_PPW_ID_r17", err)
	}
	ie.Value = uint64(v)
	return nil
}

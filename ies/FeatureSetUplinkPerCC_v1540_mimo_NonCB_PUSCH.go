package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetUplinkPerCC_v1540_mimo_NonCB_PUSCH struct {
	MaxNumberSRS_ResourcePerSet         int64 `lb:1,ub:4,madatory`
	MaxNumberSimultaneousSRS_ResourceTx int64 `lb:1,ub:4,madatory`
}

func (ie *FeatureSetUplinkPerCC_v1540_mimo_NonCB_PUSCH) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(ie.MaxNumberSRS_ResourcePerSet, &aper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger MaxNumberSRS_ResourcePerSet", err)
	}
	if err = w.WriteInteger(ie.MaxNumberSimultaneousSRS_ResourceTx, &aper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger MaxNumberSimultaneousSRS_ResourceTx", err)
	}
	return nil
}

func (ie *FeatureSetUplinkPerCC_v1540_mimo_NonCB_PUSCH) Decode(r *aper.AperReader) error {
	var err error
	var tmp_int_MaxNumberSRS_ResourcePerSet int64
	if tmp_int_MaxNumberSRS_ResourcePerSet, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger MaxNumberSRS_ResourcePerSet", err)
	}
	ie.MaxNumberSRS_ResourcePerSet = tmp_int_MaxNumberSRS_ResourcePerSet
	var tmp_int_MaxNumberSimultaneousSRS_ResourceTx int64
	if tmp_int_MaxNumberSimultaneousSRS_ResourceTx, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger MaxNumberSimultaneousSRS_ResourceTx", err)
	}
	ie.MaxNumberSimultaneousSRS_ResourceTx = tmp_int_MaxNumberSimultaneousSRS_ResourceTx
	return nil
}

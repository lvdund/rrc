package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	ServingCellConfigCommonSIB_n_TimingAdvanceOffset_Enum_n0     aper.Enumerated = 0
	ServingCellConfigCommonSIB_n_TimingAdvanceOffset_Enum_n25600 aper.Enumerated = 1
	ServingCellConfigCommonSIB_n_TimingAdvanceOffset_Enum_n39936 aper.Enumerated = 2
)

type ServingCellConfigCommonSIB_n_TimingAdvanceOffset struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *ServingCellConfigCommonSIB_n_TimingAdvanceOffset) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode ServingCellConfigCommonSIB_n_TimingAdvanceOffset", err)
	}
	return nil
}

func (ie *ServingCellConfigCommonSIB_n_TimingAdvanceOffset) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode ServingCellConfigCommonSIB_n_TimingAdvanceOffset", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}

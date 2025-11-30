package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MAC_ParametersXDD_Diff_skipUplinkTxDynamic_Enum_supported aper.Enumerated = 0
)

type MAC_ParametersXDD_Diff_skipUplinkTxDynamic struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *MAC_ParametersXDD_Diff_skipUplinkTxDynamic) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode MAC_ParametersXDD_Diff_skipUplinkTxDynamic", err)
	}
	return nil
}

func (ie *MAC_ParametersXDD_Diff_skipUplinkTxDynamic) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode MAC_ParametersXDD_Diff_skipUplinkTxDynamic", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}

package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	DummyG_maxNumberSSB_CSI_RS_ResourceOneTx_Enum_n8  aper.Enumerated = 0
	DummyG_maxNumberSSB_CSI_RS_ResourceOneTx_Enum_n16 aper.Enumerated = 1
	DummyG_maxNumberSSB_CSI_RS_ResourceOneTx_Enum_n32 aper.Enumerated = 2
	DummyG_maxNumberSSB_CSI_RS_ResourceOneTx_Enum_n64 aper.Enumerated = 3
)

type DummyG_maxNumberSSB_CSI_RS_ResourceOneTx struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *DummyG_maxNumberSSB_CSI_RS_ResourceOneTx) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode DummyG_maxNumberSSB_CSI_RS_ResourceOneTx", err)
	}
	return nil
}

func (ie *DummyG_maxNumberSSB_CSI_RS_ResourceOneTx) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode DummyG_maxNumberSSB_CSI_RS_ResourceOneTx", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}

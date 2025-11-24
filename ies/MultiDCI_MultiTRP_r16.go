package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MultiDCI_MultiTRP_r16 struct {
	MaxNumberCORESET_r16              MultiDCI_MultiTRP_r16_maxNumberCORESET_r16              `madatory`
	MaxNumberCORESETPerPoolIndex_r16  int64                                                   `lb:1,ub:3,madatory`
	MaxNumberUnicastPDSCH_PerPool_r16 MultiDCI_MultiTRP_r16_maxNumberUnicastPDSCH_PerPool_r16 `madatory`
}

func (ie *MultiDCI_MultiTRP_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.MaxNumberCORESET_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberCORESET_r16", err)
	}
	if err = w.WriteInteger(ie.MaxNumberCORESETPerPoolIndex_r16, &uper.Constraint{Lb: 1, Ub: 3}, false); err != nil {
		return utils.WrapError("WriteInteger MaxNumberCORESETPerPoolIndex_r16", err)
	}
	if err = ie.MaxNumberUnicastPDSCH_PerPool_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberUnicastPDSCH_PerPool_r16", err)
	}
	return nil
}

func (ie *MultiDCI_MultiTRP_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.MaxNumberCORESET_r16.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberCORESET_r16", err)
	}
	var tmp_int_MaxNumberCORESETPerPoolIndex_r16 int64
	if tmp_int_MaxNumberCORESETPerPoolIndex_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 3}, false); err != nil {
		return utils.WrapError("ReadInteger MaxNumberCORESETPerPoolIndex_r16", err)
	}
	ie.MaxNumberCORESETPerPoolIndex_r16 = tmp_int_MaxNumberCORESETPerPoolIndex_r16
	if err = ie.MaxNumberUnicastPDSCH_PerPool_r16.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberUnicastPDSCH_PerPool_r16", err)
	}
	return nil
}

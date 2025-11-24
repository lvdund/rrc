package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetUplinkPerCC_mimo_CB_PUSCH struct {
	MaxNumberMIMO_LayersCB_PUSCH *MIMO_LayersUL `optional`
	MaxNumberSRS_ResourcePerSet  int64          `lb:1,ub:2,madatory`
}

func (ie *FeatureSetUplinkPerCC_mimo_CB_PUSCH) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.MaxNumberMIMO_LayersCB_PUSCH != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MaxNumberMIMO_LayersCB_PUSCH != nil {
		if err = ie.MaxNumberMIMO_LayersCB_PUSCH.Encode(w); err != nil {
			return utils.WrapError("Encode MaxNumberMIMO_LayersCB_PUSCH", err)
		}
	}
	if err = w.WriteInteger(ie.MaxNumberSRS_ResourcePerSet, &uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
		return utils.WrapError("WriteInteger MaxNumberSRS_ResourcePerSet", err)
	}
	return nil
}

func (ie *FeatureSetUplinkPerCC_mimo_CB_PUSCH) Decode(r *uper.UperReader) error {
	var err error
	var MaxNumberMIMO_LayersCB_PUSCHPresent bool
	if MaxNumberMIMO_LayersCB_PUSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if MaxNumberMIMO_LayersCB_PUSCHPresent {
		ie.MaxNumberMIMO_LayersCB_PUSCH = new(MIMO_LayersUL)
		if err = ie.MaxNumberMIMO_LayersCB_PUSCH.Decode(r); err != nil {
			return utils.WrapError("Decode MaxNumberMIMO_LayersCB_PUSCH", err)
		}
	}
	var tmp_int_MaxNumberSRS_ResourcePerSet int64
	if tmp_int_MaxNumberSRS_ResourcePerSet, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
		return utils.WrapError("ReadInteger MaxNumberSRS_ResourcePerSet", err)
	}
	ie.MaxNumberSRS_ResourcePerSet = tmp_int_MaxNumberSRS_ResourcePerSet
	return nil
}

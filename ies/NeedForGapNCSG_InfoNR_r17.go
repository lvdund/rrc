package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NeedForGapNCSG_InfoNR_r17 struct {
	IntraFreq_needForNCSG_r17 NeedForNCSG_IntraFreqList_r17 `madatory`
	InterFreq_needForNCSG_r17 NeedForNCSG_BandListNR_r17    `madatory`
}

func (ie *NeedForGapNCSG_InfoNR_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.IntraFreq_needForNCSG_r17.Encode(w); err != nil {
		return utils.WrapError("Encode IntraFreq_needForNCSG_r17", err)
	}
	if err = ie.InterFreq_needForNCSG_r17.Encode(w); err != nil {
		return utils.WrapError("Encode InterFreq_needForNCSG_r17", err)
	}
	return nil
}

func (ie *NeedForGapNCSG_InfoNR_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.IntraFreq_needForNCSG_r17.Decode(r); err != nil {
		return utils.WrapError("Decode IntraFreq_needForNCSG_r17", err)
	}
	if err = ie.InterFreq_needForNCSG_r17.Decode(r); err != nil {
		return utils.WrapError("Decode InterFreq_needForNCSG_r17", err)
	}
	return nil
}

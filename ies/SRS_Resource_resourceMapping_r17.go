package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_Resource_resourceMapping_r17 struct {
	StartPosition_r17    int64                                                 `lb:0,ub:13,madatory`
	NrofSymbols_r17      SRS_Resource_resourceMapping_r17_nrofSymbols_r17      `madatory`
	RepetitionFactor_r17 SRS_Resource_resourceMapping_r17_repetitionFactor_r17 `madatory`
}

func (ie *SRS_Resource_resourceMapping_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.StartPosition_r17, &uper.Constraint{Lb: 0, Ub: 13}, false); err != nil {
		return utils.WrapError("WriteInteger StartPosition_r17", err)
	}
	if err = ie.NrofSymbols_r17.Encode(w); err != nil {
		return utils.WrapError("Encode NrofSymbols_r17", err)
	}
	if err = ie.RepetitionFactor_r17.Encode(w); err != nil {
		return utils.WrapError("Encode RepetitionFactor_r17", err)
	}
	return nil
}

func (ie *SRS_Resource_resourceMapping_r17) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_StartPosition_r17 int64
	if tmp_int_StartPosition_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 13}, false); err != nil {
		return utils.WrapError("ReadInteger StartPosition_r17", err)
	}
	ie.StartPosition_r17 = tmp_int_StartPosition_r17
	if err = ie.NrofSymbols_r17.Decode(r); err != nil {
		return utils.WrapError("Decode NrofSymbols_r17", err)
	}
	if err = ie.RepetitionFactor_r17.Decode(r); err != nil {
		return utils.WrapError("Decode RepetitionFactor_r17", err)
	}
	return nil
}

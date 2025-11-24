package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_Resource_resourceMapping struct {
	StartPosition    int64                                         `lb:0,ub:5,madatory`
	NrofSymbols      SRS_Resource_resourceMapping_nrofSymbols      `madatory`
	RepetitionFactor SRS_Resource_resourceMapping_repetitionFactor `madatory`
}

func (ie *SRS_Resource_resourceMapping) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.StartPosition, &uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("WriteInteger StartPosition", err)
	}
	if err = ie.NrofSymbols.Encode(w); err != nil {
		return utils.WrapError("Encode NrofSymbols", err)
	}
	if err = ie.RepetitionFactor.Encode(w); err != nil {
		return utils.WrapError("Encode RepetitionFactor", err)
	}
	return nil
}

func (ie *SRS_Resource_resourceMapping) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_StartPosition int64
	if tmp_int_StartPosition, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("ReadInteger StartPosition", err)
	}
	ie.StartPosition = tmp_int_StartPosition
	if err = ie.NrofSymbols.Decode(r); err != nil {
		return utils.WrapError("Decode NrofSymbols", err)
	}
	if err = ie.RepetitionFactor.Decode(r); err != nil {
		return utils.WrapError("Decode RepetitionFactor", err)
	}
	return nil
}

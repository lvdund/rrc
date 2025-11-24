package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NAICS_Capability_Entry struct {
	NumberOfNAICS_CapableCC int64                                        `lb:1,ub:5,madatory`
	NumberOfAggregatedPRB   NAICS_Capability_Entry_numberOfAggregatedPRB `madatory`
}

func (ie *NAICS_Capability_Entry) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.NumberOfNAICS_CapableCC, &uper.Constraint{Lb: 1, Ub: 5}, false); err != nil {
		return utils.WrapError("WriteInteger NumberOfNAICS_CapableCC", err)
	}
	if err = ie.NumberOfAggregatedPRB.Encode(w); err != nil {
		return utils.WrapError("Encode NumberOfAggregatedPRB", err)
	}
	return nil
}

func (ie *NAICS_Capability_Entry) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_NumberOfNAICS_CapableCC int64
	if tmp_int_NumberOfNAICS_CapableCC, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 5}, false); err != nil {
		return utils.WrapError("ReadInteger NumberOfNAICS_CapableCC", err)
	}
	ie.NumberOfNAICS_CapableCC = tmp_int_NumberOfNAICS_CapableCC
	if err = ie.NumberOfAggregatedPRB.Decode(r); err != nil {
		return utils.WrapError("Decode NumberOfAggregatedPRB", err)
	}
	return nil
}

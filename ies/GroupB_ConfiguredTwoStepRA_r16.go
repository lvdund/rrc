package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type GroupB_ConfiguredTwoStepRA_r16 struct {
	Ra_MsgA_SizeGroupA         GroupB_ConfiguredTwoStepRA_r16_ra_MsgA_SizeGroupA       `madatory`
	MessagePowerOffsetGroupB   GroupB_ConfiguredTwoStepRA_r16_messagePowerOffsetGroupB `madatory`
	NumberOfRA_PreamblesGroupA int64                                                   `lb:1,ub:64,madatory`
}

func (ie *GroupB_ConfiguredTwoStepRA_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Ra_MsgA_SizeGroupA.Encode(w); err != nil {
		return utils.WrapError("Encode Ra_MsgA_SizeGroupA", err)
	}
	if err = ie.MessagePowerOffsetGroupB.Encode(w); err != nil {
		return utils.WrapError("Encode MessagePowerOffsetGroupB", err)
	}
	if err = w.WriteInteger(ie.NumberOfRA_PreamblesGroupA, &aper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return utils.WrapError("WriteInteger NumberOfRA_PreamblesGroupA", err)
	}
	return nil
}

func (ie *GroupB_ConfiguredTwoStepRA_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Ra_MsgA_SizeGroupA.Decode(r); err != nil {
		return utils.WrapError("Decode Ra_MsgA_SizeGroupA", err)
	}
	if err = ie.MessagePowerOffsetGroupB.Decode(r); err != nil {
		return utils.WrapError("Decode MessagePowerOffsetGroupB", err)
	}
	var tmp_int_NumberOfRA_PreamblesGroupA int64
	if tmp_int_NumberOfRA_PreamblesGroupA, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return utils.WrapError("ReadInteger NumberOfRA_PreamblesGroupA", err)
	}
	ie.NumberOfRA_PreamblesGroupA = tmp_int_NumberOfRA_PreamblesGroupA
	return nil
}

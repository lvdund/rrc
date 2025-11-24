package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_Resource struct {
	Pucch_ResourceId          PUCCH_ResourceId                          `madatory`
	StartingPRB               PRB_Id                                    `madatory`
	IntraSlotFrequencyHopping *PUCCH_Resource_intraSlotFrequencyHopping `optional`
	SecondHopPRB              *PRB_Id                                   `optional`
	Format                    PUCCH_Resource_format                     `madatory`
}

func (ie *PUCCH_Resource) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.IntraSlotFrequencyHopping != nil, ie.SecondHopPRB != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Pucch_ResourceId.Encode(w); err != nil {
		return utils.WrapError("Encode Pucch_ResourceId", err)
	}
	if err = ie.StartingPRB.Encode(w); err != nil {
		return utils.WrapError("Encode StartingPRB", err)
	}
	if ie.IntraSlotFrequencyHopping != nil {
		if err = ie.IntraSlotFrequencyHopping.Encode(w); err != nil {
			return utils.WrapError("Encode IntraSlotFrequencyHopping", err)
		}
	}
	if ie.SecondHopPRB != nil {
		if err = ie.SecondHopPRB.Encode(w); err != nil {
			return utils.WrapError("Encode SecondHopPRB", err)
		}
	}
	if err = ie.Format.Encode(w); err != nil {
		return utils.WrapError("Encode Format", err)
	}
	return nil
}

func (ie *PUCCH_Resource) Decode(r *uper.UperReader) error {
	var err error
	var IntraSlotFrequencyHoppingPresent bool
	if IntraSlotFrequencyHoppingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SecondHopPRBPresent bool
	if SecondHopPRBPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Pucch_ResourceId.Decode(r); err != nil {
		return utils.WrapError("Decode Pucch_ResourceId", err)
	}
	if err = ie.StartingPRB.Decode(r); err != nil {
		return utils.WrapError("Decode StartingPRB", err)
	}
	if IntraSlotFrequencyHoppingPresent {
		ie.IntraSlotFrequencyHopping = new(PUCCH_Resource_intraSlotFrequencyHopping)
		if err = ie.IntraSlotFrequencyHopping.Decode(r); err != nil {
			return utils.WrapError("Decode IntraSlotFrequencyHopping", err)
		}
	}
	if SecondHopPRBPresent {
		ie.SecondHopPRB = new(PRB_Id)
		if err = ie.SecondHopPRB.Decode(r); err != nil {
			return utils.WrapError("Decode SecondHopPRB", err)
		}
	}
	if err = ie.Format.Decode(r); err != nil {
		return utils.WrapError("Decode Format", err)
	}
	return nil
}

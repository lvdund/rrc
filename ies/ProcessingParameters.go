package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ProcessingParameters struct {
	Fallback            ProcessingParameters_fallback             `madatory`
	DifferentTB_PerSlot *ProcessingParameters_differentTB_PerSlot `optional`
}

func (ie *ProcessingParameters) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.DifferentTB_PerSlot != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Fallback.Encode(w); err != nil {
		return utils.WrapError("Encode Fallback", err)
	}
	if ie.DifferentTB_PerSlot != nil {
		if err = ie.DifferentTB_PerSlot.Encode(w); err != nil {
			return utils.WrapError("Encode DifferentTB_PerSlot", err)
		}
	}
	return nil
}

func (ie *ProcessingParameters) Decode(r *uper.UperReader) error {
	var err error
	var DifferentTB_PerSlotPresent bool
	if DifferentTB_PerSlotPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Fallback.Decode(r); err != nil {
		return utils.WrapError("Decode Fallback", err)
	}
	if DifferentTB_PerSlotPresent {
		ie.DifferentTB_PerSlot = new(ProcessingParameters_differentTB_PerSlot)
		if err = ie.DifferentTB_PerSlot.Decode(r); err != nil {
			return utils.WrapError("Decode DifferentTB_PerSlot", err)
		}
	}
	return nil
}

package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_Choice_nothing uint64 = iota
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_Choice_Scs15_r17
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_Choice_Scs30_r17
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_Choice_Scs60_r17
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_Choice_Scs120_r17
)

type NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17 struct {
	Choice     uint64
	Scs15_r17  *NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs15_r17
	Scs30_r17  *NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17
	Scs60_r17  *NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs60_r17
	Scs120_r17 *NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17
}

func (ie *NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_Choice_Scs15_r17:
		if err = ie.Scs15_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode Scs15_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_Choice_Scs30_r17:
		if err = ie.Scs30_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode Scs30_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_Choice_Scs60_r17:
		if err = ie.Scs60_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode Scs60_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_Choice_Scs120_r17:
		if err = ie.Scs120_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode Scs120_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_Choice_Scs15_r17:
		ie.Scs15_r17 = new(NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs15_r17)
		if err = ie.Scs15_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Scs15_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_Choice_Scs30_r17:
		ie.Scs30_r17 = new(NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17)
		if err = ie.Scs30_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Scs30_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_Choice_Scs60_r17:
		ie.Scs60_r17 = new(NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs60_r17)
		if err = ie.Scs60_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Scs60_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_Choice_Scs120_r17:
		ie.Scs120_r17 = new(NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17)
		if err = ie.Scs120_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Scs120_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}

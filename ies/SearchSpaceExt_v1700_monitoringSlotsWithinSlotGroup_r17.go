package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SearchSpaceExt_v1700_monitoringSlotsWithinSlotGroup_r17_Choice_nothing uint64 = iota
	SearchSpaceExt_v1700_monitoringSlotsWithinSlotGroup_r17_Choice_SlotGroupLength4_r17
	SearchSpaceExt_v1700_monitoringSlotsWithinSlotGroup_r17_Choice_SlotGroupLength8_r17
)

type SearchSpaceExt_v1700_monitoringSlotsWithinSlotGroup_r17 struct {
	Choice               uint64
	SlotGroupLength4_r17 uper.BitString `lb:4,ub:4,madatory`
	SlotGroupLength8_r17 uper.BitString `lb:8,ub:8,madatory`
}

func (ie *SearchSpaceExt_v1700_monitoringSlotsWithinSlotGroup_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SearchSpaceExt_v1700_monitoringSlotsWithinSlotGroup_r17_Choice_SlotGroupLength4_r17:
		if err = w.WriteBitString(ie.SlotGroupLength4_r17.Bytes, uint(ie.SlotGroupLength4_r17.NumBits), &uper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
			err = utils.WrapError("Encode SlotGroupLength4_r17", err)
		}
	case SearchSpaceExt_v1700_monitoringSlotsWithinSlotGroup_r17_Choice_SlotGroupLength8_r17:
		if err = w.WriteBitString(ie.SlotGroupLength8_r17.Bytes, uint(ie.SlotGroupLength8_r17.NumBits), &uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			err = utils.WrapError("Encode SlotGroupLength8_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SearchSpaceExt_v1700_monitoringSlotsWithinSlotGroup_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SearchSpaceExt_v1700_monitoringSlotsWithinSlotGroup_r17_Choice_SlotGroupLength4_r17:
		var tmp_bs_SlotGroupLength4_r17 []byte
		var n_SlotGroupLength4_r17 uint
		if tmp_bs_SlotGroupLength4_r17, n_SlotGroupLength4_r17, err = r.ReadBitString(&uper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode SlotGroupLength4_r17", err)
		}
		ie.SlotGroupLength4_r17 = uper.BitString{
			Bytes:   tmp_bs_SlotGroupLength4_r17,
			NumBits: uint64(n_SlotGroupLength4_r17),
		}
	case SearchSpaceExt_v1700_monitoringSlotsWithinSlotGroup_r17_Choice_SlotGroupLength8_r17:
		var tmp_bs_SlotGroupLength8_r17 []byte
		var n_SlotGroupLength8_r17 uint
		if tmp_bs_SlotGroupLength8_r17, n_SlotGroupLength8_r17, err = r.ReadBitString(&uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode SlotGroupLength8_r17", err)
		}
		ie.SlotGroupLength8_r17 = uper.BitString{
			Bytes:   tmp_bs_SlotGroupLength8_r17,
			NumBits: uint64(n_SlotGroupLength8_r17),
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}

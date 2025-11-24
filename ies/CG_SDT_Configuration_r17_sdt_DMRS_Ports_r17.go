package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CG_SDT_Configuration_r17_sdt_DMRS_Ports_r17_Choice_nothing uint64 = iota
	CG_SDT_Configuration_r17_sdt_DMRS_Ports_r17_Choice_DmrsType1_r17
	CG_SDT_Configuration_r17_sdt_DMRS_Ports_r17_Choice_DmrsType2_r17
)

type CG_SDT_Configuration_r17_sdt_DMRS_Ports_r17 struct {
	Choice        uint64
	DmrsType1_r17 uper.BitString `lb:8,ub:8,madatory`
	DmrsType2_r17 uper.BitString `lb:12,ub:12,madatory`
}

func (ie *CG_SDT_Configuration_r17_sdt_DMRS_Ports_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CG_SDT_Configuration_r17_sdt_DMRS_Ports_r17_Choice_DmrsType1_r17:
		if err = w.WriteBitString(ie.DmrsType1_r17.Bytes, uint(ie.DmrsType1_r17.NumBits), &uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			err = utils.WrapError("Encode DmrsType1_r17", err)
		}
	case CG_SDT_Configuration_r17_sdt_DMRS_Ports_r17_Choice_DmrsType2_r17:
		if err = w.WriteBitString(ie.DmrsType2_r17.Bytes, uint(ie.DmrsType2_r17.NumBits), &uper.Constraint{Lb: 12, Ub: 12}, false); err != nil {
			err = utils.WrapError("Encode DmrsType2_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CG_SDT_Configuration_r17_sdt_DMRS_Ports_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CG_SDT_Configuration_r17_sdt_DMRS_Ports_r17_Choice_DmrsType1_r17:
		var tmp_bs_DmrsType1_r17 []byte
		var n_DmrsType1_r17 uint
		if tmp_bs_DmrsType1_r17, n_DmrsType1_r17, err = r.ReadBitString(&uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode DmrsType1_r17", err)
		}
		ie.DmrsType1_r17 = uper.BitString{
			Bytes:   tmp_bs_DmrsType1_r17,
			NumBits: uint64(n_DmrsType1_r17),
		}
	case CG_SDT_Configuration_r17_sdt_DMRS_Ports_r17_Choice_DmrsType2_r17:
		var tmp_bs_DmrsType2_r17 []byte
		var n_DmrsType2_r17 uint
		if tmp_bs_DmrsType2_r17, n_DmrsType2_r17, err = r.ReadBitString(&uper.Constraint{Lb: 12, Ub: 12}, false); err != nil {
			return utils.WrapError("Decode DmrsType2_r17", err)
		}
		ie.DmrsType2_r17 = uper.BitString{
			Bytes:   tmp_bs_DmrsType2_r17,
			NumBits: uint64(n_DmrsType2_r17),
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}

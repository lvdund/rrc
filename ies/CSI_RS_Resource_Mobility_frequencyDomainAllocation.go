package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_RS_Resource_Mobility_frequencyDomainAllocation_Choice_nothing uint64 = iota
	CSI_RS_Resource_Mobility_frequencyDomainAllocation_Choice_Row1
	CSI_RS_Resource_Mobility_frequencyDomainAllocation_Choice_Row2
)

type CSI_RS_Resource_Mobility_frequencyDomainAllocation struct {
	Choice uint64
	Row1   uper.BitString `lb:4,ub:4,madatory`
	Row2   uper.BitString `lb:12,ub:12,madatory`
}

func (ie *CSI_RS_Resource_Mobility_frequencyDomainAllocation) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_RS_Resource_Mobility_frequencyDomainAllocation_Choice_Row1:
		if err = w.WriteBitString(ie.Row1.Bytes, uint(ie.Row1.NumBits), &uper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
			err = utils.WrapError("Encode Row1", err)
		}
	case CSI_RS_Resource_Mobility_frequencyDomainAllocation_Choice_Row2:
		if err = w.WriteBitString(ie.Row2.Bytes, uint(ie.Row2.NumBits), &uper.Constraint{Lb: 12, Ub: 12}, false); err != nil {
			err = utils.WrapError("Encode Row2", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CSI_RS_Resource_Mobility_frequencyDomainAllocation) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_RS_Resource_Mobility_frequencyDomainAllocation_Choice_Row1:
		var tmp_bs_Row1 []byte
		var n_Row1 uint
		if tmp_bs_Row1, n_Row1, err = r.ReadBitString(&uper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode Row1", err)
		}
		ie.Row1 = uper.BitString{
			Bytes:   tmp_bs_Row1,
			NumBits: uint64(n_Row1),
		}
	case CSI_RS_Resource_Mobility_frequencyDomainAllocation_Choice_Row2:
		var tmp_bs_Row2 []byte
		var n_Row2 uint
		if tmp_bs_Row2, n_Row2, err = r.ReadBitString(&uper.Constraint{Lb: 12, Ub: 12}, false); err != nil {
			return utils.WrapError("Decode Row2", err)
		}
		ie.Row2 = uper.BitString{
			Bytes:   tmp_bs_Row2,
			NumBits: uint64(n_Row2),
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}

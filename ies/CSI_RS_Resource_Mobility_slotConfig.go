package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_RS_Resource_Mobility_slotConfig_Choice_nothing uint64 = iota
	CSI_RS_Resource_Mobility_slotConfig_Choice_Ms4
	CSI_RS_Resource_Mobility_slotConfig_Choice_Ms5
	CSI_RS_Resource_Mobility_slotConfig_Choice_Ms10
	CSI_RS_Resource_Mobility_slotConfig_Choice_Ms20
	CSI_RS_Resource_Mobility_slotConfig_Choice_Ms40
)

type CSI_RS_Resource_Mobility_slotConfig struct {
	Choice uint64
	Ms4    int64 `lb:0,ub:31,madatory`
	Ms5    int64 `lb:0,ub:39,madatory`
	Ms10   int64 `lb:0,ub:79,madatory`
	Ms20   int64 `lb:0,ub:159,madatory`
	Ms40   int64 `lb:0,ub:319,madatory`
}

func (ie *CSI_RS_Resource_Mobility_slotConfig) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 5, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_RS_Resource_Mobility_slotConfig_Choice_Ms4:
		if err = w.WriteInteger(int64(ie.Ms4), &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			err = utils.WrapError("Encode Ms4", err)
		}
	case CSI_RS_Resource_Mobility_slotConfig_Choice_Ms5:
		if err = w.WriteInteger(int64(ie.Ms5), &uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			err = utils.WrapError("Encode Ms5", err)
		}
	case CSI_RS_Resource_Mobility_slotConfig_Choice_Ms10:
		if err = w.WriteInteger(int64(ie.Ms10), &uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			err = utils.WrapError("Encode Ms10", err)
		}
	case CSI_RS_Resource_Mobility_slotConfig_Choice_Ms20:
		if err = w.WriteInteger(int64(ie.Ms20), &uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			err = utils.WrapError("Encode Ms20", err)
		}
	case CSI_RS_Resource_Mobility_slotConfig_Choice_Ms40:
		if err = w.WriteInteger(int64(ie.Ms40), &uper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			err = utils.WrapError("Encode Ms40", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CSI_RS_Resource_Mobility_slotConfig) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(5, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_RS_Resource_Mobility_slotConfig_Choice_Ms4:
		var tmp_int_Ms4 int64
		if tmp_int_Ms4, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode Ms4", err)
		}
		ie.Ms4 = tmp_int_Ms4
	case CSI_RS_Resource_Mobility_slotConfig_Choice_Ms5:
		var tmp_int_Ms5 int64
		if tmp_int_Ms5, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			return utils.WrapError("Decode Ms5", err)
		}
		ie.Ms5 = tmp_int_Ms5
	case CSI_RS_Resource_Mobility_slotConfig_Choice_Ms10:
		var tmp_int_Ms10 int64
		if tmp_int_Ms10, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			return utils.WrapError("Decode Ms10", err)
		}
		ie.Ms10 = tmp_int_Ms10
	case CSI_RS_Resource_Mobility_slotConfig_Choice_Ms20:
		var tmp_int_Ms20 int64
		if tmp_int_Ms20, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			return utils.WrapError("Decode Ms20", err)
		}
		ie.Ms20 = tmp_int_Ms20
	case CSI_RS_Resource_Mobility_slotConfig_Choice_Ms40:
		var tmp_int_Ms40 int64
		if tmp_int_Ms40, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			return utils.WrapError("Decode Ms40", err)
		}
		ie.Ms40 = tmp_int_Ms40
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}

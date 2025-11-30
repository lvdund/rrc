package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUCCH_Resource_format_Choice_nothing uint64 = iota
	PUCCH_Resource_format_Choice_Format0
	PUCCH_Resource_format_Choice_Format1
	PUCCH_Resource_format_Choice_Format2
	PUCCH_Resource_format_Choice_Format3
	PUCCH_Resource_format_Choice_Format4
)

type PUCCH_Resource_format struct {
	Choice  uint64
	Format0 *PUCCH_format0
	Format1 *PUCCH_format1
	Format2 *PUCCH_format2
	Format3 *PUCCH_format3
	Format4 *PUCCH_format4
}

func (ie *PUCCH_Resource_format) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 5, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PUCCH_Resource_format_Choice_Format0:
		if err = ie.Format0.Encode(w); err != nil {
			err = utils.WrapError("Encode Format0", err)
		}
	case PUCCH_Resource_format_Choice_Format1:
		if err = ie.Format1.Encode(w); err != nil {
			err = utils.WrapError("Encode Format1", err)
		}
	case PUCCH_Resource_format_Choice_Format2:
		if err = ie.Format2.Encode(w); err != nil {
			err = utils.WrapError("Encode Format2", err)
		}
	case PUCCH_Resource_format_Choice_Format3:
		if err = ie.Format3.Encode(w); err != nil {
			err = utils.WrapError("Encode Format3", err)
		}
	case PUCCH_Resource_format_Choice_Format4:
		if err = ie.Format4.Encode(w); err != nil {
			err = utils.WrapError("Encode Format4", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *PUCCH_Resource_format) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(5, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PUCCH_Resource_format_Choice_Format0:
		ie.Format0 = new(PUCCH_format0)
		if err = ie.Format0.Decode(r); err != nil {
			return utils.WrapError("Decode Format0", err)
		}
	case PUCCH_Resource_format_Choice_Format1:
		ie.Format1 = new(PUCCH_format1)
		if err = ie.Format1.Decode(r); err != nil {
			return utils.WrapError("Decode Format1", err)
		}
	case PUCCH_Resource_format_Choice_Format2:
		ie.Format2 = new(PUCCH_format2)
		if err = ie.Format2.Decode(r); err != nil {
			return utils.WrapError("Decode Format2", err)
		}
	case PUCCH_Resource_format_Choice_Format3:
		ie.Format3 = new(PUCCH_format3)
		if err = ie.Format3.Decode(r); err != nil {
			return utils.WrapError("Decode Format3", err)
		}
	case PUCCH_Resource_format_Choice_Format4:
		ie.Format4 = new(PUCCH_format4)
		if err = ie.Format4.Decode(r); err != nil {
			return utils.WrapError("Decode Format4", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}

package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CellsTriggeredItem_Choice_nothing uint64 = iota
	CellsTriggeredItem_Choice_PhysCellId
	CellsTriggeredItem_Choice_PhysCellIdEUTRA
	CellsTriggeredItem_Choice_PhysCellIdUTRA_FDD_r16
)

type CellsTriggeredItem struct {
	Choice                 uint64
	PhysCellId             *PhysCellId
	PhysCellIdEUTRA        *EUTRA_PhysCellId
	PhysCellIdUTRA_FDD_r16 *PhysCellIdUTRA_FDD_r16
}

func (ie *CellsTriggeredItem) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CellsTriggeredItem_Choice_PhysCellId:
		if err = ie.PhysCellId.Encode(w); err != nil {
			err = utils.WrapError("Encode PhysCellId", err)
		}
	case CellsTriggeredItem_Choice_PhysCellIdEUTRA:
		if err = ie.PhysCellIdEUTRA.Encode(w); err != nil {
			err = utils.WrapError("Encode PhysCellIdEUTRA", err)
		}
	case CellsTriggeredItem_Choice_PhysCellIdUTRA_FDD_r16:
		if err = ie.PhysCellIdUTRA_FDD_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode PhysCellIdUTRA_FDD_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CellsTriggeredItem) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CellsTriggeredItem_Choice_PhysCellId:
		ie.PhysCellId = new(PhysCellId)
		if err = ie.PhysCellId.Decode(r); err != nil {
			return utils.WrapError("Decode PhysCellId", err)
		}
	case CellsTriggeredItem_Choice_PhysCellIdEUTRA:
		ie.PhysCellIdEUTRA = new(EUTRA_PhysCellId)
		if err = ie.PhysCellIdEUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode PhysCellIdEUTRA", err)
		}
	case CellsTriggeredItem_Choice_PhysCellIdUTRA_FDD_r16:
		ie.PhysCellIdUTRA_FDD_r16 = new(PhysCellIdUTRA_FDD_r16)
		if err = ie.PhysCellIdUTRA_FDD_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PhysCellIdUTRA_FDD_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}

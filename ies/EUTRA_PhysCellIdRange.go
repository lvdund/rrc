package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EUTRA_PhysCellIdRange struct {
	Start      EUTRA_PhysCellId                  `madatory`
	Range_cell *EUTRA_PhysCellIdRange_range_cell `optional`
}

func (ie *EUTRA_PhysCellIdRange) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Range_cell != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Start.Encode(w); err != nil {
		return utils.WrapError("Encode Start", err)
	}
	if ie.Range_cell != nil {
		if err = ie.Range_cell.Encode(w); err != nil {
			return utils.WrapError("Encode Range_cell", err)
		}
	}
	return nil
}

func (ie *EUTRA_PhysCellIdRange) Decode(r *uper.UperReader) error {
	var err error
	var Range_cellPresent bool
	if Range_cellPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Start.Decode(r); err != nil {
		return utils.WrapError("Decode Start", err)
	}
	if Range_cellPresent {
		ie.Range_cell = new(EUTRA_PhysCellIdRange_range_cell)
		if err = ie.Range_cell.Decode(r); err != nil {
			return utils.WrapError("Decode Range_cell", err)
		}
	}
	return nil
}

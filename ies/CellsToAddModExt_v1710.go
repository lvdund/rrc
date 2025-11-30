package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CellsToAddModExt_v1710 struct {
	Ntn_PolarizationDL_r17 *CellsToAddModExt_v1710_ntn_PolarizationDL_r17 `optional`
	Ntn_PolarizationUL_r17 *CellsToAddModExt_v1710_ntn_PolarizationUL_r17 `optional`
}

func (ie *CellsToAddModExt_v1710) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Ntn_PolarizationDL_r17 != nil, ie.Ntn_PolarizationUL_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ntn_PolarizationDL_r17 != nil {
		if err = ie.Ntn_PolarizationDL_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ntn_PolarizationDL_r17", err)
		}
	}
	if ie.Ntn_PolarizationUL_r17 != nil {
		if err = ie.Ntn_PolarizationUL_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ntn_PolarizationUL_r17", err)
		}
	}
	return nil
}

func (ie *CellsToAddModExt_v1710) Decode(r *aper.AperReader) error {
	var err error
	var Ntn_PolarizationDL_r17Present bool
	if Ntn_PolarizationDL_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ntn_PolarizationUL_r17Present bool
	if Ntn_PolarizationUL_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Ntn_PolarizationDL_r17Present {
		ie.Ntn_PolarizationDL_r17 = new(CellsToAddModExt_v1710_ntn_PolarizationDL_r17)
		if err = ie.Ntn_PolarizationDL_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ntn_PolarizationDL_r17", err)
		}
	}
	if Ntn_PolarizationUL_r17Present {
		ie.Ntn_PolarizationUL_r17 = new(CellsToAddModExt_v1710_ntn_PolarizationUL_r17)
		if err = ie.Ntn_PolarizationUL_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ntn_PolarizationUL_r17", err)
		}
	}
	return nil
}

package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type FR_Info struct {
	ServCellIndex ServCellIndex   `madatory`
	Fr_Type       FR_Info_fr_Type `madatory`
}

func (ie *FR_Info) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.ServCellIndex.Encode(w); err != nil {
		return utils.WrapError("Encode ServCellIndex", err)
	}
	if err = ie.Fr_Type.Encode(w); err != nil {
		return utils.WrapError("Encode Fr_Type", err)
	}
	return nil
}

func (ie *FR_Info) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.ServCellIndex.Decode(r); err != nil {
		return utils.WrapError("Decode ServCellIndex", err)
	}
	if err = ie.Fr_Type.Decode(r); err != nil {
		return utils.WrapError("Decode Fr_Type", err)
	}
	return nil
}

package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type ServingCellAndBWP_Id_r17 struct {
	Servingcell_r17 ServCellIndex `madatory`
	Bwp_r17         BWP_Id        `madatory`
}

func (ie *ServingCellAndBWP_Id_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Servingcell_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Servingcell_r17", err)
	}
	if err = ie.Bwp_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Bwp_r17", err)
	}
	return nil
}

func (ie *ServingCellAndBWP_Id_r17) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Servingcell_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Servingcell_r17", err)
	}
	if err = ie.Bwp_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Bwp_r17", err)
	}
	return nil
}

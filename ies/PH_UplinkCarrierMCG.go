package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PH_UplinkCarrierMCG struct {
	Ph_Type1or3 PH_UplinkCarrierMCG_ph_Type1or3 `madatory`
}

func (ie *PH_UplinkCarrierMCG) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Ph_Type1or3.Encode(w); err != nil {
		return utils.WrapError("Encode Ph_Type1or3", err)
	}
	return nil
}

func (ie *PH_UplinkCarrierMCG) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Ph_Type1or3.Decode(r); err != nil {
		return utils.WrapError("Decode Ph_Type1or3", err)
	}
	return nil
}

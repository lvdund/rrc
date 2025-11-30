package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type DL_AM_RLC_v1610 struct {
	T_StatusProhibit_v1610 *T_StatusProhibit_v1610 `optional`
}

func (ie *DL_AM_RLC_v1610) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.T_StatusProhibit_v1610 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.T_StatusProhibit_v1610 != nil {
		if err = ie.T_StatusProhibit_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode T_StatusProhibit_v1610", err)
		}
	}
	return nil
}

func (ie *DL_AM_RLC_v1610) Decode(r *aper.AperReader) error {
	var err error
	var T_StatusProhibit_v1610Present bool
	if T_StatusProhibit_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	if T_StatusProhibit_v1610Present {
		ie.T_StatusProhibit_v1610 = new(T_StatusProhibit_v1610)
		if err = ie.T_StatusProhibit_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode T_StatusProhibit_v1610", err)
		}
	}
	return nil
}

package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MRB_PDCP_ConfigBroadcast_r17 struct {
	Pdcp_SN_SizeDL_r17    *MRB_PDCP_ConfigBroadcast_r17_pdcp_SN_SizeDL_r17   `optional`
	HeaderCompression_r17 MRB_PDCP_ConfigBroadcast_r17_headerCompression_r17 `lb:1,ub:16,madatory`
	T_Reordering_r17      *MRB_PDCP_ConfigBroadcast_r17_t_Reordering_r17     `optional`
}

func (ie *MRB_PDCP_ConfigBroadcast_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Pdcp_SN_SizeDL_r17 != nil, ie.T_Reordering_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Pdcp_SN_SizeDL_r17 != nil {
		if err = ie.Pdcp_SN_SizeDL_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Pdcp_SN_SizeDL_r17", err)
		}
	}
	if err = ie.HeaderCompression_r17.Encode(w); err != nil {
		return utils.WrapError("Encode HeaderCompression_r17", err)
	}
	if ie.T_Reordering_r17 != nil {
		if err = ie.T_Reordering_r17.Encode(w); err != nil {
			return utils.WrapError("Encode T_Reordering_r17", err)
		}
	}
	return nil
}

func (ie *MRB_PDCP_ConfigBroadcast_r17) Decode(r *aper.AperReader) error {
	var err error
	var Pdcp_SN_SizeDL_r17Present bool
	if Pdcp_SN_SizeDL_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var T_Reordering_r17Present bool
	if T_Reordering_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Pdcp_SN_SizeDL_r17Present {
		ie.Pdcp_SN_SizeDL_r17 = new(MRB_PDCP_ConfigBroadcast_r17_pdcp_SN_SizeDL_r17)
		if err = ie.Pdcp_SN_SizeDL_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Pdcp_SN_SizeDL_r17", err)
		}
	}
	if err = ie.HeaderCompression_r17.Decode(r); err != nil {
		return utils.WrapError("Decode HeaderCompression_r17", err)
	}
	if T_Reordering_r17Present {
		ie.T_Reordering_r17 = new(MRB_PDCP_ConfigBroadcast_r17_t_Reordering_r17)
		if err = ie.T_Reordering_r17.Decode(r); err != nil {
			return utils.WrapError("Decode T_Reordering_r17", err)
		}
	}
	return nil
}

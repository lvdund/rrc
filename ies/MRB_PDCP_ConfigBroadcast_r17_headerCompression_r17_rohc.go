package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MRB_PDCP_ConfigBroadcast_r17_headerCompression_r17_rohc struct {
	MaxCID_r17   int64                                                                `lb:1,ub:16,madatory`
	Profiles_r17 MRB_PDCP_ConfigBroadcast_r17_headerCompression_r17_rohc_profiles_r17 `madatory`
}

func (ie *MRB_PDCP_ConfigBroadcast_r17_headerCompression_r17_rohc) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(ie.MaxCID_r17, &aper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("WriteInteger MaxCID_r17", err)
	}
	if err = ie.Profiles_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Profiles_r17", err)
	}
	return nil
}

func (ie *MRB_PDCP_ConfigBroadcast_r17_headerCompression_r17_rohc) Decode(r *aper.AperReader) error {
	var err error
	var tmp_int_MaxCID_r17 int64
	if tmp_int_MaxCID_r17, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("ReadInteger MaxCID_r17", err)
	}
	ie.MaxCID_r17 = tmp_int_MaxCID_r17
	if err = ie.Profiles_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Profiles_r17", err)
	}
	return nil
}

package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MRB_RLC_ConfigBroadcast_r17 struct {
	LogicalChannelIdentity_r17 LogicalChannelIdentity                          `madatory`
	Sn_FieldLength_r17         *MRB_RLC_ConfigBroadcast_r17_sn_FieldLength_r17 `optional`
	T_Reassembly_r17           *T_Reassembly                                   `optional`
}

func (ie *MRB_RLC_ConfigBroadcast_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Sn_FieldLength_r17 != nil, ie.T_Reassembly_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.LogicalChannelIdentity_r17.Encode(w); err != nil {
		return utils.WrapError("Encode LogicalChannelIdentity_r17", err)
	}
	if ie.Sn_FieldLength_r17 != nil {
		if err = ie.Sn_FieldLength_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sn_FieldLength_r17", err)
		}
	}
	if ie.T_Reassembly_r17 != nil {
		if err = ie.T_Reassembly_r17.Encode(w); err != nil {
			return utils.WrapError("Encode T_Reassembly_r17", err)
		}
	}
	return nil
}

func (ie *MRB_RLC_ConfigBroadcast_r17) Decode(r *aper.AperReader) error {
	var err error
	var Sn_FieldLength_r17Present bool
	if Sn_FieldLength_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var T_Reassembly_r17Present bool
	if T_Reassembly_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.LogicalChannelIdentity_r17.Decode(r); err != nil {
		return utils.WrapError("Decode LogicalChannelIdentity_r17", err)
	}
	if Sn_FieldLength_r17Present {
		ie.Sn_FieldLength_r17 = new(MRB_RLC_ConfigBroadcast_r17_sn_FieldLength_r17)
		if err = ie.Sn_FieldLength_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sn_FieldLength_r17", err)
		}
	}
	if T_Reassembly_r17Present {
		ie.T_Reassembly_r17 = new(T_Reassembly)
		if err = ie.T_Reassembly_r17.Decode(r); err != nil {
			return utils.WrapError("Decode T_Reassembly_r17", err)
		}
	}
	return nil
}

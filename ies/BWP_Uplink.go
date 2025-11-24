package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BWP_Uplink struct {
	Bwp_Id        BWP_Id               `madatory`
	Bwp_Common    *BWP_UplinkCommon    `optional`
	Bwp_Dedicated *BWP_UplinkDedicated `optional`
}

func (ie *BWP_Uplink) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Bwp_Common != nil, ie.Bwp_Dedicated != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Bwp_Id.Encode(w); err != nil {
		return utils.WrapError("Encode Bwp_Id", err)
	}
	if ie.Bwp_Common != nil {
		if err = ie.Bwp_Common.Encode(w); err != nil {
			return utils.WrapError("Encode Bwp_Common", err)
		}
	}
	if ie.Bwp_Dedicated != nil {
		if err = ie.Bwp_Dedicated.Encode(w); err != nil {
			return utils.WrapError("Encode Bwp_Dedicated", err)
		}
	}
	return nil
}

func (ie *BWP_Uplink) Decode(r *uper.UperReader) error {
	var err error
	var Bwp_CommonPresent bool
	if Bwp_CommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Bwp_DedicatedPresent bool
	if Bwp_DedicatedPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Bwp_Id.Decode(r); err != nil {
		return utils.WrapError("Decode Bwp_Id", err)
	}
	if Bwp_CommonPresent {
		ie.Bwp_Common = new(BWP_UplinkCommon)
		if err = ie.Bwp_Common.Decode(r); err != nil {
			return utils.WrapError("Decode Bwp_Common", err)
		}
	}
	if Bwp_DedicatedPresent {
		ie.Bwp_Dedicated = new(BWP_UplinkDedicated)
		if err = ie.Bwp_Dedicated.Decode(r); err != nil {
			return utils.WrapError("Decode Bwp_Dedicated", err)
		}
	}
	return nil
}

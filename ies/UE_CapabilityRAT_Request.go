package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UE_CapabilityRAT_Request struct {
	Rat_Type                RAT_Type `madatory`
	CapabilityRequestFilter *[]byte  `optional`
}

func (ie *UE_CapabilityRAT_Request) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.CapabilityRequestFilter != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Rat_Type.Encode(w); err != nil {
		return utils.WrapError("Encode Rat_Type", err)
	}
	if ie.CapabilityRequestFilter != nil {
		if err = w.WriteOctetString(*ie.CapabilityRequestFilter, nil, false); err != nil {
			return utils.WrapError("Encode CapabilityRequestFilter", err)
		}
	}
	return nil
}

func (ie *UE_CapabilityRAT_Request) Decode(r *aper.AperReader) error {
	var err error
	var CapabilityRequestFilterPresent bool
	if CapabilityRequestFilterPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Rat_Type.Decode(r); err != nil {
		return utils.WrapError("Decode Rat_Type", err)
	}
	if CapabilityRequestFilterPresent {
		var tmp_os_CapabilityRequestFilter []byte
		if tmp_os_CapabilityRequestFilter, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode CapabilityRequestFilter", err)
		}
		ie.CapabilityRequestFilter = &tmp_os_CapabilityRequestFilter
	}
	return nil
}

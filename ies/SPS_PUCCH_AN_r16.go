package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SPS_PUCCH_AN_r16 struct {
	Sps_PUCCH_AN_ResourceID_r16 PUCCH_ResourceId `madatory`
	MaxPayloadSize_r16          *int64           `lb:4,ub:256,optional`
}

func (ie *SPS_PUCCH_AN_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.MaxPayloadSize_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Sps_PUCCH_AN_ResourceID_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sps_PUCCH_AN_ResourceID_r16", err)
	}
	if ie.MaxPayloadSize_r16 != nil {
		if err = w.WriteInteger(*ie.MaxPayloadSize_r16, &uper.Constraint{Lb: 4, Ub: 256}, false); err != nil {
			return utils.WrapError("Encode MaxPayloadSize_r16", err)
		}
	}
	return nil
}

func (ie *SPS_PUCCH_AN_r16) Decode(r *uper.UperReader) error {
	var err error
	var MaxPayloadSize_r16Present bool
	if MaxPayloadSize_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Sps_PUCCH_AN_ResourceID_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sps_PUCCH_AN_ResourceID_r16", err)
	}
	if MaxPayloadSize_r16Present {
		var tmp_int_MaxPayloadSize_r16 int64
		if tmp_int_MaxPayloadSize_r16, err = r.ReadInteger(&uper.Constraint{Lb: 4, Ub: 256}, false); err != nil {
			return utils.WrapError("Decode MaxPayloadSize_r16", err)
		}
		ie.MaxPayloadSize_r16 = &tmp_int_MaxPayloadSize_r16
	}
	return nil
}

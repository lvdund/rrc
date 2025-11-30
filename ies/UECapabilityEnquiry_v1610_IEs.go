package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UECapabilityEnquiry_v1610_IEs struct {
	Rrc_SegAllowed_r16   *UECapabilityEnquiry_v1610_IEs_rrc_SegAllowed_r16 `optional`
	NonCriticalExtension interface{}                                       `optional`
}

func (ie *UECapabilityEnquiry_v1610_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Rrc_SegAllowed_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Rrc_SegAllowed_r16 != nil {
		if err = ie.Rrc_SegAllowed_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Rrc_SegAllowed_r16", err)
		}
	}
	return nil
}

func (ie *UECapabilityEnquiry_v1610_IEs) Decode(r *aper.AperReader) error {
	var err error
	var Rrc_SegAllowed_r16Present bool
	if Rrc_SegAllowed_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Rrc_SegAllowed_r16Present {
		ie.Rrc_SegAllowed_r16 = new(UECapabilityEnquiry_v1610_IEs_rrc_SegAllowed_r16)
		if err = ie.Rrc_SegAllowed_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Rrc_SegAllowed_r16", err)
		}
	}
	return nil
}

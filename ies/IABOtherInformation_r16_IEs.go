package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type IABOtherInformation_r16_IEs struct {
	Ip_InfoType_r16          *IABOtherInformation_r16_IEs_ip_InfoType_r16 `optional`
	LateNonCriticalExtension *[]byte                                      `optional,ext`
	NonCriticalExtension     interface{}                                  `optional,ext`
}

func (ie *IABOtherInformation_r16_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Ip_InfoType_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ip_InfoType_r16 != nil {
		if err = ie.Ip_InfoType_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ip_InfoType_r16", err)
		}
	}
	return nil
}

func (ie *IABOtherInformation_r16_IEs) Decode(r *aper.AperReader) error {
	var err error
	var Ip_InfoType_r16Present bool
	if Ip_InfoType_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Ip_InfoType_r16Present {
		ie.Ip_InfoType_r16 = new(IABOtherInformation_r16_IEs_ip_InfoType_r16)
		if err = ie.Ip_InfoType_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ip_InfoType_r16", err)
		}
	}
	return nil
}

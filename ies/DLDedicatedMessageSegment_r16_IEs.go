package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type DLDedicatedMessageSegment_r16_IEs struct {
	SegmentNumber_r16               int64                                                        `lb:0,ub:4,madatory`
	Rrc_MessageSegmentContainer_r16 []byte                                                       `madatory`
	Rrc_MessageSegmentType_r16      DLDedicatedMessageSegment_r16_IEs_rrc_MessageSegmentType_r16 `madatory`
	LateNonCriticalExtension        *[]byte                                                      `optional`
	NonCriticalExtension            interface{}                                                  `optional`
}

func (ie *DLDedicatedMessageSegment_r16_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.LateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.SegmentNumber_r16, &aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger SegmentNumber_r16", err)
	}
	if err = w.WriteOctetString(ie.Rrc_MessageSegmentContainer_r16, nil, false); err != nil {
		return utils.WrapError("WriteOctetString Rrc_MessageSegmentContainer_r16", err)
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, nil, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *DLDedicatedMessageSegment_r16_IEs) Decode(r *aper.AperReader) error {
	var err error
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_SegmentNumber_r16 int64
	if tmp_int_SegmentNumber_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger SegmentNumber_r16", err)
	}
	ie.SegmentNumber_r16 = tmp_int_SegmentNumber_r16
	var tmp_os_Rrc_MessageSegmentContainer_r16 []byte
	if tmp_os_Rrc_MessageSegmentContainer_r16, err = r.ReadOctetString(nil, false); err != nil {
		return utils.WrapError("ReadOctetString Rrc_MessageSegmentContainer_r16", err)
	}
	ie.Rrc_MessageSegmentContainer_r16 = tmp_os_Rrc_MessageSegmentContainer_r16
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}
	return nil
}

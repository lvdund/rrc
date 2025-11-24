package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ULDedicatedMessageSegment_r16_IEs struct {
	SegmentNumber_r16               int64                                                        `lb:0,ub:15,madatory`
	Rrc_MessageSegmentContainer_r16 []byte                                                       `madatory`
	Rrc_MessageSegmentType_r16      ULDedicatedMessageSegment_r16_IEs_rrc_MessageSegmentType_r16 `madatory`
	LateNonCriticalExtension        *[]byte                                                      `optional`
	NonCriticalExtension            interface{}                                                  `optional`
}

func (ie *ULDedicatedMessageSegment_r16_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.LateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.SegmentNumber_r16, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger SegmentNumber_r16", err)
	}
	if err = w.WriteOctetString(ie.Rrc_MessageSegmentContainer_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("WriteOctetString Rrc_MessageSegmentContainer_r16", err)
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *ULDedicatedMessageSegment_r16_IEs) Decode(r *uper.UperReader) error {
	var err error
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_SegmentNumber_r16 int64
	if tmp_int_SegmentNumber_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger SegmentNumber_r16", err)
	}
	ie.SegmentNumber_r16 = tmp_int_SegmentNumber_r16
	var tmp_os_Rrc_MessageSegmentContainer_r16 []byte
	if tmp_os_Rrc_MessageSegmentContainer_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("ReadOctetString Rrc_MessageSegmentContainer_r16", err)
	}
	ie.Rrc_MessageSegmentContainer_r16 = tmp_os_Rrc_MessageSegmentContainer_r16
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}
	return nil
}

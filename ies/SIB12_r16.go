package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SIB12_r16 struct {
	SegmentNumber_r16    int64                     `lb:0,ub:63,madatory`
	SegmentType_r16      SIB12_r16_segmentType_r16 `madatory`
	SegmentContainer_r16 []byte                    `madatory`
}

func (ie *SIB12_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(ie.SegmentNumber_r16, &aper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("WriteInteger SegmentNumber_r16", err)
	}
	if err = ie.SegmentType_r16.Encode(w); err != nil {
		return utils.WrapError("Encode SegmentType_r16", err)
	}
	if err = w.WriteOctetString(ie.SegmentContainer_r16, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("WriteOctetString SegmentContainer_r16", err)
	}
	return nil
}

func (ie *SIB12_r16) Decode(r *aper.AperReader) error {
	var err error
	var tmp_int_SegmentNumber_r16 int64
	if tmp_int_SegmentNumber_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("ReadInteger SegmentNumber_r16", err)
	}
	ie.SegmentNumber_r16 = tmp_int_SegmentNumber_r16
	if err = ie.SegmentType_r16.Decode(r); err != nil {
		return utils.WrapError("Decode SegmentType_r16", err)
	}
	var tmp_os_SegmentContainer_r16 []byte
	if tmp_os_SegmentContainer_r16, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("ReadOctetString SegmentContainer_r16", err)
	}
	ie.SegmentContainer_r16 = tmp_os_SegmentContainer_r16
	return nil
}

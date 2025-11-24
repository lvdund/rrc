package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB17_r17 struct {
	SegmentNumber_r17    int64                     `lb:0,ub:63,madatory`
	SegmentType_r17      SIB17_r17_segmentType_r17 `madatory`
	SegmentContainer_r17 []byte                    `madatory`
}

func (ie *SIB17_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.SegmentNumber_r17, &uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("WriteInteger SegmentNumber_r17", err)
	}
	if err = ie.SegmentType_r17.Encode(w); err != nil {
		return utils.WrapError("Encode SegmentType_r17", err)
	}
	if err = w.WriteOctetString(ie.SegmentContainer_r17, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("WriteOctetString SegmentContainer_r17", err)
	}
	return nil
}

func (ie *SIB17_r17) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_SegmentNumber_r17 int64
	if tmp_int_SegmentNumber_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("ReadInteger SegmentNumber_r17", err)
	}
	ie.SegmentNumber_r17 = tmp_int_SegmentNumber_r17
	if err = ie.SegmentType_r17.Decode(r); err != nil {
		return utils.WrapError("Decode SegmentType_r17", err)
	}
	var tmp_os_SegmentContainer_r17 []byte
	if tmp_os_SegmentContainer_r17, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("ReadOctetString SegmentContainer_r17", err)
	}
	ie.SegmentContainer_r17 = tmp_os_SegmentContainer_r17
	return nil
}

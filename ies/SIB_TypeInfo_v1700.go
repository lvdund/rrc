package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SIB_TypeInfo_v1700 struct {
	SibType_r17   *SIB_TypeInfo_v1700_sibType_r17   `optional`
	ValueTag_r17  *int64                            `lb:0,ub:31,optional`
	AreaScope_r17 *SIB_TypeInfo_v1700_areaScope_r17 `optional`
}

func (ie *SIB_TypeInfo_v1700) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.SibType_r17 != nil, ie.ValueTag_r17 != nil, ie.AreaScope_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SibType_r17 != nil {
		if err = ie.SibType_r17.Encode(w); err != nil {
			return utils.WrapError("Encode SibType_r17", err)
		}
	}
	if ie.ValueTag_r17 != nil {
		if err = w.WriteInteger(*ie.ValueTag_r17, &aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Encode ValueTag_r17", err)
		}
	}
	if ie.AreaScope_r17 != nil {
		if err = ie.AreaScope_r17.Encode(w); err != nil {
			return utils.WrapError("Encode AreaScope_r17", err)
		}
	}
	return nil
}

func (ie *SIB_TypeInfo_v1700) Decode(r *aper.AperReader) error {
	var err error
	var SibType_r17Present bool
	if SibType_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ValueTag_r17Present bool
	if ValueTag_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var AreaScope_r17Present bool
	if AreaScope_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if SibType_r17Present {
		ie.SibType_r17 = new(SIB_TypeInfo_v1700_sibType_r17)
		if err = ie.SibType_r17.Decode(r); err != nil {
			return utils.WrapError("Decode SibType_r17", err)
		}
	}
	if ValueTag_r17Present {
		var tmp_int_ValueTag_r17 int64
		if tmp_int_ValueTag_r17, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode ValueTag_r17", err)
		}
		ie.ValueTag_r17 = &tmp_int_ValueTag_r17
	}
	if AreaScope_r17Present {
		ie.AreaScope_r17 = new(SIB_TypeInfo_v1700_areaScope_r17)
		if err = ie.AreaScope_r17.Decode(r); err != nil {
			return utils.WrapError("Decode AreaScope_r17", err)
		}
	}
	return nil
}

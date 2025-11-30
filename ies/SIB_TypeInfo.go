package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SIB_TypeInfo struct {
	Type_sib  SIB_TypeInfo_type_sib   `madatory`
	ValueTag  *int64                  `lb:0,ub:31,optional`
	AreaScope *SIB_TypeInfo_areaScope `optional`
}

func (ie *SIB_TypeInfo) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.ValueTag != nil, ie.AreaScope != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Type_sib.Encode(w); err != nil {
		return utils.WrapError("Encode Type_sib", err)
	}
	if ie.ValueTag != nil {
		if err = w.WriteInteger(*ie.ValueTag, &aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Encode ValueTag", err)
		}
	}
	if ie.AreaScope != nil {
		if err = ie.AreaScope.Encode(w); err != nil {
			return utils.WrapError("Encode AreaScope", err)
		}
	}
	return nil
}

func (ie *SIB_TypeInfo) Decode(r *aper.AperReader) error {
	var err error
	var ValueTagPresent bool
	if ValueTagPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var AreaScopePresent bool
	if AreaScopePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Type_sib.Decode(r); err != nil {
		return utils.WrapError("Decode Type_sib", err)
	}
	if ValueTagPresent {
		var tmp_int_ValueTag int64
		if tmp_int_ValueTag, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode ValueTag", err)
		}
		ie.ValueTag = &tmp_int_ValueTag
	}
	if AreaScopePresent {
		ie.AreaScope = new(SIB_TypeInfo_areaScope)
		if err = ie.AreaScope.Decode(r); err != nil {
			return utils.WrapError("Decode AreaScope", err)
		}
	}
	return nil
}

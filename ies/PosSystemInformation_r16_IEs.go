package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PosSystemInformation_r16_IEs struct {
	PosSIB_TypeAndInfo_r16   []PosSIB_TypeAndInfo_r16 `lb:1,ub:maxSIB,madatory`
	LateNonCriticalExtension *[]byte                  `optional`
	NonCriticalExtension     interface{}              `optional`
}

func (ie *PosSystemInformation_r16_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.LateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	tmp_PosSIB_TypeAndInfo_r16 := utils.NewSequence[*PosSIB_TypeAndInfo_r16]([]*PosSIB_TypeAndInfo_r16{}, aper.Constraint{Lb: 1, Ub: maxSIB}, false)
	for _, i := range ie.PosSIB_TypeAndInfo_r16 {
		tmp_PosSIB_TypeAndInfo_r16.Value = append(tmp_PosSIB_TypeAndInfo_r16.Value, &i)
	}
	if err = tmp_PosSIB_TypeAndInfo_r16.Encode(w); err != nil {
		return utils.WrapError("Encode PosSIB_TypeAndInfo_r16", err)
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *PosSystemInformation_r16_IEs) Decode(r *aper.AperReader) error {
	var err error
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	tmp_PosSIB_TypeAndInfo_r16 := utils.NewSequence[*PosSIB_TypeAndInfo_r16]([]*PosSIB_TypeAndInfo_r16{}, aper.Constraint{Lb: 1, Ub: maxSIB}, false)
	fn_PosSIB_TypeAndInfo_r16 := func() *PosSIB_TypeAndInfo_r16 {
		return new(PosSIB_TypeAndInfo_r16)
	}
	if err = tmp_PosSIB_TypeAndInfo_r16.Decode(r, fn_PosSIB_TypeAndInfo_r16); err != nil {
		return utils.WrapError("Decode PosSIB_TypeAndInfo_r16", err)
	}
	ie.PosSIB_TypeAndInfo_r16 = []PosSIB_TypeAndInfo_r16{}
	for _, i := range tmp_PosSIB_TypeAndInfo_r16.Value {
		ie.PosSIB_TypeAndInfo_r16 = append(ie.PosSIB_TypeAndInfo_r16, *i)
	}
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}
	return nil
}

package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	UplinkDataCompression_r17_newSetup_bufferSize_r17_Enum_kbyte2 aper.Enumerated = 0
	UplinkDataCompression_r17_newSetup_bufferSize_r17_Enum_kbyte4 aper.Enumerated = 1
	UplinkDataCompression_r17_newSetup_bufferSize_r17_Enum_kbyte8 aper.Enumerated = 2
	UplinkDataCompression_r17_newSetup_bufferSize_r17_Enum_spare1 aper.Enumerated = 3
)

type UplinkDataCompression_r17_newSetup_bufferSize_r17 struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *UplinkDataCompression_r17_newSetup_bufferSize_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode UplinkDataCompression_r17_newSetup_bufferSize_r17", err)
	}
	return nil
}

func (ie *UplinkDataCompression_r17_newSetup_bufferSize_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode UplinkDataCompression_r17_newSetup_bufferSize_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}

package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RRC_PosSystemInfoRequest_r16_IEs struct {
	RequestedPosSI_List aper.BitString `lb:maxSI_Message,ub:maxSI_Message,madatory`
	Spare               aper.BitString `lb:11,ub:11,madatory`
}

func (ie *RRC_PosSystemInfoRequest_r16_IEs) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteBitString(ie.RequestedPosSI_List.Bytes, uint(ie.RequestedPosSI_List.NumBits), &aper.Constraint{Lb: maxSI_Message, Ub: maxSI_Message}, false); err != nil {
		return utils.WrapError("WriteBitString RequestedPosSI_List", err)
	}
	if err = w.WriteBitString(ie.Spare.Bytes, uint(ie.Spare.NumBits), &aper.Constraint{Lb: 11, Ub: 11}, false); err != nil {
		return utils.WrapError("WriteBitString Spare", err)
	}
	return nil
}

func (ie *RRC_PosSystemInfoRequest_r16_IEs) Decode(r *aper.AperReader) error {
	var err error
	var tmp_bs_RequestedPosSI_List []byte
	var n_RequestedPosSI_List uint
	if tmp_bs_RequestedPosSI_List, n_RequestedPosSI_List, err = r.ReadBitString(&aper.Constraint{Lb: maxSI_Message, Ub: maxSI_Message}, false); err != nil {
		return utils.WrapError("ReadBitString RequestedPosSI_List", err)
	}
	ie.RequestedPosSI_List = aper.BitString{
		Bytes:   tmp_bs_RequestedPosSI_List,
		NumBits: uint64(n_RequestedPosSI_List),
	}
	var tmp_bs_Spare []byte
	var n_Spare uint
	if tmp_bs_Spare, n_Spare, err = r.ReadBitString(&aper.Constraint{Lb: 11, Ub: 11}, false); err != nil {
		return utils.WrapError("ReadBitString Spare", err)
	}
	ie.Spare = aper.BitString{
		Bytes:   tmp_bs_Spare,
		NumBits: uint64(n_Spare),
	}
	return nil
}

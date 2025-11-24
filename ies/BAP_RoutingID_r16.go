package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BAP_RoutingID_r16 struct {
	Bap_Address_r16 uper.BitString `lb:10,ub:10,madatory`
	Bap_PathId_r16  uper.BitString `lb:10,ub:10,madatory`
}

func (ie *BAP_RoutingID_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteBitString(ie.Bap_Address_r16.Bytes, uint(ie.Bap_Address_r16.NumBits), &uper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
		return utils.WrapError("WriteBitString Bap_Address_r16", err)
	}
	if err = w.WriteBitString(ie.Bap_PathId_r16.Bytes, uint(ie.Bap_PathId_r16.NumBits), &uper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
		return utils.WrapError("WriteBitString Bap_PathId_r16", err)
	}
	return nil
}

func (ie *BAP_RoutingID_r16) Decode(r *uper.UperReader) error {
	var err error
	var tmp_bs_Bap_Address_r16 []byte
	var n_Bap_Address_r16 uint
	if tmp_bs_Bap_Address_r16, n_Bap_Address_r16, err = r.ReadBitString(&uper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
		return utils.WrapError("ReadBitString Bap_Address_r16", err)
	}
	ie.Bap_Address_r16 = uper.BitString{
		Bytes:   tmp_bs_Bap_Address_r16,
		NumBits: uint64(n_Bap_Address_r16),
	}
	var tmp_bs_Bap_PathId_r16 []byte
	var n_Bap_PathId_r16 uint
	if tmp_bs_Bap_PathId_r16, n_Bap_PathId_r16, err = r.ReadBitString(&uper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
		return utils.WrapError("ReadBitString Bap_PathId_r16", err)
	}
	ie.Bap_PathId_r16 = uper.BitString{
		Bytes:   tmp_bs_Bap_PathId_r16,
		NumBits: uint64(n_Bap_PathId_r16),
	}
	return nil
}

package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MIB struct {
	SystemFrameNumber       aper.BitString              `lb:6,ub:6,madatory`
	SubCarrierSpacingCommon MIB_subCarrierSpacingCommon `madatory`
	Ssb_SubcarrierOffset    int64                       `lb:0,ub:15,madatory`
	Dmrs_TypeA_Position     MIB_dmrs_TypeA_Position     `madatory`
	Pdcch_ConfigSIB1        PDCCH_ConfigSIB1            `madatory`
	CellBarred              MIB_cellBarred              `madatory`
	IntraFreqReselection    MIB_intraFreqReselection    `madatory`
	Spare                   aper.BitString              `lb:1,ub:1,madatory`
}

func (ie *MIB) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteBitString(ie.SystemFrameNumber.Bytes, uint(ie.SystemFrameNumber.NumBits), &aper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
		return utils.WrapError("WriteBitString SystemFrameNumber", err)
	}
	if err = ie.SubCarrierSpacingCommon.Encode(w); err != nil {
		return utils.WrapError("Encode SubCarrierSpacingCommon", err)
	}
	if err = w.WriteInteger(ie.Ssb_SubcarrierOffset, &aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger Ssb_SubcarrierOffset", err)
	}
	if err = ie.Dmrs_TypeA_Position.Encode(w); err != nil {
		return utils.WrapError("Encode Dmrs_TypeA_Position", err)
	}
	if err = ie.Pdcch_ConfigSIB1.Encode(w); err != nil {
		return utils.WrapError("Encode Pdcch_ConfigSIB1", err)
	}
	if err = ie.CellBarred.Encode(w); err != nil {
		return utils.WrapError("Encode CellBarred", err)
	}
	if err = ie.IntraFreqReselection.Encode(w); err != nil {
		return utils.WrapError("Encode IntraFreqReselection", err)
	}
	if err = w.WriteBitString(ie.Spare.Bytes, uint(ie.Spare.NumBits), &aper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return utils.WrapError("WriteBitString Spare", err)
	}
	return nil
}

func (ie *MIB) Decode(r *aper.AperReader) error {
	var err error
	var tmp_bs_SystemFrameNumber []byte
	var n_SystemFrameNumber uint
	if tmp_bs_SystemFrameNumber, n_SystemFrameNumber, err = r.ReadBitString(&aper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
		return utils.WrapError("ReadBitString SystemFrameNumber", err)
	}
	ie.SystemFrameNumber = aper.BitString{
		Bytes:   tmp_bs_SystemFrameNumber,
		NumBits: uint64(n_SystemFrameNumber),
	}
	if err = ie.SubCarrierSpacingCommon.Decode(r); err != nil {
		return utils.WrapError("Decode SubCarrierSpacingCommon", err)
	}
	var tmp_int_Ssb_SubcarrierOffset int64
	if tmp_int_Ssb_SubcarrierOffset, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger Ssb_SubcarrierOffset", err)
	}
	ie.Ssb_SubcarrierOffset = tmp_int_Ssb_SubcarrierOffset
	if err = ie.Dmrs_TypeA_Position.Decode(r); err != nil {
		return utils.WrapError("Decode Dmrs_TypeA_Position", err)
	}
	if err = ie.Pdcch_ConfigSIB1.Decode(r); err != nil {
		return utils.WrapError("Decode Pdcch_ConfigSIB1", err)
	}
	if err = ie.CellBarred.Decode(r); err != nil {
		return utils.WrapError("Decode CellBarred", err)
	}
	if err = ie.IntraFreqReselection.Decode(r); err != nil {
		return utils.WrapError("Decode IntraFreqReselection", err)
	}
	var tmp_bs_Spare []byte
	var n_Spare uint
	if tmp_bs_Spare, n_Spare, err = r.ReadBitString(&aper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return utils.WrapError("ReadBitString Spare", err)
	}
	ie.Spare = aper.BitString{
		Bytes:   tmp_bs_Spare,
		NumBits: uint64(n_Spare),
	}
	return nil
}

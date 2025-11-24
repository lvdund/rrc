package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultRLFNR_r16_measResult_r16_rsIndexResults_r16 struct {
	ResultsSSB_Indexes_r16    *ResultsPerSSB_IndexList    `optional`
	SsbRLMConfigBitmap_r16    *uper.BitString             `lb:64,ub:64,optional`
	ResultsCSI_RS_Indexes_r16 *ResultsPerCSI_RS_IndexList `optional`
	Csi_rsRLMConfigBitmap_r16 *uper.BitString             `lb:96,ub:96,optional`
}

func (ie *MeasResultRLFNR_r16_measResult_r16_rsIndexResults_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ResultsSSB_Indexes_r16 != nil, ie.SsbRLMConfigBitmap_r16 != nil, ie.ResultsCSI_RS_Indexes_r16 != nil, ie.Csi_rsRLMConfigBitmap_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ResultsSSB_Indexes_r16 != nil {
		if err = ie.ResultsSSB_Indexes_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ResultsSSB_Indexes_r16", err)
		}
	}
	if ie.SsbRLMConfigBitmap_r16 != nil {
		if err = w.WriteBitString(ie.SsbRLMConfigBitmap_r16.Bytes, uint(ie.SsbRLMConfigBitmap_r16.NumBits), &uper.Constraint{Lb: 64, Ub: 64}, false); err != nil {
			return utils.WrapError("Encode SsbRLMConfigBitmap_r16", err)
		}
	}
	if ie.ResultsCSI_RS_Indexes_r16 != nil {
		if err = ie.ResultsCSI_RS_Indexes_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ResultsCSI_RS_Indexes_r16", err)
		}
	}
	if ie.Csi_rsRLMConfigBitmap_r16 != nil {
		if err = w.WriteBitString(ie.Csi_rsRLMConfigBitmap_r16.Bytes, uint(ie.Csi_rsRLMConfigBitmap_r16.NumBits), &uper.Constraint{Lb: 96, Ub: 96}, false); err != nil {
			return utils.WrapError("Encode Csi_rsRLMConfigBitmap_r16", err)
		}
	}
	return nil
}

func (ie *MeasResultRLFNR_r16_measResult_r16_rsIndexResults_r16) Decode(r *uper.UperReader) error {
	var err error
	var ResultsSSB_Indexes_r16Present bool
	if ResultsSSB_Indexes_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SsbRLMConfigBitmap_r16Present bool
	if SsbRLMConfigBitmap_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ResultsCSI_RS_Indexes_r16Present bool
	if ResultsCSI_RS_Indexes_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Csi_rsRLMConfigBitmap_r16Present bool
	if Csi_rsRLMConfigBitmap_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ResultsSSB_Indexes_r16Present {
		ie.ResultsSSB_Indexes_r16 = new(ResultsPerSSB_IndexList)
		if err = ie.ResultsSSB_Indexes_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ResultsSSB_Indexes_r16", err)
		}
	}
	if SsbRLMConfigBitmap_r16Present {
		var tmp_bs_SsbRLMConfigBitmap_r16 []byte
		var n_SsbRLMConfigBitmap_r16 uint
		if tmp_bs_SsbRLMConfigBitmap_r16, n_SsbRLMConfigBitmap_r16, err = r.ReadBitString(&uper.Constraint{Lb: 64, Ub: 64}, false); err != nil {
			return utils.WrapError("Decode SsbRLMConfigBitmap_r16", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_SsbRLMConfigBitmap_r16,
			NumBits: uint64(n_SsbRLMConfigBitmap_r16),
		}
		ie.SsbRLMConfigBitmap_r16 = &tmp_bitstring
	}
	if ResultsCSI_RS_Indexes_r16Present {
		ie.ResultsCSI_RS_Indexes_r16 = new(ResultsPerCSI_RS_IndexList)
		if err = ie.ResultsCSI_RS_Indexes_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ResultsCSI_RS_Indexes_r16", err)
		}
	}
	if Csi_rsRLMConfigBitmap_r16Present {
		var tmp_bs_Csi_rsRLMConfigBitmap_r16 []byte
		var n_Csi_rsRLMConfigBitmap_r16 uint
		if tmp_bs_Csi_rsRLMConfigBitmap_r16, n_Csi_rsRLMConfigBitmap_r16, err = r.ReadBitString(&uper.Constraint{Lb: 96, Ub: 96}, false); err != nil {
			return utils.WrapError("Decode Csi_rsRLMConfigBitmap_r16", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_Csi_rsRLMConfigBitmap_r16,
			NumBits: uint64(n_Csi_rsRLMConfigBitmap_r16),
		}
		ie.Csi_rsRLMConfigBitmap_r16 = &tmp_bitstring
	}
	return nil
}

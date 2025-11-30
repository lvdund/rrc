package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UE_SidelinkCapabilityAddXDD_Mode_r16 struct {
	Mac_ParametersSidelinkXDD_Diff_r16 *MAC_ParametersSidelinkXDD_Diff_r16 `optional`
}

func (ie *UE_SidelinkCapabilityAddXDD_Mode_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Mac_ParametersSidelinkXDD_Diff_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Mac_ParametersSidelinkXDD_Diff_r16 != nil {
		if err = ie.Mac_ParametersSidelinkXDD_Diff_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Mac_ParametersSidelinkXDD_Diff_r16", err)
		}
	}
	return nil
}

func (ie *UE_SidelinkCapabilityAddXDD_Mode_r16) Decode(r *aper.AperReader) error {
	var err error
	var Mac_ParametersSidelinkXDD_Diff_r16Present bool
	if Mac_ParametersSidelinkXDD_Diff_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Mac_ParametersSidelinkXDD_Diff_r16Present {
		ie.Mac_ParametersSidelinkXDD_Diff_r16 = new(MAC_ParametersSidelinkXDD_Diff_r16)
		if err = ie.Mac_ParametersSidelinkXDD_Diff_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Mac_ParametersSidelinkXDD_Diff_r16", err)
		}
	}
	return nil
}

package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Phy_ParametersFRX_Diff_mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot struct {
	SameSymbol *Phy_ParametersFRX_Diff_mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_sameSymbol `optional`
	DiffSymbol *Phy_ParametersFRX_Diff_mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_diffSymbol `optional`
}

func (ie *Phy_ParametersFRX_Diff_mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.SameSymbol != nil, ie.DiffSymbol != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SameSymbol != nil {
		if err = ie.SameSymbol.Encode(w); err != nil {
			return utils.WrapError("Encode SameSymbol", err)
		}
	}
	if ie.DiffSymbol != nil {
		if err = ie.DiffSymbol.Encode(w); err != nil {
			return utils.WrapError("Encode DiffSymbol", err)
		}
	}
	return nil
}

func (ie *Phy_ParametersFRX_Diff_mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot) Decode(r *uper.UperReader) error {
	var err error
	var SameSymbolPresent bool
	if SameSymbolPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var DiffSymbolPresent bool
	if DiffSymbolPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if SameSymbolPresent {
		ie.SameSymbol = new(Phy_ParametersFRX_Diff_mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_sameSymbol)
		if err = ie.SameSymbol.Decode(r); err != nil {
			return utils.WrapError("Decode SameSymbol", err)
		}
	}
	if DiffSymbolPresent {
		ie.DiffSymbol = new(Phy_ParametersFRX_Diff_mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_diffSymbol)
		if err = ie.DiffSymbol.Decode(r); err != nil {
			return utils.WrapError("Decode DiffSymbol", err)
		}
	}
	return nil
}

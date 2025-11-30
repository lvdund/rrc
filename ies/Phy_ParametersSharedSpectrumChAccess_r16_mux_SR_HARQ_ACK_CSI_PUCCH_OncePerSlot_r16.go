package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type Phy_ParametersSharedSpectrumChAccess_r16_mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16 struct {
	SameSymbol_r16 *Phy_ParametersSharedSpectrumChAccess_r16_mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16_sameSymbol_r16 `optional`
	DiffSymbol_r16 *Phy_ParametersSharedSpectrumChAccess_r16_mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16_diffSymbol_r16 `optional`
}

func (ie *Phy_ParametersSharedSpectrumChAccess_r16_mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.SameSymbol_r16 != nil, ie.DiffSymbol_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SameSymbol_r16 != nil {
		if err = ie.SameSymbol_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SameSymbol_r16", err)
		}
	}
	if ie.DiffSymbol_r16 != nil {
		if err = ie.DiffSymbol_r16.Encode(w); err != nil {
			return utils.WrapError("Encode DiffSymbol_r16", err)
		}
	}
	return nil
}

func (ie *Phy_ParametersSharedSpectrumChAccess_r16_mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16) Decode(r *aper.AperReader) error {
	var err error
	var SameSymbol_r16Present bool
	if SameSymbol_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DiffSymbol_r16Present bool
	if DiffSymbol_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if SameSymbol_r16Present {
		ie.SameSymbol_r16 = new(Phy_ParametersSharedSpectrumChAccess_r16_mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16_sameSymbol_r16)
		if err = ie.SameSymbol_r16.Decode(r); err != nil {
			return utils.WrapError("Decode SameSymbol_r16", err)
		}
	}
	if DiffSymbol_r16Present {
		ie.DiffSymbol_r16 = new(Phy_ParametersSharedSpectrumChAccess_r16_mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16_diffSymbol_r16)
		if err = ie.DiffSymbol_r16.Decode(r); err != nil {
			return utils.WrapError("Decode DiffSymbol_r16", err)
		}
	}
	return nil
}

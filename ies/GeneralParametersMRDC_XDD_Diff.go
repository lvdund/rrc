package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type GeneralParametersMRDC_XDD_Diff struct {
	SplitSRB_WithOneUL_Path      *GeneralParametersMRDC_XDD_Diff_splitSRB_WithOneUL_Path      `optional`
	SplitDRB_withUL_Both_MCG_SCG *GeneralParametersMRDC_XDD_Diff_splitDRB_withUL_Both_MCG_SCG `optional`
	Srb3                         *GeneralParametersMRDC_XDD_Diff_srb3                         `optional`
	Dummy                        *GeneralParametersMRDC_XDD_Diff_dummy                        `optional`
}

func (ie *GeneralParametersMRDC_XDD_Diff) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.SplitSRB_WithOneUL_Path != nil, ie.SplitDRB_withUL_Both_MCG_SCG != nil, ie.Srb3 != nil, ie.Dummy != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SplitSRB_WithOneUL_Path != nil {
		if err = ie.SplitSRB_WithOneUL_Path.Encode(w); err != nil {
			return utils.WrapError("Encode SplitSRB_WithOneUL_Path", err)
		}
	}
	if ie.SplitDRB_withUL_Both_MCG_SCG != nil {
		if err = ie.SplitDRB_withUL_Both_MCG_SCG.Encode(w); err != nil {
			return utils.WrapError("Encode SplitDRB_withUL_Both_MCG_SCG", err)
		}
	}
	if ie.Srb3 != nil {
		if err = ie.Srb3.Encode(w); err != nil {
			return utils.WrapError("Encode Srb3", err)
		}
	}
	if ie.Dummy != nil {
		if err = ie.Dummy.Encode(w); err != nil {
			return utils.WrapError("Encode Dummy", err)
		}
	}
	return nil
}

func (ie *GeneralParametersMRDC_XDD_Diff) Decode(r *aper.AperReader) error {
	var err error
	var SplitSRB_WithOneUL_PathPresent bool
	if SplitSRB_WithOneUL_PathPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SplitDRB_withUL_Both_MCG_SCGPresent bool
	if SplitDRB_withUL_Both_MCG_SCGPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Srb3Present bool
	if Srb3Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DummyPresent bool
	if DummyPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if SplitSRB_WithOneUL_PathPresent {
		ie.SplitSRB_WithOneUL_Path = new(GeneralParametersMRDC_XDD_Diff_splitSRB_WithOneUL_Path)
		if err = ie.SplitSRB_WithOneUL_Path.Decode(r); err != nil {
			return utils.WrapError("Decode SplitSRB_WithOneUL_Path", err)
		}
	}
	if SplitDRB_withUL_Both_MCG_SCGPresent {
		ie.SplitDRB_withUL_Both_MCG_SCG = new(GeneralParametersMRDC_XDD_Diff_splitDRB_withUL_Both_MCG_SCG)
		if err = ie.SplitDRB_withUL_Both_MCG_SCG.Decode(r); err != nil {
			return utils.WrapError("Decode SplitDRB_withUL_Both_MCG_SCG", err)
		}
	}
	if Srb3Present {
		ie.Srb3 = new(GeneralParametersMRDC_XDD_Diff_srb3)
		if err = ie.Srb3.Decode(r); err != nil {
			return utils.WrapError("Decode Srb3", err)
		}
	}
	if DummyPresent {
		ie.Dummy = new(GeneralParametersMRDC_XDD_Diff_dummy)
		if err = ie.Dummy.Decode(r); err != nil {
			return utils.WrapError("Decode Dummy", err)
		}
	}
	return nil
}

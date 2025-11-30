package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SDT_Config_r17 struct {
	Sdt_DRB_List_r17          []DRB_Identity                           `lb:0,ub:maxDRB,optional`
	Sdt_SRB2_Indication_r17   *SDT_Config_r17_sdt_SRB2_Indication_r17  `optional`
	Sdt_MAC_PHY_CG_Config_r17 *SDT_CG_Config_r17                       `optional,setuprelease`
	Sdt_DRB_ContinueROHC_r17  *SDT_Config_r17_sdt_DRB_ContinueROHC_r17 `optional`
}

func (ie *SDT_Config_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{len(ie.Sdt_DRB_List_r17) > 0, ie.Sdt_SRB2_Indication_r17 != nil, ie.Sdt_MAC_PHY_CG_Config_r17 != nil, ie.Sdt_DRB_ContinueROHC_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.Sdt_DRB_List_r17) > 0 {
		tmp_Sdt_DRB_List_r17 := utils.NewSequence[*DRB_Identity]([]*DRB_Identity{}, aper.Constraint{Lb: 0, Ub: maxDRB}, false)
		for _, i := range ie.Sdt_DRB_List_r17 {
			tmp_Sdt_DRB_List_r17.Value = append(tmp_Sdt_DRB_List_r17.Value, &i)
		}
		if err = tmp_Sdt_DRB_List_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sdt_DRB_List_r17", err)
		}
	}
	if ie.Sdt_SRB2_Indication_r17 != nil {
		if err = ie.Sdt_SRB2_Indication_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sdt_SRB2_Indication_r17", err)
		}
	}
	if ie.Sdt_MAC_PHY_CG_Config_r17 != nil {
		tmp_Sdt_MAC_PHY_CG_Config_r17 := utils.SetupRelease[*SDT_CG_Config_r17]{
			Setup: ie.Sdt_MAC_PHY_CG_Config_r17,
		}
		if err = tmp_Sdt_MAC_PHY_CG_Config_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sdt_MAC_PHY_CG_Config_r17", err)
		}
	}
	if ie.Sdt_DRB_ContinueROHC_r17 != nil {
		if err = ie.Sdt_DRB_ContinueROHC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sdt_DRB_ContinueROHC_r17", err)
		}
	}
	return nil
}

func (ie *SDT_Config_r17) Decode(r *aper.AperReader) error {
	var err error
	var Sdt_DRB_List_r17Present bool
	if Sdt_DRB_List_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sdt_SRB2_Indication_r17Present bool
	if Sdt_SRB2_Indication_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sdt_MAC_PHY_CG_Config_r17Present bool
	if Sdt_MAC_PHY_CG_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sdt_DRB_ContinueROHC_r17Present bool
	if Sdt_DRB_ContinueROHC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sdt_DRB_List_r17Present {
		tmp_Sdt_DRB_List_r17 := utils.NewSequence[*DRB_Identity]([]*DRB_Identity{}, aper.Constraint{Lb: 0, Ub: maxDRB}, false)
		fn_Sdt_DRB_List_r17 := func() *DRB_Identity {
			return new(DRB_Identity)
		}
		if err = tmp_Sdt_DRB_List_r17.Decode(r, fn_Sdt_DRB_List_r17); err != nil {
			return utils.WrapError("Decode Sdt_DRB_List_r17", err)
		}
		ie.Sdt_DRB_List_r17 = []DRB_Identity{}
		for _, i := range tmp_Sdt_DRB_List_r17.Value {
			ie.Sdt_DRB_List_r17 = append(ie.Sdt_DRB_List_r17, *i)
		}
	}
	if Sdt_SRB2_Indication_r17Present {
		ie.Sdt_SRB2_Indication_r17 = new(SDT_Config_r17_sdt_SRB2_Indication_r17)
		if err = ie.Sdt_SRB2_Indication_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sdt_SRB2_Indication_r17", err)
		}
	}
	if Sdt_MAC_PHY_CG_Config_r17Present {
		tmp_Sdt_MAC_PHY_CG_Config_r17 := utils.SetupRelease[*SDT_CG_Config_r17]{}
		if err = tmp_Sdt_MAC_PHY_CG_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sdt_MAC_PHY_CG_Config_r17", err)
		}
		ie.Sdt_MAC_PHY_CG_Config_r17 = tmp_Sdt_MAC_PHY_CG_Config_r17.Setup
	}
	if Sdt_DRB_ContinueROHC_r17Present {
		ie.Sdt_DRB_ContinueROHC_r17 = new(SDT_Config_r17_sdt_DRB_ContinueROHC_r17)
		if err = ie.Sdt_DRB_ContinueROHC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sdt_DRB_ContinueROHC_r17", err)
		}
	}
	return nil
}

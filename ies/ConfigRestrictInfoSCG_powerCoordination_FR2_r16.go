package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ConfigRestrictInfoSCG_powerCoordination_FR2_r16 struct {
	P_maxNR_FR2_MCG_r16 *P_Max `optional`
	P_maxNR_FR2_SCG_r16 *P_Max `optional`
	P_maxUE_FR2_r16     *P_Max `optional`
}

func (ie *ConfigRestrictInfoSCG_powerCoordination_FR2_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.P_maxNR_FR2_MCG_r16 != nil, ie.P_maxNR_FR2_SCG_r16 != nil, ie.P_maxUE_FR2_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.P_maxNR_FR2_MCG_r16 != nil {
		if err = ie.P_maxNR_FR2_MCG_r16.Encode(w); err != nil {
			return utils.WrapError("Encode P_maxNR_FR2_MCG_r16", err)
		}
	}
	if ie.P_maxNR_FR2_SCG_r16 != nil {
		if err = ie.P_maxNR_FR2_SCG_r16.Encode(w); err != nil {
			return utils.WrapError("Encode P_maxNR_FR2_SCG_r16", err)
		}
	}
	if ie.P_maxUE_FR2_r16 != nil {
		if err = ie.P_maxUE_FR2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode P_maxUE_FR2_r16", err)
		}
	}
	return nil
}

func (ie *ConfigRestrictInfoSCG_powerCoordination_FR2_r16) Decode(r *uper.UperReader) error {
	var err error
	var P_maxNR_FR2_MCG_r16Present bool
	if P_maxNR_FR2_MCG_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var P_maxNR_FR2_SCG_r16Present bool
	if P_maxNR_FR2_SCG_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var P_maxUE_FR2_r16Present bool
	if P_maxUE_FR2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if P_maxNR_FR2_MCG_r16Present {
		ie.P_maxNR_FR2_MCG_r16 = new(P_Max)
		if err = ie.P_maxNR_FR2_MCG_r16.Decode(r); err != nil {
			return utils.WrapError("Decode P_maxNR_FR2_MCG_r16", err)
		}
	}
	if P_maxNR_FR2_SCG_r16Present {
		ie.P_maxNR_FR2_SCG_r16 = new(P_Max)
		if err = ie.P_maxNR_FR2_SCG_r16.Decode(r); err != nil {
			return utils.WrapError("Decode P_maxNR_FR2_SCG_r16", err)
		}
	}
	if P_maxUE_FR2_r16Present {
		ie.P_maxUE_FR2_r16 = new(P_Max)
		if err = ie.P_maxUE_FR2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode P_maxUE_FR2_r16", err)
		}
	}
	return nil
}

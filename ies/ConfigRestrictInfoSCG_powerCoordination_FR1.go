package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type ConfigRestrictInfoSCG_powerCoordination_FR1 struct {
	P_maxNR_FR1 *P_Max `optional`
	P_maxEUTRA  *P_Max `optional`
	P_maxUE_FR1 *P_Max `optional`
}

func (ie *ConfigRestrictInfoSCG_powerCoordination_FR1) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.P_maxNR_FR1 != nil, ie.P_maxEUTRA != nil, ie.P_maxUE_FR1 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.P_maxNR_FR1 != nil {
		if err = ie.P_maxNR_FR1.Encode(w); err != nil {
			return utils.WrapError("Encode P_maxNR_FR1", err)
		}
	}
	if ie.P_maxEUTRA != nil {
		if err = ie.P_maxEUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode P_maxEUTRA", err)
		}
	}
	if ie.P_maxUE_FR1 != nil {
		if err = ie.P_maxUE_FR1.Encode(w); err != nil {
			return utils.WrapError("Encode P_maxUE_FR1", err)
		}
	}
	return nil
}

func (ie *ConfigRestrictInfoSCG_powerCoordination_FR1) Decode(r *aper.AperReader) error {
	var err error
	var P_maxNR_FR1Present bool
	if P_maxNR_FR1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var P_maxEUTRAPresent bool
	if P_maxEUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var P_maxUE_FR1Present bool
	if P_maxUE_FR1Present, err = r.ReadBool(); err != nil {
		return err
	}
	if P_maxNR_FR1Present {
		ie.P_maxNR_FR1 = new(P_Max)
		if err = ie.P_maxNR_FR1.Decode(r); err != nil {
			return utils.WrapError("Decode P_maxNR_FR1", err)
		}
	}
	if P_maxEUTRAPresent {
		ie.P_maxEUTRA = new(P_Max)
		if err = ie.P_maxEUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode P_maxEUTRA", err)
		}
	}
	if P_maxUE_FR1Present {
		ie.P_maxUE_FR1 = new(P_Max)
		if err = ie.P_maxUE_FR1.Decode(r); err != nil {
			return utils.WrapError("Decode P_maxUE_FR1", err)
		}
	}
	return nil
}

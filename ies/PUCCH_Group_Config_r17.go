package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_Group_Config_r17 struct {
	Fr1_FR1_NonSharedTDD_r17 *PUCCH_Group_Config_r17_fr1_FR1_NonSharedTDD_r17 `optional`
	Fr2_FR2_NonSharedTDD_r17 *PUCCH_Group_Config_r17_fr2_FR2_NonSharedTDD_r17 `optional`
	Fr1_FR2_NonSharedTDD_r17 *PUCCH_Group_Config_r17_fr1_FR2_NonSharedTDD_r17 `optional`
}

func (ie *PUCCH_Group_Config_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Fr1_FR1_NonSharedTDD_r17 != nil, ie.Fr2_FR2_NonSharedTDD_r17 != nil, ie.Fr1_FR2_NonSharedTDD_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Fr1_FR1_NonSharedTDD_r17 != nil {
		if err = ie.Fr1_FR1_NonSharedTDD_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Fr1_FR1_NonSharedTDD_r17", err)
		}
	}
	if ie.Fr2_FR2_NonSharedTDD_r17 != nil {
		if err = ie.Fr2_FR2_NonSharedTDD_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Fr2_FR2_NonSharedTDD_r17", err)
		}
	}
	if ie.Fr1_FR2_NonSharedTDD_r17 != nil {
		if err = ie.Fr1_FR2_NonSharedTDD_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Fr1_FR2_NonSharedTDD_r17", err)
		}
	}
	return nil
}

func (ie *PUCCH_Group_Config_r17) Decode(r *uper.UperReader) error {
	var err error
	var Fr1_FR1_NonSharedTDD_r17Present bool
	if Fr1_FR1_NonSharedTDD_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Fr2_FR2_NonSharedTDD_r17Present bool
	if Fr2_FR2_NonSharedTDD_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Fr1_FR2_NonSharedTDD_r17Present bool
	if Fr1_FR2_NonSharedTDD_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Fr1_FR1_NonSharedTDD_r17Present {
		ie.Fr1_FR1_NonSharedTDD_r17 = new(PUCCH_Group_Config_r17_fr1_FR1_NonSharedTDD_r17)
		if err = ie.Fr1_FR1_NonSharedTDD_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Fr1_FR1_NonSharedTDD_r17", err)
		}
	}
	if Fr2_FR2_NonSharedTDD_r17Present {
		ie.Fr2_FR2_NonSharedTDD_r17 = new(PUCCH_Group_Config_r17_fr2_FR2_NonSharedTDD_r17)
		if err = ie.Fr2_FR2_NonSharedTDD_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Fr2_FR2_NonSharedTDD_r17", err)
		}
	}
	if Fr1_FR2_NonSharedTDD_r17Present {
		ie.Fr1_FR2_NonSharedTDD_r17 = new(PUCCH_Group_Config_r17_fr1_FR2_NonSharedTDD_r17)
		if err = ie.Fr1_FR2_NonSharedTDD_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Fr1_FR2_NonSharedTDD_r17", err)
		}
	}
	return nil
}

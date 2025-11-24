package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MUSIM_Assistance_r17 struct {
	Musim_PreferredRRC_State_r17 *MUSIM_Assistance_r17_musim_PreferredRRC_State_r17 `optional`
	Musim_GapPreferenceList_r17  *MUSIM_GapPreferenceList_r17                       `optional`
}

func (ie *MUSIM_Assistance_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Musim_PreferredRRC_State_r17 != nil, ie.Musim_GapPreferenceList_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Musim_PreferredRRC_State_r17 != nil {
		if err = ie.Musim_PreferredRRC_State_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Musim_PreferredRRC_State_r17", err)
		}
	}
	if ie.Musim_GapPreferenceList_r17 != nil {
		if err = ie.Musim_GapPreferenceList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Musim_GapPreferenceList_r17", err)
		}
	}
	return nil
}

func (ie *MUSIM_Assistance_r17) Decode(r *uper.UperReader) error {
	var err error
	var Musim_PreferredRRC_State_r17Present bool
	if Musim_PreferredRRC_State_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Musim_GapPreferenceList_r17Present bool
	if Musim_GapPreferenceList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Musim_PreferredRRC_State_r17Present {
		ie.Musim_PreferredRRC_State_r17 = new(MUSIM_Assistance_r17_musim_PreferredRRC_State_r17)
		if err = ie.Musim_PreferredRRC_State_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Musim_PreferredRRC_State_r17", err)
		}
	}
	if Musim_GapPreferenceList_r17Present {
		ie.Musim_GapPreferenceList_r17 = new(MUSIM_GapPreferenceList_r17)
		if err = ie.Musim_GapPreferenceList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Musim_GapPreferenceList_r17", err)
		}
	}
	return nil
}

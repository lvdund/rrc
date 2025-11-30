package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type FeatureCombination_r17 struct {
	RedCap_r17           *FeatureCombination_r17_redCap_r17           `optional`
	SmallData_r17        *FeatureCombination_r17_smallData_r17        `optional`
	Nsag_r17             *NSAG_List_r17                               `optional`
	Msg3_Repetitions_r17 *FeatureCombination_r17_msg3_Repetitions_r17 `optional`
	Spare4               *FeatureCombination_r17_spare4               `optional`
	Spare3               *FeatureCombination_r17_spare3               `optional`
	Spare2               *FeatureCombination_r17_spare2               `optional`
	Spare1               *FeatureCombination_r17_spare1               `optional`
}

func (ie *FeatureCombination_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.RedCap_r17 != nil, ie.SmallData_r17 != nil, ie.Nsag_r17 != nil, ie.Msg3_Repetitions_r17 != nil, ie.Spare4 != nil, ie.Spare3 != nil, ie.Spare2 != nil, ie.Spare1 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.RedCap_r17 != nil {
		if err = ie.RedCap_r17.Encode(w); err != nil {
			return utils.WrapError("Encode RedCap_r17", err)
		}
	}
	if ie.SmallData_r17 != nil {
		if err = ie.SmallData_r17.Encode(w); err != nil {
			return utils.WrapError("Encode SmallData_r17", err)
		}
	}
	if ie.Nsag_r17 != nil {
		if err = ie.Nsag_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Nsag_r17", err)
		}
	}
	if ie.Msg3_Repetitions_r17 != nil {
		if err = ie.Msg3_Repetitions_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Msg3_Repetitions_r17", err)
		}
	}
	if ie.Spare4 != nil {
		if err = ie.Spare4.Encode(w); err != nil {
			return utils.WrapError("Encode Spare4", err)
		}
	}
	if ie.Spare3 != nil {
		if err = ie.Spare3.Encode(w); err != nil {
			return utils.WrapError("Encode Spare3", err)
		}
	}
	if ie.Spare2 != nil {
		if err = ie.Spare2.Encode(w); err != nil {
			return utils.WrapError("Encode Spare2", err)
		}
	}
	if ie.Spare1 != nil {
		if err = ie.Spare1.Encode(w); err != nil {
			return utils.WrapError("Encode Spare1", err)
		}
	}
	return nil
}

func (ie *FeatureCombination_r17) Decode(r *aper.AperReader) error {
	var err error
	var RedCap_r17Present bool
	if RedCap_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SmallData_r17Present bool
	if SmallData_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Nsag_r17Present bool
	if Nsag_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Msg3_Repetitions_r17Present bool
	if Msg3_Repetitions_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Spare4Present bool
	if Spare4Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Spare3Present bool
	if Spare3Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Spare2Present bool
	if Spare2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Spare1Present bool
	if Spare1Present, err = r.ReadBool(); err != nil {
		return err
	}
	if RedCap_r17Present {
		ie.RedCap_r17 = new(FeatureCombination_r17_redCap_r17)
		if err = ie.RedCap_r17.Decode(r); err != nil {
			return utils.WrapError("Decode RedCap_r17", err)
		}
	}
	if SmallData_r17Present {
		ie.SmallData_r17 = new(FeatureCombination_r17_smallData_r17)
		if err = ie.SmallData_r17.Decode(r); err != nil {
			return utils.WrapError("Decode SmallData_r17", err)
		}
	}
	if Nsag_r17Present {
		ie.Nsag_r17 = new(NSAG_List_r17)
		if err = ie.Nsag_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Nsag_r17", err)
		}
	}
	if Msg3_Repetitions_r17Present {
		ie.Msg3_Repetitions_r17 = new(FeatureCombination_r17_msg3_Repetitions_r17)
		if err = ie.Msg3_Repetitions_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Msg3_Repetitions_r17", err)
		}
	}
	if Spare4Present {
		ie.Spare4 = new(FeatureCombination_r17_spare4)
		if err = ie.Spare4.Decode(r); err != nil {
			return utils.WrapError("Decode Spare4", err)
		}
	}
	if Spare3Present {
		ie.Spare3 = new(FeatureCombination_r17_spare3)
		if err = ie.Spare3.Decode(r); err != nil {
			return utils.WrapError("Decode Spare3", err)
		}
	}
	if Spare2Present {
		ie.Spare2 = new(FeatureCombination_r17_spare2)
		if err = ie.Spare2.Decode(r); err != nil {
			return utils.WrapError("Decode Spare2", err)
		}
	}
	if Spare1Present {
		ie.Spare1 = new(FeatureCombination_r17_spare1)
		if err = ie.Spare1.Decode(r); err != nil {
			return utils.WrapError("Decode Spare1", err)
		}
	}
	return nil
}

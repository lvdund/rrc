package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SHR_Cause_r17 struct {
	T304_cause_r17         *SHR_Cause_r17_t304_cause_r17         `optional`
	T310_cause_r17         *SHR_Cause_r17_t310_cause_r17         `optional`
	T312_cause_r17         *SHR_Cause_r17_t312_cause_r17         `optional`
	SourceDAPS_Failure_r17 *SHR_Cause_r17_sourceDAPS_Failure_r17 `optional`
}

func (ie *SHR_Cause_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.T304_cause_r17 != nil, ie.T310_cause_r17 != nil, ie.T312_cause_r17 != nil, ie.SourceDAPS_Failure_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.T304_cause_r17 != nil {
		if err = ie.T304_cause_r17.Encode(w); err != nil {
			return utils.WrapError("Encode T304_cause_r17", err)
		}
	}
	if ie.T310_cause_r17 != nil {
		if err = ie.T310_cause_r17.Encode(w); err != nil {
			return utils.WrapError("Encode T310_cause_r17", err)
		}
	}
	if ie.T312_cause_r17 != nil {
		if err = ie.T312_cause_r17.Encode(w); err != nil {
			return utils.WrapError("Encode T312_cause_r17", err)
		}
	}
	if ie.SourceDAPS_Failure_r17 != nil {
		if err = ie.SourceDAPS_Failure_r17.Encode(w); err != nil {
			return utils.WrapError("Encode SourceDAPS_Failure_r17", err)
		}
	}
	return nil
}

func (ie *SHR_Cause_r17) Decode(r *uper.UperReader) error {
	var err error
	var T304_cause_r17Present bool
	if T304_cause_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var T310_cause_r17Present bool
	if T310_cause_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var T312_cause_r17Present bool
	if T312_cause_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SourceDAPS_Failure_r17Present bool
	if SourceDAPS_Failure_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if T304_cause_r17Present {
		ie.T304_cause_r17 = new(SHR_Cause_r17_t304_cause_r17)
		if err = ie.T304_cause_r17.Decode(r); err != nil {
			return utils.WrapError("Decode T304_cause_r17", err)
		}
	}
	if T310_cause_r17Present {
		ie.T310_cause_r17 = new(SHR_Cause_r17_t310_cause_r17)
		if err = ie.T310_cause_r17.Decode(r); err != nil {
			return utils.WrapError("Decode T310_cause_r17", err)
		}
	}
	if T312_cause_r17Present {
		ie.T312_cause_r17 = new(SHR_Cause_r17_t312_cause_r17)
		if err = ie.T312_cause_r17.Decode(r); err != nil {
			return utils.WrapError("Decode T312_cause_r17", err)
		}
	}
	if SourceDAPS_Failure_r17Present {
		ie.SourceDAPS_Failure_r17 = new(SHR_Cause_r17_sourceDAPS_Failure_r17)
		if err = ie.SourceDAPS_Failure_r17.Decode(r); err != nil {
			return utils.WrapError("Decode SourceDAPS_Failure_r17", err)
		}
	}
	return nil
}

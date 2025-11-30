package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RAN_VisibleParameters_r17 struct {
	Ran_VisiblePeriodicity_r17            *RAN_VisibleParameters_r17_ran_VisiblePeriodicity_r17 `optional`
	NumberOfBufferLevelEntries_r17        *int64                                                `lb:1,ub:8,optional`
	ReportPlayoutDelayForMediaStartup_r17 *bool                                                 `optional`
}

func (ie *RAN_VisibleParameters_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Ran_VisiblePeriodicity_r17 != nil, ie.NumberOfBufferLevelEntries_r17 != nil, ie.ReportPlayoutDelayForMediaStartup_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ran_VisiblePeriodicity_r17 != nil {
		if err = ie.Ran_VisiblePeriodicity_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ran_VisiblePeriodicity_r17", err)
		}
	}
	if ie.NumberOfBufferLevelEntries_r17 != nil {
		if err = w.WriteInteger(*ie.NumberOfBufferLevelEntries_r17, &aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode NumberOfBufferLevelEntries_r17", err)
		}
	}
	if ie.ReportPlayoutDelayForMediaStartup_r17 != nil {
		if err = w.WriteBoolean(*ie.ReportPlayoutDelayForMediaStartup_r17); err != nil {
			return utils.WrapError("Encode ReportPlayoutDelayForMediaStartup_r17", err)
		}
	}
	return nil
}

func (ie *RAN_VisibleParameters_r17) Decode(r *aper.AperReader) error {
	var err error
	var Ran_VisiblePeriodicity_r17Present bool
	if Ran_VisiblePeriodicity_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NumberOfBufferLevelEntries_r17Present bool
	if NumberOfBufferLevelEntries_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ReportPlayoutDelayForMediaStartup_r17Present bool
	if ReportPlayoutDelayForMediaStartup_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Ran_VisiblePeriodicity_r17Present {
		ie.Ran_VisiblePeriodicity_r17 = new(RAN_VisibleParameters_r17_ran_VisiblePeriodicity_r17)
		if err = ie.Ran_VisiblePeriodicity_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ran_VisiblePeriodicity_r17", err)
		}
	}
	if NumberOfBufferLevelEntries_r17Present {
		var tmp_int_NumberOfBufferLevelEntries_r17 int64
		if tmp_int_NumberOfBufferLevelEntries_r17, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode NumberOfBufferLevelEntries_r17", err)
		}
		ie.NumberOfBufferLevelEntries_r17 = &tmp_int_NumberOfBufferLevelEntries_r17
	}
	if ReportPlayoutDelayForMediaStartup_r17Present {
		var tmp_bool_ReportPlayoutDelayForMediaStartup_r17 bool
		if tmp_bool_ReportPlayoutDelayForMediaStartup_r17, err = r.ReadBoolean(); err != nil {
			return utils.WrapError("Decode ReportPlayoutDelayForMediaStartup_r17", err)
		}
		ie.ReportPlayoutDelayForMediaStartup_r17 = &tmp_bool_ReportPlayoutDelayForMediaStartup_r17
	}
	return nil
}

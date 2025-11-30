package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SearchSpaceExt_v1700 struct {
	MonitoringSlotPeriodicityAndOffset_v1710 *SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710 `lb:0,ub:31,optional`
	MonitoringSlotsWithinSlotGroup_r17       *SearchSpaceExt_v1700_monitoringSlotsWithinSlotGroup_r17       `lb:4,ub:4,optional`
	Duration_r17                             *int64                                                         `lb:4,ub:20476,optional`
	SearchSpaceType_r17                      *SearchSpaceExt_v1700_searchSpaceType_r17                      `optional`
	SearchSpaceGroupIdList_r17               []int64                                                        `lb:1,ub:3,e_lb:0,e_ub:maxNrofSearchSpaceGroups_1_r17,optional,ext`
	SearchSpaceLinkingId_r17                 *int64                                                         `lb:0,ub:maxNrofSearchSpacesLinks_1_r17,optional,ext`
}

func (ie *SearchSpaceExt_v1700) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.MonitoringSlotPeriodicityAndOffset_v1710 != nil, ie.MonitoringSlotsWithinSlotGroup_r17 != nil, ie.Duration_r17 != nil, ie.SearchSpaceType_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MonitoringSlotPeriodicityAndOffset_v1710 != nil {
		if err = ie.MonitoringSlotPeriodicityAndOffset_v1710.Encode(w); err != nil {
			return utils.WrapError("Encode MonitoringSlotPeriodicityAndOffset_v1710", err)
		}
	}
	if ie.MonitoringSlotsWithinSlotGroup_r17 != nil {
		if err = ie.MonitoringSlotsWithinSlotGroup_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MonitoringSlotsWithinSlotGroup_r17", err)
		}
	}
	if ie.Duration_r17 != nil {
		if err = w.WriteInteger(*ie.Duration_r17, &aper.Constraint{Lb: 4, Ub: 20476}, false); err != nil {
			return utils.WrapError("Encode Duration_r17", err)
		}
	}
	if ie.SearchSpaceType_r17 != nil {
		if err = ie.SearchSpaceType_r17.Encode(w); err != nil {
			return utils.WrapError("Encode SearchSpaceType_r17", err)
		}
	}
	return nil
}

func (ie *SearchSpaceExt_v1700) Decode(r *aper.AperReader) error {
	var err error
	var MonitoringSlotPeriodicityAndOffset_v1710Present bool
	if MonitoringSlotPeriodicityAndOffset_v1710Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MonitoringSlotsWithinSlotGroup_r17Present bool
	if MonitoringSlotsWithinSlotGroup_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Duration_r17Present bool
	if Duration_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SearchSpaceType_r17Present bool
	if SearchSpaceType_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if MonitoringSlotPeriodicityAndOffset_v1710Present {
		ie.MonitoringSlotPeriodicityAndOffset_v1710 = new(SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710)
		if err = ie.MonitoringSlotPeriodicityAndOffset_v1710.Decode(r); err != nil {
			return utils.WrapError("Decode MonitoringSlotPeriodicityAndOffset_v1710", err)
		}
	}
	if MonitoringSlotsWithinSlotGroup_r17Present {
		ie.MonitoringSlotsWithinSlotGroup_r17 = new(SearchSpaceExt_v1700_monitoringSlotsWithinSlotGroup_r17)
		if err = ie.MonitoringSlotsWithinSlotGroup_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MonitoringSlotsWithinSlotGroup_r17", err)
		}
	}
	if Duration_r17Present {
		var tmp_int_Duration_r17 int64
		if tmp_int_Duration_r17, err = r.ReadInteger(&aper.Constraint{Lb: 4, Ub: 20476}, false); err != nil {
			return utils.WrapError("Decode Duration_r17", err)
		}
		ie.Duration_r17 = &tmp_int_Duration_r17
	}
	if SearchSpaceType_r17Present {
		ie.SearchSpaceType_r17 = new(SearchSpaceExt_v1700_searchSpaceType_r17)
		if err = ie.SearchSpaceType_r17.Decode(r); err != nil {
			return utils.WrapError("Decode SearchSpaceType_r17", err)
		}
	}
	return nil
}

package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SearchSpace struct {
	SearchSpaceId                      SearchSpaceId                                   `madatory`
	ControlResourceSetId               *ControlResourceSetId                           `optional`
	MonitoringSlotPeriodicityAndOffset *SearchSpace_monitoringSlotPeriodicityAndOffset `lb:0,ub:1,optional`
	Duration                           *int64                                          `lb:2,ub:2559,optional`
	MonitoringSymbolsWithinSlot        *aper.BitString                                 `lb:14,ub:14,optional`
	NrofCandidates                     *SearchSpace_nrofCandidates                     `optional`
	SearchSpaceType                    *SearchSpace_searchSpaceType                    `optional`
}

func (ie *SearchSpace) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.ControlResourceSetId != nil, ie.MonitoringSlotPeriodicityAndOffset != nil, ie.Duration != nil, ie.MonitoringSymbolsWithinSlot != nil, ie.NrofCandidates != nil, ie.SearchSpaceType != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.SearchSpaceId.Encode(w); err != nil {
		return utils.WrapError("Encode SearchSpaceId", err)
	}
	if ie.ControlResourceSetId != nil {
		if err = ie.ControlResourceSetId.Encode(w); err != nil {
			return utils.WrapError("Encode ControlResourceSetId", err)
		}
	}
	if ie.MonitoringSlotPeriodicityAndOffset != nil {
		if err = ie.MonitoringSlotPeriodicityAndOffset.Encode(w); err != nil {
			return utils.WrapError("Encode MonitoringSlotPeriodicityAndOffset", err)
		}
	}
	if ie.Duration != nil {
		if err = w.WriteInteger(*ie.Duration, &aper.Constraint{Lb: 2, Ub: 2559}, false); err != nil {
			return utils.WrapError("Encode Duration", err)
		}
	}
	if ie.MonitoringSymbolsWithinSlot != nil {
		if err = w.WriteBitString(ie.MonitoringSymbolsWithinSlot.Bytes, uint(ie.MonitoringSymbolsWithinSlot.NumBits), &aper.Constraint{Lb: 14, Ub: 14}, false); err != nil {
			return utils.WrapError("Encode MonitoringSymbolsWithinSlot", err)
		}
	}
	if ie.NrofCandidates != nil {
		if err = ie.NrofCandidates.Encode(w); err != nil {
			return utils.WrapError("Encode NrofCandidates", err)
		}
	}
	if ie.SearchSpaceType != nil {
		if err = ie.SearchSpaceType.Encode(w); err != nil {
			return utils.WrapError("Encode SearchSpaceType", err)
		}
	}
	return nil
}

func (ie *SearchSpace) Decode(r *aper.AperReader) error {
	var err error
	var ControlResourceSetIdPresent bool
	if ControlResourceSetIdPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MonitoringSlotPeriodicityAndOffsetPresent bool
	if MonitoringSlotPeriodicityAndOffsetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var DurationPresent bool
	if DurationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MonitoringSymbolsWithinSlotPresent bool
	if MonitoringSymbolsWithinSlotPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NrofCandidatesPresent bool
	if NrofCandidatesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SearchSpaceTypePresent bool
	if SearchSpaceTypePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.SearchSpaceId.Decode(r); err != nil {
		return utils.WrapError("Decode SearchSpaceId", err)
	}
	if ControlResourceSetIdPresent {
		ie.ControlResourceSetId = new(ControlResourceSetId)
		if err = ie.ControlResourceSetId.Decode(r); err != nil {
			return utils.WrapError("Decode ControlResourceSetId", err)
		}
	}
	if MonitoringSlotPeriodicityAndOffsetPresent {
		ie.MonitoringSlotPeriodicityAndOffset = new(SearchSpace_monitoringSlotPeriodicityAndOffset)
		if err = ie.MonitoringSlotPeriodicityAndOffset.Decode(r); err != nil {
			return utils.WrapError("Decode MonitoringSlotPeriodicityAndOffset", err)
		}
	}
	if DurationPresent {
		var tmp_int_Duration int64
		if tmp_int_Duration, err = r.ReadInteger(&aper.Constraint{Lb: 2, Ub: 2559}, false); err != nil {
			return utils.WrapError("Decode Duration", err)
		}
		ie.Duration = &tmp_int_Duration
	}
	if MonitoringSymbolsWithinSlotPresent {
		var tmp_bs_MonitoringSymbolsWithinSlot []byte
		var n_MonitoringSymbolsWithinSlot uint
		if tmp_bs_MonitoringSymbolsWithinSlot, n_MonitoringSymbolsWithinSlot, err = r.ReadBitString(&aper.Constraint{Lb: 14, Ub: 14}, false); err != nil {
			return utils.WrapError("Decode MonitoringSymbolsWithinSlot", err)
		}
		tmp_bitstring := aper.BitString{
			Bytes:   tmp_bs_MonitoringSymbolsWithinSlot,
			NumBits: uint64(n_MonitoringSymbolsWithinSlot),
		}
		ie.MonitoringSymbolsWithinSlot = &tmp_bitstring
	}
	if NrofCandidatesPresent {
		ie.NrofCandidates = new(SearchSpace_nrofCandidates)
		if err = ie.NrofCandidates.Decode(r); err != nil {
			return utils.WrapError("Decode NrofCandidates", err)
		}
	}
	if SearchSpaceTypePresent {
		ie.SearchSpaceType = new(SearchSpace_searchSpaceType)
		if err = ie.SearchSpaceType.Decode(r); err != nil {
			return utils.WrapError("Decode SearchSpaceType", err)
		}
	}
	return nil
}

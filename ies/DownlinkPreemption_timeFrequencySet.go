package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	DownlinkPreemption_timeFrequencySet_Enum_set0 aper.Enumerated = 0
	DownlinkPreemption_timeFrequencySet_Enum_set1 aper.Enumerated = 1
)

type DownlinkPreemption_timeFrequencySet struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *DownlinkPreemption_timeFrequencySet) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode DownlinkPreemption_timeFrequencySet", err)
	}
	return nil
}

func (ie *DownlinkPreemption_timeFrequencySet) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode DownlinkPreemption_timeFrequencySet", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}

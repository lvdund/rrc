package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SSB_MTC2_periodicity_Enum_sf5    aper.Enumerated = 0
	SSB_MTC2_periodicity_Enum_sf10   aper.Enumerated = 1
	SSB_MTC2_periodicity_Enum_sf20   aper.Enumerated = 2
	SSB_MTC2_periodicity_Enum_sf40   aper.Enumerated = 3
	SSB_MTC2_periodicity_Enum_sf80   aper.Enumerated = 4
	SSB_MTC2_periodicity_Enum_spare3 aper.Enumerated = 5
	SSB_MTC2_periodicity_Enum_spare2 aper.Enumerated = 6
	SSB_MTC2_periodicity_Enum_spare1 aper.Enumerated = 7
)

type SSB_MTC2_periodicity struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SSB_MTC2_periodicity) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SSB_MTC2_periodicity", err)
	}
	return nil
}

func (ie *SSB_MTC2_periodicity) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SSB_MTC2_periodicity", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}

package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDCP_Parameters_maxNumberROHC_ContextSessions_Enum_cs2     aper.Enumerated = 0
	PDCP_Parameters_maxNumberROHC_ContextSessions_Enum_cs4     aper.Enumerated = 1
	PDCP_Parameters_maxNumberROHC_ContextSessions_Enum_cs8     aper.Enumerated = 2
	PDCP_Parameters_maxNumberROHC_ContextSessions_Enum_cs12    aper.Enumerated = 3
	PDCP_Parameters_maxNumberROHC_ContextSessions_Enum_cs16    aper.Enumerated = 4
	PDCP_Parameters_maxNumberROHC_ContextSessions_Enum_cs24    aper.Enumerated = 5
	PDCP_Parameters_maxNumberROHC_ContextSessions_Enum_cs32    aper.Enumerated = 6
	PDCP_Parameters_maxNumberROHC_ContextSessions_Enum_cs48    aper.Enumerated = 7
	PDCP_Parameters_maxNumberROHC_ContextSessions_Enum_cs64    aper.Enumerated = 8
	PDCP_Parameters_maxNumberROHC_ContextSessions_Enum_cs128   aper.Enumerated = 9
	PDCP_Parameters_maxNumberROHC_ContextSessions_Enum_cs256   aper.Enumerated = 10
	PDCP_Parameters_maxNumberROHC_ContextSessions_Enum_cs512   aper.Enumerated = 11
	PDCP_Parameters_maxNumberROHC_ContextSessions_Enum_cs1024  aper.Enumerated = 12
	PDCP_Parameters_maxNumberROHC_ContextSessions_Enum_cs16384 aper.Enumerated = 13
	PDCP_Parameters_maxNumberROHC_ContextSessions_Enum_spare2  aper.Enumerated = 14
	PDCP_Parameters_maxNumberROHC_ContextSessions_Enum_spare1  aper.Enumerated = 15
)

type PDCP_Parameters_maxNumberROHC_ContextSessions struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *PDCP_Parameters_maxNumberROHC_ContextSessions) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode PDCP_Parameters_maxNumberROHC_ContextSessions", err)
	}
	return nil
}

func (ie *PDCP_Parameters_maxNumberROHC_ContextSessions) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode PDCP_Parameters_maxNumberROHC_ContextSessions", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}

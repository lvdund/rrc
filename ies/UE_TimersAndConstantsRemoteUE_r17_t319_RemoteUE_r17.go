package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	UE_TimersAndConstantsRemoteUE_r17_t319_RemoteUE_r17_Enum_ms100  aper.Enumerated = 0
	UE_TimersAndConstantsRemoteUE_r17_t319_RemoteUE_r17_Enum_ms200  aper.Enumerated = 1
	UE_TimersAndConstantsRemoteUE_r17_t319_RemoteUE_r17_Enum_ms300  aper.Enumerated = 2
	UE_TimersAndConstantsRemoteUE_r17_t319_RemoteUE_r17_Enum_ms400  aper.Enumerated = 3
	UE_TimersAndConstantsRemoteUE_r17_t319_RemoteUE_r17_Enum_ms600  aper.Enumerated = 4
	UE_TimersAndConstantsRemoteUE_r17_t319_RemoteUE_r17_Enum_ms1000 aper.Enumerated = 5
	UE_TimersAndConstantsRemoteUE_r17_t319_RemoteUE_r17_Enum_ms1500 aper.Enumerated = 6
	UE_TimersAndConstantsRemoteUE_r17_t319_RemoteUE_r17_Enum_ms2000 aper.Enumerated = 7
)

type UE_TimersAndConstantsRemoteUE_r17_t319_RemoteUE_r17 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *UE_TimersAndConstantsRemoteUE_r17_t319_RemoteUE_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode UE_TimersAndConstantsRemoteUE_r17_t319_RemoteUE_r17", err)
	}
	return nil
}

func (ie *UE_TimersAndConstantsRemoteUE_r17_t319_RemoteUE_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode UE_TimersAndConstantsRemoteUE_r17_t319_RemoteUE_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}

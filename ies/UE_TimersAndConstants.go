package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UE_TimersAndConstants struct {
	T300 UE_TimersAndConstants_t300 `madatory`
	T301 UE_TimersAndConstants_t301 `madatory`
	T310 UE_TimersAndConstants_t310 `madatory`
	N310 UE_TimersAndConstants_n310 `madatory`
	T311 UE_TimersAndConstants_t311 `madatory`
	N311 UE_TimersAndConstants_n311 `madatory`
	T319 UE_TimersAndConstants_t319 `madatory`
}

func (ie *UE_TimersAndConstants) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.T300.Encode(w); err != nil {
		return utils.WrapError("Encode T300", err)
	}
	if err = ie.T301.Encode(w); err != nil {
		return utils.WrapError("Encode T301", err)
	}
	if err = ie.T310.Encode(w); err != nil {
		return utils.WrapError("Encode T310", err)
	}
	if err = ie.N310.Encode(w); err != nil {
		return utils.WrapError("Encode N310", err)
	}
	if err = ie.T311.Encode(w); err != nil {
		return utils.WrapError("Encode T311", err)
	}
	if err = ie.N311.Encode(w); err != nil {
		return utils.WrapError("Encode N311", err)
	}
	if err = ie.T319.Encode(w); err != nil {
		return utils.WrapError("Encode T319", err)
	}
	return nil
}

func (ie *UE_TimersAndConstants) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.T300.Decode(r); err != nil {
		return utils.WrapError("Decode T300", err)
	}
	if err = ie.T301.Decode(r); err != nil {
		return utils.WrapError("Decode T301", err)
	}
	if err = ie.T310.Decode(r); err != nil {
		return utils.WrapError("Decode T310", err)
	}
	if err = ie.N310.Decode(r); err != nil {
		return utils.WrapError("Decode N310", err)
	}
	if err = ie.T311.Decode(r); err != nil {
		return utils.WrapError("Decode T311", err)
	}
	if err = ie.N311.Decode(r); err != nil {
		return utils.WrapError("Decode N311", err)
	}
	if err = ie.T319.Decode(r); err != nil {
		return utils.WrapError("Decode T319", err)
	}
	return nil
}

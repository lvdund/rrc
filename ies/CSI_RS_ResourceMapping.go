package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CSI_RS_ResourceMapping struct {
	FrequencyDomainAllocation    CSI_RS_ResourceMapping_frequencyDomainAllocation `lb:4,ub:4,madatory`
	NrofPorts                    CSI_RS_ResourceMapping_nrofPorts                 `madatory`
	FirstOFDMSymbolInTimeDomain  int64                                            `lb:0,ub:13,madatory`
	FirstOFDMSymbolInTimeDomain2 *int64                                           `lb:2,ub:12,optional`
	Cdm_Type                     CSI_RS_ResourceMapping_cdm_Type                  `madatory`
	Density                      CSI_RS_ResourceMapping_density                   `madatory`
	FreqBand                     CSI_FrequencyOccupation                          `madatory`
}

func (ie *CSI_RS_ResourceMapping) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.FirstOFDMSymbolInTimeDomain2 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.FrequencyDomainAllocation.Encode(w); err != nil {
		return utils.WrapError("Encode FrequencyDomainAllocation", err)
	}
	if err = ie.NrofPorts.Encode(w); err != nil {
		return utils.WrapError("Encode NrofPorts", err)
	}
	if err = w.WriteInteger(ie.FirstOFDMSymbolInTimeDomain, &aper.Constraint{Lb: 0, Ub: 13}, false); err != nil {
		return utils.WrapError("WriteInteger FirstOFDMSymbolInTimeDomain", err)
	}
	if ie.FirstOFDMSymbolInTimeDomain2 != nil {
		if err = w.WriteInteger(*ie.FirstOFDMSymbolInTimeDomain2, &aper.Constraint{Lb: 2, Ub: 12}, false); err != nil {
			return utils.WrapError("Encode FirstOFDMSymbolInTimeDomain2", err)
		}
	}
	if err = ie.Cdm_Type.Encode(w); err != nil {
		return utils.WrapError("Encode Cdm_Type", err)
	}
	if err = ie.Density.Encode(w); err != nil {
		return utils.WrapError("Encode Density", err)
	}
	if err = ie.FreqBand.Encode(w); err != nil {
		return utils.WrapError("Encode FreqBand", err)
	}
	return nil
}

func (ie *CSI_RS_ResourceMapping) Decode(r *aper.AperReader) error {
	var err error
	var FirstOFDMSymbolInTimeDomain2Present bool
	if FirstOFDMSymbolInTimeDomain2Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.FrequencyDomainAllocation.Decode(r); err != nil {
		return utils.WrapError("Decode FrequencyDomainAllocation", err)
	}
	if err = ie.NrofPorts.Decode(r); err != nil {
		return utils.WrapError("Decode NrofPorts", err)
	}
	var tmp_int_FirstOFDMSymbolInTimeDomain int64
	if tmp_int_FirstOFDMSymbolInTimeDomain, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 13}, false); err != nil {
		return utils.WrapError("ReadInteger FirstOFDMSymbolInTimeDomain", err)
	}
	ie.FirstOFDMSymbolInTimeDomain = tmp_int_FirstOFDMSymbolInTimeDomain
	if FirstOFDMSymbolInTimeDomain2Present {
		var tmp_int_FirstOFDMSymbolInTimeDomain2 int64
		if tmp_int_FirstOFDMSymbolInTimeDomain2, err = r.ReadInteger(&aper.Constraint{Lb: 2, Ub: 12}, false); err != nil {
			return utils.WrapError("Decode FirstOFDMSymbolInTimeDomain2", err)
		}
		ie.FirstOFDMSymbolInTimeDomain2 = &tmp_int_FirstOFDMSymbolInTimeDomain2
	}
	if err = ie.Cdm_Type.Decode(r); err != nil {
		return utils.WrapError("Decode Cdm_Type", err)
	}
	if err = ie.Density.Decode(r); err != nil {
		return utils.WrapError("Decode Density", err)
	}
	if err = ie.FreqBand.Decode(r); err != nil {
		return utils.WrapError("Decode FreqBand", err)
	}
	return nil
}

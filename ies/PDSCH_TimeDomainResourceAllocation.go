package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PDSCH_TimeDomainResourceAllocation struct {
	K0                   *int64                                         `lb:0,ub:32,optional`
	MappingType          PDSCH_TimeDomainResourceAllocation_mappingType `madatory`
	StartSymbolAndLength int64                                          `lb:0,ub:127,madatory`
}

func (ie *PDSCH_TimeDomainResourceAllocation) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.K0 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.K0 != nil {
		if err = w.WriteInteger(*ie.K0, &aper.Constraint{Lb: 0, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode K0", err)
		}
	}
	if err = ie.MappingType.Encode(w); err != nil {
		return utils.WrapError("Encode MappingType", err)
	}
	if err = w.WriteInteger(ie.StartSymbolAndLength, &aper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
		return utils.WrapError("WriteInteger StartSymbolAndLength", err)
	}
	return nil
}

func (ie *PDSCH_TimeDomainResourceAllocation) Decode(r *aper.AperReader) error {
	var err error
	var K0Present bool
	if K0Present, err = r.ReadBool(); err != nil {
		return err
	}
	if K0Present {
		var tmp_int_K0 int64
		if tmp_int_K0, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode K0", err)
		}
		ie.K0 = &tmp_int_K0
	}
	if err = ie.MappingType.Decode(r); err != nil {
		return utils.WrapError("Decode MappingType", err)
	}
	var tmp_int_StartSymbolAndLength int64
	if tmp_int_StartSymbolAndLength, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
		return utils.WrapError("ReadInteger StartSymbolAndLength", err)
	}
	ie.StartSymbolAndLength = tmp_int_StartSymbolAndLength
	return nil
}

package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SearchSpace_searchSpaceType_common_dci_Format2_3_dummy1_Enum_sl1  aper.Enumerated = 0
	SearchSpace_searchSpaceType_common_dci_Format2_3_dummy1_Enum_sl2  aper.Enumerated = 1
	SearchSpace_searchSpaceType_common_dci_Format2_3_dummy1_Enum_sl4  aper.Enumerated = 2
	SearchSpace_searchSpaceType_common_dci_Format2_3_dummy1_Enum_sl5  aper.Enumerated = 3
	SearchSpace_searchSpaceType_common_dci_Format2_3_dummy1_Enum_sl8  aper.Enumerated = 4
	SearchSpace_searchSpaceType_common_dci_Format2_3_dummy1_Enum_sl10 aper.Enumerated = 5
	SearchSpace_searchSpaceType_common_dci_Format2_3_dummy1_Enum_sl16 aper.Enumerated = 6
	SearchSpace_searchSpaceType_common_dci_Format2_3_dummy1_Enum_sl20 aper.Enumerated = 7
)

type SearchSpace_searchSpaceType_common_dci_Format2_3_dummy1 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SearchSpace_searchSpaceType_common_dci_Format2_3_dummy1) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SearchSpace_searchSpaceType_common_dci_Format2_3_dummy1", err)
	}
	return nil
}

func (ie *SearchSpace_searchSpaceType_common_dci_Format2_3_dummy1) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SearchSpace_searchSpaceType_common_dci_Format2_3_dummy1", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}

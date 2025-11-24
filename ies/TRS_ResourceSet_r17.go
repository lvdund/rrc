package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type TRS_ResourceSet_r17 struct {
	PowerControlOffsetSS_r17        TRS_ResourceSet_r17_powerControlOffsetSS_r17 `madatory`
	ScramblingID_Info_r17           TRS_ResourceSet_r17_scramblingID_Info_r17    `lb:2,ub:2,madatory`
	FirstOFDMSymbolInTimeDomain_r17 int64                                        `lb:0,ub:9,madatory,ext`
	StartingRB_r17                  int64                                        `lb:0,ub:maxNrofPhysicalResourceBlocks_1,madatory,ext`
	NrofRBs_r17                     int64                                        `lb:24,ub:maxNrofPhysicalResourceBlocksPlus1,madatory,ext`
	Ssb_Index_r17                   SSB_Index                                    `madatory,ext`
	PeriodicityAndOffset_r17        TRS_ResourceSet_r17_periodicityAndOffset_r17 `lb:0,ub:9,madatory,ext`
	FrequencyDomainAllocation_r17   uper.BitString                               `lb:4,ub:4,madatory,ext`
	IndBitID_r17                    int64                                        `lb:0,ub:5,madatory,ext`
	NrofResources_r17               TRS_ResourceSet_r17_nrofResources_r17        `madatory,ext`
}

func (ie *TRS_ResourceSet_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.PowerControlOffsetSS_r17.Encode(w); err != nil {
		return utils.WrapError("Encode PowerControlOffsetSS_r17", err)
	}
	if err = ie.ScramblingID_Info_r17.Encode(w); err != nil {
		return utils.WrapError("Encode ScramblingID_Info_r17", err)
	}
	return nil
}

func (ie *TRS_ResourceSet_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.PowerControlOffsetSS_r17.Decode(r); err != nil {
		return utils.WrapError("Decode PowerControlOffsetSS_r17", err)
	}
	if err = ie.ScramblingID_Info_r17.Decode(r); err != nil {
		return utils.WrapError("Decode ScramblingID_Info_r17", err)
	}
	return nil
}

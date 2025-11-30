package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDCCH_ConfigCommon_sdt_SearchSpace_r17_Choice_nothing uint64 = iota
	PDCCH_ConfigCommon_sdt_SearchSpace_r17_Choice_NewSearchSpace
	PDCCH_ConfigCommon_sdt_SearchSpace_r17_Choice_ExistingSearchSpace
)

type PDCCH_ConfigCommon_sdt_SearchSpace_r17 struct {
	Choice              uint64
	NewSearchSpace      *SearchSpace
	ExistingSearchSpace *SearchSpaceId
}

func (ie *PDCCH_ConfigCommon_sdt_SearchSpace_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PDCCH_ConfigCommon_sdt_SearchSpace_r17_Choice_NewSearchSpace:
		if err = ie.NewSearchSpace.Encode(w); err != nil {
			err = utils.WrapError("Encode NewSearchSpace", err)
		}
	case PDCCH_ConfigCommon_sdt_SearchSpace_r17_Choice_ExistingSearchSpace:
		if err = ie.ExistingSearchSpace.Encode(w); err != nil {
			err = utils.WrapError("Encode ExistingSearchSpace", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *PDCCH_ConfigCommon_sdt_SearchSpace_r17) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PDCCH_ConfigCommon_sdt_SearchSpace_r17_Choice_NewSearchSpace:
		ie.NewSearchSpace = new(SearchSpace)
		if err = ie.NewSearchSpace.Decode(r); err != nil {
			return utils.WrapError("Decode NewSearchSpace", err)
		}
	case PDCCH_ConfigCommon_sdt_SearchSpace_r17_Choice_ExistingSearchSpace:
		ie.ExistingSearchSpace = new(SearchSpaceId)
		if err = ie.ExistingSearchSpace.Decode(r); err != nil {
			return utils.WrapError("Decode ExistingSearchSpace", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}

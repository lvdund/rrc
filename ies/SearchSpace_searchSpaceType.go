package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SearchSpace_searchSpaceType_Choice_nothing uint64 = iota
	SearchSpace_searchSpaceType_Choice_Common
	SearchSpace_searchSpaceType_Choice_Ue_Specific
)

type SearchSpace_searchSpaceType struct {
	Choice      uint64
	Common      *SearchSpace_searchSpaceType_common
	Ue_Specific *SearchSpace_searchSpaceType_ue_Specific
}

func (ie *SearchSpace_searchSpaceType) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SearchSpace_searchSpaceType_Choice_Common:
		if err = ie.Common.Encode(w); err != nil {
			err = utils.WrapError("Encode Common", err)
		}
	case SearchSpace_searchSpaceType_Choice_Ue_Specific:
		if err = ie.Ue_Specific.Encode(w); err != nil {
			err = utils.WrapError("Encode Ue_Specific", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SearchSpace_searchSpaceType) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SearchSpace_searchSpaceType_Choice_Common:
		ie.Common = new(SearchSpace_searchSpaceType_common)
		if err = ie.Common.Decode(r); err != nil {
			return utils.WrapError("Decode Common", err)
		}
	case SearchSpace_searchSpaceType_Choice_Ue_Specific:
		ie.Ue_Specific = new(SearchSpace_searchSpaceType_ue_Specific)
		if err = ie.Ue_Specific.Decode(r); err != nil {
			return utils.WrapError("Decode Ue_Specific", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}

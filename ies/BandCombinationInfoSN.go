package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombinationInfoSN struct {
	BandCombinationIndex BandCombinationIndex `madatory`
	RequestedFeatureSets FeatureSetEntryIndex `madatory`
}

func (ie *BandCombinationInfoSN) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.BandCombinationIndex.Encode(w); err != nil {
		return utils.WrapError("Encode BandCombinationIndex", err)
	}
	if err = ie.RequestedFeatureSets.Encode(w); err != nil {
		return utils.WrapError("Encode RequestedFeatureSets", err)
	}
	return nil
}

func (ie *BandCombinationInfoSN) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.BandCombinationIndex.Decode(r); err != nil {
		return utils.WrapError("Decode BandCombinationIndex", err)
	}
	if err = ie.RequestedFeatureSets.Decode(r); err != nil {
		return utils.WrapError("Decode RequestedFeatureSets", err)
	}
	return nil
}

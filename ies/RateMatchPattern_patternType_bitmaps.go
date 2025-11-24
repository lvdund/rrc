package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RateMatchPattern_patternType_bitmaps struct {
	ResourceBlocks         uper.BitString                                              `lb:275,ub:275,madatory`
	SymbolsInResourceBlock RateMatchPattern_patternType_bitmaps_symbolsInResourceBlock `lb:14,ub:14,madatory`
	PeriodicityAndPattern  *RateMatchPattern_patternType_bitmaps_periodicityAndPattern `lb:2,ub:2,optional`
}

func (ie *RateMatchPattern_patternType_bitmaps) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.PeriodicityAndPattern != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteBitString(ie.ResourceBlocks.Bytes, uint(ie.ResourceBlocks.NumBits), &uper.Constraint{Lb: 275, Ub: 275}, false); err != nil {
		return utils.WrapError("WriteBitString ResourceBlocks", err)
	}
	if err = ie.SymbolsInResourceBlock.Encode(w); err != nil {
		return utils.WrapError("Encode SymbolsInResourceBlock", err)
	}
	if ie.PeriodicityAndPattern != nil {
		if err = ie.PeriodicityAndPattern.Encode(w); err != nil {
			return utils.WrapError("Encode PeriodicityAndPattern", err)
		}
	}
	return nil
}

func (ie *RateMatchPattern_patternType_bitmaps) Decode(r *uper.UperReader) error {
	var err error
	var PeriodicityAndPatternPresent bool
	if PeriodicityAndPatternPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_bs_ResourceBlocks []byte
	var n_ResourceBlocks uint
	if tmp_bs_ResourceBlocks, n_ResourceBlocks, err = r.ReadBitString(&uper.Constraint{Lb: 275, Ub: 275}, false); err != nil {
		return utils.WrapError("ReadBitString ResourceBlocks", err)
	}
	ie.ResourceBlocks = uper.BitString{
		Bytes:   tmp_bs_ResourceBlocks,
		NumBits: uint64(n_ResourceBlocks),
	}
	if err = ie.SymbolsInResourceBlock.Decode(r); err != nil {
		return utils.WrapError("Decode SymbolsInResourceBlock", err)
	}
	if PeriodicityAndPatternPresent {
		ie.PeriodicityAndPattern = new(RateMatchPattern_patternType_bitmaps_periodicityAndPattern)
		if err = ie.PeriodicityAndPattern.Decode(r); err != nil {
			return utils.WrapError("Decode PeriodicityAndPattern", err)
		}
	}
	return nil
}

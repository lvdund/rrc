package ies

import "rrc/utils"

// SRS-Config ::= SEQUENCE
// Extensible
type SrsConfig struct {
	SrsResourcesettoreleaselist         *[]SrsResourcesetid `lb:1,ub:maxNrofSRSResourcesets`
	SrsResourcesettoaddmodlist          *[]SrsResourceset   `lb:1,ub:maxNrofSRSResourcesets`
	SrsResourcetoreleaselist            *[]SrsResourceid    `lb:1,ub:maxNrofSRSResources`
	SrsResourcetoaddmodlist             *[]SrsResource      `lb:1,ub:maxNrofSRSResources`
	TpcAccumulation                     *SrsConfigTpcAccumulation
	SrsRequestdci12R16                  *utils.INTEGER            `lb:0,ub:2,ext`
	SrsRequestdci02R16                  *utils.INTEGER            `lb:0,ub:2,ext`
	SrsResourcesettoaddmodlistdci02R16  *[]SrsResourceset         `lb:1,ub:maxNrofSRSResourcesets,ext`
	SrsResourcesettoreleaselistdci02R16 *[]SrsResourcesetid       `lb:1,ub:maxNrofSRSResourcesets,ext`
	SrsPosresourcesettoreleaselistR16   *[]SrsPosresourcesetidR16 `lb:1,ub:maxNrofSRSPosresourcesetsR16,ext`
	SrsPosresourcesettoaddmodlistR16    *[]SrsPosresourcesetR16   `lb:1,ub:maxNrofSRSPosresourcesetsR16,ext`
	SrsPosresourcetoreleaselistR16      *[]SrsPosresourceidR16    `lb:1,ub:maxNrofSRSPosresourcesR16,ext`
	SrsPosresourcetoaddmodlistR16       *[]SrsPosresourceR16      `lb:1,ub:maxNrofSRSPosresourcesR16,ext`
}

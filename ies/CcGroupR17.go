package ies

// CC-Group-r17 ::= SEQUENCE
type CcGroupR17 struct {
	ServcellindexlowerR17  Servcellindex
	ServcellindexhigherR17 *Servcellindex
	DefaultdcLocationR17   DefaultdcLocationR17
	OffsettodefaultR17     *CcGroupR17OffsettodefaultR17 `lb:1,ub:maxNrofReqComDCLocationR17`
}

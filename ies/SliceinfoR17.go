package ies

// SliceInfo-r17 ::= SEQUENCE
type SliceinfoR17 struct {
	NsagIdentityinfoR17               NsagIdentityinfoR17
	NsagCellreselectionpriorityR17    *Cellreselectionpriority
	NsagCellreselectionsubpriorityR17 *Cellreselectionsubpriority
	SlicecelllistnrR17                *SliceinfoR17SlicecelllistnrR17
}

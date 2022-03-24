package datastructure

type Trier interface {
	Compare(Trier) CompareResult
	GetChildOf(Trier) (Trier, error)
	CanBeRoot() bool
	Less(Trier) bool
	Add(Trier) (Trier, error)
}

type CompareResult int64

const (
	Equal CompareResult = iota
	IsParent
	IsAncestor
	IsChild
	IsDescendent
	Unrelated
)

func (CompareResult CompareResult) ToString() string {
	return [...]string{"Equal", "IsParent", "IsAncestor", "IsChild", "IsDescendent", "Unrelated"}[CompareResult]
}

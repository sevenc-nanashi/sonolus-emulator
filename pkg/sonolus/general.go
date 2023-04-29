package sonolus

type ItemDetails[T any] struct {
  Item T `json:"item"`
  Description string `json:"description"`
  Recommended []T `json:"recommended"`
}


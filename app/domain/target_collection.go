package domain

//easyjson:json
type TargetCollection []Target

func (c TargetCollection) Verify() bool {
	var ok bool
	for _, target := range c {
		if ok = target.Verify(); !ok {
			return false
		}
	}

	return true
}

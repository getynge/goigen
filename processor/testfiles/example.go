package testfiles

type Example struct{}

func (e Example) NotIncluded() {}
func (e *Example) Included()   {}
func (*Example) AlsoIncluded() {}

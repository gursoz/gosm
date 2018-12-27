package trigger

type trigger struct {
	id    int
	value string
}

func (tr1 trigger) Compare(tr2 Trigger) bool {
	if tr21, ok := tr2.(trigger); ok {
		if tr1.id == tr21.id {
			return true
		}
	}
	return false
}

func (tr1 trigger) NewTrigger(id int, value string) *Trigger {
	newTr := new(trigger)
	newTr.id = id
	newTr.value = value
	interfaceTrigger := (Trigger)(*newTr)
	return &interfaceTrigger
}


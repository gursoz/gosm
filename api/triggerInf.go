package statemachine

// Trigger Trigger which may cause a state transisition to happen
type Trigger interface {
	Compare(trigger Trigger) bool
}

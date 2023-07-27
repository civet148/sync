package sync

type NoCopy struct{}

func (n *NoCopy) Lock()   {}
func (n *NoCopy) Unlock() {}

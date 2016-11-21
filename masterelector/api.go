package masterelector

type Cluster struct {
	peers []string
	ch chan struct{}
}

// Init creates a new master election cluster. It doesn't do anything until Run is called.
func Init(secret string, peers []string) (*Cluster, error) {
	return &Cluster{peers}, nil
}

// Run starts connecting to others and never returns unless Stop is called.
func (c *Cluster) Run() {
	for {
		c.WaitForMastership()
		c.ch = make(chan struct{})
		// wait till we lose mastership
		close(c.ch)
	}
}

// Stop lets Run return.
func (c *Cluster) Stop() {
}

// WaitForMastership blocks until this server becomes the master.
func (c *Cluster) WaitForMastership() {
}

// ElectedChannel returns a channel that blocks as long as we're master and gets closed when we lose mastership.
func (c *Cluster) ElectedChannel() <-chan struct{} {
	return c.ch
}

func (c *Cluster) SendToMaster(msg interface{}) {
}

func (c *Cluster) SendToMinions(msg interface{}) {
}

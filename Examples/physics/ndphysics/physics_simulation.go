package ndphysics

type PhysicsSimulation struct {
	physicsBodies []*PhysicsBody //store all our physics bodies
}

func NewSimulation() PhysicsSimulation {
	return PhysicsSimulation{physicsBodies: make([]*PhysicsBody, 0, 100)}
}

func (ps *PhysicsSimulation) AddPhysicsBody(pb *PhysicsBody) {
	ps.physicsBodies = append(ps.physicsBodies, pb)
}

func (ps *PhysicsSimulation) Simualte() {

	for i, pb := range ps.physicsBodies {
		pb.PhysicsUpdate()
		for j, pbOther := range ps.physicsBodies {
			if i == j {
				continue
			}
			pb.CheckIntersection(pbOther)
		}
	}
}

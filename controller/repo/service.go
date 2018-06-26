package repo

import "github.com/itsubaki/interstellar/controller"

type ServiceRepository struct {
	Reposotory []*controller.Service
}

func (r *ServiceRepository) Insert(s *controller.Service) {
	r.Reposotory = append(r.Reposotory, s)
}

func (r *ServiceRepository) SelectAll() []*controller.Service {
	return r.Reposotory
}

func (r *ServiceRepository) FindByName(name string) (*controller.Service, bool) {
	// TODO O(N)
	for i := range r.Reposotory {
		if r.Reposotory[i].Name != name {
			continue
		}
		return r.Reposotory[i], true
	}

	return nil, false
}

func (r *ServiceRepository) FindByInstanceID(id string) (*controller.Service, bool) {
	// TODO O(N)
	for i := range r.Reposotory {
		if r.Reposotory[i].ServiceID != id {
			continue
		}
		return r.Reposotory[i], true
	}

	return nil, false
}

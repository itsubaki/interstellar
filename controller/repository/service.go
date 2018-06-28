package repository

import "github.com/itsubaki/interstellar/controller"

type ServiceRepository struct {
	Repository []*controller.Service
}

func (r *ServiceRepository) Insert(s *controller.Service) {
	r.Repository = append(r.Repository, s)
}

func (r *ServiceRepository) SelectAll() []*controller.Service {
	return r.Repository
}

func (r *ServiceRepository) FindByName(name string) (*controller.Service, bool) {
	// TODO O(N)
	for i := range r.Repository {
		if r.Repository[i].Name != name {
			continue
		}
		return r.Repository[i], true
	}

	return nil, false
}

func (r *ServiceRepository) FindByInstanceID(id string) (*controller.Service, bool) {
	// TODO O(N)
	for i := range r.Repository {
		if r.Repository[i].ServiceID != id {
			continue
		}
		return r.Repository[i], true
	}

	return nil, false
}

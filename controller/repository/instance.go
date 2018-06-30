package repository

import "github.com/itsubaki/interstellar/controller"

type InstanceRepository struct {
	Repository []*controller.Instance
}

func (r *InstanceRepository) Insert(i *controller.Instance) {
	r.Repository = append(r.Repository, i)
}

func (r *InstanceRepository) SelectAll() []*controller.Instance {
	return r.Repository
}

func (r *InstanceRepository) FindByID(id string) (*controller.Instance, bool) {
	// TODO O(N)
	for i := range r.Repository {
		if r.Repository[i].InstanceID != id {
			continue
		}
		return r.Repository[i], true
	}

	return nil, false
}

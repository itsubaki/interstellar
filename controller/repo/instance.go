package repo

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

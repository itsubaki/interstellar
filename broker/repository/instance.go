package repository

import (
	"fmt"

	"github.com/itsubaki/interstellar/broker"
)

type InstanceRepository struct {
	Repository []*broker.Instance
}

func (r *InstanceRepository) Insert(i *broker.Instance) {
	r.Repository = append(r.Repository, i)
}

func (r *InstanceRepository) Update(i *broker.Instance) error {
	exist, ok := r.FindByID(i.InstanceID)
	if !ok {
		return fmt.Errorf("instance=%s not found", i.InstanceID)
	}

	exist.Status = i.Status
	exist.Output = i.Output
	return nil
}

func (r *InstanceRepository) SelectAll() []*broker.Instance {
	return r.Repository
}

func (r *InstanceRepository) FindByID(id string) (*broker.Instance, bool) {
	// TODO O(N)
	for i := range r.Repository {
		if r.Repository[i].InstanceID != id {
			continue
		}
		return r.Repository[i], true
	}

	return nil, false
}

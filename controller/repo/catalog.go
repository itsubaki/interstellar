package repo

import (
	"github.com/itsubaki/interstellar/broker"
)

type CatalogRepository struct {
	Reposotory []*broker.Catalog
}

func (r *CatalogRepository) Insert(catalog *broker.Catalog) {
	r.Reposotory = append(r.Reposotory, catalog)
}

func (r *CatalogRepository) FindByName(name string) (*broker.Catalog, bool) {
	// TODO O(N)
	for i := range r.Reposotory {
		if r.Reposotory[i].Name != name {
			continue
		}
		return r.Reposotory[i], true
	}

	return nil, false
}

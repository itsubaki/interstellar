package repo

import "github.com/itsubaki/interstellar/broker"

type CatalogRepository struct {
	Repository []*broker.Catalog
}

func (r *CatalogRepository) Insert(catalog *broker.Catalog) {
	r.Repository = append(r.Repository, catalog)
}

func (r *CatalogRepository) SelectAll() []*broker.Catalog {
	return r.Repository
}

func (r *CatalogRepository) FindByName(name string) (*broker.Catalog, bool) {
	// TODO O(N)
	for i := range r.Repository {
		if r.Repository[i].Name != name {
			continue
		}
		return r.Repository[i], true
	}

	return nil, false
}

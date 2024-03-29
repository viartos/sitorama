package core

import (
	"github.com/art-sitedesign/sitorama/app/core/builder"
	"github.com/art-sitedesign/sitorama/app/models"
)

func (c *Core) CreateBuilders(model *models.ProjectCreate) []builder.Builder {
	result := make([]builder.Builder, 0, 3)

	ws := c.createWebServer(model)
	if ws != nil {
		result = append(result, ws)
	}

	db := c.createDatabase(model)
	if db != nil {
		result = append(result, db)
	}

	ch := c.createCache(model)
	if ch != nil {
		result = append(result, ch)
	}

	return result
}

func (c *Core) createWebServer(model *models.ProjectCreate) builder.Builder {
	switch model.WebServer {
	case builder.BuilderNginxPHPFPM:
		return builder.NewNginxPHPFPM(c.docker, model.Domain, model.EntryPoint)
	case builder.BuilderApache:
		return builder.NewApache(c.docker, model.Domain, model.EntryPoint)
	default:
		return nil
	}
}

func (c *Core) createDatabase(model *models.ProjectCreate) builder.Builder {
	switch model.Database {
	case builder.BuilderPostgres:
		return builder.NewPostgres(c.docker, model.Domain)
	case builder.BuilderMySQL:
		return builder.NewMySQL(c.docker, model.Domain)
	default:
		return nil
	}
}

func (c *Core) createCache(model *models.ProjectCreate) builder.Builder {
	switch model.Cache {
	case builder.BuilderRedis:
		return nil
	case builder.BuilderMemcached:
		return nil
	default:
		return nil
	}
}

package profile

import (
	"github.com/go-ozzo/ozzo-routing/v2"
	"github.com/inouttt/go-crud/internal/errors"
	"github.com/inouttt/go-crud/pkg/log"
	"github.com/inouttt/go-crud/pkg/pagination"
	"net/http"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger) {
	res := resource{service, logger}

	r.Post("/profiles", res.create)
	
	r.Use(authHandler)
	
	// the following endpoints require a valid JWT
	r.Get("/profiles/<id>", res.get)
	r.Get("/profiles", res.query)
	r.Put("/profiles/<id>", res.update)
	r.Delete("/profiles/<id>", res.delete)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) get(c *routing.Context) error {
	profile, err := r.service.Get(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(profile)
}

func (r resource) query(c *routing.Context) error {
	ctx := c.Request.Context()
	count, err := r.service.Count(ctx)
	if err != nil {
		return err
	}
	pages := pagination.NewFromRequest(c.Request, count)
	profiles, err := r.service.Query(ctx, pages.Offset(), pages.Limit())
	if err != nil {
		return err
	}
	pages.Items = profiles
	return c.Write(pages)
}

func (r resource) create(c *routing.Context) error {
	var input CreateProfileRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	profile, err := r.service.Create(c.Request.Context(), input)
	if err != nil {
		return err
	}

	return c.WriteWithStatus(profile, http.StatusCreated)
}

func (r resource) update(c *routing.Context) error {
	var input UpdateProfileRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}

	profile, err := r.service.Update(c.Request.Context(), c.Param("id"), input)
	if err != nil {
		return err
	}

	return c.Write(profile)
}

func (r resource) delete(c *routing.Context) error {
	profile, err := r.service.Delete(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(profile)
}

package actions

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/x/responder"

	"short_urls/models"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Domain)
// DB Table: Plural (domains)
// Resource: Plural (Domains)
// Path: Plural (/domains)
// View Template Folder: Plural (/templates/domains/)

// DomainsResource is the resource for the Domain model
type DomainsResource struct {
	buffalo.Resource
}

// List gets all Domains. This function is mapped to the path
// GET /domains
func (v DomainsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	domains := &models.Domains{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Domains from the DB
	if err := q.All(domains); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// Add the paginator to the context so it can be used in the template.
		c.Set("pagination", q.Paginator)

		c.Set("domains", domains)
		return c.Render(http.StatusOK, r.HTML("domains/index.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(domains))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(domains))
	}).Respond(c)
}

// Show gets the data for one Domain. This function is mapped to
// the path GET /domains/{domain_id}
func (v DomainsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Domain
	domain := &models.Domain{}

	// To find the Domain the parameter domain_id is used.
	if err := tx.Find(domain, c.Param("domain_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		c.Set("domain", domain)

		return c.Render(http.StatusOK, r.HTML("domains/show.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(domain))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(domain))
	}).Respond(c)
}

// New renders the form for creating a new Domain.
// This function is mapped to the path GET /domains/new
func (v DomainsResource) New(c buffalo.Context) error {
	c.Set("domain", &models.Domain{})

	return c.Render(http.StatusOK, r.HTML("domains/new.plush.html"))
}

// Create adds a Domain to the DB. This function is mapped to the
// path POST /domains
func (v DomainsResource) Create(c buffalo.Context) error {
	// Allocate an empty Domain
	domain := &models.Domain{}

	// Bind domain to the html form elements
	if err := c.Bind(domain); err != nil {
		return err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(domain)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("html", func(c buffalo.Context) error {
			// Make the errors available inside the html template
			c.Set("errors", verrs)

			// Render again the new.html template that the user can
			// correct the input.
			c.Set("domain", domain)

			return c.Render(http.StatusUnprocessableEntity, r.HTML("domains/new.plush.html"))
		}).Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a success message
		c.Flash().Add("success", T.Translate(c, "domain.created.success"))

		// and redirect to the show page
		return c.Redirect(http.StatusSeeOther, "/domains/%v", domain.ID)
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.JSON(domain))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.XML(domain))
	}).Respond(c)
}

// Edit renders a edit form for a Domain. This function is
// mapped to the path GET /domains/{domain_id}/edit
func (v DomainsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Domain
	domain := &models.Domain{}

	if err := tx.Find(domain, c.Param("domain_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("domain", domain)
	return c.Render(http.StatusOK, r.HTML("domains/edit.plush.html"))
}

// Update changes a Domain in the DB. This function is mapped to
// the path PUT /domains/{domain_id}
func (v DomainsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Domain
	domain := &models.Domain{}

	if err := tx.Find(domain, c.Param("domain_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	// Bind Domain to the html form elements
	if err := c.Bind(domain); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(domain)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("html", func(c buffalo.Context) error {
			// Make the errors available inside the html template
			c.Set("errors", verrs)

			// Render again the edit.html template that the user can
			// correct the input.
			c.Set("domain", domain)

			return c.Render(http.StatusUnprocessableEntity, r.HTML("domains/edit.plush.html"))
		}).Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a success message
		c.Flash().Add("success", T.Translate(c, "domain.updated.success"))

		// and redirect to the show page
		return c.Redirect(http.StatusSeeOther, "/domains/%v", domain.ID)
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(domain))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(domain))
	}).Respond(c)
}

// Destroy deletes a Domain from the DB. This function is mapped
// to the path DELETE /domains/{domain_id}
func (v DomainsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Domain
	domain := &models.Domain{}

	// To find the Domain the parameter domain_id is used.
	if err := tx.Find(domain, c.Param("domain_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy(domain); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a flash message
		c.Flash().Add("success", T.Translate(c, "domain.destroyed.success"))

		// Redirect to the index page
		return c.Redirect(http.StatusSeeOther, "/domains")
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(domain))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(domain))
	}).Respond(c)
}

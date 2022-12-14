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
// Model: Singular (Visit)
// DB Table: Plural (visits)
// Resource: Plural (Visits)
// Path: Plural (/visits)
// View Template Folder: Plural (/templates/visits/)

// VisitsResource is the resource for the Visit model
type VisitsResource struct {
	buffalo.Resource
}

// List gets all Visits. This function is mapped to the path
// GET /visits
func (v VisitsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	visits := &models.Visits{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Visits from the DB
	if err := q.All(visits); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// Add the paginator to the context so it can be used in the template.
		c.Set("pagination", q.Paginator)

		c.Set("visits", visits)
		return c.Render(http.StatusOK, r.HTML("visits/index.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(visits))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(visits))
	}).Respond(c)
}

// Show gets the data for one Visit. This function is mapped to
// the path GET /visits/{visit_id}
func (v VisitsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Visit
	visit := &models.Visit{}

	// To find the Visit the parameter visit_id is used.
	if err := tx.Find(visit, c.Param("visit_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		c.Set("visit", visit)

		return c.Render(http.StatusOK, r.HTML("visits/show.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(visit))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(visit))
	}).Respond(c)
}

// New renders the form for creating a new Visit.
// This function is mapped to the path GET /visits/new
func (v VisitsResource) New(c buffalo.Context) error {
	c.Set("visit", &models.Visit{})

	return c.Render(http.StatusOK, r.HTML("visits/new.plush.html"))
}

// Create adds a Visit to the DB. This function is mapped to the
// path POST /visits
func (v VisitsResource) Create(c buffalo.Context) error {
	// Allocate an empty Visit
	visit := &models.Visit{}

	// Bind visit to the html form elements
	if err := c.Bind(visit); err != nil {
		return err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(visit)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("html", func(c buffalo.Context) error {
			// Make the errors available inside the html template
			c.Set("errors", verrs)

			// Render again the new.html template that the user can
			// correct the input.
			c.Set("visit", visit)

			return c.Render(http.StatusUnprocessableEntity, r.HTML("visits/new.plush.html"))
		}).Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a success message
		c.Flash().Add("success", T.Translate(c, "visit.created.success"))

		// and redirect to the show page
		return c.Redirect(http.StatusSeeOther, "/visits/%v", visit.ID)
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.JSON(visit))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.XML(visit))
	}).Respond(c)
}

// Edit renders a edit form for a Visit. This function is
// mapped to the path GET /visits/{visit_id}/edit
func (v VisitsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Visit
	visit := &models.Visit{}

	if err := tx.Find(visit, c.Param("visit_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("visit", visit)
	return c.Render(http.StatusOK, r.HTML("visits/edit.plush.html"))
}

// Update changes a Visit in the DB. This function is mapped to
// the path PUT /visits/{visit_id}
func (v VisitsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Visit
	visit := &models.Visit{}

	if err := tx.Find(visit, c.Param("visit_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	// Bind Visit to the html form elements
	if err := c.Bind(visit); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(visit)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("html", func(c buffalo.Context) error {
			// Make the errors available inside the html template
			c.Set("errors", verrs)

			// Render again the edit.html template that the user can
			// correct the input.
			c.Set("visit", visit)

			return c.Render(http.StatusUnprocessableEntity, r.HTML("visits/edit.plush.html"))
		}).Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a success message
		c.Flash().Add("success", T.Translate(c, "visit.updated.success"))

		// and redirect to the show page
		return c.Redirect(http.StatusSeeOther, "/visits/%v", visit.ID)
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(visit))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(visit))
	}).Respond(c)
}

// Destroy deletes a Visit from the DB. This function is mapped
// to the path DELETE /visits/{visit_id}
func (v VisitsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Visit
	visit := &models.Visit{}

	// To find the Visit the parameter visit_id is used.
	if err := tx.Find(visit, c.Param("visit_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy(visit); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a flash message
		c.Flash().Add("success", T.Translate(c, "visit.destroyed.success"))

		// Redirect to the index page
		return c.Redirect(http.StatusSeeOther, "/visits")
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(visit))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(visit))
	}).Respond(c)
}

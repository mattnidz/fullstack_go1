package actions

import "github.com/gobuffalo/buffalo"

type NewsResource struct {
	buffalo.Resource
}

// List default implementation.
func (v NewsResource) List(c buffalo.Context) error {
	return c.Render(200, r.String("News#List"))
}

// Show default implementation.
func (v NewsResource) Show(c buffalo.Context) error {
	return c.Render(200, r.String("News#Show"))
}

// New default implementation.
func (v NewsResource) New(c buffalo.Context) error {
	return c.Render(200, r.String("News#New"))
}

// Create default implementation.
func (v NewsResource) Create(c buffalo.Context) error {
	return c.Render(200, r.String("News#Create"))
}

// Edit default implementation.
func (v NewsResource) Edit(c buffalo.Context) error {
	return c.Render(200, r.String("News#Edit"))
}

// Update default implementation.
func (v NewsResource) Update(c buffalo.Context) error {
	return c.Render(200, r.String("News#Update"))
}

// Destroy default implementation.
func (v NewsResource) Destroy(c buffalo.Context) error {
	return c.Render(200, r.String("News#Destroy"))
}

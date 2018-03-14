package controller

import "api/model"

// Base is the base of controllers.
type Base struct {
	Routes model.Routes
	Prefix string
}

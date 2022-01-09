package herald

import (
	"github.com/pkg/errors"
	"github.com/shelld1t/core/app"
	"github.com/shelld1t/core/httpServer"
	controller "github.com/shelld1t/herald/internal/herald/http"
)

// Root main app entity
type Root struct {
	app       *app.App
	Container *container
}

// container of dependency
type container struct {
	// ...
}

// New create root
func New() (*Root, error) {
	a, err := app.New()
	if err != nil {
		return nil, err
	}
	c, err := initContainer()
	if err != nil {
		return nil, errors.Wrap(err, "error init dependency")
	}
	err = initRouters(c, a)
	if err != nil {
		return nil, errors.Wrap(err, "error init dependency")
	}
	return &Root{
		app:       a,
		Container: c,
	}, nil
}

// initContainer init dependency container
func initContainer() (*container, error) {
	// .... create your dependency
	return &container{}, nil
}

// initRouters init server routers (http or grpc)
func initRouters(c *container, app *app.App) error {
	err := app.InitHttpHandlers(func(server *httpServer.Server) error {
		server.AddEndpoints(controller.NewHealthController().HealthEndpoints())
		dataController := controller.NewDataController()
		server.AddEndpointsGroup("api/v1", dataController.DataEndpoints())
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

// Run application
func (at *Root) Run() error {
	return at.app.Run()
}

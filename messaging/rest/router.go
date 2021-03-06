package rest

import (
	"github.com/go-chi/chi"

	"github.com/crusttech/crust/internal/auth"
	"github.com/crusttech/crust/messaging/rest/handlers"
)

func MountRoutes() func(chi.Router) {
	// Initialize handlers & controllers.
	return func(r chi.Router) {
		r.Group(func(r chi.Router) {
			handlers.NewAttachment(Attachment{}.New()).MountRoutes(r)
		})

		r.Group(func(r chi.Router) {
			handlers.NewPermissions(Permissions{}.New()).MountRoutes(r)
		})

		// Protect all _private_ routes
		r.Group(func(r chi.Router) {
			r.Use(auth.MiddlewareValidOnly)
			r.Use(middlewareAllowedAccess)

			handlers.NewActivity(Activity{}.New()).MountRoutes(r)
			handlers.NewChannel(Channel{}.New()).MountRoutes(r)
			handlers.NewMessage(Message{}.New()).MountRoutes(r)
			handlers.NewSearch(Search{}.New()).MountRoutes(r)
			handlers.NewStatus(Status{}.New()).MountRoutes(r)
			handlers.NewCommands(Commands{}.New()).MountRoutes(r)
		})
	}
}

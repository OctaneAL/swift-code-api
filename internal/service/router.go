package service

import (
	"github.com/OctaneAL/swift-code-api/internal/config"
	"github.com/OctaneAL/swift-code-api/internal/data/pg"
	"github.com/OctaneAL/swift-code-api/internal/service/handlers"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router(cfg config.Config) chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxSwiftCodesQ(pg.NewSwiftCodesQ(cfg.DB())),
		),
	)
	r.Route("/v1/swift-codes", func(r chi.Router) {
		// configure endpoints here
		r.Route("/{swiftCode}", func(r chi.Router) {
			r.Get("/", handlers.GetBySwiftCode)
			r.Delete("/", handlers.DeleteSwiftCode)
		})
	})

	return r
}

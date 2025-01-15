package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"post-tech-challenge-10soat/internal/adapters/repository"
	c "post-tech-challenge-10soat/internal/application/core/usecases/client"
	"post-tech-challenge-10soat/internal/application/core/usecases/order"
	"post-tech-challenge-10soat/internal/application/core/usecases/payment"
	p "post-tech-challenge-10soat/internal/application/core/usecases/product"
	h "post-tech-challenge-10soat/internal/handler"
	"post-tech-challenge-10soat/internal/infra/config"
	"post-tech-challenge-10soat/internal/infra/logger"
	"post-tech-challenge-10soat/internal/infra/storage/postgres"

	_ "post-tech-challenge-10soat/docs"
)

//	@title			POS-Tech API
//	@version		1.0
//	@description	API em Go para o desafio na pos-tech fiap de Software Architecture.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/v1

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	conf, err := config.New()
	if err != nil {
		slog.Error("Error loading environment variables", "error", err)
		os.Exit(1)
	}
	logger.Set(conf.App)
	slog.Info("Starting the application", "app", conf.App.Name, "env", conf.App.Env)

	ctx := context.Background()
	db, err := postgres.New(ctx, conf.DB)
	if err != nil {
		slog.Error("Error initializing database connection", "error", err)
		os.Exit(1)
	}

	slog.Info("Successfully connected to the database", "db", conf.DB.Connection)

	err = db.Migrate()
	if err != nil {
		slog.Error("Error migrating database", "error", err)
		os.Exit(1)
	}

	defer db.Close()

	// adds health handler
	healthHandler := h.NewHealthHandler()

	// adds client handler
	cr := repository.NewClientRepository(db)
	ch := h.NewClientHandler(c.NewCreateClientUsecase(cr), c.NewGetClientUseCase(cr))

	// adds product handler
	pr := repository.NewProductRepository(db)
	ctr := repository.NewCategoryRepository(db)

	ph := h.NewProductHandler(
		p.NewCreateProductUsecase(pr, ctr),
		p.NewDeleteProductUsecase(pr),
		p.NewListProductsUsecase(pr, ctr),
		p.NewUpdateProductUsecase(pr, ctr),
	)

	// adds order handler
	or := repository.NewOrderRepository(db)
	opr := repository.NewOrderProductRepository(db)
	puc := payment.NewPaymentCheckoutUsecase(repository.NewPaymentRepository(db))
	oh := h.NewOrderHandler(order.NewCreateOrderUsecase(pr, cr, or, opr, puc), order.NewListOrdersUsecase(or))

	router, err := h.NewRouter(conf.HTTP, *healthHandler, *ch, *ph, *oh)
	if err != nil {
		slog.Error("Error initializing router", "error", err)
		os.Exit(1)
	}

	listenAddress := fmt.Sprintf("%s:%s", conf.HTTP.URL, conf.HTTP.Port)
	slog.Info("Starting the HTTP server", "listen_address", listenAddress)
	err = router.Run(listenAddress)
	if err != nil {
		slog.Error("Error starting the HTTP server", "error", err)
		os.Exit(1)
	}
}

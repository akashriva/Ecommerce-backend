package routes

import (
	"net/http"

	"github.com/akashriva/gin_framework/controller"
	"github.com/akashriva/gin_framework/helper"
)

//health check service

var healthCheckRoutes = Routes{
	Route{"Health cheak", http.MethodGet, helper.HealthCheckRoute, controller.HealthCheck},
}

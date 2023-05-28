package main

import (
	"github.com/myypo/btcinform/internal/router"
)

func main() {
	router.NewRouterImpl().Serve()
}

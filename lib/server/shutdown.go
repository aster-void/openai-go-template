package server

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

// don't forget to prepend port with ":", as in :3000
func StartWithProperShutdown(server *echo.Echo, port string) error {
	var errch chan error

	go func() {
		err := server.Start(port)
		if !errors.Is(err, http.ErrServerClosed) {
			errch <- err
		}
		errch <- nil
	}()

	intrpt, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	select {
	case err := <-errch:
		return err
	case <-intrpt.Done():
		timeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		return server.Shutdown(timeout)
	}
}

package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/zeinababbasi/timeoff/internal/app/timeoff/calendar"
	"github.com/zeinababbasi/timeoff/internal/app/timeoff/handler"
)

// nolint:funlen,wrapcheck
func main() {
	echoEngine := echo.New()

	baseCal := calendar.SetupBaseCalendar()
	holidayHandler := handler.NewCalendarHandler(baseCal)

	echoEngine.GET("/calendar/holidays", holidayHandler.GetHolidays)
	echoEngine.POST("/calendar/holiday", holidayHandler.AddHoliday)
	echoEngine.PUT("/calendar/holiday", holidayHandler.UpdateHoliday)
	echoEngine.DELETE("/calendar/holiday", holidayHandler.RemoveHoliday)

	echoEngine.POST("/calendar/is/holiday", holidayHandler.IsHoliday)

	echoEngine.GET("/healthz", func(c echo.Context) error { return c.NoContent(http.StatusNoContent) })

	hijriEngine := calendar.NewHolidayEngine(baseCal)
	if err := hijriEngine.Start(); err != nil {
		logrus.Fatalf("failed to start hijri engine: %s", err.Error())
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := echoEngine.Start(":1234"); err != nil {
			logrus.Fatalf("failed to start timeoff server: %s", err.Error())
		}
	}()

	logrus.Info("timeoff server started!")

	s := <-sig

	logrus.Infof("signal %s received", s)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	echoEngine.Server.SetKeepAlivesEnabled(false)

	if err := echoEngine.Shutdown(ctx); err != nil {
		logrus.Errorf("failed to shutdown timeoff server: %s", err.Error())
	}
}

// Register registers server command for timeoff binary.
func Register(root *cobra.Command) {
	root.AddCommand(
		&cobra.Command{
			Use:   "server",
			Short: "Run TimeOff server component",
			Run: func(cmd *cobra.Command, args []string) {
				main()
			},
		},
	)
}

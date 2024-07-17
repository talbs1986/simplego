package app

import "context"

// ScenarioJobFunc defines the job scenario processing function
type ScenarioJobFunc func(context.Context) error

// StartJob starts the application as job scenario
func StartJob(app *App, proc ScenarioJobFunc, services ...CloseableService) {
	app.slog.Info("simplego: starting job scenario")
	app.WithCloseableServices(services...)
	if err := proc(app.CTX); err != nil {
		app.slog.Fatal(err, "simplego: failed to process job scenario")
		return
	}
	app.WaitForShutodwn()
	app.slog.Info("simplego: finished job scenario successfully")
}

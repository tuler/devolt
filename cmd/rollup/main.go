package main

import (
	"context"
	"log/slog"
	"github.com/rollmelette/rollmelette"
)

type MyApplication struct{}

func (a *MyApplication) Advance(
	env rollmelette.Env,
	metadata rollmelette.Metadata,
	deposit rollmelette.Deposit,
	payload []byte,
) error {
	// Handle advance input
	return nil
}

func (a *MyApplication) Inspect(env rollmelette.EnvInspector, payload []byte) error {
	slog.Info("Inspect", "payload", string(payload))
	return nil
}

func main() {
	ctx := context.Background()
	opts := rollmelette.NewRunOpts()
	app := new(MyApplication)
	err := rollmelette.Run(ctx, opts, app)
	if err != nil {
		slog.Error("application error", "error", err)
	}
}
package main

import "log/slog"

type SimpleNotifier struct {
}

func (s SimpleNotifier) NotifyUserCreated(user user) error {
	slog.Info("Created user with simple notifier", "user_name", user.Username, "user_email", user.Email)
	return nil
}

type BetterNotifier struct {
}

func (s BetterNotifier) NotifyUserCreated(user user) error {
	slog.Info("Created user with better notifier", "user_name", user.Username, "user_email", user.Email)
	return nil
}

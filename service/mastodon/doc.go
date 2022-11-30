/*
Package mastodon provides message notification integration for Mastodon (Message Service).

Usage:

	package main

	import (
		"context"
		"log"

		"github.com/nikoksr/notify"
		"github.com/nikoksr/notify/service/mastodon"
	)

	func main() {
		mastodonSvc, err := mastodon.New("server_url", "client_id", "client_secret", "account_token")
		if err != nil {
			log.Fatalf("mastodon.New() failed: %s", err.Error())
		}

		mastodonSvc.AddReceivers("recipient username")

		notifier := notify.New()
		notifier.UseServices(mastodonSvc)

		err = notifier.Send(context.Background(), "subject", "message")
		if err != nil {
			log.Fatalf("notifier.Send() failed: %s", err.Error())
		}

		log.Println("notification sent")
	}
*/
package mastodon

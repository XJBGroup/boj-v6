package main

import "github.com/Myriad-Dreamin/artisan"

func StdReply(description ...interface{}) artisan.ReplyObject {
	return artisan.Reply(
		codeField,
		artisan.Param("data", description...),
	)
}

package main

import "github.com/Myriad-Dreamin/artisan/artisan-core"

func StdReply(description ...interface{}) artisan_core.ReplyObject {
	return artisan_core.Reply(
		codeField,
		artisan_core.Param("data", description...),
	)
}

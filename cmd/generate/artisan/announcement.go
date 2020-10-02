package main

import (
	"github.com/Myriad-Dreamin/artisan/artisan-core"
	"github.com/Myriad-Dreamin/boj-v6/abstract/announcement"
)

type AnnouncementCategories struct {
	artisan_core.VirtualService
	List       artisan_core.Category
	Count      artisan_core.Category
	Post       artisan_core.Category
	GetContent artisan_core.Category
	IdGroup    artisan_core.Category
}

func DescribeAnnouncementController() artisan_core.ProposingService {
	var announcementModel = new(announcement.Announcement)
	var _announcementModel = new(announcement.Announcement)

	var listParams = []interface{}{
		artisan_core.Param("page", artisan_core.Int),
		artisan_core.Param("page_size", artisan_core.Int),
	}

	var announcementFilter = artisan_core.Object(
		append(listParams, "ListAnnouncementRequest")...)

	controller := &AnnouncementCategories{
		List: artisan_core.Ink().
			Path("announcement-list").
			Method(artisan_core.GET, "ListAnnouncement",
				artisan_core.Request(announcementFilter),
				artisan_core.Reply(
					codeField,
					artisan_core.ArrayParam(artisan_core.Param("data", _announcementModel)),
				),
			),
		Count: artisan_core.Ink().
			Path("announcement-count").
			Method(artisan_core.GET, "CountAnnouncement",
				artisan_core.Request(),
				artisan_core.Reply(
					codeField,
					artisan_core.Param("data", artisan_core.Int64),
				),
			),
		Post: artisan_core.Ink().
			Path("announcement").
			Method(artisan_core.POST, "PostAnnouncement", artisan_core.AuthMeta("~"),
				artisan_core.Request(
					artisan_core.SnakeParam(&announcementModel.Title, required),
					artisan_core.SnakeParam(&announcementModel.Content, required),
				),
				artisan_core.Reply(
					codeField,
					artisan_core.Param("data", &announcementModel),
				),
			),
		IdGroup: artisan_core.Ink().
			Path("announcement/:aid").Meta(&Meta{artisan_core.RouterMeta{
			RuntimeRouterMeta: "announcement:aid",
		}}).
			Method(artisan_core.GET, "GetAnnouncement",
				artisan_core.Request(),
				artisan_core.Reply(
					codeField,
					artisan_core.Param("data", &announcementModel),
				)).
			Method(artisan_core.PUT, "PutAnnouncement", artisan_core.AuthMeta("~"),
				artisan_core.Request(
					artisan_core.SPsC(&announcementModel.Title, &announcementModel.Content),
				),
				artisan_core.Reply(codeField),
			).
			Method(artisan_core.DELETE, "DeleteAnnouncement", artisan_core.AuthMeta("~"),
				artisan_core.Request(), artisan_core.Reply(codeField),
			),
	}
	controller.Name("AnnouncementController").
		UseModel(artisan_core.Model(artisan_core.Name("announcement"), &announcementModel))
	return controller
}

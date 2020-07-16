package main

import (
	"github.com/Myriad-Dreamin/artisan"
	"github.com/Myriad-Dreamin/boj-v6/abstract/announcement"
)

type AnnouncementCategories struct {
	artisan.VirtualService
	List       artisan.Category
	Count      artisan.Category
	Post       artisan.Category
	GetContent artisan.Category
	IdGroup    artisan.Category
}

func DescribeAnnouncementService() artisan.ProposingService {
	var announcementModel = new(announcement.Announcement)
	var _announcementModel = new(announcement.Announcement)

	var listParams = []interface{}{
		artisan.Param("page", artisan.Int),
		artisan.Param("page_size", artisan.Int),
	}

	var announcementFilter = artisan.Object(
		append(listParams, "AnnouncementFilter")...)

	svc := &AnnouncementCategories{
		List: artisan.Ink().
			Path("announcement-list").
			Method(artisan.GET, "ListAnnouncements",
				artisan.Request(announcementFilter),
				artisan.Reply(
					codeField,
					artisan.ArrayParam(artisan.Param("data", _announcementModel)),
				),
			),
		Count: artisan.Ink().
			Path("announcement-count").
			Method(artisan.GET, "CountAnnouncement",
				artisan.Reply(
					codeField,
					artisan.Param("data", artisan.Int64),
				),
			),
		Post: artisan.Ink().
			Path("announcement").
			Method(artisan.POST, "PostAnnouncement", artisan.AuthMeta("~"),
				artisan.Request(
					artisan.SnakeParam(&announcementModel.Title, required),
					artisan.SnakeParam(&announcementModel.Content, required),
				),
				artisan.Reply(
					codeField,
					artisan.Param("data", &announcementModel),
				),
			),
		IdGroup: artisan.Ink().
			Path("announcement/:aid").Meta(&Meta{artisan.RouterMeta{
			RuntimeRouterMeta: "announcement:aid",
		}}).
			Method(artisan.GET, "GetAnnouncement",
				artisan.Reply(
					codeField,
					artisan.Param("data", &announcementModel),
				)).
			Method(artisan.PUT, "PutAnnouncement", artisan.AuthMeta("~"),
				artisan.Request(
					artisan.SPsC(&announcementModel.Title, &announcementModel.Content),
				)).
			Method(artisan.DELETE, "Delete", artisan.AuthMeta("~")),
	}
	svc.Name("AnnouncementService").
		UseModel(artisan.Model(artisan.Name("announcement"), &announcementModel))
	return svc
}

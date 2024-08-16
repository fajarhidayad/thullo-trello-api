package routes

import (
	"github.com/fajarhidayad/thullo-trello-api/handler/auth"
	"github.com/fajarhidayad/thullo-trello-api/handler/board"
	"github.com/fajarhidayad/thullo-trello-api/handler/card"
	"github.com/fajarhidayad/thullo-trello-api/handler/comment"
	"github.com/fajarhidayad/thullo-trello-api/handler/label"
	"github.com/fajarhidayad/thullo-trello-api/handler/list"
	"github.com/fajarhidayad/thullo-trello-api/handler/user"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Register(api fiber.Router, db *gorm.DB) {
	// Auth routes
	auths := api.Group("/auth")
	auths.Post("/register", auth.RegisterUser(db))
	auths.Post("/login", auth.SignInUser(db))

	// User routes
	users := api.Group("/users")
	users.Get("/:id", user.GetUserInfo(db))
	users.Put("/:id", user.UpdateUser(db))

	// Board routes
	boards := api.Group("/boards")
	boards.Get("/", board.GetAllBoards(db))
	boards.Get("/:id", board.GetBoard(db))
	boards.Post("/", board.CreateBoard(db))
	boards.Put("/:id", board.UpdateBoard(db))
	boards.Delete("/:id", board.DeleteBoard(db))
	boards.Get("/:id/lists", list.GetListsFromBoard(db))
	boards.Post("/:boardId/lists", list.CreateList(db))

	// Board member routes in board group
	boards.Get("/:id/members", board.GetAllBoardMembers(db))
	boards.Post("/:id/members", board.AddBoardMember(db))
	boards.Delete("/:id/members/:memberId", board.DeleteBoardMember(db))

	// Label route in board group
	boards.Get("/:id/labels", label.GetAllLabelsInBoard(db))
	boards.Post("/:id/labels", label.CreateNewLabel(db))

	// List routes
	lists := api.Group("/lists")
	lists.Get("/:id", list.GetListDetails(db))
	lists.Put("/:id", list.UpdateList(db))
	lists.Delete("/:id", list.DeleteList(db))
	lists.Get("/:id/cards", card.GetAllCardsFromList(db))

	// Card routes
	cards := api.Group("/cards")
	cards.Post("/", card.CreateCard(db))
	cards.Get("/:id", card.GetCardDetails(db))
	cards.Put("/:id", card.UpdateCard(db))
	cards.Delete("/:id", card.DeleteCard(db))

	// Attachment routes in card group
	cards.Get("/:id/attachments", card.GetAllAttachmentsInCard(db))
	cards.Post("/:id/attachments", card.AddAttachmentInCard(db))
	cards.Delete("/:id/attachments/:attachmentId", card.DeleteAttachmentInCard(db))

	// Comment routes in card group
	cards.Get("/:cardId/comments", comment.GetAllCommentsFromCard(db))
	cards.Post("/:cardId/comments", comment.CreateComment(db))

	// Member routes in card group
	cards.Get("/:cardId/members", card.GetAllMembersInCard(db))
	cards.Post("/:cardId/members", card.AddMemberInCard(db))
	cards.Delete("/:cardId/members/:memberId", card.DeleteMemberInCard(db))

	// Standalone route for deleting
	api.Delete("/comments/:commentId", comment.DeleteComment(db))
	api.Delete("/labels/:labelId", label.DeleteLabel(db))

}

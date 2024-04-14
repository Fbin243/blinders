package practiceapi

import (
	"log"
	"math/rand"
	"strings"

	"blinders/packages/auth"
	"blinders/packages/logging"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var DefaultLanguageLocale = "en"

func (s Service) HandleSuggestLanguageUnit(ctx *fiber.Ctx) error {
	authUser := ctx.Locals(auth.UserAuthKey).(*auth.UserAuth)
	if authUser == nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot get user auth information"})
	}

	userOID, err := primitive.ObjectIDFromHex(authUser.ID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user ID"})
	}

	var rsp logging.SuggestPracticeUnitResponse
	loggedEvent, err := s.Logger.GetSuggestPracticeUnitEventLogByUserID(userOID)
	if err != nil || len(loggedEvent) == 0 {
		log.Printf("practice: cannot get log event from Logger, err: %v, event num: %v\n", err, len(loggedEvent))
		// we could return some pre-defined document here.
		rsp = s.GetRandomPracticeUnitForUser(userOID)
	} else {
		// currently, randomly return practice unit to user
		idx := rand.Intn(len(loggedEvent))
		rsp = loggedEvent[idx].Response
	}

	return ctx.Status(fiber.StatusOK).JSON(rsp)
}

func (s Service) GetRandomPracticeUnitForUser(userID primitive.ObjectID) logging.SuggestPracticeUnitResponse {
	// user's learning language code with RFC-5646 format
	lang := ""

	usr, err := s.Db.Matches.GetMatchInfoByUserID(userID)
	if err == nil {
		// we could index the most 'active' language with index 0 that mark that specific is currently actively learning by user.
		lang = strings.Split(usr.Learnings[0], "-")[0] // we only take the Two-character primary language subtags (ex: "en-US" => "en")
	}

	units, ok := DefaultLanguageUnit[lang]
	if !ok {
		// use english as default
		units = DefaultLanguageUnit[DefaultLanguageLocale]
	}

	idx := rand.Intn(len(units))
	return units[idx]
}

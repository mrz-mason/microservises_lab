package main

import "github.com/gofiber/fiber/v2"

type Phone struct {
	TypeID      int `json:"type_id,omitempty"`
	CountryCode int `json:"country_code,omitempty"`
	Operator    int `json:"operator,omitempty"`
	Number      int `json:"number,omitempty"`
}

type Contact struct {
	ID         int      `json:"id,omitempty"`
	Username   string   `json:"username,omitempty"`
	GivenName  string   `json:"given_name,omitempty"`
	FamilyName string   `json:"family_name,omitempty"`
	Phones     []Phone  `json:"phones,omitempty"`
	Emails     []string `json:"emails,omitempty"`
	Birthdate  string   `json:"birthdate,omitempty"`
}

type Group struct {
	ID          int    `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Contacts    []int  `json:"contacts,omitempty"`
}

func NewGroupEnt() *Group {
	return &Group{}
}

func NewGroup(ctx *fiber.Ctx) error {
	var groupResponse Group

	if err := ctx.BodyParser(&groupResponse); err != nil {
		ctx.SendStatus(fiber.StatusBadRequest)
	}
	return ctx.JSON(groupResponse)
}

func GroupRead(ctx *fiber.Ctx) error {
	id := ctx.QueryInt("id")
	EntGroup := NewGroupEnt()
	EntGroup.ID = id
	return ctx.JSON(EntGroup)
}

func GroupChange(ctx *fiber.Ctx) error {
	id := ctx.QueryInt("id")
	EntGroup := NewGroupEnt()
	EntGroup.ID = id
	return ctx.JSON(EntGroup)
}
func GroupDelete(ctx *fiber.Ctx) error {
	id := ctx.QueryInt("id")
	EntGroup := NewGroupEnt()
	EntGroup.ID = id
	return ctx.JSON(EntGroup)
}

func NewContactEnt() *Contact {
	return &Contact{}
}

func NewContact(ctx *fiber.Ctx) error {
	var contactResponse Contact
	if err := ctx.BodyParser(&contactResponse); err != nil {
		ctx.SendStatus(fiber.StatusBadRequest)
	}
	return ctx.JSON(contactResponse)
}

func ContactRead(ctx *fiber.Ctx) error {
	id := ctx.QueryInt("id")
	EntContact := NewContactEnt()
	EntContact.ID = id
	return ctx.JSON(EntContact)
}

func ContactChange(ctx *fiber.Ctx) error {
	id := ctx.QueryInt("id")
	testContact := NewContactEnt()
	testContact.ID = id
	return ctx.JSON(testContact)
}

func ContactDelete(ctx *fiber.Ctx) error {
	id := ctx.QueryInt("id")
	EntContact := NewContactEnt()
	EntContact.ID = id
	return ctx.JSON(EntContact)
}

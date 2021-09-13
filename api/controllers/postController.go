package controllers

import (
	"github.com/Kamva/mgm/v2"
	"github.com/YonchevSimeon/fiber-mongodb-vue/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// GET /api/posts
func GetAll(ctx *fiber.Ctx) error {
	collection := mgm.Coll(&models.Post{})

	posts := []models.Post{}

	err := collection.SimpleFind(&posts, bson.D{})

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		// return
	}

	return ctx.JSON(fiber.Map{
		"ok":    true,
		"posts": posts,
	})
}

// GET /api/posts/:id
func GetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	post := &models.Post{}

	collection := mgm.Coll(post)

	err := collection.FindByID(id, post)

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Post not found.",
		})
	}

	return ctx.JSON(fiber.Map{
		"ok":   true,
		"post": post,
	})
}

// POST /api/posts
func Create(ctx *fiber.Ctx) error {
	params := new(struct {
		Title   string
		Content string
	})

	ctx.BodyParser(&params)

	if len(params.Title) == 0 || len(params.Content) == 0 {
		return ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Title or content not specified.",
		})
	}

	post := models.CreatePost(params.Title, params.Content)

	err := mgm.Coll(post).Create(post)

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"ok":   true,
		"post": post,
	})
}

// PATCH /api/posts/:id
func Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	post := &models.Post{}

	collection := mgm.Coll(post)

	err := collection.FindByID(id, post)

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Post not found",
		})
	}

	params := new(struct {
		Content string
	})

	ctx.BodyParser(&params)

	if len(params.Content) == 0 {
		return ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Content not specified.",
		})
	}

	post.Content = params.Content

	err = collection.Update(post)

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"ok":   true,
		"post": post,
	})
}

// DELETE /api/posts/:id
func Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	post := &models.Post{}

	collection := mgm.Coll(post)

	err := collection.FindByID(id, post)

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Post not found",
		})
	}

	err = collection.Delete(post)

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"ok":   true,
		"post": post,
	})
}

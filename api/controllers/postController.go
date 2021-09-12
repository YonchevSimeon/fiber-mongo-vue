package controllers

import (
	"github.com/Kamva/mgm/v2"
	"github.com/YonchevSimeon/fiber-mongodb-vue/models"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
)

// GET /api/posts
func GetAll(ctx *fiber.Ctx) {
	collection := mgm.Coll(&models.Post{})

	posts := []models.Post{}

	err := collection.SimpleFind(&posts, bson.D{})

	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(fiber.Map{
		"ok":    true,
		"posts": posts,
	})
}

// GET /api/posts/:id
func GetById(ctx *fiber.Ctx) {
	id := ctx.Params("id")

	post := &models.Post{}

	collection := mgm.Coll(post)

	err := collection.FindByID(id, post)

	if err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Post not found.",
		})
		return
	}

	ctx.JSON(fiber.Map{
		"ok":   true,
		"post": post,
	})
}

// POST /api/posts
func Create(ctx *fiber.Ctx) {
	params := new(struct {
		Title   string
		Content string
	})

	ctx.BodyParser(&params)

	if len(params.Title) == 0 || len(params.Content) == 0 {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Title or content not specified.",
		})
		return
	}

	post := models.Create(params.Title, params.Content)

	err := mgm.Coll(post).Create(post)

	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(fiber.Map{
		"ok":   true,
		"post": post,
	})
}

// PATCH /api/posts/:id
func Update(ctx *fiber.Ctx) {
	id := ctx.Params("id")

	post := &models.Post{}

	collection := mgm.Coll(post)

	err := collection.FindByID(id, post)

	if err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Post not found",
		})
		return
	}

	params := new(struct {
		Content string
	})

	ctx.BodyParser(&params)

	if len(params.Content) == 0 {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Content not specified.",
		})
		return
	}

	post.Content = params.Content

	err = collection.Update(post)

	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(fiber.Map{
		"ok":   true,
		"post": post,
	})
}

// DELETE /api/posts/:id
func Delete(ctx *fiber.Ctx) {
	id := ctx.Params("id")

	post := &models.Post{}

	collection := mgm.Coll(post)

	err := collection.FindByID(id, post)

	if err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Post not found",
		})
		return
	}

	err = collection.Delete(post)

	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(fiber.Map{
		"ok":   true,
		"post": post,
	})
}

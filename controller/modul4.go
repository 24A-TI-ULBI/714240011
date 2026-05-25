package controller

import (
	"context"
	"gocroot/config"
	"gocroot/model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// === CONTROLLER JADWAL ===

// GetAllJadwal retrieves all schedules
func GetAllJadwal(c *fiber.Ctx) error {
	var jadwals []model.Jadwal
	cursor, err := config.Mongoconn.Collection("jadwal").Find(context.Background(), bson.M{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var j model.Jadwal
		if err := cursor.Decode(&j); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": err.Error(),
			})
		}
		jadwals = append(jadwals, j)
	}

	if jadwals == nil {
		jadwals = []model.Jadwal{}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   jadwals,
	})
}

// GetJadwalByID retrieves a specific schedule by its ObjectID hex
func GetJadwalByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid ID format",
		})
	}

	var j model.Jadwal
	err = config.Mongoconn.Collection("jadwal").FindOne(context.Background(), bson.M{"_id": id}).Decode(&j)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Schedule not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   j,
	})
}

// CreateJadwal inserts a new schedule
func CreateJadwal(c *fiber.Ctx) error {
	var j model.Jadwal
	if err := c.BodyParser(&j); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	j.ID = primitive.NewObjectID()
	_, err := config.Mongoconn.Collection("jadwal").InsertOne(context.Background(), j)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   j,
	})
}

// UpdateJadwal updates an existing schedule by ID
func UpdateJadwal(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid ID format",
		})
	}

	var j model.Jadwal
	if err := c.BodyParser(&j); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	j.ID = id
	filter := bson.M{"_id": id}
	update := bson.M{"$set": j}

	_, err = config.Mongoconn.Collection("jadwal").UpdateOne(context.Background(), filter, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   j,
	})
}

// DeleteJadwal deletes a schedule by ID
func DeleteJadwal(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid ID format",
		})
	}

	_, err = config.Mongoconn.Collection("jadwal").DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   "Schedule deleted successfully",
	})
}

// === CONTROLLER RUANGAN ===

// GetAllRuangan retrieves all rooms
func GetAllRuangan(c *fiber.Ctx) error {
	var ruangans []model.Ruangan
	cursor, err := config.Mongoconn.Collection("ruangan").Find(context.Background(), bson.M{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var r model.Ruangan
		if err := cursor.Decode(&r); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": err.Error(),
			})
		}
		ruangans = append(ruangans, r)
	}

	if ruangans == nil {
		ruangans = []model.Ruangan{}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   ruangans,
	})
}

// CreateRuangan inserts a new room
func CreateRuangan(c *fiber.Ctx) error {
	var r model.Ruangan
	if err := c.BodyParser(&r); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	r.ID = primitive.NewObjectID()
	_, err := config.Mongoconn.Collection("ruangan").InsertOne(context.Background(), r)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   r,
	})
}

// GetRuanganByKode retrieves a room by its Kode to check availability
func GetRuanganByKode(c *fiber.Ctx) error {
	kode := c.Params("kode")
	var r model.Ruangan
	err := config.Mongoconn.Collection("ruangan").FindOne(context.Background(), bson.M{"kode": kode}).Decode(&r)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Room not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   r,
	})
}

// UpdateRuanganByKode updates a room by its Kode
func UpdateRuanganByKode(c *fiber.Ctx) error {
	kode := c.Params("kode")
	var r model.Ruangan
	if err := c.BodyParser(&r); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	filter := bson.M{"kode": kode}
	update := bson.M{"$set": bson.M{
		"nama":         r.Nama,
		"kapasitas":    r.Kapasitas,
		"ketersediaan": r.Ketersediaan,
	}}

	_, err := config.Mongoconn.Collection("ruangan").UpdateOne(context.Background(), filter, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	// Fetch updated room to return
	err = config.Mongoconn.Collection("ruangan").FindOne(context.Background(), filter).Decode(&r)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   r,
	})
}

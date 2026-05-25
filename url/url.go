package url

import (
	"gocroot/controller"

	"github.com/gofiber/fiber/v2"
)

func Web(page *fiber.App) {
	page.Get("/", controller.Homepage)
	page.Get("/ip", controller.GetIPServer)
	page.Get("/whatsauth/refreshtoken", controller.RefreshWAToken)

	page.Post("/whatsauth/webhook", controller.WhatsAuthReceiver)

	page.Get("/auth/phonenumber/:login", controller.GetPhoneNumber)

	// Route Modul 4 — Jadwal & Ruangan
	page.Get("/jadwal", controller.GetAllJadwal)
	page.Get("/jadwal/:id", controller.GetJadwalByID)
	page.Post("/jadwal", controller.CreateJadwal)
	page.Put("/jadwal/:id", controller.UpdateJadwal)
	page.Delete("/jadwal/:id", controller.DeleteJadwal)

	page.Get("/ruangan", controller.GetAllRuangan)
	page.Post("/ruangan", controller.CreateRuangan)
	page.Get("/ruangan/:kode", controller.GetRuanganByKode)
	page.Put("/ruangan/:kode", controller.UpdateRuanganByKode)
}

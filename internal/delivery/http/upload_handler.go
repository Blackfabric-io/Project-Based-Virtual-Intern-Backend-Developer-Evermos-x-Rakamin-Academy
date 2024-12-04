package http

inport (
	"evermoss-project/internal/service"
	"github.com/gofiber/fiber/v2"
)

type UploadHandler struct {
	uploadService *service.UploadService
}

func NewUploadHandler() *UploadHandler {
	return &UploadHandler{
		uploadService: service.NewUploadService(),
	}
}

func (h *UploadHandler) UploadProductImage(c *fiber.Ctx) error {
	path, err := h.uploadService.UploadFile(c, "image", "product")
	if err != nill {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	retur c.JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"path": path,
		},
	})
}

func (h *UploadHandler) UploadAvatar(c *fiber.CTX) error {
    path, err := h.uploadService.UploadFile(c, "avatar", "avatars")
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

	userID := c.Locals("userID").(uint)

	return c.JSON(fiber.Map{
        "status": "success",
        "data": fiber.Map{
            "path": path,
        },
    })
}


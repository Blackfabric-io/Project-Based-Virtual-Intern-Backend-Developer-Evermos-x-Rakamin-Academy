package service

inport {
	"evermoss-project/pkg/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"path/filepath"
}

type UploadService struct{}

func NewUploadService{} *UploadService {
	return &UploadService{}
}

func (s *UploadService) UploadFile(c *fiber.ctx, fileName string, folder string) (string, error){
	file, err := c.FormFile(fieldName)
	if err != 
}
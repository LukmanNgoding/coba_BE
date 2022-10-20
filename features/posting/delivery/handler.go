package delivery

import (
	"context"
	"net/http"
	"os"
	"strings"

	jwt "github.com/ALTA-Group-Project-Social-Media-Apps/Social-Media-Apps/features/JWT"
	"github.com/ALTA-Group-Project-Social-Media-Apps/Social-Media-Apps/features/posting/domain"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type postHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := postHandler{srv: srv}
	e.GET("", handler.ShowAllPost())
	e.POST("/post", handler.AddPost(), middleware.JWT([]byte("R4hs!!a@")))
	e.PUT("/post/update", handler.UpdatePost(), middleware.JWT([]byte("R4hs!!a@")))
	e.DELETE("/post", handler.DeleteByID(), middleware.JWT([]byte("R4hs!!a@")))
}

func (us *postHandler) DeleteByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := jwt.ExtractToken(c)
		if id == 0 {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "cannot validate token",
			})
		}
		err := us.srv.Delete(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, SuccessDelete("success delete user"))
	}
}

func (us *postHandler) AddPost() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input AddFormat
		input.ID = jwt.ExtractToken(c)
		if input.ID == 0 {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "cannot validate token",
			})
		}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}
		cnv := ToDomain(input)
		res, err := us.srv.AddPost(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("berhasil register", ToResponse(res, "reg")))
	}

}

func (us *postHandler) UpdatePost() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input UpdateFormat
		input.Username = c.FormValue("username")
		input.Content = c.FormValue("content")

		file, err := c.FormFile("photo")
		if err != nil {
			return err
		}

		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		s3Config := &aws.Config{
			Region:      aws.String(os.Getenv("AWS_REGION")),
			Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_USER"), os.Getenv("AWS_KEY"), ""),
		}

		s3Session := session.New(s3Config)

		uploader := s3manager.NewUploader(s3Session)
		inputData := &s3manager.UploadInput{
			Bucket: aws.String("alifproject"),              // bucket's name
			Key:    aws.String("myfiles/" + file.Filename), // files destination location
			Body:   src,                                    // content of the file
			// ACL:    aws.String("Objects - List"),
			// ContentType: aws.String("image/png"), // content type
		}
		_, _ = uploader.UploadWithContext(context.Background(), inputData)
		input.Photo = "s3://alifproject/myfiles/" + strings.ReplaceAll(file.Filename, " ", "+")
		input.ID = jwt.ExtractToken(c)
		if input.ID == 0 {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "cannot validate token",
			})
		}

		cnv := ToDomain(input)
		res, err := us.srv.UpdatePost(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("berhasil update", ToResponse(res, "upd")))
	}

}

func (us *postHandler) ShowAllPost() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input AddFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}
		cnv := ToDomain(input)
		res, err := us.srv.AddPost(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusCreated, SuccessResponse("Success get book", ToResponse(res, "all")))
	}
}

package handlers

import (
	"domain-app/internal/store/cms_db"
	"domain-app/internal/store/db"
	"fmt"
	"github.com/labstack/echo"
)

type HealthCheckHandler struct{}

type Response struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func (handler HealthCheckHandler) ServeHTTP(c echo.Context) (*Response, error) {
	dbSearchParam := c.QueryParam("db")

	switch dbSearchParam {
	case "test":
		_, err := db.Connect_db()
		if err != nil {
			fmt.Println("Error connecting to DB: ", err)
			return nil, err
		}

		response := Response{Ok: true, Message: "Database is alive!"}
		return &response, nil
	case "cms":
		database, err := cms_db.Connect_cms_db()
		if err != nil {
			fmt.Println("Error connecting to DB: ", err)
			echo.NewHTTPError(500, err)
			return nil, err

		}

		err = cms_db.Disconnect_cms_db(database)
		if err != nil {
			fmt.Println("Error disconnecting from DB: ", err)
			return nil, err
		}

		response := Response{Ok: true, Message: "CMS Database is alive!"}
		return &response, nil
	default:
		response := Response{Ok: true, Message: "All good here!"}
		return &response, nil
	}
}

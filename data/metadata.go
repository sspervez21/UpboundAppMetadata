package data

import (
	"UpboundAppMetadata/models"
	"UpboundAppMetadata/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

type maintainer struct {
	name  string
	email string
}

type appMetadata struct {
	title       string
	version     string
	maintainers []maintainer
	company     string
	website     string
	source      string
	license     string
	description string
}

func makeModelAppObject(app appMetadata) *models.AppObject {
	appObject := &models.AppObject{
		Title:       &app.title,
		Version:     &app.version,
		Company:     &app.company,
		Website:     &app.website,
		Source:      &app.source,
		License:     &app.license,
		Description: &app.description,
	}

	for _, element := range app.maintainers {
		// TODO: make append more performant
		appObject.Maintainers = append(appObject.Maintainers, &models.Maintainer{Name: element.name, Email: element.email})
	}

	return appObject
}

// GetAppMetadata function
func GetAppMetadata(params operations.GetAppParams) middleware.Responder {
	app, ok := retrieveData(*params.Title)

	if !ok {
		return operations.NewGetAppNotFound().WithPayload(&models.NotFound{
			Code:    int64(operations.GetAppNotFoundCode),
			Message: "App with this title does not exist.",
		})
	}

	return operations.NewGetAppOK().WithPayload(makeModelAppObject(app))
}

// AddAppMetadata function
func AddAppMetadata(params operations.PostAppParams) middleware.Responder {
	app := appMetadata{
		title:       *params.AppMetadata.Title,
		version:     *params.AppMetadata.Version,
		company:     *params.AppMetadata.Company,
		website:     *params.AppMetadata.Website,
		source:      *params.AppMetadata.Source,
		license:     *params.AppMetadata.License,
		description: *params.AppMetadata.Description,
	}

	for _, element := range params.AppMetadata.Maintainers {
		// TODO: make append more performant
		app.maintainers = append(app.maintainers, maintainer{name: element.Name, email: element.Email})
	}

	ok := storeData(app)

	if !ok {
		return operations.NewPostAppBadRequest().WithPayload(&models.BadRequest{
			Code:    int64(operations.PostAppBadRequestCode),
			Message: "Server could not store the data provided.",
		})
	}

	return operations.NewPostAppOK()
}

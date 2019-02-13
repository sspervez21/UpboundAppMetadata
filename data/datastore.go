package data

var allApps []appMetadata

func storeData(app appMetadata) bool {
	ok := createIndexEntry(app.title, app.version, app)

	if !ok {
		return false
	}

	// TODO: make append more performant
	allApps = append(allApps, app)
	return true
}

func retrieveData(title string, version string) (appMetadata, bool) {
	app, ok := retrieveIndexEntry(title, version)

	// not found
	if !ok {
		return appMetadata{}, false
	}

	return app, true
}

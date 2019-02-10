package data

var allApps []appMetadata

func storeData(app appMetadata) bool {
	// TODO: make append more performant
	allApps = append(allApps, app)
	return true
}

func retrieveData(title string) (appMetadata, bool) {
	for _, element := range allApps {
		if element.title == title {
			return element, true
		}
	}

	// not found
	return appMetadata{}, false
}

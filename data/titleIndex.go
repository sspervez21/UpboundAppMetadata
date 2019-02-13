package data

var titleIndex = make(map[string]map[string]*appMetadata)

func createIndexEntry(title string, version string, app appMetadata) bool {
	// Check if the app already exists. Return false if it does
	_, ok := titleIndex[title][version]

	if ok {
		return false
	}

	if titleIndex[title] == nil {
		titleIndex[title] = make(map[string]*appMetadata)
	}

	titleIndex[title][version] = &app
	return true
}

func retrieveIndexEntry(title string, version string) (appMetadata, bool) {
	app, ok := titleIndex[title][version]

	// not found
	if !ok {
		return appMetadata{}, false
	}

	return *app, ok
}

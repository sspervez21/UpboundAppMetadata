package data

import (
	"testing"
)

func TestEmptyIndex(t *testing.T) {
	app, ok := retrieveIndexEntry("app1", "1.0.0")

	if ok {
		t.Fatalf("app found?: %s:%s\n", app.title, app.version)
	}
}

func TestCreateIndexEntry(t *testing.T) {
	app := appMetadata{
		title:       "Valid app1",
		version:     "1.0.0",
		company:     "Random Inc.",
		website:     "https://website.com",
		source:      "https://github.com/random/repo",
		license:     "Apache-2.0",
		description: "Some application content, and description",
	}

	app.maintainers = append(app.maintainers, maintainer{name: "firstmaintainer app1", email: "firstmaintainer@hotmail.com"}, maintainer{name: "secondmaintainer app1", email: "secondmaintainer@gmail.com"})

	createIndexEntry(app.title, app.version, app)

	app1Indexed, ok := retrieveIndexEntry(app.title, app.version)

	if !ok {
		t.Fatalf("app not found: %s:%s\n", app.title, app.version)
	}

	if !compareAppMetadata(app, app1Indexed) {
		t.Fatalf("app not correct: %v\n%v\n", app, app1Indexed)
	}
}

func TestCreateMultipleIndexEntries(t *testing.T) {
	app1 := appMetadata{
		title:       "Valid app1",
		version:     "1.0.0",
		company:     "Random Inc.",
		website:     "https://website.com",
		source:      "https://github.com/random/repo",
		license:     "Apache-2.0",
		description: "Some application content, and description",
	}

	app1.maintainers = append(app1.maintainers, maintainer{name: "firstmaintainer app1", email: "firstmaintainer@hotmail.com"}, maintainer{name: "secondmaintainer app1", email: "secondmaintainer@gmail.com"})

	createIndexEntry(app1.title, app1.version, app1)

	app2 := appMetadata{
		title:       "Valid app2",
		version:     "1.0.0",
		company:     "Random Inc.",
		website:     "https://website.com",
		source:      "https://github.com/random/repo",
		license:     "Apache-2.0",
		description: "Some application content, and description",
	}

	app2.maintainers = append(app2.maintainers, maintainer{name: "firstmaintainer app2", email: "firstmaintainer@hotmail.com"}, maintainer{name: "secondmaintainer app2", email: "secondmaintainer@gmail.com"})

	createIndexEntry(app2.title, app2.version, app2)

	app3 := appMetadata{
		title:       "Valid app1",
		version:     "1.0.1",
		company:     "Random Inc.",
		website:     "https://website.com",
		source:      "https://github.com/random/repo",
		license:     "Apache-2.0",
		description: "Some application content, and description",
	}

	app3.maintainers = append(app3.maintainers, maintainer{name: "firstmaintainer app1", email: "firstmaintainer@hotmail.com"}, maintainer{name: "secondmaintainer app1", email: "secondmaintainer@gmail.com"})

	createIndexEntry(app3.title, app3.version, app3)

	app3Indexed, ok3 := retrieveIndexEntry(app3.title, app3.version)
	if !ok3 {
		t.Fatalf("app not found: %s:%s\n", app3.title, app3.version)
	}

	app2Indexed, ok2 := retrieveIndexEntry(app2.title, app2.version)
	if !ok2 {
		t.Fatalf("app not found: %s:%s\n", app2.title, app2.version)
	}

	app1Indexed, ok1 := retrieveIndexEntry(app1.title, app1.version)
	if !ok1 {
		t.Fatalf("app not found: %s:%s\n", app1.title, app1.version)
	}

	if !compareAppMetadata(app1, app1Indexed) {
		t.Fatalf("app not correct: %v\n%v\n", app1, app1Indexed)
	}

	if !compareAppMetadata(app2, app2Indexed) {
		t.Fatalf("app not correct: %v\n%v\n", app2, app2Indexed)
	}

	if !compareAppMetadata(app3, app3Indexed) {
		t.Fatalf("app not correct: %v\n%v\n", app3, app3Indexed)
	}
}

func compareAppMetadata(app1 appMetadata, app1Indexed appMetadata) bool {
	if app1.title != app1Indexed.title {
		return false
	}

	if app1.version != app1Indexed.version {
		return false
	}

	if len(app1.maintainers) != len(app1Indexed.maintainers) {
		return false
	}

	for idx, element := range app1.maintainers {
		if element.name != app1Indexed.maintainers[idx].name {
			return false
		}

		if element.email != app1Indexed.maintainers[idx].email {
			return false
		}
	}

	if app1.company != app1Indexed.company {
		return false
	}

	if app1.website != app1Indexed.website {
		return false
	}

	if app1.source != app1Indexed.source {
		return false
	}

	if app1.license != app1Indexed.license {
		return false
	}

	if app1.description != app1Indexed.description {
		return false
	}

	return true
}

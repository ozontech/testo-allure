package allure

// Reflection holds information about current allure state.
type Reflection struct {
	UUID              UUID
	TestCaseID        string
	HistoryID         string
	FullName          string
	Description       string
	Parameters        []Parameter
	Labels            []Label
	Links             []Link
	Severity          Severity
	Status            Status
	StatusDetails     StatusDetails
	Categories        []Category
	Title             string
	MaxAttachmentSize int64
}

// Reflect returns [Reflection] for the underlying allure state.
func Reflect(t Interface) Reflection {
	a := t.unwrapAllure()

	return Reflection{
		UUID:              a.uuid,
		TestCaseID:        a.testCaseID(),
		HistoryID:         a.historyID(),
		FullName:          a.fullName(),
		Description:       a.description.Load(),
		Parameters:        a.rawParameters.Load(),
		Labels:            a.labels(),
		Links:             a.rawLinks.Load().ClonedSlice(),
		Severity:          a.severity.Load(),
		Status:            a.status(),
		StatusDetails:     a.statusDetails(),
		Categories:        a.categories,
		Title:             a.title(),
		MaxAttachmentSize: a.maxAttachmentSize,
	}
}

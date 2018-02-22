package check

// FileSummary contains the filename, location of the file
// on GitHub, and all of the errors related to the file
type FileSummary struct {
	Filename string `json:"filename"`
	FileURL  string `json:"file_url"`
}

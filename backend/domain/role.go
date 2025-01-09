package domain

type Role struct {
	ResourceID  string
	Name        string
	Description string
	Permissions map[string]bool
	// Output only
	CreatorID int
}

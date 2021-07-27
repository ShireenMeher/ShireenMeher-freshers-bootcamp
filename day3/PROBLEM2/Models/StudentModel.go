//Models/StudentModel.go
package Models

type Student struct {
	Id      int64   `json:"id"`
	Name    string `json:"name"`
	LastName    string `json:"last-name"`
	DOB    int64 `json:"dob"`
	Address string `json:"address"`
	Subject string `json:"subject"`
	Marks int64 `json:"marks"`
}
func (b *Student) TableName() string {
	return "students"
}

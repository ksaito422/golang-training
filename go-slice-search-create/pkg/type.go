package pkg

type Company struct {
	ID   uint
	Name string
}

type User struct {
	ID        uint
	CompanyID uint
	Name      string
	Age       uint
}

type CompanyUsers struct {
	ID     uint
	Name   string
	UserID []uint
}

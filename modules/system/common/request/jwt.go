package request

import (
	"github.com/gofrs/uuid/v5"
	jwt "github.com/golang-jwt/jwt/v4"
)

// Custom claims structure
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

type BaseClaims struct {
	UUID        uuid.UUID
	ID          uint
	Username    string
	NickName    string
	AuthorityId uint
	CompanyInfo CompanyInfo
}

type CompanyInfo struct {
	CompanyID    uint   `json:"company_id"`
	Company      string `json:"company"`
	DepartmentId uint   `json:"department_id"`
	Department   string `json:"department"`
}

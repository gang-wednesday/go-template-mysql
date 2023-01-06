// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gqlmodels

type Author struct {
	ID                 string  `json:"id"`
	UserName           *string `json:"userName"`
	Email              *string `json:"email"`
	Name               *string `json:"name"`
	Active             *bool   `json:"active"`
	Address            *string `json:"address"`
	LastLogin          *string `json:"lastLogin"`
	LastPasswordChange *string `json:"lastPasswordChange"`
	Token              *string `json:"token"`
	Role               *Role   `json:"role"`
	CreatedAt          *int    `json:"createdAt"`
	UpdatedAt          *int    `json:"updatedAt"`
	DeletedAt          *int    `json:"deletedAt"`
	Posts              []*Post `json:"posts"`
}

type AuthorCreateInput struct {
	UserName *string `json:"userName"`
	Email    *string `json:"email"`
	Name     *string `json:"name"`
	Password string  `json:"password"`
	Address  *string `json:"address"`
	RoleID   string  `json:"roleId"`
	Active   *bool   `json:"active"`
}

type AuthorDeleteInput struct {
	ID string `json:"id"`
}

type AuthorFilter struct {
	Search *string `json:"search"`
}

type AuthorPagination struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type AuthorUpdateInput struct {
	ID       string  `json:"id"`
	UserName *string `json:"userName"`
	Email    *string `json:"email"`
	Name     *string `json:"name"`
	Address  *string `json:"address"`
}

type AuthorsPayload struct {
	Authors []*Author `json:"authors"`
	Total   int       `json:"total"`
}

type BooleanFilter struct {
	IsTrue  *bool `json:"isTrue"`
	IsFalse *bool `json:"isFalse"`
	IsNull  *bool `json:"isNull"`
}

type ChangePasswordResponse struct {
	Ok bool `json:"ok"`
}

type FloatFilter struct {
	EqualTo           *float64  `json:"equalTo"`
	NotEqualTo        *float64  `json:"notEqualTo"`
	LessThan          *float64  `json:"lessThan"`
	LessThanOrEqualTo *float64  `json:"lessThanOrEqualTo"`
	MoreThan          *float64  `json:"moreThan"`
	MoreThanOrEqualTo *float64  `json:"moreThanOrEqualTo"`
	In                []float64 `json:"in"`
	NotIn             []float64 `json:"notIn"`
}

type IDFilter struct {
	EqualTo    *string  `json:"equalTo"`
	NotEqualTo *string  `json:"notEqualTo"`
	In         []string `json:"in"`
	NotIn      []string `json:"notIn"`
}

type IntFilter struct {
	EqualTo           *int  `json:"equalTo"`
	NotEqualTo        *int  `json:"notEqualTo"`
	LessThan          *int  `json:"lessThan"`
	LessThanOrEqualTo *int  `json:"lessThanOrEqualTo"`
	MoreThan          *int  `json:"moreThan"`
	MoreThanOrEqualTo *int  `json:"moreThanOrEqualTo"`
	In                []int `json:"in"`
	NotIn             []int `json:"notIn"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type Post struct {
	ID        string  `json:"id"`
	Title     *string `json:"title"`
	Content   *string `json:"content"`
	Author    *Author `json:"author"`
	CreatedAt *int    `json:"createdAt"`
	UpdatedAt *int    `json:"updatedAt"`
	DeletedAt *int    `json:"deletedAt"`
}

type PostCreateInput struct {
	Title   *string `json:"title"`
	Content *string `json:"content"`
}

type PostDeleteInput struct {
	ID string `json:"id"`
}

type PostFilterByTitle struct {
	Title *string `json:"title"`
}

type PostPagination struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type PostPayload struct {
	Posts     []*Post `json:"posts"`
	PostCount *int    `json:"postCount"`
}

type PostUpdateInput struct {
	ID      string  `json:"id"`
	Title   *string `json:"title"`
	Content *string `json:"content"`
}

type PostsPayload struct {
	Posts []*Post `json:"posts"`
	Total *int    `json:"total"`
}

type RefreshTokenResponse struct {
	Token string `json:"token"`
}

type Role struct {
	ID          string    `json:"id"`
	AccessLevel int       `json:"accessLevel"`
	Name        string    `json:"name"`
	UpdatedAt   *int      `json:"updatedAt"`
	DeletedAt   *int      `json:"deletedAt"`
	CreatedAt   *int      `json:"createdAt"`
	Authors     []*Author `json:"authors"`
}

type RoleCreateInput struct {
	AccessLevel int    `json:"accessLevel"`
	Name        string `json:"name"`
}

type RoleDeletePayload struct {
	ID string `json:"id"`
}

type RolePagination struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type RolePayload struct {
	Role *Role `json:"role"`
}

type RoleUpdateInput struct {
	AccessLevel *int    `json:"accessLevel"`
	Name        *string `json:"name"`
	UpdatedAt   *int    `json:"updatedAt"`
	DeletedAt   *int    `json:"deletedAt"`
	CreatedAt   *int    `json:"createdAt"`
}

type RolesCreateInput struct {
	Roles []*RoleCreateInput `json:"roles"`
}

type RolesDeletePayload struct {
	Ids []string `json:"ids"`
}

type RolesPayload struct {
	Roles []*Role `json:"roles"`
}

type RolesUpdatePayload struct {
	Ok bool `json:"ok"`
}

type StringFilter struct {
	EqualTo            *string  `json:"equalTo"`
	NotEqualTo         *string  `json:"notEqualTo"`
	In                 []string `json:"in"`
	NotIn              []string `json:"notIn"`
	StartWith          *string  `json:"startWith"`
	NotStartWith       *string  `json:"notStartWith"`
	EndWith            *string  `json:"endWith"`
	NotEndWith         *string  `json:"notEndWith"`
	Contain            *string  `json:"contain"`
	NotContain         *string  `json:"notContain"`
	StartWithStrict    *string  `json:"startWithStrict"`
	NotStartWithStrict *string  `json:"notStartWithStrict"`
	EndWithStrict      *string  `json:"endWithStrict"`
	NotEndWithStrict   *string  `json:"notEndWithStrict"`
	ContainStrict      *string  `json:"containStrict"`
	NotContainStrict   *string  `json:"notContainStrict"`
}

package schemas

var (
	Users string
	Roles string
)

func SetTableName() {
	Users = "role_based_access.users"
	Roles = "role_based_access.roles"
}

package repository

const (
	// profile table mutation
	saveProfileQuery       = "insert into profile (name, phone, password) values ($1, $2, $3) returning id"
	updateLoginCountQuery  = "update profile set login_count = $1 where id = $2"
	updateProfileByIDQuery = "update profile set name = $1, phone = $2 where id = $3"

	// profile queries
	profileSelectAll       = "select id, name, phone, password, login_count, created_at, updated_at from profile "
	getProfileByPhoneQuery = profileSelectAll + "where phone = $1"
	getProfileByIDQuery    = profileSelectAll + "where id = $1"
	// end of profile table query
)

package repository

import (
	"context"
	"database/sql"
)

// save user profile and return the profile id
func (r Repository) SaveProfile(ctx context.Context, user User) (id int64, err error) {
	err = r.Db.QueryRowContext(
		ctx,
		saveProfileQuery,
		user.Name,
		user.Phone,
		user.Password,
	).Scan(&id)
	return
}

// convert given sql.Row to user profile struct
func (r Repository) scanProfileRow(row *sql.Row, user *User) error {
	return row.Scan(
		&user.ID,
		&user.Name,
		&user.Phone,
		&user.Password,
		&user.LoginCount,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
}

// get the user profile by phone number
func (r Repository) GetProfileByPhone(ctx context.Context, phone string) (user User, err error) {
	err = r.scanProfileRow(
		r.Db.QueryRowContext(ctx, getProfileByPhoneQuery, phone),
		&user,
	)
	return
}

// get the user profile by the profile id
func (r Repository) GetProfileByID(ctx context.Context, id int64) (user User, err error) {
	err = r.scanProfileRow(
		r.Db.QueryRowContext(ctx, getProfileByIDQuery, id),
		&user,
	)
	return
}

// update the user login_count by profile id
func (r Repository) UpdateLoginCount(ctx context.Context, userID int64, loginCount int) (err error) {
	_, err = r.Db.ExecContext(ctx, updateLoginCountQuery, loginCount, userID)
	return
}

// update the user name and phone by profile id
func (r Repository) UpdateUserByID(ctx context.Context, user User) (err error) {
	_, err = r.Db.ExecContext(ctx, updateProfileByIDQuery, user.Name, user.Phone, user.ID)
	return
}

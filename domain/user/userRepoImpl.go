package user

import (
	"api_server/model"
	"api_server/utils"
	"database/sql"
)


type userRepoImple struct{
	db *sql.DB
}

func  NewUserRepoImple(db *sql.DB) IUserRepo {
	return &userRepoImple{
		db: db,
	}
}
func (u userRepoImple) CreateUser(user *model.User)error{
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.CREATE_USER)
	defer stmt.Close()
	if err != nil{
		_ = tx.Rollback()
		return err
	}
	_, err = stmt.Exec(user.UserID, user.IDCard, user.Username, user.DateOfBirth, user.Job.JobID, user.Education.EducationID)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (u userRepoImple) ReadUser(i int, i2 int) ([]*model.User, error) {
	users := make([]*model.User, 0)
	stmt, err := u.db.Prepare(utils.SELECT_USERS)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(i, i2)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := model.User{}
		err := rows.Scan(
			&user.UserID, &user.IDCard, &user.Username, &user.DateOfBirth, &user.Education.EducationID, &user.Education.EducationLabel,
			&user.Job.JobID, &user.Job.JobLabel, &user.UserStatus, &user.CreatedDate, &user.UpdatedDate)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (u userRepoImple) CountUser() (int, error) {
	var totalData int
	stmt, err := u.db.Prepare(utils.SELECT_COUNT_DATA_USER)
	if err != nil {
		return totalData, err
	}
	defer stmt.Close()
	err = stmt.QueryRow().Scan(&totalData)
	if err != nil {
		return totalData, err
	}
	return totalData, nil
}

func (u userRepoImple) ReadUserById(s string) (*model.User, error) {
	user := model.User{}
	stmt, err := u.db.Prepare(utils.SELECT_USER_BY_ID)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(s).Scan(
		&user.UserID, &user.IDCard, &user.Username, &user.DateOfBirth, &user.Education.EducationID, &user.Education.EducationLabel,
		&user.Job.JobID, &user.Job.JobLabel, &user.UserStatus, &user.CreatedDate, &user.UpdatedDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return &user, nil
		}else{
			return &user, err
		}
	}
	return &user, nil
}

func (u userRepoImple) UpdateUser(user *model.User) error {
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.UPDATE_USER)

	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(user.IDCard, user.Username, user.DateOfBirth, user.Job.JobID, user.Education.EducationID, user.UserID)
	if err != nil {
		tx.Rollback()
		return err
	}
	stmt.Close()
	return tx.Commit()
}

func (u userRepoImple) DeleteUser(s string) error {
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.DELETE_USER)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(s)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (u userRepoImple) ReadJob() ([]*model.Job, error) {
	jobList := make([]*model.Job, 0)
	stmt, err := u.db.Prepare(utils.SELECT_PEKERJAAN)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		job := model.Job{}
		err := rows.Scan(&job.JobID, &job.JobLabel)
		if err != nil {
			return nil, err
		}
		jobList = append(jobList, &job)
	}
	return jobList, nil
}

func (u userRepoImple) ReadEducation() ([]*model.Education, error) {
	educationList := make([]*model.Education, 0)
	stmt, err := u.db.Prepare(utils.SELECT_PENDIDIKAN)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		education := model.Education{}
		err := rows.Scan(&education.EducationID, &education.EducationLabel)
		if err != nil {
			return nil, err
		}
		educationList = append(educationList, &education)
	}
	return educationList, nil
}
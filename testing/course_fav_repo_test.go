package testing

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/repository"
	"github.com/stretchr/testify/suite"
)

type CourseFavouriteRepoTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    repository.IUserCoursesFavouriteRepository
}

var payload = entity.UserCoursesFavourite{
	ID:       1,
	UserID:   "1",
	CourseID: "1",
}

func (s *CourseFavouriteRepoTestSuite) TestAddOrRemoveCourseFavourite_Added() {
	sqlmock.NewRows([]string{"id", "user_id", "course_id"}).AddRow(1, payload.UserID, payload.CourseID)

	s.mockSql.ExpectQuery(`SELECT \* FROM "user_courses_favourites" WHERE \(user_id = \$1 AND course_id = \$2\)`).
		WithArgs(payload.UserID, payload.CourseID).WillReturnError(gorm.ErrRecordNotFound)

	s.mockSql.ExpectBegin()

	s.mockSql.ExpectExec(`INSERT INTO "user_courses_favourites" \(user_id, course_id\) VALUES \(\$1, \$2\)`).
		WithArgs(payload.UserID, payload.CourseID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	s.mockSql.ExpectCommit()

	err, result := s.repo.AddOrRemoveToFavourite(payload)

	s.NoError(err)
	s.Equal("Add to Favourite Executed", result)

	err = s.mockSql.ExpectationsWereMet()
	s.NoError(err)
}

func (s *CourseFavouriteRepoTestSuite) TestAddOrRemoveCourseFavourite_Removed() {
	s.mockSql.ExpectQuery(`SELECT \* FROM "user_courses_favourites" WHERE \(user_id = \$1 AND course_id = \$2\)`).
		WithArgs(payload.UserID, payload.CourseID).
		WillReturnRows(sqlmock.NewRows([]string{"user_id", "course_id"}).AddRow(payload.UserID, payload.CourseID))

	s.mockSql.ExpectBegin()

	s.mockSql.ExpectExec(`DELETE FROM "user_courses_favourites" WHERE \(user_id = \$1 AND course_id = \$2\)`).
		WithArgs(payload.UserID, payload.CourseID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	s.mockSql.ExpectCommit()

	err, result := s.repo.AddOrRemoveToFavourite(payload)

	s.NoError(err)
	s.Equal("Removed from Favourite Executed", result)

	err = s.mockSql.ExpectationsWereMet()
	s.NoError(err)
}

func (s *CourseFavouriteRepoTestSuite) SetupTest() {
	s.mockDb, s.mockSql, _ = sqlmock.New()
	gormDb, err := gorm.Open("postgres", s.mockDb)
	if err != nil {
		panic(err)
	}

	s.repo = repository.NewUserCoursesFavouriteRepository(gormDb)
}

func TestCourseFavouriteRepoTestSuite(t *testing.T) {
	suite.Run(t, new(CourseFavouriteRepoTestSuite))
}

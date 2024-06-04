package testing

import (
	"database/sql"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/entity/dto"
	"github.com/kelompok-2/ilmu-padi/repository"
	"github.com/kelompok-2/ilmu-padi/shared/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CourseRepositoryTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    repository.CourseRepository
}

var exampleCourse = entity.Course{
	Title:           "Test",
	Description:     "Test",
	Category:        "Test",
	Video_URL:       "Test",
	Duration:        "0",
	Instructor_Name: "0",
	Rating:          0,
	UserId:          "0",
	ID:              1,
}

func (s *CourseRepositoryTestSuite) SetupTest() {
	s.mockDb, s.mockSql, _ = sqlmock.New()
	gormDb, err := gorm.Open("postgres", s.mockDb)
	if err != nil {
		panic(err)
	}
	s.repo = repository.NewCourseRepository(gormDb)
}

func (s *CourseRepositoryTestSuite) TestGetAll_Success() {
	page := 1
	rowsPerPage := 2
	offset := (page - 1) * rowsPerPage

	expectedCourse := []entity.Course{
		{
			Title:           "Test",
			Description:     "Test",
			Category:        "Test",
			Video_URL:       "Test",
			Duration:        "0",
			Instructor_Name: "0",
			Rating:          0,
			UserId:          "0",
			ID:              1,
		},
		{
			Title:           "Test",
			Description:     "Test",
			Category:        "Test",
			Video_URL:       "Test",
			Duration:        "0",
			Instructor_Name: "0",
			Rating:          0,
			UserId:          "0",
			ID:              2,
		},
	}

	expectedPaging := model.Paging{
		Page:        page,
		RowsPerPage: rowsPerPage,
		TotalRows:   2,
		TotalPages:  1,
	}

	rows := sqlmock.NewRows([]string{}).AddRow().AddRow()

	s.mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM courses limit $1 offset $2`)).WithArgs(rowsPerPage, offset).WillReturnRows(rows)

	totalRows := sqlmock.NewRows([]string{"count"}).AddRow(2)

	s.mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(*) FROM courses`)).WillReturnRows(totalRows)

	actualCourses, actualPaging, actualErr := s.repo.FindAll(page, rowsPerPage)
	assert.NoError(s.T(), actualErr)
	assert.Equal(s.T(), expectedCourse, actualCourses)
	assert.Equal(s.T(), expectedPaging, actualPaging)
}

func (s *CourseRepositoryTestSuite) TestCreate_Success() {
	s.mockSql.ExpectBegin()

	s.mockSql.ExpectExec(regexp.QuoteMeta(`INSERT INTO courses (title, description, category, video_url, duration, instructor_name, rating) VALUES ($1, $2, $3, $4, $5, $6, $7)`)).
		WithArgs().
		WillReturnResult(sqlmock.NewResult(0, 1))

	s.mockSql.ExpectCommit()

	err := s.repo.Create(exampleCourse)
	s.NoError(err)
}

func (s *CourseRepositoryTestSuite) TestCreate_Failed() {
	s.mockSql.ExpectBegin()

	s.mockSql.ExpectExec(regexp.QuoteMeta(`INSERT INTO courses (title, description, category, video_url, duration, instructor_name, rating) VALUES ($1, $2, $3, $4, $5, $6, $7)`)).
		WithArgs().
		WillReturnError(fmt.Errorf("something wrong"))

	s.mockSql.ExpectCommit()

	err := s.repo.Create(exampleCourse)
	s.Error(err)
}

func (s *CourseRepositoryTestSuite) TestGetById_Success() {

	rows := sqlmock.NewRows([]string{"id", "title", "description", "category", "video_url", "duration", "instructor_name", "rating"}).AddRow(exampleCourse.ID, exampleCourse.Title, exampleCourse.Description, exampleCourse.Category, exampleCourse.Video_URL, exampleCourse.Duration, exampleCourse.Instructor_Name, exampleCourse.Rating)

	s.mockSql.ExpectQuery(regexp.QuoteMeta(`select * from courses where id=$1`)).WithArgs(exampleCourse.ID).WillReturnRows(rows)

	actualCourse, actualErr := s.repo.FindByID(dto.CourseIDDto{ID: exampleCourse.ID})

	assert.Nil(s.T(), actualErr)
	assert.NoError(s.T(), actualErr)
	assert.Equal(s.T(), exampleCourse.ID, actualCourse.ID)
}

func (s *CourseRepositoryTestSuite) TestGetById_Failed() {
	sqlmock.NewRows([]string{"id", "title", "description", "category", "video_url", "duration", "instructor_name", "rating"}).AddRow(exampleCourse.ID, exampleCourse.Title, exampleCourse.Description, exampleCourse.Category, exampleCourse.Video_URL, exampleCourse.Duration, exampleCourse.Instructor_Name, exampleCourse.Rating)

	s.mockSql.ExpectQuery(regexp.QuoteMeta(`select * from courses where id=$1`)).WithArgs(exampleCourse.ID).WillReturnError(fmt.Errorf("something wrong"))

	actualCourse, actualErr := s.repo.FindByID(dto.CourseIDDto{ID: exampleCourse.ID})

	assert.Error(s.T(), actualErr)
	assert.Equal(s.T(), exampleCourse.ID, actualCourse.ID)
}

func (s *CourseRepositoryTestSuite) TestUpdate_Success() {
	s.mockSql.ExpectBegin()

	s.mockSql.ExpectExec(regexp.QuoteMeta(`UPDATE courses SET title=$1, description=$2, category=$3, video_url=$4, duration=$5, instructor_name=$6, rating=$7 WHERE id=$8`)).WithArgs(exampleCourse.Title, exampleCourse.Description, exampleCourse.Category, exampleCourse.Video_URL, exampleCourse.Duration, exampleCourse.Instructor_Name, exampleCourse.Rating, exampleCourse.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	s.mockSql.ExpectCommit()

	err := s.repo.Update(exampleCourse)
	s.NoError(err)
}

func (s *CourseRepositoryTestSuite) TestUpdate_Failed() {
	s.mockSql.ExpectBegin()

	s.mockSql.ExpectExec(regexp.QuoteMeta(`UPDATE courses SET title=$1, description=$2, category=$3, video_url=$4, duration=$5, instructor_name=$6, rating=$7 WHERE id=$8`)).WithArgs(exampleCourse.Title, exampleCourse.Description, exampleCourse.Category, exampleCourse.Video_URL, exampleCourse.Duration, exampleCourse.Instructor_Name, exampleCourse.Rating, exampleCourse.ID).WillReturnError(fmt.Errorf("something wrong"))

	s.mockSql.ExpectCommit()

	err := s.repo.Update(exampleCourse)
	s.Error(err)
}

func (s *CourseRepositoryTestSuite) TestDelete_Success() {
	s.mockSql.ExpectBegin()

	s.mockSql.ExpectExec(regexp.QuoteMeta(`UPDATE courses SET deleted_at=$1 WHERE id=$2`)).WithArgs(time.Now(), exampleCourse.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	s.mockSql.ExpectCommit()

	err := s.repo.Delete(dto.CourseIDDto{ID: exampleCourse.ID})
	s.NoError(err)
}

func (s *CourseRepositoryTestSuite) TestDelete_Failed() {
	s.mockSql.ExpectBegin()

	s.mockSql.ExpectExec(regexp.QuoteMeta(`UPDATE courses SET deleted_at=$1 WHERE id=$2`)).WithArgs(time.Now(), exampleCourse.ID).WillReturnError(fmt.Errorf("something wrong"))

	s.mockSql.ExpectCommit()

	err := s.repo.Delete(dto.CourseIDDto{ID: exampleCourse.ID})
	s.Error(err)
}

func TestCourseRepoTestSuite(t *testing.T) {
	suite.Run(t, new(CourseRepositoryTestSuite))
}
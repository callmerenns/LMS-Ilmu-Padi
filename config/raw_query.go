package config

const (
	SelectCourseList = `SELECT id, title, description, content_url, category, is_free, created_at, updated_at, deleted_at FROM courses WHERE user_id = $3 ORDER BY created_at DESC LIMIT $1 OFFSET $2`

	SelectCourseListFull = `SELECT id, title, description, content_url, category, is_free, created_at, updated_at, deleted_at FROM courses WHERE title BETWEEN $3 AND $4 AND user_id = $5 ORDER BY created_at DESC LIMIT $1 OFFSET $2`

	SelectCourseByID = `SELECT id, title, description, content_url, category, is_free, created_at, updated_at, deleted_at FROM courses WHERE id = $1`

	SelectCourseByCategoryType = `SELECT id, title, description, content_url, category, is_free, created_at, updated_at, deleted_at FROM courses WHERE category=$1 AND user_id = $2 ORDER BY created_at DESC`

	SelectCountCourse = `SELECT COUNT(*) FROM courses WHERE user_id=$1`

	InsertCourse = `INSERT INTO courses (title, description, content_url, category, user_id, is_free, created_at, updated_at deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, title, created_at`
)

package course

import "time"

// Course represents a course.
type Course struct {
	CourseID          string    `bson:"courseId" json:"courseId"`
	CourseName        string    `bson:"courseName" json:"courseName"`
	TeacherID         string    `bson:"teacherId" json:"teacherId"`
	CourseDescription string    `bson:"courseDescription" json:"courseDescription"`
	Lessons           []Lesson  `bson:"lessons" json:"lessons"`
	Creation          time.Time `bson:"creation" json:"creation"`
}

// Lesson represents a lesson of course.
type Lesson struct {
	LessonName        string    `bson:"lessonName" json:"lessonName"`
	LessonID          string    `bson:"lessonId" json:"lessonId"`
	LessonDescription string    `bson:"lessonDescription" json:"lessonDescription"`
	Creation          time.Time `bson:"creation" json:"creation"`
}

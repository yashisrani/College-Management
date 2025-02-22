package store

import (
	"github.com/google/uuid"
	"github.com/yashisrani/Go-Backend/model"
	"github.com/yashisrani/Go-Backend/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgress struct {
	DB *gorm.DB
}

func (store *Postgress) NewStore() error {
	dsn := "host=localhost user=yash1 password=password1 dbname=project1 port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		utils.Log(model.LogLevelError, model.Store, model.NewStore, "error while connecting database", err)
		return err
	} else {
		utils.Log(model.LogLevelInfo, model.Store, model.NewStore, "Database connected successfully", nil)
		store.DB = db
	}

	// apply automigration on model
	err = db.AutoMigrate(
		model.User{},
		model.College{},
		model.Class{},
		model.Teacher{},
		model.Book{},
		model.ComputerLab{},
		model.Course{},
	)
	if err != nil {
		utils.Log(model.LogLevelError, model.Store, model.NewStore, "error while automigration", err)
		return err
	}

	return nil
}

//  store (store.go) -> controller(server.go) -> cmd(main.go)

type StoreOperation interface {
	NewStore() error
	CreateUser(user *model.User) error
	GetUsers() ([]model.User, error)
	GetUserByID(userID uuid.UUID) (*model.User, error)
	GetUserByFilter(filter map[string]interface{}) ([]model.User, error)
	DeleteUser(userID string) error
	UpdateUser(user *model.User) error
	SignUp(user *model.User) error
	SignIn(userSignIn model.UserSignIn) (*model.User, error)

	CreateCollege(College *model.College) error
	GetColleges() ([]model.College, error)
	GetCollegeByID(CollegeID uuid.UUID) (*model.College, error)
	GetCollegeByFilter(filter map[string]interface{}) ([]model.College, error)
	DeleteCollege(CollegeID string) error
	UpdateCollege(College *model.College) error

	CreateClass(Class *model.Class) error
	GetClass() ([]model.Class, error)
	GetClassByID(ClassID uuid.UUID) (*model.Class, error)
	GetClassByFilter(filter map[string]interface{}) ([]model.Class, error)
	DeleteClass(ClassID string) error
	UpdateClass(Class *model.Class) error

	CreateTeacher(Teacher *model.Teacher) error
	GetTeacher() ([]model.Teacher, error)
	GetTeacherByID(TeacherID uuid.UUID) (*model.Teacher, error)
	GetTeacherByFilter(filter map[string]interface{}) ([]model.Teacher, error)
	DeleteTeacher(TeacherID string) error
	UpdateTeacher(Teacher *model.Teacher) error

	CreateBook(Book *model.Book) error
	GetBook() ([]model.Book, error)
	GetBookByID(BookID uuid.UUID) (*model.Book, error)
	GetBookByFilter(filter map[string]interface{}) ([]model.Book, error)
	DeleteBook(BookID string) error
	UpdateBook(Book *model.Book) error

	CreateComputerLab(ComputerLab *model.ComputerLab) error
	GetComputerLab() ([]model.ComputerLab, error)
	GetComputerLabByID(ComputerLabID uuid.UUID) (*model.ComputerLab, error)
	GetComputerLabByFilter(filter map[string]interface{}) ([]model.ComputerLab, error)
	DeleteComputerLab(ComputerLabID string) error
	UpdateComputerLab(ComputerLab *model.ComputerLab) error

	CreateCourse(Course *model.Course) error
	GetCourse() ([]model.Course, error)
	GetCourseByID(CourseID uuid.UUID) (*model.Course, error)
	GetCourseByFilter(filter map[string]interface{}) ([]model.Course, error)
	DeleteCourse(CourseID string) error
	UpdateCourse(Course *model.Course) error
}

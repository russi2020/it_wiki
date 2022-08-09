package models

import (
	"github.com/jinzhu/gorm"
	u "it_wiki/utils"
	"log"
)

type Question struct {
	gorm.Model
	Title   string `json:"title" ;sql:"title"`
	Text    string `json:"text" ;sql:"text"`
	ThemeId uint   `json:"theme_id" ;sql:"theme_id" ;gorm:"foreignKey:ThemeRefer"`
	UserId  uint   `json:"user_id" ;sql:"user_id" ;gorm:"foreignKey:UserRefer"`
}

type Answer struct {
	gorm.Model
	Text       string `json:"text" ;sql:"text"`
	QuestionId uint   `json:"question_id" ;sql:"question_id" ;gorm:"foreignKey:QuestionRefer"`
	UserId     uint   `json:"user_id" ;sql:"user_id" ;gorm:"foreignKey:UserRefer"`
}

type Category struct {
	gorm.Model
	Topic  string `json:"topic" ;sql:"topic"`
	UserId uint   `json:"user_id" ;sql:"user_id" ;gorm:"foreignKey:UserRefer"`
}

type Theme struct {
	gorm.Model
	Title      string `json:"title" ;sql:"title"`
	CategoryId uint   `json:"category_id" ;sql:"category_id" ;gorm:"foreignKey:CategoryRefer"`
	UserId     uint   `json:"user_id" ;sql:"user_id" ;gorm:"foreignKey:UserRefer"`
}

func (c *Category) Validate() (map[string]interface{}, bool) {

	if c.Topic == "" {
		return u.Message(false, "Category topic should be on the payload"), false
	}

	if c.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	//Все обязательные параметры присутствуют
	return u.Message(true, "success"), true
}

func (c *Category) CreateCategoryRecord() map[string]interface{} {

	if resp, ok := c.Validate(); !ok {
		return resp
	}

	GetDB().Create(c)

	if c.ID <= 0 {
		return u.Message(false, "Failed to create account, connection error.")
	}

	response := u.Message(true, "Category has been created")
	response["categories"] = c
	return response
}

func GetCategoryByName(topic string) map[string]interface{} {

	category := Category{}
	err := GetDB().Table("categories").Where("topic = ?", topic).First(category).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Category not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	resp := u.Message(true, "Get category")
	resp["categories"] = category

	return resp
}

func GetAllCategories() []*Category {
	categories := make([]*Category, 0)
	err := GetDB().Table("categories").Find(&categories).Error
	if err != nil {
		log.Println(err)
		return nil
	}

	return categories
}

func GetCategory(id uint) *Category {
	category := Category{}
	err := GetDB().Table("categories").Where("id = ?", id).Find(&category).Error
	if err != nil {
		log.Println(err)
		return nil
	}

	return &category
}

func (t *Theme) Validate() (map[string]interface{}, bool) {

	if t.Title == "" {
		return u.Message(false, "Theme title should be on the payload"), false
	}

	if t.CategoryId <= 0 {
		return u.Message(false, "Category is not recognized"), false
	}

	if t.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	//Все обязательные параметры присутствуют
	return u.Message(true, "success"), true
}

func (t *Theme) CreateThemeRecord() map[string]interface{} {

	if resp, ok := t.Validate(); !ok {
		return resp
	}

	GetDB().Create(t)

	if t.ID <= 0 {
		return u.Message(false, "Failed to create theme entry, connection error.")
	}

	response := u.Message(true, "Theme has been created")
	response["themes"] = t
	return response
}

func GetAllThemes() []*Theme {
	themes := make([]*Theme, 0)
	err := GetDB().Table("themes").Find(&themes).Error
	if err != nil {
		log.Println(err)
		return nil
	}

	return themes
}

func GetTheme(id uint) *Theme {
	theme := Theme{}
	err := GetDB().Table("themes").Where("id = ?", id).First(&theme).Error
	if err != nil {
		log.Println(err)
		return nil
	}

	return &theme
}

func (question *Question) Validate() (map[string]interface{}, bool) {

	if question.Title == "" {
		return u.Message(false, "Theme title should be on the payload"), false
	}

	if question.Text == "" {
		return u.Message(false, "Text is not recognized"), false
	}

	if question.ThemeId <= 0 {
		return u.Message(false, "Theme is not recognized"), false
	}

	if question.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	//Все обязательные параметры присутствуют
	return u.Message(true, "success"), true
}

func (question *Question) CreateQuestionRecord() map[string]interface{} {

	if resp, ok := question.Validate(); !ok {
		return resp
	}

	if question.ID <= 0 {
		return u.Message(false, "Failed to create question entry, connection error.")
	}

	GetDB().Create(question)

	response := u.Message(true, "Question has been created")
	response["questions"] = question
	return response
}

func GetAllQuestions() []*Question {
	questions := make([]*Question, 0)
	err := GetDB().Table("questions").Find(&questions).Error
	if err != nil {
		log.Println(err)
		return nil
	}

	return questions
}

func GetQuestion(id uint) *Question {
	question := &Question{}
	err := GetDB().Table("questions").Where("id = ?", id).First(&question).Error
	if err != nil {
		log.Println(err)
		return nil
	}

	return question
}

func (answer *Answer) Validate() (map[string]interface{}, bool) {

	if answer.Text == "" {
		return u.Message(false, "Text is not recognized"), false
	}

	if answer.QuestionId <= 0 {
		return u.Message(false, "Theme is not recognized"), false
	}

	if answer.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	return u.Message(true, "success"), true
}

func (answer *Answer) CreateAnswerRecord() map[string]interface{} {

	if resp, ok := answer.Validate(); !ok {
		return resp
	}

	if answer.ID <= 0 {
		return u.Message(false, "Failed to create answer entry, connection error.")
	}

	GetDB().Create(answer)

	response := u.Message(true, "Answer has been created")
	response["answers"] = answer
	return response
}

func GetAllAnswers() []*Answer {
	answers := make([]*Answer, 0)
	err := GetDB().Table("answers").Find(&answers).Error
	if err != nil {
		log.Println(err)
		return nil
	}

	return answers
}

func GetAnswer(id uint) *Answer {
	answer := &Answer{}
	err := GetDB().Table("answers").Where("id = ?", id).First(&answer).Error
	if err != nil {
		log.Println(err)
		return nil
	}

	return answer
}

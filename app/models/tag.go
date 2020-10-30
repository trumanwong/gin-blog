package models

import "time"

type Tag struct {
	Model

	Name string `json:"name"`
	State int `json:"state"`
	Creator int `json:"creator"`
	Modifier int `json:"modifier"`
}

func (Tag) TableName() string {
	return "blog_tags"
}

func (this *Tag) GetOneByName(name string) (tag Tag) {
	db.Select("*").Where("name = ?", name).First(&tag)
	return
}

func (this *Tag) Add(params map[string]interface{}) {
	tag := &Tag{
		Name:     params["name"].(string),
		State:    params["state"].(int),
		//Creator:  params["creator"].(int),
		//Modifier: params["modifier"].(int),
	}
	tag.CreatedAt = time.Now()
	tag.UpdatedAt = time.Now()
	db.Create(tag)
}
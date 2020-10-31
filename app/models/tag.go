package models

type Tag struct {
	Model

	Name string `json:"name"`
	State int `json:"state"`
	Creator int `json:"creator"`
	Modifier int `json:"modifier"`
}

func (this *Tag) Paginate(where interface{}, page int, pageSize int) (data []Tag, count int) {
	db.Model(this).Where(where).Offset(page).Limit(pageSize).Find(&data)
	db.Model(this).Where(where).Count(&count)
	return
}

func (Tag) TableName() string {
	return "tags"
}

func (this *Tag) GetOneByName(name string) {
	db.Where("name = ?", name).First(&this)
}

func (this *Tag) Find(id int)  {
	db.Where("id = ?", id).First(&this)
}

func (this *Tag) Create() {
	db.Model(&this).Create(&this)
}

func (this *Tag) Update() {
	db.Model(&this).Where("id = ?", this.ID).Update(&this)
}

func (this *Tag) Delete(id int)  {
	db.Where("id = ?", id).Delete(&this)
}
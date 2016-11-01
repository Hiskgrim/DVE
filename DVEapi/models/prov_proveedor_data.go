package models

type ProvProveedorData struct {
	IdData     int    `orm:"column(id_data)"`
	IdForm     int    `orm:"column(id_form);null"`
	RegData    int    `orm:"column(reg_data);null"`
	InputName  string `orm:"column(input_name)"`
	InputValue string `orm:"column(input_value)"`
	StatusData int    `orm:"column(status_data)"`
}

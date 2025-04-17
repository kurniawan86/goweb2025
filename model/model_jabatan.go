package model

import "linklist/node"

var DaftarJabatan node.ListJabatan

// CREATE
func CreateJabatan(jab node.Jabatan) bool {
	tempLL := node.ListJabatan{
		Data: jab,
		Link: nil,
	}
	if DaftarJabatan.Link == nil {
		DaftarJabatan.Link = &tempLL
		return true
	} else {
		temp := &DaftarJabatan
		for temp.Link != nil {
			temp = temp.Link
		}
		temp.Link = &tempLL
		return true
	}
	return false
}

// READ
func ReadJabatan() []node.Jabatan {
	daftarJabatan := []node.Jabatan{}
	temp := &DaftarJabatan
	for temp.Link != nil {
		daftarJabatan = append(daftarJabatan, temp.Link.Data)
		temp = temp.Link
	}
	return daftarJabatan
}

// UPDATE
func UpdateJabatan(jab node.Jabatan, id int) bool {
	temp := DaftarJabatan.Link
	for temp != nil {
		if temp.Data.IdJabatan == id {
			temp.Data = jab
			return true
		}
		temp = temp.Link
	}
	return false
}

// DELETE
func DeleteJabatan(id int) bool {
	temp := &DaftarJabatan
	for temp.Link != nil {
		if temp.Link.Data.IdJabatan == id {
			temp.Link = temp.Link.Link
			return true
		}
		temp = temp.Link
	}
	return false
}

func SearchJabatan(id int) bool {
	temp := &DaftarJabatan
	for temp.Link != nil {
		if temp.Link.Data.IdJabatan == id {
			return true
		}
		temp = temp.Link
	}
	return false
}

func GetNamaJabatan(id int) string {
	temp := &DaftarJabatan
	for temp.Link != nil {
		if temp.Link.Data.IdJabatan == id {
			return temp.Link.Data.NamaJabatan
		}
		temp = temp.Link
	}
	return ""
}

package model

import "linklist/node"

var DaftarPegawai node.ListPegawai

func CreatePegawai(emp node.Pegawai) bool {
	tempLL := node.ListPegawai{
		Data: emp,
		Link: nil,
	}
	if DaftarPegawai.Link == nil {
		DaftarPegawai.Link = &tempLL
		return true
	} else {
		temp := &DaftarPegawai
		for temp.Link != nil {
			temp = temp.Link
		}
		temp.Link = &tempLL
		return true
	}
	return false
}

func ReadPegawai() []node.Pegawai {
	daftarPegawai := []node.Pegawai{}
	temp := &DaftarPegawai
	for temp.Link != nil {
		daftarPegawai = append(daftarPegawai, temp.Link.Data)
		temp = temp.Link
	}
	return daftarPegawai
}

// UpdatePegawai memperbarui data pegawai berdasarkan id
func UpdatePegawai(emp node.Pegawai, id int) bool {
	temp := DaftarPegawai.Link
	for temp != nil {
		if temp.Data.ID == id {
			temp.Data = emp
			return true
		}
		temp = temp.Link
	}
	return false
}

// DeletePegawai menghapus pegawai berdasarkan id
func DeletePegawai(id int) bool {
	temp := &DaftarPegawai
	for temp.Link != nil {
		if temp.Link.Data.ID == id {
			// Hapus node dengan menghubungkan ke node berikutnya
			temp.Link = temp.Link.Link
			return true
		}
		temp = temp.Link
	}
	return false
}

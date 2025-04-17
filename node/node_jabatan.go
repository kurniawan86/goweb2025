package node

type Jabatan struct {
	IdJabatan   int
	NamaJabatan string
	Gaji        int
}

type ListJabatan struct {
	Data Jabatan
	Link *ListJabatan
}

package node

type Address struct {
	Jalan, Kota string
	Nomer       int
}

type Pegawai struct {
	ID      int
	Nama    string
	Alamat  Address
	NoTelp  string
	Email   string
	Jabatan int
}

type ListPegawai struct {
	Data Pegawai
	Link *ListPegawai
}

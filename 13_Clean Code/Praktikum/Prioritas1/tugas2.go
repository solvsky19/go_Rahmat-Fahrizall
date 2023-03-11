package main

type Kendaraan struct {
	TotalRoda       int
	KecepatanPerJam int
}

type Mobil struct {
	Kendaraan
}

func (m *Mobil) Berjalan() {
	m.TambahKecepatan(10)
}

func (m *Mobil) TambahKecepatan(kecepatanBaru int) {
	m.KecepatanPerJam += kecepatanBaru
}

func main() {
	mobilCepat := Mobil{}
	mobilCepat.Berjalan()
	mobilCepat.Berjalan()
	mobilCepat.Berjalan()

	mobilLambat := Mobil{}
	mobilLambat.Berjalan()
}

// Struct kendaraan diubah menjadi Kendaraan agar mengikuti standar penamaan dalam bahasa Go.
// Field totalroda dan kecepatanperjam diubah menjadi TotalRoda dan KecepatanPerJam agar
// mengikuti standar penamaan dalam bahasa Go.
// Struct mobil diubah menjadi Mobil agar mengikuti standar penamaan dalam bahasa Go.
// Method berjalan diubah menjadi Berjalan agar mengikuti standar penamaan dalam bahasa Go.
// Method tambahkecepatan diubah menjadi TambahKecepatan agar mengikuti standar penamaan dalam bahasa Go.
// Penambahan variabel kecepatanbaru pada method TambahKecepatan menggunakan operator += agar lebih efisien.
// Variabel mobilcepat dan mobillamban diubah menjadi mobilCepat dan mobilLambat agar mengikuti standar
// penamaan dalam bahasa Go.

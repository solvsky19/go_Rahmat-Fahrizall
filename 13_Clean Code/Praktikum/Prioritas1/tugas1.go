package main

type user struct {
	id       int
	username int
	password int
}

type userservice struct {
	t []user
}

func (u userservice) getallusers() []user {
	return u.t
}

func (u userservice) getuserbyid(id int) user {
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}

	return user{}
}

// Kekurangannya adalah:
// Penamaan struktur user tidak jelas karena hanya berisi tiga tipe data dengan nama yang tidak spesifik.
// Tipe data username dan password adalah int, sedangkan keduanya harus berupa string.
// Demi keamanan, tipe data ID harus berupa string atau uuid.UUID. agar lebih aman dari segi keamanan. Integer
// dapat menyebabkan kolisi ID jika Anda tidak berhati-hati.
// Variabel t pada struct userservice kurang jelas dan sebaiknya diganti dengan nama yang lebih deskriptif.Selain itu,
// ipe data t harus menggunakan bagian pointer struct user([]*user) agar lebih efisien.
// Method getallusers() dan getuserbyid() pada struct userservice seharusnya menerima pointer receiver (*userservice)
// untuk menghindari penggandaan data ketika receiver dilewatkan sebagai parameter.
// tidak perlu menggunakan variabel r dalam metode getuserbyid() dan harus diganti dengan _ karena kita
// tidak memerlukan nilai ini dalam loop. Fungsi getuserbyid() mengembalikan nilai user{}, sementara
// seharusnya mengembalikan nil untuk menunjukkan bahwa pengguna tidak ditemukan. Jadi, tipe kembalian untuk
// method tersebut seharusnya adalah *user atau nil.
// Setiap method pada struct userservice seharusnya diawali dengan huruf kapital untuk dapat diakses oleh package lainnya.

//  Berikut adalah alasan untuk setiap kekurangan tersebut:

// Penamaan yang buruk dapat membuat kode sulit dipahami oleh orang lain atau oleh Anda sendiri
// ketika Anda kembali ke kode itu nanti. Ada baiknya memberi nama struct sesuai dengan apa yang
// dikandungnya. Misalnya, jika struct mengandung detail pengguna, maka sebaiknya dinamai UserDetail
// atau UserInformation.

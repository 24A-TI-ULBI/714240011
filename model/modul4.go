package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Ruangan struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Nama         string             `json:"nama,omitempty" bson:"nama,omitempty"`
	Kode         string             `json:"kode,omitempty" bson:"kode,omitempty"`
	Kapasitas    int                `json:"kapasitas,omitempty" bson:"kapasitas,omitempty"`
	Ketersediaan bool               `json:"ketersediaan" bson:"ketersediaan"`
}

type Jadwal struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	MataKuliah string             `json:"mata_kuliah,omitempty" bson:"mata_kuliah,omitempty"`
	Dosen      string             `json:"dosen,omitempty" bson:"dosen,omitempty"`
	Prodi      string             `json:"prodi,omitempty" bson:"prodi,omitempty"`
	Hari       string             `json:"hari,omitempty" bson:"hari,omitempty"`
	JamMulai   string             `json:"jam_mulai,omitempty" bson:"jam_mulai,omitempty"`
	JamSelesai string             `json:"jam_selesai,omitempty" bson:"jam_selesai,omitempty"`
	Ruangan    Ruangan            `json:"ruangan,omitempty" bson:"ruangan,omitempty"`
}

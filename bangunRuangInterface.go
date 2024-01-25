package main

import "math"

// Definisikan interface
type BangunRuang interface {
	luas() float64
	keliling() float64
	volume() float64
}

// Struktur untuk Balok
type Balok struct {
	Panjang float64
	Lebar   float64
	Tinggi  float64
}

// Implementasikan metode-metode interface untuk Balok
func (b Balok) luas() float64 {
	return 2 * (b.Panjang*b.Lebar + b.Panjang*b.Tinggi + b.Lebar*b.Tinggi)
}

func (b Balok) keliling() float64 {
	return 4 * (b.Panjang + b.Lebar + b.Tinggi)
}

func (b Balok) volume() float64 {
	return b.Panjang * b.Lebar * b.Tinggi
}

// Struktur untuk Kubus
type Kubus struct {
	Sisi float64
}

// Implementasikan metode-metode interface untuk Kubus
func (k Kubus) luas() float64 {
	return 6 * k.Sisi * k.Sisi
}

func (k Kubus) keliling() float64 {
	return 12 * k.Sisi
}

func (k Kubus) volume() float64 {
	return math.Pow(k.Sisi, 3)
}

// Struktur untuk Tabung
type Tabung struct {
	JariJari float64
	Tinggi   float64
}

// Implementasikan metode-metode interface untuk Tabung
func (t Tabung) luas() float64 {
	return 2 * math.Pi * t.JariJari * (t.JariJari + t.Tinggi)
}

func (t Tabung) keliling() float64 {
	return 2 * math.Pi * t.JariJari
}

func (t Tabung) volume() float64 {
	return math.Pi * math.Pow(t.JariJari, 2) * t.Tinggi
}

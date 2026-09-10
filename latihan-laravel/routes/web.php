<?php

use Illuminate\Support\Facades\Route;

Route::get('/salam', function () {
return 'Selamat datang di Pemrograman Web II';
});

Route::get('/Windi Sulaiman/{H1H024005}', function (string $nim) {
return 'Data mahasiswa dengan NIM ' . $nim;
});

Route::get('/praktikum pemweb2/{TK245010}', function (?string $kode = null) {
if ($kode === null) {
return 'Menampilkan seluruh matakuliah';
}
return 'Menampilkan matakuliah kode ' . $kode;
});

Route::get('/semester/{angka}', function (string $angka) {
    return 'Semester ke ' . $angka;
})->whereNumber('angka');

use App\Http\Controllers\MahasiswaController;

Route::get('/data-mahasiswa', [MahasiswaController::class, 'index'])->name('mahasiswa.index');
Route::get('/data-mahasiswa/{nim}', [MahasiswaController::class, 'show'])->name('mahasiswa.show');

Route::get('/cari-mahasiswa', [MahasiswaController::class, 'cari']);

use App\Http\Controllers\MatakuliahController;

Route::get('/matakuliah', [MatakuliahController::class, 'index'])->name('matakuliah.index');
Route::get('/matakuliah/{kode}', [MatakuliahController::class, 'show'])->name('matakuliah.show');
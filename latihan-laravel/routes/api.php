<?php

use Illuminate\Support\Facades\Route;
use App\Http\Controllers\Api\MahasiswaController;
use App\Http\Controllers\Api\MatakuliahController;
use App\Http\Controllers\Api\AuthController;

// 1. Rute Publik (BEBAS TANPA TOKEN / AUTH)
Route::get('/status', function () {
    return response()->json([
        'sukses' => true,
        'pesan' => 'API Pemweb II aktif',
        'waktu' => now()->toIso8601String(),
    ]);
});

Route::post('/auth/register', [AuthController::class, 'register']);
Route::post('/auth/login', [AuthController::class, 'login'])->middleware('throttle:5,1');

// 2. Rute Terproteksi Token (HARUS PAKAI TOKEN)
Route::middleware('auth:sanctum')->group(function () {
    
    // Auth & Profil
    Route::get('/auth/profil', [AuthController::class, 'profil']);
    Route::put('/auth/password', [AuthController::class, 'updatePassword']);
    Route::post('/auth/logout', [AuthController::class, 'logout']);
    Route::post('/auth/logout-semua', [AuthController::class, 'logoutSemua']);

    // Relasi Program Studi
    Route::get('/program-studi/{id}/mahasiswa', [MahasiswaController::class, 'byProgramStudi']);

    // Mahasiswa
    Route::get('/mahasiswa', [MahasiswaController::class, 'index'])->middleware('ability:mahasiswa:baca');
    Route::get('/mahasiswa/{id}', [MahasiswaController::class, 'show'])->middleware('ability:mahasiswa:baca');
    Route::post('/mahasiswa', [MahasiswaController::class, 'store'])->middleware('ability:mahasiswa:tulis');
    Route::put('/mahasiswa/{id}', [MahasiswaController::class, 'update'])->middleware('ability:mahasiswa:tulis');
    Route::delete('/mahasiswa/{id}', [MahasiswaController::class, 'destroy'])->middleware(['ability:mahasiswa:tulis', 'peran.admin']);

    // Matakuliah
    Route::apiResource('matakuliah', MatakuliahController::class);
});
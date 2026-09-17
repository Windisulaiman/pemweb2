<?php

namespace Database\Seeders;

use App\Models\MataKuliah;
use Illuminate\Database\Seeder;

class MataKuliahSeeder extends Seeder
{
    /**
     * Run the database seeds.
     */
    public function run(): void
    {
        $daftar = [
            [
                'kode' => 'TK101',
                'nama' => 'Pemrograman Dasar',
                'sks' => 3,
                'semester' => 1,
            ],
            [
                'kode' => 'TK102',
                'nama' => 'Algoritma',
                'sks' => 3,
                'semester' => 2,
            ],
            [
                'kode' => 'TK201',
                'nama' => 'Pemrograman Web',
                'sks' => 3,
                'semester' => 3,
            ],
            [
                'kode' => 'TK202',
                'nama' => 'Basis Data',
                'sks' => 3,
                'semester' => 3,
            ],
            [
                'kode' => 'TK501',
                'nama' => 'Internet of Things',
                'sks' => 3,
                'semester' => 4,
            ],
            [
                'kode' => 'TK302',
                'nama' => 'Jaringan Komputer',
                'sks' => 3,
                'semester' => 4,
            ],
            [
                'kode' => 'TK601',
                'nama' => 'Sistem Operasi',
                'sks' => 3,
                'semester' => 5,
            ],
            [
                'kode' => 'TK402',
                'nama' => 'Keamanan Jaringan',
                'sks' => 3,
                'semester' => 5,
            ],
        ];

        foreach ($daftar as $item) {
            MataKuliah::create($item);
        }
    }
}
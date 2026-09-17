<?php

namespace Database\Seeders;

use App\Models\Mahasiswa;
use App\Models\MataKuliah;
use Illuminate\Database\Seeder;

class MahasiswaMataKuliahSeeder extends Seeder
{
    public function run(): void
    {
        $mahasiswas = Mahasiswa::all();
        $matakuliahs = MataKuliah::all();

        foreach ($mahasiswas as $mahasiswa) {
            $jumlahMataKuliah = rand(3, 5);

            $mataKuliahTerpilih = $matakuliahs
                ->random(min($jumlahMataKuliah, $matakuliahs->count()));

            foreach ($mataKuliahTerpilih as $matakuliah) {
                $mahasiswa->mataKuliahs()->attach(
                    $matakuliah->id,
                    [
                        'nilai' => rand(25, 40) / 10,
                    ]
                );
            }
        }
    }
}
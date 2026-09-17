<?php

namespace Database\Seeders;

use App\Models\Mahasiswa;
use Illuminate\Database\Seeder;
use Illuminate\Database\Console\Seeds\WithoutModelEvents;

class DatabaseSeeder extends Seeder
{
    use WithoutModelEvents;

    /**
     * Seed the application's database.
     */
    public function run(): void
    {
        $this->call(ProgramStudiSeeder::class);

         $this->call(MataKuliahSeeder::class);
         Mahasiswa::factory()->count(30)->create();
         
         $this->call(MahasiswaMataKuliahSeeder::class);
    }
}
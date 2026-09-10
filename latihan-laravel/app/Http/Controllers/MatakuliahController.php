<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class MatakuliahController extends Controller
{
    private function getMatakuliahData()
    {
        return [
            ['kode' => 'MK001', 'nama' => 'Pemrograman Web II', 'sks' => 3],
            ['kode' => 'MK002', 'nama' => 'Algoritma & Struktur Data', 'sks' => 2],
            ['kode' => 'MK003', 'nama' => 'Pancasila', 'sks' => 2],
            ['kode' => 'MK004', 'nama' => 'Basis Data', 'sks' => 3],
            ['kode' => 'MK005', 'nama' => 'Bahasa Inggris', 'sks' => 2],
        ];
    }

    public function index(Request $request)
    {
        $matakuliah = $this->getMatakuliahData();
        $q = $request->query('q', '');

        // Fitur pencarian berdasarkan kode atau nama matakuliah
        if ($q !== '') {
            $matakuliah = array_filter($matakuliah, function ($item) use ($q) {
                return stripos($item['nama'], $q) !== false || stripos($item['kode'], $q) !== false;
            });
        }

        return view('matakuliah.index', [
            'daftarMatakuliah' => $matakuliah,
            'q' => $q
        ]);
    }

    public function show(string $kode)
    {
        $matakuliah = $this->getMatakuliahData();
        $detail = null;

        foreach ($matakuliah as $item) {
            if ($item['kode'] === $kode) {
                $detail = $item;
                break;
            }
        }

        return view('matakuliah.show', [
            'kode' => $kode,
            'matakuliah' => $detail
        ]);
    }
}
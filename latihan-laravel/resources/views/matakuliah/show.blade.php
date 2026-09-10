@extends('layouts.app')

@section('judul', 'Detail Matakuliah')

@section('konten')
<h1 class="h3 mb-4">Detail Matakuliah</h1>

<div class="card" style="max-width: 500px;">
    <div class="card-body">
        @if ($matakuliah)
            <h5 class="card-title">{{ $matakuliah['nama'] }}</h5>
            <p class="card-text mb-1"><strong>Kode:</strong> {{ $matakuliah['kode'] }}</p>
            <p class="card-text"><strong>Jumlah SKS:</strong> <x-badge-sks :sks="$matakuliah['sks']" /></p>
        @else
            <p class="text-danger mb-0">Matakuliah dengan kode <strong>{{ $kode }}</strong> tidak ditemukan.</p>
        @endif
    </div>
</div>

<a href="{{ route('matakuliah.index') }}" class="btn btn-secondary mt-3">Kembali</a>
@endsection
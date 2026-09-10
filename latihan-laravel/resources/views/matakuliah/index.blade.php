@extends('layouts.app')

@section('judul', 'Daftar Matakuliah')

@section('konten')
<h1 class="h3 mb-4">Daftar Matakuliah</h1>

<!-- Form Pencarian -->
<form action="{{ route('matakuliah.index') }}" method="GET" class="mb-3">
    <div class="input-group" style="max-width: 400px;">
        <input type="text" name="q" class="form-control" placeholder="Cari kode atau nama..." value="{{ $q }}">
        <button type="submit" class="btn btn-primary">Cari</button>
        @if($q)
            <a href="{{ route('matakuliah.index') }}" class="btn btn-outline-secondary">Reset</a>
        @endif
    </div>
</form>

<table class="table table-bordered bg-white">
    <thead>
        <tr>
            <th>Kode</th>
            <th>Nama Matakuliah</th>
            <th>SKS</th>
            <th>Aksi</th>
        </tr>
    </thead>
    <tbody>
        @forelse ($daftarMatakuliah as $mk)
        <tr>
            <td>{{ $mk['kode'] }}</td>
            <td>{{ $mk['nama'] }}</td>
            <td><x-badge-sks :sks="$mk['sks']" /></td>
            <td>
                <a href="{{ route('matakuliah.show', $mk['kode']) }}" class="btn btn-sm btn-primary">
                    Detail
                </a>
            </td>
        </tr>
        @empty
        <tr>
            <td colspan="4" class="text-center">Data matakuliah tidak ditemukan.</td>
        </tr>
        @endforelse
    </tbody>
</table>
@endsection
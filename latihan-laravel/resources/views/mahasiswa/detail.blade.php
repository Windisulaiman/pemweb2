<!DOCTYPE html>
<html lang="id">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">

    <title>Detail Mahasiswa</title>

    <style>
        body {
            font-family: Arial, sans-serif;
            margin: 40px;
        }

        h1 {
            margin-bottom: 20px;
        }

        .info {
            margin-bottom: 30px;
        }

        .info p {
            margin: 8px 0;
        }

        table {
            width: 100%;
            border-collapse: collapse;
            margin-top: 15px;
        }

        th, td {
            border: 1px solid #ccc;
            padding: 10px;
            text-align: left;
        }

        th {
            background-color: #f2f2f2;
        }

        .nilai {
            text-align: center;
        }
    </style>
</head>

<body>

    <h1>Detail Mahasiswa</h1>

    <div class="info">
        <p>
            <strong>NIM:</strong>
            {{ $mahasiswa->nim }}
        </p>

        <p>
            <strong>Nama:</strong>
            {{ $mahasiswa->nama }}
        </p>

        <p>
            <strong>Email:</strong>
            {{ $mahasiswa->email }}
        </p>

        <p>
            <strong>Program Studi:</strong>
            {{ $mahasiswa->programStudi->nama }}
        </p>

        <p>
            <strong>Angkatan:</strong>
            {{ $mahasiswa->angkatan }}
        </p>

        <p>
            <strong>IPK:</strong>
            {{ $mahasiswa->ipk }}
        </p>
    </div>

    <h2>Daftar Mata Kuliah</h2>

    <table>
        <thead>
            <tr>
                <th>No</th>
                <th>Kode</th>
                <th>Mata Kuliah</th>
                <th>SKS</th>
                <th>Semester</th>
                <th>Nilai</th>
            </tr>
        </thead>

        <tbody>
            @forelse ($mahasiswa->mataKuliahs as $index => $matakuliah)
                <tr>
                    <td>{{ $index + 1 }}</td>
                    <td>{{ $matakuliah->kode }}</td>
                    <td>{{ $matakuliah->nama }}</td>
                    <td>{{ $matakuliah->sks }}</td>
                    <td>{{ $matakuliah->semester }}</td>
                    <td class="nilai">
                        {{ $matakuliah->pivot->nilai }}
                    </td>
                </tr>
            @empty
                <tr>
                    <td colspan="6">
                        Belum ada mata kuliah yang diambil.
                    </td>
                </tr>
            @endforelse
        </tbody>
    </table>

</body>
</html>
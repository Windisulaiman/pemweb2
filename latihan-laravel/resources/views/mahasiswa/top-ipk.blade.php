<!DOCTYPE html>
<html lang="id">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">

    <title>10 Mahasiswa IPK Tertinggi</title>
</head>

<body>

    <h1>10 Mahasiswa dengan IPK Tertinggi</h1>

    <p>
        Program Studi: <strong>Teknik Komputer</strong>
    </p>

    <table border="1" cellpadding="10" cellspacing="0">

        <thead>
            <tr>
                <th>No</th>
                <th>NIM</th>
                <th>Nama</th>
                <th>Angkatan</th>
                <th>IPK</th>
            </tr>
        </thead>

        <tbody>

            @forelse ($daftarMahasiswa as $index => $mahasiswa)

                <tr>
                    <td>{{ $index + 1 }}</td>
                    <td>{{ $mahasiswa->nim }}</td>
                    <td>{{ $mahasiswa->nama }}</td>
                    <td>{{ $mahasiswa->angkatan }}</td>
                    <td>{{ $mahasiswa->ipk }}</td>
                </tr>

            @empty

                <tr>
                    <td colspan="5">
                        Tidak ada data mahasiswa Teknik Komputer.
                    </td>
                </tr>

            @endforelse

        </tbody>

    </table>

</body>

</html>
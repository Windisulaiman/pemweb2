import { useState } from 'react'
import { useNavigate, Link } from 'react-router-dom'
import { simpanMahasiswa } from '../api/layananMahasiswa'

export default function TambahMahasiswa() {
  const [programStudiId, setProgramStudiId] = useState(1)
  const [nim, setNim] = useState('')
  const [nama, setNama] = useState('')
  const [email, setEmail] = useState('')
  const [angkatan, setAngkatan] = useState(2025)
  const [ipk, setIpk] = useState(3.5)
  const [pesanGalat, setPesanGalat] = useState('')
  const [galatValidasi, setGalatValidasi] = useState([])
  const [sedangMemuat, setSedangMemuat] = useState(false)
  const navigasi = useNavigate()

  async function tanganiSimpan(peristiwa) {
    peristiwa.preventDefault()
    setPesanGalat('')
    setGalatValidasi([])
    setSedangMemuat(true)

    try {
      await simpanMahasiswa({
        program_studi_id: Number(programStudiId),
        nim: nim,
        nama: nama,
        email: email,
        angkatan: Number(angkatan),
        ipk: Number(ipk),
      })
      navigasi('/mahasiswa')
    } catch (galat) {
      if (galat.response) {
        const data = galat.response.data
        setPesanGalat(data.pesan || data.message || 'Gagal menyimpan data')

        if (Array.isArray(data.galat)) {
          setGalatValidasi(data.galat.map((item) => `${item.field}: ${item.pesan}`))
        } else if (data.errors && typeof data.errors === 'object') {
          setGalatValidasi(Object.values(data.errors).flat())
        }
      } else {
        setPesanGalat('Tidak dapat menghubungi server')
      }
    } finally {
      setSedangMemuat(false)
    }
  }

  return (
    <div style={{ maxWidth: 480, margin: '40px auto' }}>
      <h1>Tambah Mahasiswa</h1>
      <Link to="/mahasiswa">← Kembali ke Daftar</Link>

      {pesanGalat && <p style={{ color: 'crimson', marginTop: 12 }}>{pesanGalat}</p>}
      {galatValidasi.length > 0 && (
        <ul style={{ color: 'crimson' }}>
          {galatValidasi.map((pesan, indeks) => (
            <li key={indeks}>{pesan}</li>
          ))}
        </ul>
      )}

      <form onSubmit={tanganiSimpan} style={{ marginTop: 16 }}>
        <div style={{ marginBottom: 10 }}>
          <label>ID Program Studi</label><br />
          <input
            type="number"
            value={programStudiId}
            onChange={(e) => setProgramStudiId(e.target.value)}
            required
          />
        </div>
        <div style={{ marginBottom: 10 }}>
          <label>NIM</label><br />
          <input
            type="text"
            value={nim}
            onChange={(e) => setNim(e.target.value)}
            placeholder="Contoh: H1A125001"
            required
          />
        </div>
        <div style={{ marginBottom: 10 }}>
          <label>Nama Lengkap</label><br />
          <input
            type="text"
            value={nama}
            onChange={(e) => setNama(e.target.value)}
            required
          />
        </div>
        <div style={{ marginBottom: 10 }}>
          <label>Email</label><br />
          <input
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
        </div>
        <div style={{ marginBottom: 10 }}>
          <label>Angkatan</label><br />
          <input
            type="number"
            value={angkatan}
            onChange={(e) => setAngkatan(e.target.value)}
            required
          />
        </div>
        <div style={{ marginBottom: 10 }}>
          <label>IPK</label><br />
          <input
            type="number"
            step="0.01"
            value={ipk}
            onChange={(e) => setIpk(e.target.value)}
            required
          />
        </div>
        <button type="submit" disabled={sedangMemuat}>
          {sedangMemuat ? 'Menyimpan...' : 'Simpan Mahasiswa'}
        </button>
      </form>
    </div>
  )
}
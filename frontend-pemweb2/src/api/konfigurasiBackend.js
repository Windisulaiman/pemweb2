// Ubah nilai ini ke 'laravel' atau 'fiber' untuk berpindah backend dalam satu pengaturan
export const BACKEND_AKTIF = 'laravel'

const KONFIGURASI = {
  laravel: {
    baseURL: import.meta.env.VITE_API_LARAVEL,
    fieldSandi: 'password',
    ambilHalaman: (meta) => meta?.current_page ?? 1,
    ambilTotalHalaman: (meta) => meta?.last_page ?? 1,
  },
  fiber: {
    baseURL: import.meta.env.VITE_API_FIBER,
    fieldSandi: 'kata_sandi',
    ambilHalaman: (meta) => meta?.halaman ?? 1,
    ambilTotalHalaman: (meta) => meta?.total_halaman ?? 1,
  },
}

export const konfigurasiAktif = KONFIGURASI[BACKEND_AKTIF]
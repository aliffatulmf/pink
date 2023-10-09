# Pink (Ping)

Pink adalah layanan backend yang dirancang untuk memeriksa apakah sebuah situs web dilarang oleh pemerintah Indonesia. Ini ditulis dalam Go dan menggunakan setup server dan router.

## Inspirasi

Pink terinspirasi oleh [Indiwtf](https://indiwtf.upset.dev/) yang dibuat oleh [Upset](https://upset.dev/). Indiwtf adalah layanan yang memeriksa apakah sebuah situs web dilarang oleh pemerintah Indonesia.

## Environment variables

- `PINK_HOST`: Host untuk server.
- `PINK_PORT`: Port untuk server.
- `PINK_ENV`: Lingkungan untuk server. Ini bisa "development" atau "production".

## Endpoint API

File `router.go` mendefinisikan endpoint API berikut:

### `/api/v1/check`

- **Method:** GET
- **Fungsi Handler:** `VerifyDomain`
- **Deskripsi:** Endpoint ini memeriksa apakah domain valid atau tidak. Detail apa yang diperiksa ditentukan dalam fungsi `VerifyDomain`.

### `/api/v1/list`

- **Method:** GET
- **Fungsi Handler:** `ListDomains`
- **Deskripsi:** Endpoint ini mengembalikan daftar domain. Detail apa yang dikembalikan ditentukan dalam fungsi `ListDomains`.

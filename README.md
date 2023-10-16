# Pink (Ping)

Pink adalah layanan backend yang dirancang untuk memeriksa apakah sebuah situs web dilarang oleh pemerintah Indonesia. Ini ditulis dalam Go dan menggunakan setup server dan router.

## Inspirasi

Pink terinspirasi oleh [Indiwtf](https://indiwtf.upset.dev/) yang dibuat oleh [Upset](https://upset.dev/). Indiwtf adalah layanan yang memeriksa apakah sebuah situs web dilarang oleh pemerintah Indonesia.

## Arguments

- `-host`: Host untuk server.
- `-port`: Port untuk server.
- `-env`: Lingkungan untuk server. Ini bisa "development" atau "production".

## Endpoint API

File `router.go` mendefinisikan endpoint API berikut:
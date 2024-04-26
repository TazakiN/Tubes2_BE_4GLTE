
# Wiki Race Path Finder - Backend

Repositori ini berisi kode backend dari aplikasi Wiki Race Path Finder yang dibuat dalam bahasa Go dengan bantuan Framework [Gin](https://gin-gonic.com/). Aplikasi ini memungkinkan pengguna untuk mencari jalur terpendek antara dua artikel di Wikipedia dengan menggunakan algoritma BFS dan IDS.

## Table of Contents

1. [Penjelasan Algoritma](#penjelasan-algoritma)
    - [BFS (Breadth First Search)](#bfs-breadth-first-search)
    - [IDS (Iterative Deepening Search)](#ids-iterative-deepening-search)
2. [Cara Menjalankan Aplikasi](#cara-menjalankan-aplikasi)
    - [Menjalankan tanpa Docker](#menjalankan-tanpa-docker)
    - [Menjalankan dengan Docker](#menjalankan-dengan-docker)
3. [Contoh Penggunaan Aplikasi](#contoh-penggunaan-aplikasi)
4. [Kontributor (4GLTE)](#kontributor-4glte)

## Penjelasan Algoritma

### BFS (Breadth First Search)

BFS adalah algoritma pencarian yang mengunjungi simpul secara berurutan, dimulai dari simpul awal, kemudian simpul tetangga dari simpul tersebut, dan seterusnya. Algoritma ini akan mengunjungi semua simpul yang terhubung dengan simpul awal sebelum mengunjungi simpul yang lebih jauh.

### IDS (Iterative Deepening Search)

IDS adalah algoritma pencarian yang menggabungkan konsep DFS dan BFS. IDS akan melakukan DFS dengan kedalaman maksimum sesuai masukan pengguna, kemudian akan meningkatkan kedalaman maksimum jika tidak ditemukan solusi. IDS akan terus melakukan DFS dengan kedalaman yang semakin dalam jika masih belum ditemukan solusi.

## Cara Menjalankan Aplikasi

### Menjalankan tanpa Docker

1. Clone repositori ini dan repositori frontend di [sini](https://github.com/TazakiN/Tubes2_FE_4GLTE)
2. Buka terminal dan arahkan ke direktori `src`.
3. - Jalankan perintah `go run main.go` untuk menginstal semua dependensi.
   - Jalankan perintah `./main.exe` jika menggunakan Windows dan ingin menggunakan program yang sudah di-build.
4. Jalankan aplikasi frontend dengan mengikuti petunjuk di [sini](https://github.com/TazakiN/Tubes2_FE_4GLTE)
5. Buka browser dengan alamat `localhost:3000`.
6. Aplikasi siap digunakan.

### Menjalankan dengan Docker

1. Clone repositori ini dan repositori frontend di [sini](https://github.com/TazakiN/Tubes2_FE_4GLTE) pada direktori yang sama (pada contoh ini, direktori `root`).
2. Buat sebuah file bernama docker-compose.yml dengan isi sebagai berikut:

    ```yml
    version: '3.8'

    services:
    backend:
        build: ./Tubes2_BE_4GLTE/src
        ports:
        - "3321:3321"

    frontend:
        build: ./Tubes2_FE_4GLTE/src
        ports:
        - "3000:3000"
    ```

    struktur direktori saat ini adalah sebagai berikut

    ```markdown
    root/
    ├── Tubes2_BE_4GLTE/
    ├── Tubes2_FE_4GLTE/
    └── docker-compose.yml
    ```

3. Buka terminal dan arahkan ke direktori root.
4. Jalankan perintah `docker-compose up` untuk menjalankan aplikasi.
5. Tunggu hingga proses selesai dan buka browser dengan alamat `localhost:3000`.

## Contoh Penggunaan Aplikasi

Berikut adalah tampilan awal aplikasi:

![Tampilan Awal](doc/Tampilan%20awal.png)

Berikut adalah contoh tampilan hasil:

![Tampilan Hasil](doc/Tampilan%20hasil.png)

## Kontributor (4GLTE)

- 10023485 -  [Lidya Rahmatul Fitri](https://github.com/Lidyarf24)
- 13522032 - [Tazkia Nizami](https://github.com/TazakiN)
- 13522097 - [Ellijah Darrellshane Suryanegara](https://github.com/HenryofSkalitz1202)

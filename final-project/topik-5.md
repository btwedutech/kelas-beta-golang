
## Topik #5 : *CLI App & Web Service : Database Auto Restore*

### Deskripsi Project 

Anda adalah seorang system engineer yang bekerja pada sebuah perusahaan multinasional yang bergerak dalam pengembangan sistem enterprise. Saat ini perusahaan anda telah meng-handle lebih dari 100 client mulai dari sistem hingga infrastruktur yang mereka gunakan. 

Suatu hari technical lead di perusahaan anda, meminta teman untuk membuat service app untuk melakukan automation backup database pada setiap project client setiap harinya. Yang dimana setiap database yang telah di *dump* lalu di compress dengan ZIP lalu di upload ke sebuah web service yang fungsi utamanya untuk menerima backup database yang sudah di zip/compress.

Kemudian anda diminta untuk membuat tool yang dapat melakukan restore dari file zip database yang sudah diupload, dengan cara mendownload dari Web Service yang telah dibuat lalu, melakukan `unzip`,  kemudian melakukan import db ke database kosongan.


## Topik #4 : *CLI App & Web Service : Database Auto Backup*

### Deskripsi Project 

Anda adalah seorang sistem engineer yang bekerja pada sebuah perusahaan multinasional yang bergerak dalam pengembangan sistem enterprise. Saat ini perushaan anda telah meng-handle lebih dari 100 client mulai dari sistem hingga infrastruktur yang mereka gunakan. 

Suatu hari technical lead di perusahaan anda, meminta anda untuk membuat service app untuk melakukan automation backup database pada setiap project client setiap harinya. Yang dimana setiap database yang telah di *dump* lalu di compress dengan ZIP lalu di upload ke sebuah web service yang fungsi utamanya untuk menerima backup database yang sudah di zip/compress.

### Project Specification

![Skema](resources/topik-4.jpg)

Pada project ini akan ada 2 aplikasi/service yang ada harus buat

#### 1. Web Service

Sesuai dengan informasi yang diberikan oleh tech lead anda, bahwa akan ada web service yang tugas nya untuk menerima data upload file zip, yang dimana spesifikasi lengkap sebagai berikut.

1. Web Service memiliki beberapa route seperti tabel dibawah
    
   | Path | Method | Req Param | Body | Response | Deskripsi |
   |:-----|:-------|:----------|:-----|:---|:---|
   | /| GET |-|-| [Response 1](#response-1)|Memberikan list database yang dibackup terakhir| 
   | /{db_name}|GET | `db_name` : adalah nama database yang pernah di backup |-|[Response 2](#response-2)| Memberikan list history sebuah database file zip yang pernah di upload |
   | /{db_name}|POST | `multipart/form-data` body dari file zip dan `db_name` sebagai parameter ||| Melakukan upload file zip ke web service|
   |/{id_file}/download| GET | |||Melakukan download file|
   ###### Response 1
   ```json
   {
    "data": [
            {
                "database_name": "pt_xyz",
                "latest_backup": {
                    "id": 34,
                    "file_name": "mysql-2023-10-31-00-00-00-pt_xyz-0f69a75b-9fdf-48c6-8ec9-2277934b7bb8.sql.zip",
                    "timestamp": "2023-10-31 00:00:00"
                }
            },
            {
                "database_name": "pt_abc",
                "latest_backup": {
                    "id": 100,
                    "file_name": "mysql-2023-10-31-00-00-00-pt_abc-c584d377-0082-4473-a32b-29510d922fde.zip",
                    "timestamp": "2023-10-31 00:00:00"
                }
            },
            {
                "database_name": "pt_limajari",
                "latest_backup": {
                    "id": 201,
                    "file_name": "mysql-2023-10-31-00-00-00-pt_limajari-6077b272-9f9f-4fe2-9924-b92b2cbbc2d5.sql.zip",
                    "timestamp": "2023-10-31 00:00:00"
                }
            },
            {
                "database_name": "cv_kucing_oren",
                "latest_backup": {
                    "id": 303,
                    "file_name": "mysql-2023-10-31-00-00-00-cv_kucing_oren-0f69a75b-9fdf-48c6-8ec9-2277934b7bb8.sql.zip",
                    "timestamp": "2023-10-31 00:00:00"
                }
            },
    ],
    "message": "Success"
   }
   ```

   ###### Response 2
   ```json
    {
        "data": {
            "database_name": "cv_kucing_oren",
            "histories": [
                {
                    "id" : 303,
                    "file_name": "mysql-2023-10-31-00-00-00-cv_kucing_oren-0f69a75b-9fdf-48c6-8ec9-2277934b7bb8.sql.zip",
                    "timestamp": "2023-10-31 00:00:00"
                },
                {
                    "id": 299,
                    "file_name": "mysql-2023-10-30-00-00-00-cv_kucing_oren-7634bf3f-23b5-45a7-8b78-fe9b1a3bcf66.sql.zip",
                    "timestamp": "2023-10-30 00:00:00"
                },
                {
                    "id": 295,
                    "file_name": "mysql-2023-10-29-00-00-00-cv_kucing_oren-8634bf3f-23b5-45a7-8b78-fe9b1a3bcf66.sql.zip",
                    "timestamp": "2023-10-29 00:00:00"
                }
            ]
        },
        "message": "success"
    }
   ```
2. 
3. 



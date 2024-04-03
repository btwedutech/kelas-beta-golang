
## Topik #2 : *CLI App : Data Validation and Convert*

### Deskripsi Project 

Anda adalah seorang engineer yang bekerja pada suatu startup. Suatu hari, technical lead startup anda meminta anda untuk mengembangkan tool yang dapat mempermudah proses konversi data CSV ke JSON, sekaligus memastikan integritas dan validitas data tersebut. 

Anda diharuskan untuk membangun sebuah tool yang mampu mengonversi file CSV ke JSON. Tool tersebut juga harus melakukan validasi data berdasarkan skema atau aturan yang ditentukan sebelumnya. Aplikasi ini harus efisien dalam menangani file dengan data hingga ratusan ribu dan dapat menampilkan error akan data yang tidak valid atau masalah yang ditemukan selama proses konversi.


### Project Specification

Pada project ini, anda akan membuat 1 aplikasi berbentuk CLI (Command Line Interface) dengan tambahan support flag, dimana spesifikasinya adalah sebagai berikut :

1. CLI jika dijalankan tanpa flag (```go run main.go```), user akan diminta untuk memasukan lokasi file yang ingin diproses

```bash
Masukan input file : C:\sample\cars.csv
```

2. Output file nanti akan berupa nama input file dengan ekstensi yang diganti sesuai output convertnya, semisal :

```bash
Masukan input file : C:\sample\cars.csv
=========== PROSES COMPLETE ===========
File berhasil divalidasi dan konversi : C:\sample\cars.json
```

3. CLI juga mendukung proses melalui flag (```go run main.go --input xxxx```), menggunakan bantuan package bawaan ```flag```. Contoh implementasinya seperti berikut

```go
	inputFile := flag.String("input", "", "Set input file")
	outputFile := flag.String("output", "", "Set output file (optional)")

	flag.Parse()
```


4. Terdapat 2 flag yang harus diimplementasi :

```bash
--input <input_file> --output <output_file>
```

5. Flag output bersifat opsional, jika tidak diisi, maka akan menyimpan file sesuai lokasi dan nama dari input file, misalnya 

```bash
    Input File : cars.csv

    Output File : cars.json
```

6. Sebelum proses konversi, ketika user telah menginputkan file yang akan diproses ataupun jika dari flag (```go run main.go --input cars.csv```), kita perlu mengambil ekstensi dari file tersebut, dan nama filenya, menggunakan bantuan package bawaan ```path/filepath```

```go
    fileExtension := filepath.Ext(inputFile)
    fileExtension = fileExtension[1:]

    fileName := filepath.Base(inputFile)
    fileName = fileName[:len(fileName)-len(fileExtension)-1]
```

7. Setelah itu mengecek ekstensi inputan file apakah sesuai atau tidak, jika tidak, maka return error, misalnya

```bash
Input file : C:\sample\cars.csx is not a valid CSV file
```

8. Melakukan proses validasi data, dimana kondisi berikut harus terpenuhi :

- File CSV bisa diload dan diparse
- Untuk file CSV, harus terdapat headernya di index pertama
- Jika terdapat header email, harus melakukan validasi akan valuenya merupakan format email valid ```x@x.x```
- Jika terdapat header seperti phone, no, telp, hp harus melakukan validasi tidak ada huruf didalamnya

9. Melakukan proses konversi dari CSV ke JSON, key JSONnya diambil dari header di index pertama dari data didalam CSV tersebut, misalnya 

```csv
id,name,phone,email
1,Ani,628713131331,ani@workmail.com
2,Budy,62813131331,budy@workmail.com
```

sehingga menjadi

```json
[
    {
        "id": 1,
        "name": "Ani",
        "phone": "628713131331",
        "email": "ani@workmail.com"
    },
    {
        "id": 2,
        "name": "Budy",
        "phone": "62813131331",
        "email": "budy@workmail.com"
    }
]
```

10. Jika terjadi error, menampilkan errornya kedalam console/terminal

```bash
Masukan input file : C:\sample\cars.csv
=========== PROSES FAILED ===========
Terjadi error : index is not valid 
```

11. Mengimplementasikan Goroutine dan channeling, agar dapat memproses setidaknya 5 data csv secara paralel saat proses validasi dan konversi json
12. Sample data bisa dicek di resources/sample_datas_small.csv dan resources/sample_datas.csv

### Optional :
Mengimplementasikan package third-party ```github.com/schollz/progressbar/v3``` untuk menunjukan sejauh mana proses yang sedang berjalan

```go
    import "github.com/schollz/progressbar/v3"

	bar := progressbar.Default(int64(csvData), ("Memproses Data"))

	for _, value := range csvData {
        // processing logic
        bar.Add(1)
    }
    bar.Clear()
```

Output di terminal :

```bash
Masukan input file : C:\sample\cars.csv
Memproses Data 100% |███████████████████| (1231/1231, 558 it/s)
=========== PROSES COMPLETE ===========
```
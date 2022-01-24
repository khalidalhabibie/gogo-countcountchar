# tahap:

1.(data collecting) mengamil data artikel, asumsi data artikel adalah atribut title dan paragraph

2.(data preprocessing) selanjutnya data cleaning dari char yang tidak diinginkan, hanya alphabet dikarnakan hanya ingin mengetahui jumlah alphabet. Setelah hanya alphabet maka semua char dijadikan Upper untuk normalisasi agar mempermudah proses perhitungan

3.(processing) pada tahap ini perhitungan akan dipermudahkan dengan konsep map pada golang yaitu variabel memiliki nilai key:value sebagai counter, key akan menyimpan char yang dihitung sedangkan value adalah jumlah chart tersebut. data akan di sort dengan package sort untuk di urutkan dari tertinggi ke terendah.

usage (with make)
```make
make lets-dep
make lets-go
```

usage (without make)
```golang
go mod tidy
go run main.go
```

lets go  :)
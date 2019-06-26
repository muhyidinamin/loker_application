# loker_application
Usage:
● Run the program : 
    go run main.go
● Initial the size of locker. example : 
    init 5
● Input the card identity use type Identity and Identity Number. example : 
    input SIM 12345
● Check the list of locker status. 
    status
● For empty the locker use number of locker you want. example : 
    leave 2 
● To find the Locker number by identity number. example :
    find 12345
● To search all list of identity number by type of identity. example : 
    search SIM
● For exit the program.
    exit

Penggunaan: 
● menjalankan program : 
    go run main.go
● untuk membuat jumlah loker :
    init [jumlah loker]
● untuk menampilkan status dari masing - masing nomor loker
    status
● untuk memasukkan kartu identitas
    input [tipe identitas] [nomor identitas]
● untuk mengosongkan loker
    leave [nomor loker]
● untuk menampilkan nomor loker berdasar nomor identitas
    find [nomor identitas]
● untuk menampilkan daftar nomor identitas sesuai tipe identitas yang dicari
    search [tipe identitas]
● untuk mengakhiri program
    exit 
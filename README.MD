API : Aplication Programming Interface
REST : Representational State Transfer : sebuah arsitektur untuk berkomunikasi antara dua buah system
Use : HTTP Protocol
HTTP request method :
    GET     -> mendapatkan data tanpa parameter indikasi
    POST    -> mengirimkan data beserta parameter indikasi
    PUT     -> mengakses fungsi edit (ubah data)  
    DELETE  ->  mengakses fungsi Delete (hapus data)
    HEAD
    OPTION
    PATCH
Request dan Response format :
    JSON : JavaScript Object Notation (key - value)
    XML
    SOAP

HTTP Response Code :
    200 OK
    201 CREATED
    400 BAD REQUEST
    404 NOT FOUND
    401 UNAUTHORIZED
    405 METHOD NOT ALLOWED
    500 INTERNAL SERVER ERROR

Gunakan Kata Benda ketika membuat endpoint
Naming Using Plural / jamak Nouns
ex : 
    GET /books/123
    DELETE /books/123
    POST /books
    PUT /books/123
    GET /cars

Use Resource Nesting to show Relations or hierarchy
    /users - ambil semua data user
    /users/123 - ambil data user tertentu sesuai dengan parameter
    /users/123/orders - ambil seluruh data order milik user 123
    /users/123/orders/0001 - ambil data order ke 0001 milik user 123

Open API adalah sebuah standar spesifikasi dari dokumentasi RestFul APIs
OPEN API dapat di implementasikan di banyak pemrograman
OPEN API dibuat untuk visualisasi dengan tujuan presenting restful API , mudah untuk di baca.
OPEN APi bisa di tulis menggunakan JSON atau YAML documents
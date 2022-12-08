# <a style="font-family:cursive">Hacktiv8 Final Project 4</a>
<p align="justify"><b>Toko Belanja</b> is an application where the admin have the authority to carry out CRUD processes for categories and products. Meanwhile, customers can process product purchases and view purchase transaction history. Customers also can top-up their balance to buy some products.</p>

# Installation
Requires <b>[Golang](https://go.dev/dl/)</b> and <b>[MySQL](https://sourceforge.net/projects/xampp/files/XAMPP%20Windows/8.1.12/)</b>

Config the .env first to connect into database

- **Clone repository**
```bash
git clone https://github.com/rickyfazaa/tokobelanja-kelompok7
```
- **Change directory**
```sh
cd tokobelanja-kelompok7
```
- **Run "main.go" file**
```sh
go run main.go
```


# Project Structure
> <i>Press to switch into the folder you want to go</i><br>
 
 📦[tokobelanja-kelompok7](https://github.com/rickyfazaa/tokobelanja-kelompok7)<br>
 ┣ 📂[config](https://github.com/rickyfazaa/tokobelanja-kelompok7/tree/master/config)<br>
 ┃ ┣ 📜[db.go](https://github.com/rickyfazaa/tokobelanja-kelompok7/blob/master/config/db.go)<br>
 ┃ ┗ 📜[db_test.go](https://github.com/rickyfazaa/tokobelanja-kelompok7/blob/master/config/db_test.go)<br>
 ┣ 📂[controller](https://github.com/rickyfazaa/tokobelanja-kelompok7/tree/master/controller)<br>
 ┃ ┣ 📜[category_controller.go](https://github.com/rickyfazaa/tokobelanja-kelompok7/blob/master/controller/category_controller.go)<br>
 ┃ ┣ 📜[product_controller.go](https://github.com/rickyfazaa/tokobelanja-kelompok7/blob/master/controller/product_controller.go)<br>
 ┃ ┣ 📜[transaction_controller.go](https://github.com/rickyfazaa/tokobelanja-kelompok7/blob/master/controller/transaction_controller.go)<br>
 ┃ ┗ 📜[user_controller.go](https://github.com/rickyfazaa/tokobelanja-kelompok7/blob/master/controller/user_controller.go)<br>
 ┣ 📂[helper](https://github.com/rickyfazaa/tokobelanja-kelompok7/tree/master/helper)<br>
 ┃ ┣ 📜[error.go](https://github.com/rickyfazaa/tokobelanja-kelompok7/blob/master/helper/error.go)<br>
 ┃ ┣ 📜[error_test.go](https://github.com/rickyfazaa/tokobelanja-kelompok7/blob/master/helper/error_test.go)<br>
 ┃ ┣ 📜[response.go](https://github.com/rickyfazaa/tokobelanja-kelompok7/blob/master/helper/response.go)<br>
 ┃ ┗ 📜[response_test.go](https://github.com/rickyfazaa/tokobelanja-kelompok7/blob/master/helper/response_test.go)<br>
 ┣ 📂[middleware](https://github.com/rickyfazaa/tokobelanja-kelompok7/tree/master/middleware)<br>
 ┃ ┣ 📜[jwt.go](https://github.com/rickyfazaa/tokobelanja-kelompok7/blob/master/middleware/jwt.go)<br>
 ┃ ┣ 📜[jwt_test.go](https://github.com/rickyfazaa/tokobelanja-kelompok7/blob/master/middleware/jwt_test.go)<br>
 ┃ ┣ 📜[middleware.go](https://github.com/rickyfazaa/tokobelanja-kelompok7/blob/master/middleware/middleware.go)<br>
 ┃ ┗ 📜[middleware_test.go](https://github.com/rickyfazaa/tokobelanja-kelompok7/blob/master/middleware/middleware_test.go)<br>
 ┣ 📂[model](https://github.com/rickyfazaa/tokobelanja-kelompok7/tree/master/model)<br>
 ┃ ┣ 📂[entity](https://github.com/rickyfazaa/tokobelanja-kelompok7/tree/master/model/entity)<br>
 ┃ ┃ ┣ 📜[category.go](https://github.com/rickyfazaa/tokobelanja-kelompok7/blob/master/model/entity/category.go)<br>
 ┃ ┃ ┣ 📜[product.go](https://github.com/rickyfazaa/tokobelanja-kelompok7/blob/master/model/entity/product.go)<br>
 ┃ ┃ ┣ 📜[transaction_history.go](https://github.com/rickyfazaa/tokobelanja-kelompok7/blob/master/model/entity/transaction_history.go)<br>
 ┃ ┃ ┗ 📜[user.go](https://github.com/rickyfazaa/tokobelanja-kelompok7/blob/master/model/entity/user.go)<br>
 ┃ ┣ 📂[input](https://github.com/rickyfazaa/tokobelanja-kelompok7/tree/master/model/input)<br>
 ┃ ┃ ┣ 📜[category_input.go](https://github.com/rickyfazaa/tokobelanja-kelompok7/blob/master/model/input/category_input.go)<br>
 ┃ ┃ ┣ 📜[product_input.go](https://github.com/rickyfazaa/tokobelanja-kelompok7/blob/master/model/input/product_input.go)<br>
 ┃ ┃ ┣ 📜[transaction_history_input.go](https://github.com/rickyfazaa/tokobelanja-kelompok7/blob/master/model/input/transaction_history_input.go)<br>
 ┃ ┃ ┗ 📜[user_input.go](https://github.com/rickyfazaa/tokobelanja-kelompok7/blob/master/model/input/user_input.go)<br>
 ┃ ┗ 📂[response](https://github.com/rickyfazaa/tokobelanja-kelompok7/tree/master/model/response)<br>
 ┃ ┃ ┣ 📜[category_response.go](https://github.com/rickyfazaa/tokobelanja-kelompok7/blob/master/model/response/category_response.go)<br>
 ┃ ┃ ┣ 📜[product_response.go](https://github.com/rickyfazaa/tokobelanja-kelompok7/blob/master/model/response/product_response.go)<br>
 ┃ ┃ ┣ 📜[transaction_history_response.go](https://github.com/rickyfazaa/tokobelanja-kelompok7/blob/master/model/response/transaction_history_response.go)<br>
 ┃ ┃ ┗ 📜[user_response.go](https://github.com/rickyfazaa/tokobelanja-kelompok7/blob/master/model/response/user_response.go)<br>
 ┣ 📂[repository](https://github.com/rickyfazaa/tokobelanja-kelompok7/tree/master/repository)<br>
 ┃ ┣ 📜[category_repository.go](https://github.com/rickyfazaa/tokobelanja-kelompok7/blob/master/repository/category_repository.go)<br>
 ┃ ┣ 📜[product_repository.go](https://github.com/rickyfazaa/tokobelanja-kelompok7/blob/master/repository/product_repository.go)<br>
 ┃ ┣ 📜[transaction_repository.go](https://github.com/rickyfazaa/tokobelanja-kelompok7/blob/master/repository/transaction_repository.go)<br>
 ┃ ┗ 📜[user_repository.go](https://github.com/rickyfazaa/tokobelanja-kelompok7/blob/master/repository/user_repository.go)<br>
 ┣ 📂[service](https://github.com/rickyfazaa/tokobelanja-kelompok7/tree/master/service)<br>
 ┃ ┣ 📜[category_service.go](https://github.com/rickyfazaa/tokobelanja-kelompok7/blob/master/service/category_service.go)<br>
 ┃ ┣ 📜[product_service.go](https://github.com/rickyfazaa/tokobelanja-kelompok7/blob/master/service/product_service.go)<br>
 ┃ ┣ 📜[transaction_service.go](https://github.com/rickyfazaa/tokobelanja-kelompok7/blob/master/service/transaction_service.go)<br>
 ┃ ┗ 📜[user_service.go](https://github.com/rickyfazaa/tokobelanja-kelompok7/blob/master/service/user_service.go)<br>
 ┣ 📜[.env](https://github.com/rickyfazaa/tokobelanja-kelompok7/blob/master/.env)<br>
 ┣ 📜[Hacktiv8-TokoBelanja-Kelompok7.postman_collection.json](https://documenter.getpostman.com/view/23401248/2s8YzQX4Wy)<br>
 ┣ 📜[README.md](https://github.com/rickyfazaa/tokobelanja-kelompok7/blob/master/README.md)<br>
 ┣ 📜[go.mod](https://github.com/rickyfazaa/tokobelanja-kelompok7/blob/master/go.mod)<br>
 ┣ 📜[go.sum](https://github.com/rickyfazaa/tokobelanja-kelompok7/blob/master/go.sum)<br>
 ┣ 📜[main.go](https://github.com/rickyfazaa/tokobelanja-kelompok7/blob/master/main.go)<br>
 ┗ 📜[tokobelanja_environments.postman_environment.json](https://github.com/rickyfazaa/tokobelanja-kelompok7/blob/master/tokobelanja_environments.postman_environment.json)<br>

## Postman Documentation Publish Version 🚀
**`LINK`** https://documenter.getpostman.com/view/23401248/2s8YzQX4Wy

# Endpoint
## 1. User
### Create Admin Account
> <i>digunakan untuk membuat akun dengan role Admin.</i>
- Method: **`POST`**
- Endpoint:
```
http://localhost:8080/users/admin
```
- Request:
    - Request Body:
    ```json
    {
        "full_name": "string",
        "email": "string",
        "password": "string"
    }
    ```
- Response Body:
  - Status: 201,
  - Message: "created",
  - Body:
  ```json
  {
    "data": {
        "id": "integer",
        "full_name": "string",
        "email": "string",
        "password": "string",
        "balance": "integer",
        "created_at": "date"
  	}
  }
  ```
<p align="justify">Notes: Kalian diminta untuk melakukan seeding data admin terlebih dahulu. Jumlah admin nya cukup satu saja.</p>

### Create User Account
> <i>digunakan untuk membuat akun dengan role User.</i>
- Method: **`POST`**
- Endpoint:
```
http://localhost:8080/users/register
```
- Request:
    - Request Body:
    ```json
    {
        "full_name": "string",
        "email": "string",
        "password": "string"
    }
    ```
- Response Body:
  - Status: 201,
  - Message: "created",
  - Body:
  ```json
  {
    "data": {
        "id": "integer",
        "full_name": "string",
        "email": "string",
        "password": "string",
        "balance": "integer",
        "created_at": "date"
  	}
  }
  ```
<p align="justify">Notes: Untuk endpoint ini, role dari data user akan otomatis menjadi customer. Boleh langsung diharcode di controllernya sebelum disimpan ke dalam database. Untuk balance juga otomatis akan dimulai dari angka 0.</p>

### Login Account
> <i>digunakan untuk melakukan login atau autentikasi Member/Admin.</i>
- Method: **`POST`**
- Endpoint:
```
http://localhost:8080/users/login
```
- Request:
    - Request Body:
    ```json
    {
        "email": "string",
        "password": "string"
    }
    ```
- Response Body:
  - Status: 200,
  - Message: "ok",
  - Body:
  ```json
  {
    "data": {
        "token": "jwt string"
  	}
  }
  ```
<p align="justify">Notes: Pada endpoint ini, wajib melakukan logika user login yang dimana harus melakukan pengecekan email dan password user. Pengecekan password wajib dilakukan dengan bantuan library/package <b>Bcrypt</b>.</p>

### Update Account Data
> <i>digunakan untuk melakukan penambahan saldo / balance.</i>
- Method: **`PATCH`**
- Endpoint:
```
http://localhost:8080/users/topup
```
- Request:
    - Headers: Authorization (Bearer Token)
    - Request Body:
    ```json
    {
        "balance": "integer"
    }
    ```
- Response Body:
  - Status: 200,
  - Message: "ok",
  - Body:
  ```json
  {
    "data": {
        "message": "Your balance has been successfully updated to Rp <current balance>"
  	}
  }
  ```
<p align="justify">Notes: Endpoint ini memerlukan proses autentikasi. Proses autentikasi wajib dilakukan dengan package/library JsonWebToken. Endpoint ini berguna untuk menambahkan balance dari user.Pastikan untuk menjumlahkan balance yang di input dengan balance yang dimiliki oleh user tersebut baru kemudian di update.</p>


## 2. Categories <i>(Must an admin)</i>
<p align="justify">Notes: Seluruh endpoint untuk mengakses endpoint categories memerlukan proses autentikasi menggunakan package JsonWebToken dan memerlukan proses autorisasi. Autorisasi diperlukan karena yang boleh mengakses endpoint categories adalah user dengan role admin.</p>

### Create Categories Type <i>(must an admin)</i>
> <i>digunakan untuk membuat tipe kategori.</i>
- Method: **`POST`**
- Endpoint:
```
http://localhost:8080/categories
```
- Request:
    - Headers: Authorization (Bearer Token)
    - Request Body:
    ```json
    {
        "type": "string"
    }
    ```
- Response Body:
  - Status: 201,
  - Message: "created",
  - Body:
  ```json
  {
    "data": {
        "id": "integer",
        "type": "string",
        "sold_product_amount": "integer",
        "created_at": "date"
  	}
  }
  ```
<p align="justify"><b>Notes: sold_product_amount pada awalnya akan otomatis menjadi 0. Nantinya sold_product_amount akan ditambahkan ketika ada pembelian product oleh user. Nanti akan dibahas pada saat pembahasan endpoint products.</b></p>

### Show All Categories Type <i>(must an admin)</i>
> <i>digunakan untuk menampilkan semua tipe kategori.</i>
- Method: **`GET`**
- Endpoint:
```
http://localhost:8080/categories
```
- Request:
    - Headers: Authorization (Bearer Token)
- Response Body:
  - Status: 200,
  - Message: "ok",
  - Body:
  ```json
  {
    "data": {
        "id": "integer",
        "type": "string",
        "sold_product_amount": "integer",
        "updated_at": "date",
        "created_at": "date",
        "Product": [
                {
            "id": "integer",
            "title": "string",
            "price": "integer",
            "stock": "integer",
            "created_at": "date",
            "updated_at": "date"
                },
                {
            "id": "integer",
            "title": "string",
            "price": "integer",
            "stock": "integer",
            "created_at": "date",
            "updated_at": "date"
                }
            ]
        }
  }
  ```

### Update Categories Type <i>(must an admin)</i>
> <i>digunakan untuk melakukan perubahan tipe kategori.</i>
- Method: **`PATCH`**
- Endpoint:
```
http://localhost:8080/categories/:categoryId
```
- Request:
    - Headers: Authorization (Bearer Token)
    - Params: categoryId (integer)
    - Request Body:
    ```json
    {
        "type": "string"
    }
    ```
- Response Body:
  - Status: 200,
  - Message: "ok",
  - Body:
  ```json
  {
    "data": {
        "id": "integer",
        "type": "string",
        "sold_product_amount": "integer",
        "updated_at": "date"
  	}
  }
  ```

### Delete Categories Type <i>(must an admin)</i>
> <i>digunakan untuk melakukan penghapusan tipe kategori.</i>
- Method: **`DELETE`**
- Endpoint:
```
http://localhost:8080/categories/:categoryId
```
- Request:
    - Headers: Authorization (Bearer Token)
    - Params: categoryId (integer)
- Response Body:
  - Status: 200,
  - Message: "ok",
  - Body:
  ```json
  {
    "data": {
        "message": "Category has been successfully deleted"
  	}
  }
  ```


## 3. Products
<p align="justify">Notes: Seluruh endpoint untuk mengakses endpoint products memerlukan proses autentikasi menggunakan package JsonWebToken. Lalu untuk endpoint selain GET/products memerlukan proses autorisasi.  Autorisasi diperlukan karena yang boleh mengakses endpoint selain GET/products adalah user dengan role admin.</p>

### Create a Product <i>(must an admin)</i>
> <i>digunakan untuk membuat produk.</i>
- Method: **`POST`**
- Endpoint:
```
http://localhost:8080/products
```
- Request:
    - Headers: Authorization (Bearer Token)
    - Request Body:
    ```json
    {
        "title": "string",
        "price": "integer",
        "stock": "integer",
        "category_id": "integer"
    }
    ```
- Response Body:
  - Status: 201,
  - Message: "created",
  - Body:
  ```json
  {
    "data": {
        "id": "integer",
        "title": "string",
        "price": "integer",
        "stock": "integer",
        "category_Id": "integer",
        "created_at": "date"
  	}
  }
  ```
<p align="justify"><b>Notes: Pada endpoint ini, harus dilakukan pengecekkan jika data category dengan id yang diberikan pada request body dengan field categoryId ada atau tidak pada database. Jika ada maka boleh disimpan ke database namun jika tidak ada maka harus melempar error.</b></p>

### Show All Products
> <i>digunakan untuk menampilkan semua produk.</i>
- Method: **`GET`**
- Endpoint:
```
http://localhost:8080/products
```
- Request:
    - Headers: Authorization (Bearer Token)
- Response Body:
  - Status: 200,
  - Message: "ok",
  - Body:
  ```json
  {
    "data": {
            [
                {
            "id": "integer",
            "title": "string",
            "price": "integer",
            "stock": "integer",
            "category_id": "integer",
            "created_at": "date"
                },
                {
            "id": "integer",
            "title": "string",
            "price": "integer",
            "stock": "integer",
            "category_id": "integer",
            "created_at": "date"
                }
            ]
        }
  }
  ```

### Update or Make a New Product <i>(must an admin)</i>
> <i>digunakan untuk melakukan perubahan atau membuat baru suatu produk.</i>
- Method: **`PUT`**
- Endpoint:
```
http://localhost:8080/products/:productId
```
- Request:
    - Headers: Authorization (Bearer Token)
    - Params: productId (integer)
    - Request Body:
    ```json
    {
        "title": "string",
        "price": "integer",
        "stock": "integer",
        "category_id": "integer"
    }
    ```
- Response Body:
  - Status: 200,
  - Message: "ok",
  - Body:
  ```json
  {
    "data": {
        "id": "integer",
        "title": "string",
        "price": "integer",
        "stock": "integer",
        "category_Id": "integer",
        "created_at": "date",
        "updated_at": "date"
  	}
  }
  ```
<p align="justify"><b>Pada endpoint ini, harus dilakukan pengecekkan jika data category dengan id yang diberikan pada request body dengan field categoryId ada atau tidak pada database. Jika ada maka boleh disimpan ke database namun jika tidak ada maka harus melempar error.</b></p>

### Delete Product <i>(must an admin)</i>
> <i>digunakan untuk melakukan penghapusan suatu produk.</i>
- Method: **`DELETE`**
- Endpoint:
```
http://localhost:8080/products/:productId
```
- Request:
    - Headers: Authorization (Bearer Token)
    - Params: productId (integer)
- Response Body:
  - Status: 200,
  - Message: "ok",
  - Body:
  ```json
  {
    "data": {
        "message": "Product has been successfully deleted"
  	}
  }
  ```


## 4. TransactionsHisotries
<p align="justify">Notes: Seluruh endpoint untuk mengakses endpoint transactions memerlukan proses autentikasi menggunakan package JsonWebToken.</p>

### Buy Some Product
> <i>digunakan untuk user membeli suatu produk yang sudah terdaftar.</i>
- Method: **`POST`**
- Endpoint:
```
http://localhost:8080/transactions
```
- Request:
    - Headers: Authorization (Bearer Token)
    - Request Body:
    ```json
    {
        "product_id": "integer",
        "quantity": "integer"
    }
    ```
- Response Body:
  - Status: 201,
  - Message: "created",
  - Body:
  ```json
  {
    "data": {
        "message": "You have succesfully purchased the product",
        "transaction_bill": {
            "total_price": "integer",
            "quantity": "integer",
            "product_title": "string"
        }
    }
  }
  ```
<p align="justify">Notes: Endpoint ini digunakan untuk pembelian barang oleh user. Berikut merupakan proses yang harus dilakukan:

- Pengecekan data product jika data product yang dibeli ada atau tidak. Jika ada maka proses dapat dilanjut dan jika tidak ada maka harus melempar error.

- Pengecekan stock product. Jika quantity tidak melebihi stock product maka proses dapat dilanjutkan, jika tidak maka harus melempar error.

- Pengecekan balance user. Jika user yang membeli barang mempunyai balance yang cukup maka proses dapat dilanjut, jika tidak cukup maka harus melempar error.

- Setelah seluruh pengecekkan selesai maka field stock dari product harus dikurangi dengan quantity yang dibeli, dan balance dari user juga harus di kurangi. Setelah itu field sold_product_amount pada category harus ditambahkan sesuai dengan quantity product yang dibeli. 

- Setelah point-point diatas sudah dilakukan baru bisa dilanjutkan untuk membuat data transaction history nya.</p>

### My Transactions History
> <i>digunakan untuk menampilkan history transaksi yang pernah dilakukan oleh User itu sendiri (yang sedang login).</i>
- Method: **`GET`**
- Endpoint:
```
http://localhost:8080/transactions/my-transactions
```
- Request:
    - Headers: Authorization (Bearer Token)
- Response Body:
  - Status: 200,
  - Message: "ok",
  - Body:
  ```json
  {
    "data": {
      [
            {
            "id": "integer",
            "product_id": "integer",
            "user_id": "integer",
            "quantity": "integer",
            "total_price": "integer",
            "Product": {
                "id": "integer",
                "title": "string",
                "price": "integer",
                "stock": "integer",
                "category_id": "integer",
                "created_at": "date",
                "updated_at": "date"
                      }
                }
            ]
  	  }
  }
  ```
<p align="justify">Notes: Endpoint ini  berguna untuk user yang sedang login mendapatkan data-data transaksi pembeliannya.</p>

### Show All User Transactions History <i>(must an admin)</i>
> <i>digunakan untuk Admin melihat seluruh data transaksi pembelian product oleh user.</i>
- Method: **`GET`**
- Endpoint:
```
http://localhost:8080/transactions/user-transactions
```
- Request:
    - Headers: Authorization (Bearer Token)
- Response Body:
  - Status: 200,
  - Message: "ok",
  - Body:
  ```json
  {
    "data": {
          [
              {
            "id": "integer",
            "product_id": "integer",
            "user_id": "integer",
            "quantity": "integer",
            "total_price": "integer",
            "Product": {
                "id": "integer",
                "title": "string",
                "price": "integer",
                "stock": "integer",
                "category_id": "integer",
                "created_at": "date",
                "updated_at": "date"
                    },
            "User": {
                "id": "integer",
                "email": "tstring",
                "full_name": "string",
                "balance": "integer",
                "created_at": "date",
                "updated_at": "date"
                    }
                }
            ]
  	  }
  }
  ```
<p align="justify">Notes: Endpoint ini berguna untuk admin melihat seluruh data transaksi pembelian product oleh user. Endpoint ini memerlukan proses autorisasi karena hanya admin yang boleh mengakses endpoint ini.</p>
<br>

# Group 7
<pre style="font-family:verdana">

<p>1. <a href="https://github.com/alrico11"><b>Alrico Rizki Wibowo</b></a>&emsp; —&emsp;GLNG-KS04-017</p>
<p>2. <a href="https://github.com/rickyfazaa"><b>Ricky Khairul Faza</b></a>&emsp; —&emsp;GLNG-KS04-022</p>
<p>3. <a href="https://github.com/mrafid01"><b>Muhammad Rafid</b></a>&emsp; —&emsp;GLNG-KS04-024</p></pre>

## Pembagian Tugas
### Alrico Rizki Wibowo
Alrico Rizki Wibowo mengerjakan beberapa hal berikut :
- ``Endpoint``	: POST /categories
- ``Endpoint``	: GET /categories
- ``Endpoint``	: PATCH /categories/:categoryId
- ``Endpoint``	: DELETE /categories/:categoryId
- ``Endpoint``	: GET /transactions/user-transactions
- ``Additional``	: Unit Test, README.md, dan .env

### Muhammad Rafid
Muhammad Rafid mengerjakan beberapa hal berikut :
- ``Endpoint``	: POST /users/admin
- ``Endpoint``	: POST /products
- ``Endpoint``	: GET /products
- ``Endpoint``	: PUT /products/:productId
- ``Endpoint``	: DELETE /products/:productId
- ``Endpoint``	: POST /transactions
- ``Helper`` 	: Generate & Validate Token, Validator, dan Middleware

### Ricky Khairul Faza
Ricky Khairul Faza mengerjakan beberapa hal berikut :
- ``Endpoint``	: POST /users/register
- ``Endpoint``	: POST /users/login
- ``Endpoint``	: PATCH /users/topup
- ``Endpoint``	: GET /transactions/my-transactions
- ``Helper``	: Generate Password, Verify Password, Fix Bugs
- ``Additional``	: Postman Collection, README.md, dan Deploy API
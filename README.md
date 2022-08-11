
# API Mini POS - Golang

API Mini POS ini terdiri dari 6 Resource : 
- Users
- Outlets
- Categories
- Products
- Outlet Products
- Transactions


## 1. Set Up Project

 - Clone this repository
 - Pindahkan file `.env.example` ke dalam folder cmd
 - Sesuaikan konfigurasi yang terdapat pada file `.env.example`
 - Rename file `.env.example` menjadi `.env`
 - Jalankan perintah
    ```
    go mod tidy
    ```

## 2. Run Project

 - Go to root folder
 - Jalankan perintah
    ```
    go run ./cmd/main.go
    ```
- Server akan berjalan pada `http://localhost:8000`

## 3. ERD Mini POS 

```mermaid
erDiagram
    USER ||--|{ OUTLET : has
    USER {
        varchar id PK
        varchar name
        varchar email
        varchar password
        enum role "Fill with 'owner' or 'staff'"
        varchar outlet_id "id outlet for staff"
        datetime created_at
        datetime updated_at
        datetime deleted_at
        string created_by
        string updated_by
        string deleted_by
    }
    OUTLET ||--|{ OUTLET_PRODUCT : has
    OUTLET {
        varchar id PK
        varchar name
        bigint phone
        varchar address
        datetime created_at
        datetime updated_at
        datetime deleted_at
        string created_by
        string updated_by
        string deleted_by
    }
    CATEGORY ||--o{ PRODUCT : has
    CATEGORY {
        varchar id PK
        varchar name
        varchar description
        datetime created_at
        datetime updated_at
        datetime deleted_at
        string created_by
        string updated_by
        string deleted_by
    }
    PRODUCT ||--|{ OUTLET_PRODUCT : has
    PRODUCT {
        varchar id PK
        varchar name
        varchar bigint
        varchar category_id FK
        varchar description
        datetime created_at
        datetime updated_at
        datetime deleted_at
        string created_by
        string updated_by
        string deleted_by
    }
    OUTLET_PRODUCT
    OUTLET_PRODUCT {
        varchar id PK
        varchar outlet_id FK
        varchar product_id FK
        double price
        int stock
        datetime created_at
        datetime updated_at
        datetime deleted_at
        string created_by
        string updated_by
        string deleted_by
    }
```

## 4. Architecture

### 4.1 Main Layer
```mermaid
flowchart LR
    A[CMD/main.go] --> 
    B(ROUTES) --> 
    C(HANDLERS) --> 
    D(SERVICES) --> 
    E(REPOSITORY)
```

- **CMD/main.go:** merupakan file utama yang dieksekusi untuk menjalankan golang
- **ROUTES:** layer yang bertugas untuk mengatur / mapping endpoint
- **HANDLERS:** layer yang bertugas untuk menerima input (request data) dan memberikan output (response data)
- **SERVICES:** layer yang bertugas untuk mengolah bisnis logic berdasarkan request data yang diterima dari layer handler
- **REPOSITORIES:** layer yang bertugas untuk mengolah data dari dan ke dalam database

### 4.2 Support Layer

- **HELPERS:** Kumpulan function global yang bisa dipanggil dari semua Main Layer
- **MIDDLEWARES:** Kumpulan function yang dieksekusi sebelum menjalankan layer Handlers
- **PKG:** Kumpulan function sistem yang tidak digunakan pada Main Layer
- **PORTS:** Kumpulan interface untuk abstraksi manipulasi data pada database
- **MODELS:** Kumpulan struct yang merepresentasikan struktur table pada database
- **SCHEMAS:** Kumpulan struct yang merepresentasikan request payload dan response payload



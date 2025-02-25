# README

## Giới thiệu
Đây là một dự án được viết bằng Golang, sử dụng phiên bản **Go 1.23** và được triển khai trên môi trường **Render**.

## Yêu cầu hệ thống
- **Go 1.23** hoặc mới hơn
- Công cụ quản lý mã nguồn **Git**
- **Tài khoản Render** để triển khai ứng dụng

## Cài đặt và chạy local
### 1. Clone repository
```bash
git clone https://github.com/sonpnts/Todo
cd Todo
```

### 2. Cài đặt các dependencies
```bash
go mod tidy
```

### 3. Chạy ứng dụng
```bash
go run main.go
```

Ứng dụng sẽ chạy trên cổng mặc định (ví dụ: `http://localhost:8080`).

## Triển khai trên Render
### 1. Tạo tài khoản và repository
- Truy cập **[Render](https://render.com/)**
- Đăng nhập và liên kết với repository trên GitHub

### 2. Tạo dịch vụ web trên Render
1. Chọn **New Web Service**
2. Kết nối với repository
3. Chọn **Environment** là `Go` và nhập phiên bản `Go 1.23`
4. Cấu hình command build:
   ```bash
   go build -o app .
   ```
5. Cấu hình command chạy:
   ```bash
   ./app
   ```
6. Nhấn **Create Web Service** để hoàn tất

### 3. Kiểm tra trạng thái
- Sau khi deploy thành công, bạn có thể truy cập ứng dụng qua domain được Render cung cấp.

## Cấu trúc thư mục
```
/To-do
│── main.go
│── go.mod
│── go.sum
│── configs/
│── handlers/
│── models/
│── middleware/
│── responsitory/
│── services/
│── README.md
```

---
**Liên hệ:** Nếu có bất kỳ câu hỏi nào, vui lòng liên hệ qua email hoặc mở issue trên GitHub.

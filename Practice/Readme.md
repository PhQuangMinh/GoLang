# API CRUD và xuất file
## 1. CRUD danh sách các cuộc gọi, database tự thiết kế, thực thể cuộc gọi bao gồm các trường (có thể thêm vài trường khác nữa):
    - ID
    - Số điện thoại
    - Thông tin metadata của cuộc gọi (trữ dưới dạng json)
    - Kết quả cuộc gọi (INIT, QUEUEING, SUCCESS, FAIL, NOT_ANSWERED, CANT_CONNECT...)
    - Thời gian tạo (create_at)
    - Thời gian cập nhật (update_at)
    - Thời gian gọi
    - Thời gian nhận kết quả cuộc gọi
    - Kết quả cuộc gọi 
    - Thời gian nhấc máy (nullable)
    - Thời gian cúp máy (nullable)

* Note: Tất cả thời gian trong đây đều lưu trữ và truyền đi dưới dạng milisecond timestamp, không dùng text

### 1.1. Yêu cầu với API tạo (POST JSON)
- Tạo cuộc gọi, ghi xuống database, đẩy thông tin cuộc gọi vào một queue trong rabbitmq, tất cả thông tin đều đựng trong json request lên, ko dùng multipath/form.

- Tạo một consumer khác có nhiệm vụ nhận thông tin cuộc gọi từ queue bên trên, consumer này nhận, sleep khoảng 2-5s, sau đó đẩy kết quả cuộc gọi vào một queue gọi là queue kết quả (consumer này fake luồng xử lý cuộc gọi thật, và trả về thông tin giả).

- Tạo một consumer khác nữa có nhiệm vụ nhận thông tin kết quả từ queue bên trên, thực hiện update bản ghi cuộc gọi theo id nếu nhận được kết quả.

**Note:** Các API, consumer này đều code chung 1 source code

### 1.2. Yêu cầu với API List
- Cần filter được theo tên khách hàng, theo số điện thoại, theo thời gian tạo trong khoảng (ví dụ từ 0h ngày 1 tháng 7 đến 10h ngày 1 tháng 7, thời gian này cho phép truyền vào khi gọi API list với start_at, end_at), API List được viết bằng phuong thức GET và tất cả tham số được truyền vào = query params, đều có thể nullable.

- Thêm một trường là "metadata_display_field" cho phép truyền vào tên một trường cụ thể muốn lấy ra từ trường metadata, ví dụ: trong db có trường metadata = {"key1": "value1", "key2": "value2"}, nếu không có metadata_display_field, dữ liệu trả về là {"key1": "value1", "key2": "value2"}, tức là full dữ liệu metadata, nếu metadata_display_field=key1, dữ liệu metadata trả về sẽ là {"key1": "value1"}, nếu metadata_display_field=key3, dữ liệu metadata trả về sẽ là {} tức là rỗng.

### 1.3. Yêu cầu với API update, delete, getOne
- Đều định danh theo ID bản ghi, tức là update, delete, getOne sẽ bắt buộc phải truyền ID vào để xác định chính xác bản ghi cần xóa, sửa, ...

## 2. Yêu cầu về tech stack
### 2.1. Sử dụng MySQL làm database, rabbitmq làm queue

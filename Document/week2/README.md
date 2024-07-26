# RabbitMQ và MySQL

## I. RabbitMQ
### 1. Tổng quan
- **RabbitMQ** là một **message broker mã nguồn mở**, được sử dụng rộng rãi để truyền tải và quản lý tin nhắn giữa các ứng dụng khác nhau trong một hệ thống phân tán.
- **RabbitMQ** đóng vai trò như một trung gian, nhận các tin nhắn từ **một ứng dụng (producer)** và chuyển tiếp chúng đến **một hoặc nhiều ứng dụng khác (consumer)** theo một cách thức được định nghĩa trước.
- Ví dụ để dễ hình dung:
    - **Producer:** Người gửi thư (ứng dụng tạo ra tin nhắn).
    - **RabbitMQ:** Bưu điện (trung tâm xử lý và phân phối thư).
    - **Queue:** Hộp thư (nơi lưu trữ tạm thời các tin nhắn).
    - **Consumer:** Người nhận thư (ứng dụng tiêu thụ tin nhắn).
### 2. Lý do sử dụng RabbitMQ
- **Khả năng mở rộng:** RabbitMQ có thể xử lý lượng lớn tin nhắn và dễ dàng mở rộng để đáp ứng nhu cầu tăng trưởng của hệ thống.
- **Đảm bảo tin cậy:** Tin nhắn sẽ được lưu trữ cho đến khi được tiêu thụ thành công, **đảm bảo không có tin nhắn nào bị mất.**
- **Linh hoạt:** RabbitMQ hỗ trợ nhiều giao thức nhắn tin khác nhau, cho phép tích hợp với nhiều hệ thống khác nhau.
- **Dễ sử dụng:** RabbitMQ cung cấp một giao diện quản lý trực quan và các thư viện client cho nhiều ngôn ngữ lập trình khác nhau.
### 3. Một số thành phần chính
- **Exchange:** là một đối tượng nhận các tin nhắn từ **producer** và chuyển tiếp chúng đến các **queue** dựa trên **routing key.**
- **Queue:** Là mọt đối tượng lưu trữ các tin nhắn cho đến khi được tiêu thụ.
- **Binding:** Là một liên kết giữa **exchange và queue**, xác định cách thứ **routing key** sẽ được sử dụng chuyển tiếp tin nhắn.
- **Routing key:** Là một chuỗi ký tự được sử dụng để xác định cách thức một tin nhắn được gửi đến **queue**.
- Virtual host: Một **container ảo** để phân cách các ứng dụng khác nhau sử dụng cùng một **instant RabbitMQ.**
### 4. Một số loại Exchange khác nhau trong RabbitMQ
#### a. Fanout Exchange
- Phân phối tin nhắn đến tất cả các queue đã được bind với nó.
- Giống như phát sóng một tin nhắn đến tất cả người đăng kí.
#### b. Direct Exchange
- Chỉ gửi tin nhắn đến các queue có routing key trùng khớp hoàn toàn với routing key của tin nhắn.
- Thường được sử dụng để gửi tin nhắn đến một queue cụ thể.
#### c. Topic Exchange
- Sẽ routing tin nhắn đến các queue dựa trên pattern matching của routing key.
- Cho phép tạo ra các routing key dạng wildcard để phân phối tin nhắn linh hoạt hơn
> Wildcard là các kí tự hoặc chuỗi ký tự được sử dụng để thay thế cho một hoặc nhiều ký tự khác trong các mẫu tìm kiếm.
#### d. Headers Exchange
- Routing tin nhắn dựa trên các header trong tin nhắn thay vì routing key.
- Cung cấp khả năng routing phức tạp.
#### e. Dead-Letter Exchange
- Một cấu hình đặc biệt để xử lý các tin nhắn không thể được routing hoặc tiêu thụ.
### 5. Json
- **JSON (JavaScript Object Notation)** là một định dạng dữ liệu nhẹ, dễ đọc và viết, được sử dụng rộng rãi để trao đổi dữ liệu giữa máy chủ và ứng dụng web dưới dạng văn bản. 
- Một JSON request thường được sử dụng trong các **giao tiếp HTTP/HTTPS giữa client và server.**
- Ví dụ:

    ```json
    {
        "name": "Minh",
        "age": 30,
        "email": "Minh@example.com",
        "isActive": true,
        "score": 95.5,
        "nickname": null,
        "address": {
            "street": "123 Main St",
            "city": "Anytown",
            "state": "CA",
            "postalCode": "12345"
        },
        "hobbies": ["reading", "gaming", "coding"],
        "preferences": {
            "newsletter": true,
            "notifications": false
        }
    }
    ```
## II. MySQL
### 1. Kiểu dữ liệu
#### a. Kiểu dữ liệu số(Numeric Data Types)
- Số nguyên
    - **TINYINT:** Số nguyên nhỏ, từ -128 đến 127.
    - **SMALLINT:** Số nguyên nhỏ hơn, từ -32768 đến 32767.
    - **MEDIUMINT:** Số nguyên trung bình.
    - **INT:** Số nguyên tiêu chuẩn, từ -2^31-1 đến 2^31.
    - **BIGINT:** Số nguyên lớn.
- Số thực:
    - **FLOAT:** Số thực có độ chính xác hơn.
    - **DOUBLE:** Số thực độ chính xác kép.
    - **DECIMAL(M, D):** Số thập phân với M là tổng số chữ số và D là số chữa sô sau dấu phẩy.
#### b. Kiểu dữ liệu chuỗi(String Data Types)
- **CHAR(M):** Chuỗi có độ dài cố định M kí tự.
- **VARCHAR(M):** Chuỗi có độ dài biến đổi, tối đa M ký tự.
- **TEXT:** Chuỗi văn bản có độ dài lớn hơn VARCHAR.
    - **TINYTEXT:** TEXT kích thước nhỏ.
    - **TEXT:** TEXT kích thước trung bình.
    - **MEDIUMTEXT:** TEXT kích thước lớn.
    - **LONGTEXT:** TEXT kích thước rất lớn.
- **BINARY(M):** Chuỗi nhị phân có độ dài cố định M byte
- **VARBINARY(M):** Chuỗi nhị phân có độ dài biến đổi, tối đa M byte.
- **BLOB:** Dữ liệu nhị phân lớn.
    - **TINYBLOB:** BLOB kích thước nhỏ.
    - **BLOB:** BLOB kích thước trung bình.
    - **MEDIUMBLOB:** BLOB kích thước lớn.
    - **LONGBLOB:** BLOB kích thước rất lớn.
- **ENUM:** Kiểu liệt kê, chỉ cho phép các giá trị xác định trước.
- **SET:** Tập hợp các giá trị, mỗi giá trị có thể được chọn hoặc không.
#### c. Kiểu dữ liệu ngày và thời gian(Date and Time Data Types)
- **DATE:** Ngày(YYYY-MM-DD).
- **TIME:** Thời gian(HH:MM:SS)
- **DATETIME:** Ngày và thời gian kết hợp(YYYY-MM-DD HH:MM:SS)
- **TIMESTAMP:** Dấu thời gian, thường được tự động cập nhật khi mọt hàng được chèn hoặc cập nhật.
- **YEAR:** Năm(YYYY).
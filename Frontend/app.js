const apiUrlLogin = 'http://localhost:8080/v2/users/login'
const apiUrlLogout = 'http://localhost:8080/v2/users/logout'
const apiUrlGetUser = 'http://localhost:8080/v2/users'
const apiUrl = 'http://localhost:8080/v2/users'
const apiUpdatePassword = 'http://localhost:8080/v2/users/password'
let accessToken = ''
const idUser = ''


document.getElementById("loginForm").addEventListener("submit", function (event) {
    event.preventDefault(); // ngăn chặn việc gửi form khi người dùng ấn nút

    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;

    fetch(apiUrlLogin, {
        method: "POST",
        headers: {
            "content-Type": "application/json",
        },
        body: JSON.stringify({
            email: email,
            password: password
        })
    })
        .then(response => { //Trả về một cái reponse gì đó
            if (!response.ok) {
                alert("Your email or password is not correct!");
                console.log(response.status)
                throw new Error('Network response was not ok');
            }
            else {
                if (response.status == 200) {
                    alert("Login successful!");
                    document.getElementById("loginSection").style.display = "none";
                    document.getElementById("navButtons").style.display = "block";
                    return response.json();
                }
            }
        })
        .then(data => { // lấy reponse từ data rồi in ra
            console.log(data);
            accessToken = data.access_token
            idUser = data.id_user
        })
        .catch(error => { //Nếu có lỗi thì in ra lỗi
            console.error('Error:', error);
        });
});

document.getElementById("logoutButton").addEventListener("click", function () {
    if (!accessToken) {
        console.error('No access token available for logout');
        return;
    }
    fetch(apiUrlLogout, {
        method: "POST",
        headers: {
            "content-Type": "application/json",
            'Authorization': `Bearer ${accessToken}`
        }
    })
        .then(response => { //Trả về một cái reponse gì đó
            if (!response.ok) {
                console.log(response.status)
                throw new Error('Network response was not ok');
            }
            else {
                if (response.status == 200) {
                    alert("Logged out!");
                    document.getElementById("loginSection").style.display = "block";
                    document.getElementById("navButtons").style.display = "none";
                    return response.json();
                }
            }
        })
        .then(data => { // lấy reponse từ data rồi in ra
            console.log(data);
        })
        .catch(error => { //Nếu có lỗi thì in ra lỗi
            console.error('Error:', error);
        });
});

document.getElementById("showGetUser").addEventListener("click", function () {
    if (!accessToken) {
        console.error('No access token available for logout');
        return;
    }
    document.getElementById("loginSection").style.display = "none";
    document.getElementById("navButtons").style.display = "none";
    document.getElementById("getUserSection").style.display = "block";
});

document.getElementById("getUserForm").addEventListener("submit", function (event) {
    event.preventDefault(); // Ngăn chặn form gửi đi theo cách thông thường

    const userId = document.getElementById("userId").value;

    if (!accessToken) {
        alert("You must be logged in to get user information.");
        return;
    }
    fetch(`${apiUrlGetUser}/${userId}`, {
        method: "GET",
        headers: {
            "content-Type": "application/json",
            "Authorization": `Bearer ${accessToken}`
        }
    })
    .then(response => {
        if (!response.ok) {
            console.log(response.status);
            throw new Error('unauthorized to access this resource');
        } else {
            return response.json();
        }
    })
    .then(data => {
        if (typeof data.error === 'undefined') {
            // Hiển thị thông tin người dùng trong phần #userInfo
            const userInfoDiv = document.getElementById("userInfo");
            userInfoDiv.innerHTML = `
                <h4>User Info</h4>
                <p>First Name: ${data.first_name}</p>
                <p>Last Name: ${data.last_name}</p>
                <p>UserName: ${data.email}</p>
                <p>Phone Number: ${data.phone_number}</p>
                <p>Email: ${data.email}</p>
                <p>UserType: ${data.user_type}</p>
            `;
        } else{
            alert(data.error)
        }
    })
    .catch(error => {
        console.error('Error:', error);
        alert(error);
    });
});

document.getElementById("showGetAllUsers").addEventListener("click", function () {
    if (!accessToken) {
        console.error('No access token available for logout');
        return;
    }
    document.getElementById("loginSection").style.display = "none";
    document.getElementById("navButtons").style.display = "none";
    document.getElementById("getAllUsersSection").style.display = "block";
});

document.getElementById("getAllUsersButton").addEventListener("click", function () {
    if (!accessToken) {
        alert("You must be logged in to view all users.");
        return;
    }

    fetch(apiUrlGetUser, {
        method: "GET",
        headers: {
            "content-Type": "application/json",
            "Authorization": `Bearer ${accessToken}`
        }
    })
    .then(response => {
        return response.json();
    })
    .then(data => {
        if (typeof data.error === 'undefined') {
            const allUsersInfoDiv = document.getElementById("allUsersInfo");
            allUsersInfoDiv.innerHTML = "<h4>All Users</h4>";
     
            const users = data.users;
            if (Array.isArray(users)) {
                users.forEach(user => {
                    allUsersInfoDiv.innerHTML += `
                        <div class="user">
                            <p><strong>First Name:</strong> ${user.first_name || 'N/A'}</p>
                            <p><strong>Last Name:</strong> ${user.last_name || 'N/A'}</p>
                            <p><strong>User Name:</strong> ${user.user_name || 'N/A'}</p>
                            <p><strong>User Type:</strong> ${user.user_type || 'N/A'}</p>
                            <p><strong>Email:</strong> ${user.email || 'N/A'}</p>
                            <p><strong>Phone Number:</strong> ${user.phone_number || 'N/A'}</p>
                            <hr>
                        </div>
                    `;
                });
            } else {
                console.error('Expected users to be an array, but got:', users);
            }
        } else{
            alert(data.error)
        }
    })
    .catch(error => {
        console.error('Error:', error);
        alert("Failed to fetch all users. Please try again.");
    });
});

document.getElementById("showUpdateUser").addEventListener("click", function () {
    if (!accessToken) {
        console.error('No access token available for logout');
        return;
    }
    document.getElementById("loginSection").style.display = "none";
    document.getElementById("navButtons").style.display = "none";
    document.getElementById("updateUserSection").style.display = "block";
});

document.getElementById("updateUserSection").addEventListener("submit", function (event) {
    event.preventDefault(); // Ngăn chặn gửi form mặc định

    const userIdUpdate = document.getElementById("userIdUpdate").value;
    console.log("User: " + userIdUpdate)
    const firstName = document.getElementById("firstName").value;
    const lastName = document.getElementById("lastName").value;
    const email = document.getElementById("emailUpdate").value;
    const phoneNumber = document.getElementById("phoneNumber").value;

    fetch(`${apiUrl}/${userIdUpdate}`, {
        method: "PUT", // Hoặc "PATCH" nếu API yêu cầu
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${accessToken}`
        },
        body: JSON.stringify({
            first_name: firstName,
            last_name: lastName,
            email: email,
            phone_number: phoneNumber
        })
    })
    .then(response => {
        return response.json();
    })
    .then(data => {
        if (typeof data.error === 'undefined') {
            alert("updated user successfully!");
        } else{
            alert(data.error)
        }
        console.log(data);
        // Xử lý thêm nếu cần
    })
    .catch(error => {
        console.error('Error:', error);
        alert(error);
    });
});

document.getElementById("showUpdatePassword").addEventListener("click", function () {
    if (!accessToken) {
        console.error('No access token available for logout');
        return;
    }
    document.getElementById("loginSection").style.display = "none";
    document.getElementById("navButtons").style.display = "none";
    document.getElementById("updatePasswordSection").style.display = "block";
});

document.getElementById("updatePasswordSection").addEventListener("submit", function (event) {
    event.preventDefault(); // Ngăn chặn gửi form mặc định

    const userIdPassword = document.getElementById("userIdPassword").value;
    const oldPassword = document.getElementById("oldPassword").value;
    const newPassword = document.getElementById("newPassword").value;
    let api = `${apiUpdatePassword}/${userIdPassword}`

    fetch(api, {
        method: "PUT", // Hoặc "PATCH" nếu API yêu cầu
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${accessToken}`
        },
        body: JSON.stringify({
            old_password: oldPassword,
            new_password: newPassword,
        })
    })
    .then(response => {
        return response.json();
    })
    .then(data => {
        if (typeof data.error === 'undefined') {
            alert("Password updated successfully!");
        } else{
            alert(data.error)
        }
        console.log(data);
        // Xử lý thêm nếu cần
    })
    .catch(error => {
        console.error('Error:', error);
        alert(error);
    });
});

document.getElementById("backPassword").addEventListener("click", function() {
  
    document.getElementById("updatePasswordSection").style.display = "none";

    document.getElementById("navButtons").style.display = "block";
});

document.getElementById("backGetUser").addEventListener("click", function() {
  
    document.getElementById("getUserSection").style.display = "none";

    document.getElementById("navButtons").style.display = "block";
});

document.getElementById("backGetUsers").addEventListener("click", function() {
  
    document.getElementById("getAllUsersSection").style.display = "none";

    document.getElementById("navButtons").style.display = "block";
});

document.getElementById("backUpdateUser").addEventListener("click", function() {
  
    document.getElementById("updateUserSection").style.display = "none";

    document.getElementById("navButtons").style.display = "block";
});


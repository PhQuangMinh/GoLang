const apiUrlLogin = 'http://localhost:8080/v2/users/login'
const apiUrlLogout = 'http://localhost:8080/v2/users/logout'
const apiUrlGetUserById = 'http://localhost:8080/v2/users/1'
let accessToken = ''


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
        console.log(response.status)
        throw new Error('Network response was not ok');
      }
      else {
        if (response.status == 200) {
          alert("Login successful!");
          document.getElementById("loginForm").style.display = "none";
          document.getElementById("logoutButton").style.display = "block";
          return response.json();
        }
      }
    })
    .then(data => { // lấy reponse từ data rồi in ra
      console.log(data);
      accessToken = data.access_token
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
          document.getElementById("loginForm").style.display = "block";
          document.getElementById("logoutButton").style.display = "none";
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
  
  fetch(apiUrlGetUserById, {
    method: "GET",
    headers: {
      "content-Type": "application/json",
      'Authorization': `Bearer ${token}`
    }
  })
    .then(response => { //Trả về một cái reponse gì đó
      if (!response.ok) {
        console.log(response.status)
        throw new Error('Network response was not ok');
      }
      // console.log(response.json())
      return response.json();
    })
    .then(data => { // lấy reponse từ data rồi in ra
      console.log(data);
    })
    .catch(error => { //Nếu có lỗi thì in ra lỗi
      console.error('Error:', error);
    });
document.addEventListener('DOMContentLoaded', function() {
    const authForm = document.getElementById('authForm');

    authForm.addEventListener('submit', function(event) {
        if (event.submitter.id === 'signupButton') {
            event.preventDefault(); // Prevent the form from submitting normally

            const username = document.getElementById('username').value;
            const password = document.getElementById('password').value;

            // Send the signup request
            fetch('http://localhost:8080/auth/signup/', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    username: username,
                    password: password
                }),
            })
            .then(response => {
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                return response.json();
            })
            .then(data => {
                // Handle the response, e.g., show a success message or redirect the user
                console.log('Signup successful:', data);
                authForm.elements.username.value = '';
                authForm.elements.password.value = '';
            })
            .catch(error => console.error('Error:', error));
        } else if (event.submitter.id === 'loginButton') {
            event.preventDefault(); // Prevent the form from submitting normally

            const username = document.getElementById('username').value;
            const password = document.getElementById('password').value;

            // Send the login request
            fetch('http://localhost:8080/auth/login/', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    username: username,
                    password: password
                }),
            })
            .then(response => {
                if (!response.ok) {
                    console.error('Response status:', response.status)
                    return response.text().then(text => {
                        throw new Error('Network response was not ok: ' + text);
                    });
                }

                const contentType = response.headers.get("content-type");
                if (!contentType || !contentType.includes("application/json")) {
                    // Redirect to the home page
                    window.location.href = "http://localhost:8080/";
                    return; // Exit the promise chain
                }
                return response.json();
            })
            .then(data => {
                console.log('Login successful, token stored.');
                authForm.elements.username.value = '';
                authForm.elements.password.value = '';
            })
            .catch(error => { 
                console.error('Error:', error);
                alert('Login failed: ' + error.message);
            });
        }
    });
});

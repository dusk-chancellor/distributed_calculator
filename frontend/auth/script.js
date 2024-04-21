document.addEventListener('DOMContentLoaded', function() {
    const authForm = document.getElementById('authForm');

    authForm.addEventListener('submit', function(event) {
        if (event.submitter.id === 'signupButton') {
            event.preventDefault(); // Prevent the form from submitting normally

            const username = document.getElementById('username').value;
            const password = document.getElementById('password').value;

            // Send the signup request
            fetch('http://localhost:8080/signup/', {
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
                if (response.headers.get('Content-Type') !== 'application/json') {
                    throw new Error('Response is not JSON');
                }
                return response.json();
            })
            .then(data => {
                // Handle the response, e.g., show a success message or redirect the user
                console.log('Signup successful:', data);
                // Optionally, redirect the user to a different page
                // window.location.href = '/dashboard';
                authForm.elements.username.value = '';
                authForm.elements.password.value = '';
            })
            .catch(error => console.error('Error:', error));
        } else if (event.submitter.id === 'loginButton') {
            event.preventDefault(); // Prevent the form from submitting normally

            const username = document.getElementById('username').value;
            const password = document.getElementById('password').value;

            // Send the login request
            fetch('http://localhost:8080/login/', {
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
                if (response.headers.get('Content-Type') !== 'application/json') {
                    throw new Error('Response is not JSON');
                }
                return response.json();
            })
            .then(data => {
                // Store the token in local storage
                localStorage.setItem('token', data.token);
                console.log('Login successful, token stored.');
                // Optionally, redirect the user to a different page
                // window.location.href = '/dashboard';
                authForm.elements.username.value = '';
                authForm.elements.password.value = '';
            })
            .catch(error => console.error('Error:', error));
        }
    });
});

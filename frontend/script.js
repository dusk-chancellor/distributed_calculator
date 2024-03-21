document.addEventListener('DOMContentLoaded', function() {
    const form = document.getElementById('expressionInput');

    form.addEventListener('submit', function(event) {
        event.preventDefault(); // Prevent the form from submitting normally

        const expression = form.elements.expression.value; // Get the expression from the input

        // Send the expression to the backend
        fetch('http://localhost:8080/expressions', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ expression: expression }),
        })
        .then(response => response.json())
        .then(data => {
            console.log('Success:', data);
            // Handle the response from the server
        })
        .catch((error) => {
            console.error('Error:', error);
        });
    });
});
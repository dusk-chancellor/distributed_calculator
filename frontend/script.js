document.addEventListener('DOMContentLoaded', function() {
    const form = document.getElementById('expressionInput');

    form.addEventListener('submit', function(event) {
        event.preventDefault(); // Prevent the form from submitting normally

        const expression = form.elements.expression.value; // Get the expression from the input

        // Send the expression to the backend
        fetch('http://localhost:8080/expression/', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ expression: expression }),
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
            console.log('Success:', data);
        })
        .catch((error) => {
            console.error('Error:', error);
        });
    });
});

document.addEventListener('DOMContentLoaded', function() {
    fetch('http://localhost:8080/expression/', {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
        },
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.json();
    })
    .then(data => {
        const expressionsDiv = document.getElementById('savedExpressions');
        data.forEach(expression => {
            const p = document.createElement('p');
            p.textContent = `ID: ${expression.id},
                Expression: ${expression.expression},
                Date: ${expression.date},
                Status: ${expression.status}`;
            expressionsDiv.appendChild(p);
        });
    })
    .catch(error => {
        console.error('There has been a problem with your fetch operation:', error);
    });
})
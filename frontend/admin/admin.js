document.getElementById('addAuthorForm').addEventListener('submit', function(event) {
    event.preventDefault();

    const firstName = document.getElementById('authorFirstName').value;
    const lastName = document.getElementById('authorLastName').value;

    fetch('../../backend/admin_add_author.php', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: `authorFirstName=${encodeURIComponent(firstName)}&authorLastName=${encodeURIComponent(lastName)}`,
    })
    .then(response => response.json())
    .then(data => {
        const statusElement = document.getElementById('addAuthorStatus');
        if (data.success) {
            statusElement.textContent = data.message;
            statusElement.style.color = 'green';
        } else {
            statusElement.textContent = data.message;
            statusElement.style.color = 'red';
        }
    })
    .catch(error => {
        console.error('Ошибка:', error);
        const statusElement = document.getElementById('addAuthorStatus');
        statusElement.textContent = 'Произошла ошибка при добавлении автора.';
        statusElement.style.color = 'red';
    });
});

// Добавить обработчик для ссылки '+ Добавить автора'
document.addEventListener('DOMContentLoaded', function() {
    const addAuthorLink = document.getElementById('add-author-link');
    if (addAuthorLink) {
        addAuthorLink.addEventListener('click', function() {
            window.open('addAuthor.html', '_blank');
        });
    }
});
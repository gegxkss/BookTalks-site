document.addEventListener('DOMContentLoaded', function() {
    fetch('/BookTalks-site/backend/check_auth.php', { credentials: 'same-origin' })
        .then(res => res.json())
        .then(data => {
            if (data.authenticated && data.username) {
                const welcome = document.createElement('div');
                welcome.className = 'welcome-user';
                welcome.textContent = `Привет, ${data.username}!`;
                const container = document.querySelector('.container');
                if (container) {
                    container.prepend(welcome);
                } else {
                    document.body.prepend(welcome);
                }
            }
        });
});

// Динамическая загрузка книг на главную страницу
function renderBooks(books, containerSelector) {
    const container = document.querySelector(containerSelector);
    if (!container) return;
    container.innerHTML = '';
    books.forEach(book => {
        const bookDiv = document.createElement('div');
        bookDiv.className = 'book';
        bookDiv.innerHTML = `
            <p class="name">${book.name}<br>${book.first_name ? book.first_name + ' ' : ''}${book.last_name || ''}</p>
            <img src="${book.coverimage_filename ? '/BookTalks-site/frontend/' + book.coverimage_filename : 'styles/images/book_about.png'}" alt="Books" class="book_about">
        `;
        bookDiv.onclick = () => {
            window.location.href = `/BookTalks-site/frontend/addBook.html?id=${book.id}`;
        };
        container.appendChild(bookDiv);
    });
}

function loadBooks() {
    fetch('/BookTalks-site/backend/books_list.php')
        .then(res => res.json())
        .then(data => {
            if (data.books) {
                // Первые 4 книги — "Бестселлеры", следующие 4 — "Новинки"
                renderBooks(data.books.slice(0, 4), '.bestsellers');
                renderBooks(data.books.slice(4, 8), '.new');
            }
        });
}

// Поиск книг по названию
function setupSearch() {
    const input = document.querySelector('.search-input');
    if (!input) return;
    input.addEventListener('input', function() {
        const query = input.value.trim();
        const dropdown = document.getElementById('searchDropdown');
        const results = document.getElementById('searchResults');
        if (!query) {
            dropdown.style.display = 'none';
            results.innerHTML = '';
            return;
        }
        fetch(`/BookTalks-site/backend/books_list.php`)
            .then(res => res.json())
            .then(data => {
                results.innerHTML = '';
                if (data.books && data.books.length > 0) {
                    const filtered = data.books.filter(book => book.name.toLowerCase().includes(query.toLowerCase()));
                    filtered.forEach(book => {
                        const li = document.createElement('li');
                        li.textContent = `${book.name} (${book.first_name ? book.first_name + ' ' : ''}${book.last_name || ''})`;
                        li.onclick = () => {
                            window.location.href = `/BookTalks-site/frontend/addBook.html?id=${book.id}`;
                        };
                        results.appendChild(li);
                    });
                    dropdown.style.display = filtered.length > 0 ? 'block' : 'none';
                } else {
                    dropdown.style.display = 'none';
                }
            });
    });
}

window.addEventListener('DOMContentLoaded', function() {
    loadBooks();
    setupSearch();
});

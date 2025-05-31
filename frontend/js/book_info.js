document.addEventListener('DOMContentLoaded', function() {
    const params = new URLSearchParams(window.location.search);
    const id = params.get('id');
    if (!id) return;

    // Загрузка информации о книге
    fetch(`/BookTalks-site/backend/book_info.php?id=${id}`)
        .then(res => res.json())
        .then(data => {
            if (data.book) {
                document.getElementById('book-title').textContent = data.book.name;
                document.getElementById('book-author').textContent = `${data.book.first_name || ''} ${data.book.last_name || ''}`;
                if (data.book.coverimage_filename) {
                    document.getElementById('book-cover').src = '/BookTalks-site/frontend/' + data.book.coverimage_filename;
                }
                document.getElementById('book-created').textContent = data.book.created_at ? `Добавлено: ${data.book.created_at}` : '';
            }
        });

    // Загрузка цитат
    fetch(`/BookTalks-site/backend/book_quotes.php?book_id=${id}`)
        .then(res => res.json())
        .then(data => {
            const quotesDiv = document.getElementById('book-quotes');
            quotesDiv.innerHTML = '<h3>Цитаты</h3>';
            if (data.quotes && data.quotes.length > 0) {
                data.quotes.forEach(q => {
                    const p = document.createElement('p');
                    p.textContent = q.text;
                    quotesDiv.appendChild(p);
                });
            } else {
                quotesDiv.innerHTML += '<p>Нет цитат</p>';
            }
        });

    // Загрузка рецензий
    fetch(`/BookTalks-site/backend/book_reviews.php?book_id=${id}`)
        .then(res => res.json())
        .then(data => {
            const reviewsDiv = document.getElementById('book-reviews');
            reviewsDiv.innerHTML = '<h3>Рецензии</h3>';
            if (data.reviews && data.reviews.length > 0) {
                data.reviews.forEach(r => {
                    const p = document.createElement('p');
                    p.textContent = r.text;
                    reviewsDiv.appendChild(p);
                });
            } else {
                reviewsDiv.innerHTML += '<p>Нет рецензий</p>';
            }
        });

    // Добавление цитаты
    document.getElementById('add-quote-btn').onclick = function() {
        const text = document.getElementById('quote-text').value.trim();
        if (!text) return;
        fetch('/BookTalks-site/backend/add_quote.php', {
            method: 'POST',
            headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
            body: `book_id=${id}&text=${encodeURIComponent(text)}`
        })
        .then(res => res.json())
        .then(data => {
            if (data.success) {
                location.reload();
            } else {
                alert(data.error || 'Ошибка добавления цитаты');
            }
        });
    };

    // Добавление рецензии
    document.getElementById('add-review-btn').onclick = function() {
        const text = document.getElementById('review-text').value.trim();
        if (!text) return;
        fetch('/BookTalks-site/backend/add_review.php', {
            method: 'POST',
            headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
            body: `book_id=${id}&text=${encodeURIComponent(text)}`
        })
        .then(res => res.json())
        .then(data => {
            if (data.success) {
                location.reload();
            } else {
                alert(data.error || 'Ошибка добавления рецензии');
            }
        });
    };

    // Поиск книг (в шапке)
    const searchInput = document.querySelector('.search-input');
    if (searchInput) {
        searchInput.addEventListener('input', function() {
            const query = searchInput.value.trim();
            const dropdown = document.getElementById('searchDropdown');
            const results = document.getElementById('searchResults');
            if (!query) {
                if (dropdown) dropdown.style.display = 'none';
                if (results) results.innerHTML = '';
                return;
            }
            fetch(`/BookTalks-site/backend/books_search.php?q=${encodeURIComponent(query)}`)
                .then(res => res.json())
                .then(data => {
                    if (results) results.innerHTML = '';
                    if (data.books && data.books.length > 0) {
                        data.books.forEach(book => {
                            const li = document.createElement('li');
                            li.textContent = `${book.name} (${book.first_name ? book.first_name + ' ' : ''}${book.last_name || ''})`;
                            li.onclick = () => {
                                window.location.href = `/BookTalks-site/frontend/book_info.html?id=${book.id}`;
                            };
                            if (results) results.appendChild(li);
                        });
                        if (dropdown) dropdown.style.display = 'block';
                    } else {
                        if (dropdown) dropdown.style.display = 'none';
                    }
                });
        });
    }
});

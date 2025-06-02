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

    // Проверка авторизации
    let isAuthenticated = false;
    fetch('/BookTalks-site/backend/check_auth.php', { credentials: 'same-origin' })
        .then(res => res.json())
        .then(data => {
            isAuthenticated = !!data.authenticated;
            if (!isAuthenticated) {
                document.getElementById('add-review-btn').onclick = showAuthWarning;
                document.getElementById('add-quote-btn').onclick = showAuthWarning;
                document.getElementById('add-rating-btn').onclick = showAuthWarning;
            } else {
                document.getElementById('add-review-btn').onclick = showReviewPopup;
                document.getElementById('add-quote-btn').onclick = showQuotePopup;
                document.getElementById('add-rating-btn').onclick = showRatingPopup;
            }
        });

    function showAuthWarning() {
        showPopup('Только для зарегистрированных пользователей!');
    }

    function showPopup(html) {
        const bg = document.getElementById('popup-bg');
        const popup = document.getElementById('popup');
        popup.innerHTML = html + '<br><button onclick="closePopup()">Закрыть</button>';
        bg.style.display = popup.style.display = 'block';
    }
    function closePopup() {
        document.getElementById('popup-bg').style.display = 'none';
        document.getElementById('popup').style.display = 'none';
    }

    function showReviewPopup() {
        showPopup('<h3>Добавить рецензию</h3><textarea id="review-text-popup" rows="5" style="width:100%"></textarea><br><button onclick="submitReview()">Сохранить</button>');
    }
    function showQuotePopup() {
        showPopup('<h3>Добавить цитату</h3><textarea id="quote-text-popup" rows="3" style="width:100%"></textarea><br><button onclick="submitQuote()">Сохранить</button>');
    }
    function showRatingPopup() {
        showPopup('<h3>Поставить оценку</h3><input id="rating-value" type="number" min="1" max="10" style="width:60px"> <button onclick="submitRating()">Сохранить</button>');
    }

    window.closePopup = closePopup;
    window.submitReview = function() {
        const text = document.getElementById('review-text-popup').value.trim();
        if (!text) return;
        fetch('/BookTalks-site/backend/add_review.php', {
            method: 'POST',
            headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
            body: `book_id=${id}&text=${encodeURIComponent(text)}`
        })
        .then(res => res.json())
        .then(data => {
            if (data.success) location.reload();
            else alert(data.error || 'Ошибка добавления рецензии');
        });
    }
    window.submitQuote = function() {
        const text = document.getElementById('quote-text-popup').value.trim();
        if (!text) return;
        fetch('/BookTalks-site/backend/add_quote.php', {
            method: 'POST',
            headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
            body: `book_id=${id}&text=${encodeURIComponent(text)}`
        })
        .then(res => res.json())
        .then(data => {
            if (data.success) location.reload();
            else alert(data.error || 'Ошибка добавления цитаты');
        });
    }
    window.submitRating = function() {
        const value = document.getElementById('rating-value').value;
        if (!value || value < 1 || value > 10) return alert('Оценка от 1 до 10');
        fetch('/BookTalks-site/backend/add_rating.php', {
            method: 'POST',
            headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
            body: `book_id=${id}&amount=${encodeURIComponent(value)}`
        })
        .then(res => res.json())
        .then(data => {
            if (data.success) location.reload();
            else alert(data.error || 'Ошибка добавления оценки');
        });
    }

    // Загрузка цитат и рецензий с авторами
    fetch(`/BookTalks-site/backend/book_quotes.php?book_id=${id}`)
        .then(res => res.json())
        .then(data => renderCarousel(data.quotes, 'quotes-carousel', 'quote'));
    fetch(`/BookTalks-site/backend/book_reviews.php?book_id=${id}`)
        .then(res => res.json())
        .then(data => renderCarousel(data.reviews, 'reviews-carousel', 'review'));

    // Карусель цитат и рецензий
    function renderCarousel(items, containerId, type) {
        const cont = document.getElementById(containerId);
        if (!items || !items.length) {
            cont.innerHTML = '<p>Нет данных</p>';
            return;
        }
        let html = '<div class="carousel-inner">';
        items.forEach(item => {
            html += `<div class="carousel-item"><div class="carousel-text">${item.text}</div><div class="carousel-author">${item.nickname || 'Аноним'}</div></div>`;
        });
        html += '</div>';
        cont.innerHTML = html;
    }

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

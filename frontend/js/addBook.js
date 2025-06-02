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
            }
        });

    // Проверка авторизации
    let isAuthenticated = false;
    fetch('/BookTalks-site/backend/check_auth.php', { credentials: 'same-origin' })
        .then(res => res.json())
        .then(data => {
            isAuthenticated = !!data.authenticated;
            if (!isAuthenticated) {
                document.querySelector('[data-popup="add-review"]').onclick = showAuthWarning;
                document.querySelector('[data-popup="add-quote"]').onclick = showAuthWarning;
                document.querySelector('[data-popup="rate-book"]').onclick = showAuthWarning;
            } else {
                document.querySelector('[data-popup="add-review"]').onclick = showReviewPopup;
                document.querySelector('[data-popup="add-quote"]').onclick = showQuotePopup;
                document.querySelector('[data-popup="rate-book"]').onclick = showRatingPopup;
            }
        });

    function showAuthWarning() {
        showPopup('Только для зарегистрированных пользователей!');
    }

    function showPopup(html) {
        // Используем встроенные попапы страницы
        let popup, overlay;
        if (html.includes('review-text-popup')) {
            popup = document.getElementById('add-review-popup');
        } else if (html.includes('quote-text-popup')) {
            popup = document.getElementById('add-quote-popup');
        } else if (html.includes('rating-value')) {
            popup = document.getElementById('rate-book-popup');
        } else {
            alert(html); return;
        }
        overlay = document.querySelector('.overlay');
        popup.style.display = 'block';
        overlay.style.display = 'block';
    }
    function closePopup() {
        document.querySelectorAll('.popup-menu').forEach(p => p.style.display = 'none');
        document.querySelector('.overlay').style.display = 'none';
    }

    // Кнопки закрытия попапов
    document.querySelectorAll('.close-popup-button').forEach(btn => {
        btn.addEventListener('click', closePopup);
    });

    function showReviewPopup() {
        document.getElementById('add-review-popup').style.display = 'block';
        document.querySelector('.overlay').style.display = 'block';
    }
    function showQuotePopup() {
        document.getElementById('add-quote-popup').style.display = 'block';
        document.querySelector('.overlay').style.display = 'block';
    }
    function showRatingPopup() {
        document.getElementById('rate-book-popup').style.display = 'block';
        document.querySelector('.overlay').style.display = 'block';
    }

    // Сохранение рецензии
    document.querySelector('#add-review-popup .save').onclick = function() {
        const text = document.querySelector('#add-review-popup textarea').value.trim();
        if (!text) return alert('Введите текст рецензии');
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
    };
    // Сохранение цитаты
    document.querySelector('#add-quote-popup .save').onclick = function() {
        const text = document.querySelector('#add-quote-popup textarea').value.trim();
        if (!text) return alert('Введите текст цитаты');
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
    };
    // Сохранение оценки (звёзды)
    document.querySelector('#rate-book-popup .save').onclick = function() {
        const value = document.querySelector('#rate-book-popup .star.selected')?.getAttribute('data-rating');
        if (!value) return alert('Выберите оценку');
        fetch('/BookTalks-site/backend/add_rating.php', {
            method: 'POST',
            headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
            body: `book_id=${id}&amount=${encodeURIComponent(value)}`
        })
        .then(res => res.json())
        .then(data => {
            if (data.success) {
                closePopup();
                // Обновить рейтинг без перезагрузки
                fetch(`/BookTalks-site/backend/rating.php?book_id=${id}`)
                    .then(res => res.json())
                    .then(data => {
                        const ratingBlock = document.getElementById('book-rating');
                        if (data && typeof data.rating !== 'undefined' && data.rating !== null) {
                            ratingBlock.textContent = `★ ${data.rating}`;
                        } else {
                            ratingBlock.textContent = '';
                        }
                    });
            } else {
                alert(data.error || 'Ошибка добавления оценки');
            }
        });
    };
    // Выбор звёзд (фикс: выделять все до выбранной)
    document.querySelectorAll('#rate-book-popup .star').forEach(star => {
        star.onclick = function() {
            const rating = parseInt(this.getAttribute('data-rating'));
            document.querySelectorAll('#rate-book-popup .star').forEach(s => {
                const sRating = parseInt(s.getAttribute('data-rating'));
                if (sRating <= rating) {
                    s.classList.add('selected');
                } else {
                    s.classList.remove('selected');
                }
            });
        };
    });

    // Загрузка цитат и рецензий с авторами (исправленные пути)
    fetch(`/BookTalks-site/backend/quotes_list.php?book_id=${id}`)
        .then(res => res.json())
        .then(data => {
            console.log('quotes:', data);
            renderCarousel(data.quotes, 'quote-slider', 'quote');
        });
    fetch(`/BookTalks-site/backend/reviews_list.php?book_id=${id}`)
        .then(res => res.json())
        .then(data => {
            console.log('reviews:', data);
            renderCarousel(data.reviews, 'review-slider', 'review');
        });

    // Загрузка и вывод рейтинга книги
    fetch(`/BookTalks-site/backend/rating.php?book_id=${id}`)
        .then(res => res.json())
        .then(data => {
            const ratingBlock = document.getElementById('book-rating');
            if (data && typeof data.rating !== 'undefined' && data.rating !== null) {
                ratingBlock.textContent = `★ ${data.rating}`;
            } else {
                ratingBlock.textContent = '';
            }
        });

    // Карусель цитат и рецензий с подробной отладкой
    function renderCarousel(items, containerId, type) {
        const cont = document.querySelector(`#${containerId} .slider-wrapper`);
        if (!cont) return;
        if (!items || !items.length) {
            cont.innerHTML = '<p>Нет данных</p>';
            return;
        }
        let html = '';
        items.forEach(item => {
            // Если есть user_id и nickname, делаем ссылку и иконку
            let authorHtml = '';
            if (item.user_id && (item.nickname || item.user_nickname)) {
                authorHtml = `<a href="profile.html?id=${item.user_id}" class="carousel-author-link" title="Профиль пользователя"><img src="/BookTalks-site/uploads/profile_${item.user_id}.png" class="profile-icon" style="width:22px;height:22px;border-radius:50%;vertical-align:middle;margin-right:6px;object-fit:cover;" onerror=\"this.src='styles/images/book_about.png'\">${item.nickname || item.user_nickname}</a>`;
            } else {
                authorHtml = `<span class="carousel-author">${item.nickname || item.user_nickname || 'Аноним'}</span>`;
            }
            html += `<div class="slide"><div class="carousel-text">${item.text}</div>${authorHtml}</div>`;
        });
        cont.innerHTML = html;
        // Для отладки
        console.log('renderCarousel', containerId, items, cont.innerHTML);
        // Слайдер: показываем только первый
        const slides = cont.querySelectorAll('.slide');
        let current = 0;
        function showSlide(idx) {
            slides.forEach((s, i) => s.style.display = i === idx ? '' : 'none');
        }
        showSlide(current);
        const prevBtn = document.querySelector(`#${containerId} .slider-button.prev`);
        const nextBtn = document.querySelector(`#${containerId} .slider-button.next`);
        if (prevBtn && nextBtn) {
            prevBtn.onclick = () => {
                current = (current - 1 + slides.length) % slides.length;
                showSlide(current);
            };
            nextBtn.onclick = () => {
                current = (current + 1) % slides.length;
                showSlide(current);
            };
        }
    }
});

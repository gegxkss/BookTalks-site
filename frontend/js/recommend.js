document.addEventListener('DOMContentLoaded', function() {
    // Динамическая загрузка книг для страницы "Что почитать"
    function renderBooks(books) {
        const container = document.querySelector('.top');
        if (!container) return;
        container.innerHTML = '';
        books.forEach(book => {
            const bookDiv = document.createElement('div');
            bookDiv.className = 'book';
            const ratingId = `book-rating-${book.id}`;
            bookDiv.innerHTML = `
                <p class="name">${book.name}<br>${book.first_name ? book.first_name + ' ' : ''}${book.last_name || ''}</p>
                <img src="${book.coverimage_filename ? '/BookTalks-site/frontend/' + book.coverimage_filename : 'styles/images/book_about.png'}" alt="Books" class="book_about">
                <div class="rating" id="${ratingId}">Загрузка...</div>
            `;
            bookDiv.onclick = () => {
                window.location.href = `/BookTalks-site/frontend/addBook.html?id=${book.id}`;
            };
            container.appendChild(bookDiv);
            // Подгружаем рейтинг отдельно
            fetch(`/BookTalks-site/backend/rating.php?book_id=${book.id}`)
                .then(res => res.json())
                .then(data => {
                    const ratingBlock = document.getElementById(ratingId);
                    if (data && typeof data.rating !== 'undefined' && data.rating !== null) {
                        ratingBlock.textContent = `★ ${data.rating}`;
                    } else {
                        ratingBlock.textContent = '5 ★';
                    }
                })
                .catch(() => {
                    const ratingBlock = document.getElementById(ratingId);
                    if (ratingBlock) ratingBlock.textContent = 'Ошибка рейтинга';
                });
        });
    }

    fetch('/BookTalks-site/backend/books_list.php')
        .then(res => res.json())
        .then(data => {
            if (data.books) {
                renderBooks(data.books);
            }
        });
});
